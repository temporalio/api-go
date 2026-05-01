package main

import (
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

const experimentalFieldNumber = 77001

type methodInfo struct {
	Method       string
	RequestName  string
	ResponseName string
}

type protoFieldInfo struct {
	Name   string
	Number int
}

type messageDef struct {
	Name   string
	Fields []protoFieldInfo
}

type descriptorSnapshot struct {
	WorkflowServiceMethods map[string]methodInfo
	WorkflowMessages       map[string]messageDef
	WorkflowDescriptors    map[string]*descriptorpb.DescriptorProto
	WorkflowMessageFiles   map[string]*descriptorpb.FileDescriptorProto
	WorkflowEnums          map[string]map[string]int32
	WorkflowPackage        string
	WorkflowRequestFile    *descriptorpb.FileDescriptorProto
	DescriptorFiles        []*descriptorpb.FileDescriptorProto
}

// experimentalOptionValue extracts the experimental option value (field number 77001)
// from the unknown-field bytes of an Options message. The option is not linked into
// this binary, so proto preserves it as unknown fields.
func experimentalOptionValue(opts proto.Message) string {
	if opts == nil {
		return ""
	}
	unknown := opts.ProtoReflect().GetUnknown()
	for len(unknown) > 0 {
		num, typ, n := protowire.ConsumeTag(unknown)
		if n < 0 {
			break
		}
		unknown = unknown[n:]
		if num == experimentalFieldNumber && typ == protowire.BytesType {
			val, n := protowire.ConsumeBytes(unknown)
			if n < 0 {
				break
			}
			return string(val)
		}
		n = protowire.ConsumeFieldValue(num, typ, unknown)
		if n < 0 {
			break
		}
		unknown = unknown[n:]
	}
	return ""
}

// detectStableRoot looks for workflowservice/v1/request_response.proto with or without
// a "temporal/api/" prefix in the descriptor set, and returns the prefix.
func detectStableRoot(fds *descriptorpb.FileDescriptorSet) string {
	for _, candidate := range []string{"temporal/api", ""} {
		path := "workflowservice/v1/request_response.proto"
		if candidate != "" {
			path = filepath.ToSlash(filepath.Join(candidate, path))
		}
		for _, f := range fds.GetFile() {
			if f.GetName() == path {
				return candidate
			}
		}
	}
	return ""
}

// loadSnapshotsFromAnnotations builds a source snapshot (full) and a base snapshot
// (filtered to exclude items annotated with the given variant) from a FileDescriptorSet.
func loadSnapshotsFromAnnotations(fds *descriptorpb.FileDescriptorSet, stableRoot string, variant string) (base, source descriptorSnapshot, err error) {
	source, err = snapshotFromDescriptorSet(fds, stableRoot)
	if err != nil {
		return descriptorSnapshot{}, descriptorSnapshot{}, err
	}
	base, err = filterSnapshotToBase(source, fds, stableRoot, variant)
	if err != nil {
		return descriptorSnapshot{}, descriptorSnapshot{}, err
	}
	return base, source, nil
}

// filterSnapshotToBase builds a base snapshot by removing all items annotated
// with the given variant from the source snapshot.
func filterSnapshotToBase(source descriptorSnapshot, fds *descriptorpb.FileDescriptorSet, stableRoot string, variant string) (descriptorSnapshot, error) {
	base := descriptorSnapshot{
		WorkflowServiceMethods: make(map[string]methodInfo),
		WorkflowMessages:       make(map[string]messageDef),
		WorkflowDescriptors:    make(map[string]*descriptorpb.DescriptorProto),
		WorkflowMessageFiles:   make(map[string]*descriptorpb.FileDescriptorProto),
		WorkflowEnums:          make(map[string]map[string]int32),
		WorkflowPackage:        source.WorkflowPackage,
		DescriptorFiles:        cloneFileDescriptors(source.DescriptorFiles),
	}
	if source.WorkflowRequestFile != nil {
		base.WorkflowRequestFile = proto.Clone(source.WorkflowRequestFile).(*descriptorpb.FileDescriptorProto)
	}

	// Build lookup maps from the raw descriptor set for annotation checking.
	// Key: method name -> *descriptorpb.MethodDescriptorProto
	rawMethods := make(map[string]*descriptorpb.MethodDescriptorProto)
	// Key: message name -> *descriptorpb.DescriptorProto
	rawMessages := make(map[string]*descriptorpb.DescriptorProto)
	// Key: "messageName.fieldName" -> *descriptorpb.FieldDescriptorProto
	rawFields := make(map[string]*descriptorpb.FieldDescriptorProto)
	// Key: "enumName.valueName" -> *descriptorpb.EnumValueDescriptorProto
	rawEnumValues := make(map[string]*descriptorpb.EnumValueDescriptorProto)

	for _, file := range fds.GetFile() {
		if strings.HasSuffix(file.GetName(), "workflowservice/v1/service.proto") {
			for _, svc := range file.GetService() {
				if svc.GetName() != "WorkflowService" {
					continue
				}
				for _, m := range svc.GetMethod() {
					rawMethods[m.GetName()] = m
				}
			}
		}
		for _, msg := range file.GetMessageType() {
			rawMessages[msg.GetName()] = msg
			for _, f := range msg.GetField() {
				rawFields[msg.GetName()+"."+f.GetName()] = f
			}
		}
		for _, enum := range file.GetEnumType() {
			for _, v := range enum.GetValue() {
				rawEnumValues[enum.GetName()+"."+v.GetName()] = v
			}
		}
	}

	// Copy methods that are NOT annotated with this variant.
	for name, method := range source.WorkflowServiceMethods {
		if raw, ok := rawMethods[name]; ok {
			if experimentalOptionValue(raw.GetOptions()) == variant {
				continue
			}
		}
		base.WorkflowServiceMethods[name] = method
	}

	// Copy messages that are NOT annotated with this variant;
	// for copied messages, remove fields annotated with this variant.
	for msgName, msgDef := range source.WorkflowMessages {
		if raw, ok := rawMessages[msgName]; ok {
			if experimentalOptionValue(raw.GetOptions()) == variant {
				continue
			}
		}
		// Build filtered message def (remove fields annotated with variant).
		filteredFields := make([]protoFieldInfo, 0, len(msgDef.Fields))
		for _, field := range msgDef.Fields {
			key := msgName + "." + field.Name
			if rawField, ok := rawFields[key]; ok {
				if experimentalOptionValue(rawField.GetOptions()) == variant {
					continue
				}
			}
			filteredFields = append(filteredFields, field)
		}
		base.WorkflowMessages[msgName] = messageDef{
			Name:   msgDef.Name,
			Fields: filteredFields,
		}
		if sourceFile, ok := source.WorkflowMessageFiles[msgName]; ok {
			base.WorkflowMessageFiles[msgName] = proto.Clone(sourceFile).(*descriptorpb.FileDescriptorProto)
		}
		// Also copy descriptor with filtered fields and nested types.
		if rawMsg, ok := rawMessages[msgName]; ok {
			cloneMsg := proto.Clone(rawMsg).(*descriptorpb.DescriptorProto)
			// Collect nested type names referenced by stripped fields so we can remove them too.
			strippedNestedTypes := make(map[string]struct{})
			filteredDescFields := cloneMsg.GetField()[:0]
			for _, f := range cloneMsg.GetField() {
				key := msgName + "." + f.GetName()
				if rawField, ok2 := rawFields[key]; ok2 {
					if experimentalOptionValue(rawField.GetOptions()) == variant {
						// Track any nested type name (map entry) this field references.
						if f.GetTypeName() != "" {
							strippedNestedTypes[trimDescriptorName(f.GetTypeName())] = struct{}{}
						}
						continue
					}
				}
				filteredDescFields = append(filteredDescFields, f)
			}
			cloneMsg.Field = filteredDescFields
			// Remove nested types that belonged only to stripped fields.
			if len(strippedNestedTypes) > 0 {
				filteredNested := cloneMsg.GetNestedType()[:0]
				for _, nested := range cloneMsg.GetNestedType() {
					if _, strip := strippedNestedTypes[nested.GetName()]; !strip {
						filteredNested = append(filteredNested, nested)
					}
				}
				cloneMsg.NestedType = filteredNested
			}
			base.WorkflowDescriptors[msgName] = cloneMsg
		} else if desc, ok2 := source.WorkflowDescriptors[msgName]; ok2 {
			base.WorkflowDescriptors[msgName] = proto.Clone(desc).(*descriptorpb.DescriptorProto)
		}
	}

	// Copy enum values that are NOT annotated with this variant.
	for enumName2, sourceValues := range source.WorkflowEnums {
		filteredValues := make(map[string]int32)
		for valueName, number := range sourceValues {
			key := enumName2 + "." + valueName
			if rawVal, ok := rawEnumValues[key]; ok {
				if experimentalOptionValue(rawVal.GetOptions()) == variant {
					continue
				}
			}
			filteredValues[valueName] = number
		}
		base.WorkflowEnums[enumName2] = filteredValues
	}

	return base, nil
}

func snapshotFromDescriptorSet(set *descriptorpb.FileDescriptorSet, stableRoot string) (descriptorSnapshot, error) {
	snapshot := descriptorSnapshot{
		WorkflowServiceMethods: make(map[string]methodInfo),
		WorkflowMessages:       make(map[string]messageDef),
		WorkflowDescriptors:    make(map[string]*descriptorpb.DescriptorProto),
		WorkflowMessageFiles:   make(map[string]*descriptorpb.FileDescriptorProto),
		WorkflowEnums:          make(map[string]map[string]int32),
		DescriptorFiles:        cloneFileDescriptors(set.File),
	}

	serviceName := filepath.ToSlash(filepath.Join(stableRoot, "workflowservice/v1/service.proto"))
	requestResponseName := filepath.ToSlash(filepath.Join(stableRoot, "workflowservice/v1/request_response.proto"))

	var serviceFile, requestFile *descriptorpb.FileDescriptorProto
	for _, file := range set.File {
		switch file.GetName() {
		case serviceName:
			serviceFile = file
		case requestResponseName:
			requestFile = file
		}
	}
	if serviceFile != nil {
		for _, service := range serviceFile.GetService() {
			if service.GetName() != "WorkflowService" {
				continue
			}
			for _, method := range service.GetMethod() {
				name := method.GetName()
				snapshot.WorkflowServiceMethods[name] = methodInfo{
					Method:       name,
					RequestName:  trimDescriptorName(method.GetInputType()),
					ResponseName: trimDescriptorName(method.GetOutputType()),
				}
			}
		}
	}

	if requestFile != nil {
		snapshot.WorkflowPackage = requestFile.GetPackage()
		snapshot.WorkflowRequestFile = proto.Clone(requestFile).(*descriptorpb.FileDescriptorProto)
	}
	for _, file := range set.File {
		if !strings.HasPrefix(file.GetPackage(), "temporal.api.") {
			continue
		}
		for _, msg := range file.GetMessageType() {
			snapshot.WorkflowMessages[msg.GetName()] = descriptorMessageDef(msg)
			snapshot.WorkflowDescriptors[msg.GetName()] = proto.Clone(msg).(*descriptorpb.DescriptorProto)
			snapshot.WorkflowMessageFiles[msg.GetName()] = proto.Clone(file).(*descriptorpb.FileDescriptorProto)
		}
		for _, enum := range file.GetEnumType() {
			values := make(map[string]int32, len(enum.GetValue()))
			for _, value := range enum.GetValue() {
				values[value.GetName()] = value.GetNumber()
			}
			snapshot.WorkflowEnums[enum.GetName()] = values
		}
	}

	return snapshot, nil
}

func descriptorMessageDef(msg *descriptorpb.DescriptorProto) messageDef {
	fields := make([]protoFieldInfo, 0, len(msg.GetField()))
	for _, field := range msg.GetField() {
		fields = append(fields, protoFieldInfo{
			Name:   field.GetName(),
			Number: int(field.GetNumber()),
		})
	}
	return messageDef{
		Name:   msg.GetName(),
		Fields: fields,
	}
}

func trimDescriptorName(name string) string {
	name = strings.TrimPrefix(name, ".")
	if idx := strings.LastIndex(name, "."); idx >= 0 {
		return name[idx+1:]
	}
	return name
}

func cloneFileDescriptors(files []*descriptorpb.FileDescriptorProto) []*descriptorpb.FileDescriptorProto {
	cloned := make([]*descriptorpb.FileDescriptorProto, 0, len(files))
	for _, file := range files {
		cloned = append(cloned, proto.Clone(file).(*descriptorpb.FileDescriptorProto))
	}
	return cloned
}
