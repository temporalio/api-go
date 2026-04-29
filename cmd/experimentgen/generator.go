package main

import (
	"fmt"
	"path/filepath"
	"slices"
	"strings"
)

type generator struct {
	resolveStableVersion func(string) (moduleVersion, error)
	skipGoModTidy        bool
}

type moduleVersion struct {
	Version   string
	GoVersion string
}

func (g generator) generate(apiRepo string, sourceSHA string, variant string, outDir string) error {
	resolvedSHA, err := g.runOutput("", "git", "-C", apiRepo, "rev-parse", sourceSHA)
	if err != nil {
		return err
	}
	resolvedSHA = strings.TrimSpace(resolvedSHA)

	module, err := g.resolveStableVersion(resolvedSHA)
	if err != nil {
		return err
	}

	baseSHA, err := g.baseSHA(apiRepo, resolvedSHA)
	if err != nil {
		return err
	}

	if err := clearOutputDir(outDir); err != nil {
		return err
	}
	if err := writeTemplate(filepath.Join(outDir, "go.mod"), goModTemplate, renderData{
		ModulePath:    fmt.Sprintf("github.com/temporalio/api-go/experimental/%s", variant),
		StableVersion: module.Version,
		GoVersion:     module.GoVersion,
	}); err != nil {
		return err
	}

	stableRoot, err := g.detectStableRoot(apiRepo, resolvedSHA)
	if err != nil {
		return err
	}

	baseSnapshot, err := g.loadSnapshot(apiRepo, baseSHA, stableRoot)
	if err != nil {
		return err
	}
	sourceSnapshot, err := g.loadSnapshot(apiRepo, resolvedSHA, stableRoot)
	if err != nil {
		return err
	}

	changes := collectChanges(baseSnapshot, sourceSnapshot)
	if len(changes.Methods) == 0 && len(changes.Overlays) == 0 && len(changes.Enums) == 0 {
		return fmt.Errorf("no additive API changes found between %s and %s", baseSHA, resolvedSHA)
	}

	methodMessageNames := make(map[string]struct{}, len(changes.Methods)*2)
	for _, method := range changes.Methods {
		methodMessageNames[method.RequestName] = struct{}{}
		methodMessageNames[method.ResponseName] = struct{}{}
	}

	pbGoFiles, err := g.generateProtoGoFiles(
		outDir,
		sourceSnapshot,
		baseSnapshot,
		changes.Methods,
		methodMessageNames,
		changes.Overlays,
		variant,
	)
	if err != nil {
		return err
	}

	slices.SortFunc(changes.Overlays, func(a messageOverlay, b messageOverlay) int {
		if cmp := strings.Compare(a.StableMessage, b.StableMessage); cmp != 0 {
			return cmp
		}
		return a.FieldNumber - b.FieldNumber
	})
	slices.SortFunc(changes.Enums, func(a enumInfo, b enumInfo) int {
		if cmp := strings.Compare(a.StableEnum, b.StableEnum); cmp != 0 {
			return cmp
		}
		if cmp := strings.Compare(a.ValueName, b.ValueName); cmp != 0 {
			return cmp
		}
		return a.ValueNumber - b.ValueNumber
	})

	if err := g.writeWorkflowFiles(outDir, changes.Overlays); err != nil {
		return err
	}
	if err := g.writeEnumFiles(outDir, changes.Enums); err != nil {
		return err
	}
	if err := g.formatGeneratedFiles(outDir, changes.Overlays, changes.Enums, pbGoFiles); err != nil {
		return err
	}
	if !g.skipGoModTidy {
		if err := g.run(outDir, "go", "mod", "tidy"); err != nil {
			return err
		}
	}
	return nil
}

func (g generator) writeWorkflowFiles(
	outDir string,
	overlays []messageOverlay,
) error {
	if len(overlays) == 0 {
		return nil
	}
	data := renderData{
		Service:       workflowService,
		Overlays:      overlays,
		OverlayGroups: groupMessageOverlays(overlays),
	}
	return writeTemplate(filepath.Join(outDir, "workflowservice", "v1", "overlay.go"), overlayTemplate, data)
}

func (g generator) writeEnumFiles(
	outDir string,
	enums []enumInfo,
) error {
	if len(enums) == 0 {
		return nil
	}
	return writeTemplate(filepath.Join(outDir, "enums", "v1", "enum.go"), enumTemplate, renderData{
		Service: enumPackage,
		Enums:   enums,
	})
}

func (g generator) formatGeneratedFiles(
	outDir string,
	overlays []messageOverlay,
	enums []enumInfo,
	pbGoFiles []string,
) error {
	gofmtPaths := make([]string, 0, len(pbGoFiles)+4)
	if len(overlays) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("workflowservice", "v1", "overlay.go"))
	}
	if len(enums) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("enums", "v1", "enum.go"))
	}
	gofmtPaths = append(gofmtPaths, pbGoFiles...)
	return g.run(outDir, "gofmt", append([]string{"-w"}, gofmtPaths...)...)
}
