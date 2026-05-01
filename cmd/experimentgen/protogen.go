package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

func (g generator) generateProtoGoFiles(
	outDir string,
	source descriptorSnapshot,
	base descriptorSnapshot,
	methods []methodInfo,
	excluded map[string]struct{},
	overlays []messageOverlay,
	variant string,
) ([]string, error) {
	goFilesToGenerate, _, syntheticFiles, err := buildSyntheticFiles(source, base, methods, excluded, overlays, variant)
	if err != nil {
		return nil, err
	}

	// Replace the request_response and service files in the descriptor list with
	// base-only versions: strip experimental-only messages/methods so they don't
	// conflict with the synthetic files that re-declare those types/methods.
	experimentalMessageNames := make(map[string]struct{})
	for name := range source.WorkflowMessages {
		if _, inBase := base.WorkflowMessages[name]; !inBase {
			experimentalMessageNames[name] = struct{}{}
		}
	}
	experimentalMethodNames := make(map[string]struct{})
	for name := range source.WorkflowServiceMethods {
		if _, inBase := base.WorkflowServiceMethods[name]; !inBase {
			experimentalMethodNames[name] = struct{}{}
		}
	}
	serviceFileName := strings.TrimSuffix(source.WorkflowRequestFile.GetName(), "request_response.proto") + "service.proto"
	baseDescFiles := applyBaseToDescriptorFiles(
		cloneFileDescriptors(source.DescriptorFiles),
		source.WorkflowRequestFile.GetName(),
		serviceFileName,
		experimentalMessageNames,
		experimentalMethodNames,
		base.WorkflowDescriptors,
	)
	descriptorFiles := append(baseDescFiles, syntheticFiles...)
	var generated []string
	if len(goFilesToGenerate) > 0 {
		files, err := g.runProtoPlugin(outDir, "protoc-gen-go", "paths=source_relative", descriptorFiles, goFilesToGenerate)
		if err != nil {
			return nil, err
		}
		generated = append(generated, files...)
	}
	// gRPC stubs are generated via the service template (see writeServiceFiles),
	// not protoc-gen-go-grpc, to avoid redeclaring stable types in the same package.
	return generated, nil
}


func (g generator) runProtoPlugin(
	outDir string,
	plugin string,
	parameter string,
	descriptorFiles []*descriptorpb.FileDescriptorProto,
	filesToGenerate []string,
) ([]string, error) {
	request := &pluginpb.CodeGeneratorRequest{
		FileToGenerate: filesToGenerate,
		ProtoFile:      descriptorFiles,
		Parameter:      proto.String(parameter),
	}
	blob, err := proto.Marshal(request)
	if err != nil {
		return nil, err
	}
	var stdin bytes.Buffer
	stdin.Write(blob)

	cmd := exec.Command(plugin)
	cmd.Dir = outDir
	cmd.Env = append(slices.Clone(os.Environ()), "GOWORK=off")
	cmd.Stdin = &stdin
	out, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("%s failed: %w\n%s", plugin, err, out)
	}

	var response pluginpb.CodeGeneratorResponse
	if err := proto.Unmarshal(out, &response); err != nil {
		return nil, err
	}
	if response.GetError() != "" {
		return nil, fmt.Errorf("%s: %s", plugin, response.GetError())
	}
	generated := make([]string, 0, len(response.File))
	for _, file := range response.File {
		name := filepath.Clean(file.GetName())
		if err := os.MkdirAll(filepath.Dir(filepath.Join(outDir, name)), 0o755); err != nil {
			return nil, err
		}
		if err := os.WriteFile(filepath.Join(outDir, name), []byte(file.GetContent()), 0o644); err != nil {
			return nil, err
		}
		generated = append(generated, name)
	}
	return generated, nil
}

