package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"go.temporal.io/api/internal/strcase"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type generator struct {
	resolveStableVersion func(apiRepo string) (moduleVersion, error)
	skipModTidy          bool
}

// moduleVersion holds version info for the stable api module.
type moduleVersion struct {
	Tag       string // e.g. "v1.2.3"
	GoVersion string // e.g. "1.21"
}

func (g generator) generate(data []byte, variant string, outDir string) error {
	expDir := filepath.Join(outDir, "experimental")
	removeStaleFiles(outDir, variant)

	var fds descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(data, &fds); err != nil {
		return fmt.Errorf("unmarshal descriptor set: %w", err)
	}

	stableRoot := detectStableRoot(&fds)

	base, source, err := loadSnapshotsFromAnnotations(&fds, stableRoot, variant)
	if err != nil {
		return err
	}

	changes := collectChanges(base, source)
	if len(changes.Methods) == 0 && len(changes.Overlays) == 0 && len(changes.Enums) == 0 {
		return fmt.Errorf("no additive API changes found for variant %q", variant)
	}

	methodMessageNames := make(map[string]struct{}, len(changes.Methods)*2)
	for _, method := range changes.Methods {
		methodMessageNames[method.RequestName] = struct{}{}
		methodMessageNames[method.ResponseName] = struct{}{}
	}

	// Resolve stable module version for go.mod generation.
	resolve := g.resolveStableVersion
	if resolve == nil {
		resolve = resolveLocalModuleVersion
	}
	modVer, err := resolve(outDir)
	if err != nil {
		return fmt.Errorf("resolve stable version: %w", err)
	}

	// Write the experimental/go.mod before generating Go files.
	if err := g.writeGoMod(expDir, modVer); err != nil {
		return err
	}

	pbGoFiles, err := g.generateProtoGoFiles(
		expDir,
		source,
		base,
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

	if err := g.writeWorkflowFiles(expDir, changes.Overlays, variant); err != nil {
		return err
	}
	if err := g.writeEnumFiles(expDir, changes.Enums, variant); err != nil {
		return err
	}
	if err := g.writeServiceFiles(expDir, changes.Methods, variant, source.WorkflowPackage); err != nil {
		return err
	}
	if err := g.formatGeneratedFiles(expDir, changes.Overlays, changes.Enums, changes.Methods, pbGoFiles, variant); err != nil {
		return err
	}
	if !g.skipModTidy {
		if err := run(expDir, "go", "mod", "tidy"); err != nil {
			return fmt.Errorf("go mod tidy: %w", err)
		}
	}
	return nil
}

func (g generator) writeGoMod(expDir string, ver moduleVersion) error {
	if err := os.MkdirAll(expDir, 0o755); err != nil {
		return err
	}
	return writeTemplate(filepath.Join(expDir, "go.mod"), goModTemplate, renderData{
		StableVersion: ver.Tag,
		GoVersion:     ver.GoVersion,
	})
}

func (g generator) writeWorkflowFiles(
	expDir string,
	overlays []messageOverlay,
	variant string,
) error {
	if len(overlays) == 0 {
		return nil
	}
	data := renderData{
		Service:       workflowService,
		Overlays:      overlays,
		OverlayGroups: groupMessageOverlays(overlays),
	}
	outFile := filepath.Join(expDir, "workflowservice", "v1", variant+"_overlay.go")
	return writeTemplate(outFile, overlayTemplate, data)
}

// removeStaleFiles deletes any previously generated file for this variant so
// re-running the generator after removing an experimental item doesn't leave
// orphaned files behind.
func removeStaleFiles(outDir, variant string) {
	// Old paths (build-tag-gated files in main module root).
	oldCandidates := []string{
		filepath.Join(outDir, "workflowservice", "v1", variant+"_messages_experimental.pb.go"),
		filepath.Join(outDir, "workflowservice", "v1", variant+"_service_experimental.go"),
		filepath.Join(outDir, "workflowservice", "v1", variant+"_service_experimental_grpc.pb.go"), // legacy name
		filepath.Join(outDir, "workflowservice", "v1", variant+"_overlay_experimental.go"),
		filepath.Join(outDir, "enums", "v1", variant+"_enum_experimental.go"),
	}
	for _, f := range oldCandidates {
		_ = os.Remove(f)
	}

	// New paths (under experimental/ subdirectory, no _experimental suffix).
	expDir := filepath.Join(outDir, "experimental")
	newCandidates := []string{
		filepath.Join(expDir, "workflowservice", "v1", variant+"_messages.pb.go"),
		filepath.Join(expDir, "workflowservice", "v1", variant+"_service.go"),
		filepath.Join(expDir, "workflowservice", "v1", variant+"_overlay.go"),
		filepath.Join(expDir, "enums", "v1", variant+"_enum.go"),
	}
	for _, f := range newCandidates {
		_ = os.Remove(f)
	}
}

func (g generator) writeServiceFiles(
	expDir string,
	methods []methodInfo,
	variant string,
	protoPackage string,
) error {
	if len(methods) == 0 {
		return nil
	}
	outFile := filepath.Join(expDir, "workflowservice", "v1", variant+"_service.go")
	return writeTemplate(outFile, serviceTemplate, renderData{
		Service:      workflowService,
		VariantTitle: strcase.ToCamel(variant),
		ProtoPackage: protoPackage,
		Methods:      methods,
	})
}

func (g generator) writeEnumFiles(
	expDir string,
	enums []enumInfo,
	variant string,
) error {
	if len(enums) == 0 {
		return nil
	}
	outFile := filepath.Join(expDir, "enums", "v1", variant+"_enum.go")
	return writeTemplate(outFile, enumTemplate, renderData{
		Service: enumPackage,
		Enums:   enums,
	})
}

func (g generator) formatGeneratedFiles(
	expDir string,
	overlays []messageOverlay,
	enums []enumInfo,
	methods []methodInfo,
	pbGoFiles []string,
	variant string,
) error {
	gofmtPaths := make([]string, 0, len(pbGoFiles)+4)
	if len(overlays) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("workflowservice", "v1", variant+"_overlay.go"))
	}
	if len(enums) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("enums", "v1", variant+"_enum.go"))
	}
	if len(methods) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("workflowservice", "v1", variant+"_service.go"))
	}
	gofmtPaths = append(gofmtPaths, pbGoFiles...)
	return run(expDir, "gofmt", append([]string{"-w"}, gofmtPaths...)...)
}

