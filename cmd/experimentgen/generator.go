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

type generator struct{}

func (g generator) generate(data []byte, variant string, outDir string) error {
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

	pbGoFiles, err := g.generateProtoGoFiles(
		outDir,
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

	if err := g.writeWorkflowFiles(outDir, changes.Overlays, variant); err != nil {
		return err
	}
	if err := g.writeEnumFiles(outDir, changes.Enums, variant); err != nil {
		return err
	}
	if err := g.writeServiceFiles(outDir, changes.Methods, variant, source.WorkflowPackage); err != nil {
		return err
	}
	if err := g.formatGeneratedFiles(outDir, changes.Overlays, changes.Enums, changes.Methods, pbGoFiles, variant); err != nil {
		return err
	}
	return nil
}

func (g generator) writeWorkflowFiles(
	outDir string,
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
	outFile := filepath.Join(outDir, "workflowservice", "v1", variant+"_overlay_experimental.go")
	return writeTemplate(outFile, overlayTemplate, data)
}

// removeStaleFiles deletes any previously generated file for this variant so
// re-running the generator after removing an experimental item doesn't leave
// orphaned files behind.
func removeStaleFiles(outDir, variant string) {
	candidates := []string{
		filepath.Join(outDir, "workflowservice", "v1", variant+"_messages_experimental.pb.go"),
		filepath.Join(outDir, "workflowservice", "v1", variant+"_service_experimental.go"),
		filepath.Join(outDir, "workflowservice", "v1", variant+"_service_experimental_grpc.pb.go"), // legacy name
		filepath.Join(outDir, "workflowservice", "v1", variant+"_overlay_experimental.go"),
		filepath.Join(outDir, "enums", "v1", variant+"_enum_experimental.go"),
	}
	for _, f := range candidates {
		_ = os.Remove(f)
	}
}

func (g generator) writeServiceFiles(
	outDir string,
	methods []methodInfo,
	variant string,
	protoPackage string,
) error {
	if len(methods) == 0 {
		return nil
	}
	outFile := filepath.Join(outDir, "workflowservice", "v1", variant+"_service_experimental.go")
	return writeTemplate(outFile, serviceTemplate, renderData{
		Service:      workflowService,
		VariantTitle: strcase.ToCamel(variant),
		ProtoPackage: protoPackage,
		Methods:      methods,
	})
}

func (g generator) writeEnumFiles(
	outDir string,
	enums []enumInfo,
	variant string,
) error {
	if len(enums) == 0 {
		return nil
	}
	outFile := filepath.Join(outDir, "enums", "v1", variant+"_enum_experimental.go")
	return writeTemplate(outFile, enumTemplate, renderData{
		Service: enumPackage,
		Enums:   enums,
	})
}

func (g generator) formatGeneratedFiles(
	outDir string,
	overlays []messageOverlay,
	enums []enumInfo,
	methods []methodInfo,
	pbGoFiles []string,
	variant string,
) error {
	gofmtPaths := make([]string, 0, len(pbGoFiles)+4)
	if len(overlays) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("workflowservice", "v1", variant+"_overlay_experimental.go"))
	}
	if len(enums) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("enums", "v1", variant+"_enum_experimental.go"))
	}
	if len(methods) > 0 {
		gofmtPaths = append(gofmtPaths, filepath.Join("workflowservice", "v1", variant+"_service_experimental.go"))
	}
	gofmtPaths = append(gofmtPaths, pbGoFiles...)
	return run(outDir, "gofmt", append([]string{"-w"}, gofmtPaths...)...)
}