func buildSyntheticFiles(
	source descriptorSnapshot,
	base descriptorSnapshot,
	methods []methodInfo,
	excluded map[string]struct{},
	overlays []messageOverlay,
	variant string,
) ([]string, []string, []*descriptorpb.FileDescriptorProto, error) {
	goPackage := "go.temporal.io/api/experimental/workflowservice/v1;workflowservice"
	experimentalPackage := source.WorkflowPackage

	workflowFileName := filepath.ToSlash(filepath.Join("workflowservice", "v1", fmt.Sprintf("%s_messages.proto", variant)))
	goFilesToGenerate := make([]string, 0, 1)
	grpcFilesToGenerate := make([]string, 0, 1)
	syntheticFiles := make([]*descriptorpb.FileDescriptorProto, 0, 2)
	localNames := make(map[string]struct{})
	var serviceMethods []*descriptorpb.MethodDescriptorProto
	var messageNames []string

	for _, method := range methods {
		for _, name := range []string{method.RequestName, method.ResponseName} {
			if _, ok := source.WorkflowDescriptors[name]; !ok {
				return nil, nil, nil, fmt.Errorf("missing workflow message descriptor %q", name)
			}
			if _, ok := localNames[name]; ok {
				continue
			}
			localNames[name] = struct{}{}
			messageNames = append(messageNames, name)
		}

		serviceMethods = append(serviceMethods, &descriptorpb.MethodDescriptorProto{
			Name:       proto.String(method.Method),
			InputType:  proto.String("." + experimentalPackage + "." + method.RequestName),
			OutputType: proto.String("." + experimentalPackage + "." + method.ResponseName),
		})
	}

	for name := range source.WorkflowMessages {
		if _, exists := base.WorkflowMessages[name]; exists {
			continue
		}
		if _, exists := excluded[name]; exists {
			continue
		}
		if _, exists := localNames[name]; exists {
			continue
		}
		localNames[name] = struct{}{}
		messageNames = append(messageNames, name)
	}
	slices.Sort(messageNames)

	messages, err := buildWorkflowMessages(source, messageNames, localNames, experimentalPackage)
	if err != nil {
		return nil, nil, nil, err
	}
	overlayMessages, err := buildOverlayMessages(source, overlays, localNames, experimentalPackage)
	if err != nil {
		return nil, nil, nil, err
	}
	messages = append(messages, overlayMessages...)
	if len(messages) > 0 {
		syntheticFiles = append(syntheticFiles, syntheticWorkflowFile(workflowFileName, experimentalPackage, goPackage, workflowDependencies(source), messages...))
		goFilesToGenerate = append(goFilesToGenerate, workflowFileName)
	}
	if len(serviceMethods) > 0 {
		name := filepath.ToSlash(filepath.Join("workflowservice", "v1", fmt.Sprintf("%s_service.proto", variant)))
		syntheticFiles = append(syntheticFiles, syntheticServiceFile(name, source.WorkflowPackage, goPackage, []string{workflowFileName}, serviceMethods))
		grpcFilesToGenerate = append(grpcFilesToGenerate, name)
	}

	return goFilesToGenerate, grpcFilesToGenerate, syntheticFiles, nil
}

func buildWorkflowMessages(
	source descriptorSnapshot,
	messageNames []string,
	localNames map[string]struct{},
	experimentalPackage string,
) ([]*descriptorpb.DescriptorProto, error) {
	messages := make([]*descriptorpb.DescriptorProto, 0, len(messageNames))
	for _, name := range messageNames {
		desc, ok := source.WorkflowDescriptors[name]
		if !ok {
			return nil, fmt.Errorf("missing workflow message descriptor %q", name)
		}
		clone := proto.Clone(desc).(*descriptorpb.DescriptorProto)
		rewriteLocalTypeNames(clone, localNames, experimentalPackage)
		messages = append(messages, clone)
	}
	return messages, nil
}