// resolveLocalModuleVersion returns the version tag and Go version for the
// stable api-go module. It first tries `git describe` to get the latest tag,
// then falls back to reading go.mod for the Go version.
func resolveLocalModuleVersion(apiRepo string) (moduleVersion, error) {
	goVersion := "1.21" // sensible default
	// Try to read go version from go.mod in apiRepo.
	if ver, err := readGoVersionFromMod(apiRepo); err == nil && ver != "" {
		goVersion = ver
	}

	tag, err := runOutput(apiRepo, "git", "-C", apiRepo, "describe", "--tags", "--match", "v*", "--abbrev=0")
	if err != nil {
		// Fall back to a placeholder if git is not available or no tags exist.
		return moduleVersion{Tag: "v0.0.0", GoVersion: goVersion}, nil
	}
	tag = strings.TrimSpace(tag)
	return moduleVersion{Tag: tag, GoVersion: goVersion}, nil
}

// readGoVersionFromMod reads the `go X.Y` directive from go.mod in dir.
func readGoVersionFromMod(dir string) (string, error) {
	data, err := os.ReadFile(filepath.Join(dir, "go.mod"))
	if err != nil {
		return "", err
	}
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "go ") {
			return strings.TrimPrefix(line, "go "), nil
		}
	}
	return "", nil
}