func buildOverlayMessages(
	source descriptorSnapshot,
	overlays []messageOverlay,
	localNames map[string]struct{},
	experimentalPackage string,
) ([]*descriptorpb.DescriptorProto, error) {
	groups := groupMessageOverlays(overlays)
	messages := make([]*descriptorpb.DescriptorProto, 0, len(groups))
	for _, group := range groups {
		msg := &descriptorpb.DescriptorProto{
			Name: proto.String(group.OverlayMessage),
		}
		for _, overlay := range group.Fields {
			sourceMsg, ok := source.WorkflowDescriptors[overlay.SourceMessage]
			if !ok {
				return nil, fmt.Errorf("missing workflow message descriptor %q", overlay.SourceMessage)
			}
			sourceField, mapEntry, ok := findSourceField(sourceMsg, overlay.SourceName, overlay.SourceNumber)
			if !ok {
				return nil, fmt.Errorf("missing workflow field descriptor %q.%s", overlay.SourceMessage, overlay.SourceName)
			}

			field := proto.Clone(sourceField).(*descriptorpb.FieldDescriptorProto)
			field.Number = proto.Int32(int32(overlay.FieldNumber))
			if mapEntry != nil {
				nested := proto.Clone(mapEntry).(*descriptorpb.DescriptorProto)
				rewriteLocalTypeNames(nested, localNames, experimentalPackage)
				msg.NestedType = append(msg.NestedType, nested)
				field.TypeName = proto.String("." + experimentalPackage + "." + group.OverlayMessage + "." + nested.GetName())
			} else {
				rewriteFieldTypeName(field, localNames, experimentalPackage)
			}
			msg.Field = append(msg.Field, field)
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func findSourceField(
	msg *descriptorpb.DescriptorProto,
	name string,
	number int,
) (*descriptorpb.FieldDescriptorProto, *descriptorpb.DescriptorProto, bool) {
	var mapEntry *descriptorpb.DescriptorProto
	nestedTypes := make(map[string]*descriptorpb.DescriptorProto, len(msg.GetNestedType()))
	for _, nested := range msg.GetNestedType() {
		nestedTypes[nested.GetName()] = nested
	}
	for _, field := range msg.GetField() {
		if field.GetName() != name || int(field.GetNumber()) != number {
			continue
		}
		if field.GetLabel() == descriptorpb.FieldDescriptorProto_LABEL_REPEATED &&
			field.GetType() == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
			if nested, ok := nestedTypes[trimDescriptorName(field.GetTypeName())]; ok && nested.GetOptions().GetMapEntry() {
				mapEntry = nested
			}
		}
		return field, mapEntry, true
	}
	return nil, nil, false
}

func syntheticWorkflowFile(
	name string,
	pkg string,
	goPackage string,
	dependencies []string,
	messages ...*descriptorpb.DescriptorProto,
) *descriptorpb.FileDescriptorProto {
	return &descriptorpb.FileDescriptorProto{
		Name:       proto.String(name),
		Package:    proto.String(pkg),
		Syntax:     proto.String("proto3"),
		Dependency: slices.Clone(dependencies),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String(goPackage),
		},
		MessageType: messages,
	}
}

func syntheticServiceFile(
	name string,
	pkg string,
	goPackage string,
	dependencies []string,
	methods []*descriptorpb.MethodDescriptorProto,
) *descriptorpb.FileDescriptorProto {
	return &descriptorpb.FileDescriptorProto{
		Name:       proto.String(name),
		Package:    proto.String(pkg),
		Syntax:     proto.String("proto3"),
		Dependency: slices.Clone(dependencies),
		Options: &descriptorpb.FileOptions{
			GoPackage: proto.String(goPackage),
		},
		Service: []*descriptorpb.ServiceDescriptorProto{
			{
				Name:   proto.String("WorkflowService"),
				Method: methods,
			},
		},
	}
}

func rewriteLocalTypeNames(
	msg *descriptorpb.DescriptorProto,
	localNames map[string]struct{},
	targetPackage string,
) {
	for _, field := range msg.GetField() {
		rewriteFieldTypeName(field, localNames, targetPackage)
	}
	for _, nested := range msg.GetNestedType() {
		rewriteLocalTypeNames(nested, localNames, targetPackage)
	}
}

func rewriteFieldTypeName(
	field *descriptorpb.FieldDescriptorProto,
	localNames map[string]struct{},
	targetPackage string,
) {
	typeName := field.GetTypeName()
	if typeName == "" {
		return
	}
	trimmed := trimDescriptorName(typeName)
	if _, ok := localNames[trimmed]; ok {
		field.TypeName = proto.String("." + targetPackage + "." + trimmed)
	}
}

// applyBaseToDescriptorFiles replaces the request_response and service files in
// the descriptor list with base-only versions: strips experimental-only messages/methods
// so they don't conflict with the synthetic files that re-declare those types/methods.
func applyBaseToDescriptorFiles(
	files []*descriptorpb.FileDescriptorProto,
	requestFileName string,
	serviceFileName string,
	experimentalMessageNames map[string]struct{},
	experimentalMethodNames map[string]struct{},
	baseDescriptors map[string]*descriptorpb.DescriptorProto,
) []*descriptorpb.FileDescriptorProto {
	for _, file := range files {
		switch file.GetName() {
		case requestFileName:
			kept := make([]*descriptorpb.DescriptorProto, 0, len(file.GetMessageType()))
			for _, msg := range file.GetMessageType() {
				if _, strip := experimentalMessageNames[msg.GetName()]; strip {
					// Skip experimental-only messages entirely.
					continue
				}
				// Use the base descriptor (which has experimental fields removed).
				if baseDesc, ok := baseDescriptors[msg.GetName()]; ok {
					kept = append(kept, proto.Clone(baseDesc).(*descriptorpb.DescriptorProto))
				} else {
					kept = append(kept, msg)
				}
			}
			file.MessageType = kept
		case serviceFileName:
			// Remove the WorkflowService service definition entirely to prevent a
			// naming conflict with the synthetic service file (which defines
			// WorkflowService in the same proto package).
			kept := file.GetService()[:0]
			for _, svc := range file.GetService() {
				if svc.GetName() != "WorkflowService" {
					kept = append(kept, svc)
				}
			}
			file.Service = kept
		}
	}
	return files
}

func workflowDependencies(source descriptorSnapshot) []string {
	dependencies := make([]string, 0, len(source.WorkflowRequestFile.GetDependency())+1)
	seen := make(map[string]struct{}, len(source.WorkflowRequestFile.GetDependency())+1)
	for _, dependency := range append([]string{source.WorkflowRequestFile.GetName()}, source.WorkflowRequestFile.GetDependency()...) {
		if dependency == "" {
			continue
		}
		if _, ok := seen[dependency]; ok {
			continue
		}
		// experimental.proto contributes no Go types; including it causes protoc-gen-go
		// to emit an import for its go_package root path which has no Go files.
		if dependency == "temporal/api/experimental.proto" {
			continue
		}
		seen[dependency] = struct{}{}
		dependencies = append(dependencies, dependency)
	}
	return dependencies
}
