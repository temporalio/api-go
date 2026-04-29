package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

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
	WorkflowEnums          map[string]map[string]int32
	WorkflowPackage        string
	WorkflowRequestFile    *descriptorpb.FileDescriptorProto
	DescriptorFiles        []*descriptorpb.FileDescriptorProto
}

func (g generator) loadSnapshot(apiRepo string, sourceSHA string, stableRoot string) (descriptorSnapshot, error) {
	worktreeDir, err := os.MkdirTemp("", "experimentgen-api-*")
	if err != nil {
		return descriptorSnapshot{}, err
	}
	cleanup := func() {
		_ = g.run("", "git", "-C", apiRepo, "worktree", "remove", "--force", worktreeDir)
		_ = os.RemoveAll(worktreeDir)
	}
	defer cleanup()

	if err := g.run("", "git", "-C", apiRepo, "worktree", "add", "--detach", worktreeDir, sourceSHA); err != nil {
		return descriptorSnapshot{}, err
	}

	targets := []string{
		filepath.ToSlash(filepath.Join(stableRoot, "workflowservice/v1/service.proto")),
		filepath.ToSlash(filepath.Join(stableRoot, "workflowservice/v1/request_response.proto")),
		filepath.ToSlash(filepath.Join(stableRoot, "enums/v1/workflow.proto")),
	}

	descPath := filepath.Join(worktreeDir, "experimental-descriptor.pb")
	args := append([]string{
		"-I", worktreeDir,
		"--include_imports",
		"--descriptor_set_out=" + descPath,
	}, targets...)
	if err := g.run(worktreeDir, "protoc", args...); err != nil {
		return descriptorSnapshot{}, err
	}

	blob, err := os.ReadFile(descPath)
	if err != nil {
		return descriptorSnapshot{}, err
	}
	var set descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(blob, &set); err != nil {
		return descriptorSnapshot{}, err
	}
	return snapshotFromDescriptorSet(&set, stableRoot)
}

func snapshotFromDescriptorSet(set *descriptorpb.FileDescriptorSet, stableRoot string) (descriptorSnapshot, error) {
	snapshot := descriptorSnapshot{
		WorkflowServiceMethods: make(map[string]methodInfo),
		WorkflowMessages:       make(map[string]messageDef),
		WorkflowDescriptors:    make(map[string]*descriptorpb.DescriptorProto),
		WorkflowEnums:          make(map[string]map[string]int32),
		DescriptorFiles:        cloneFileDescriptors(set.File),
	}

	serviceName := filepath.ToSlash(filepath.Join(stableRoot, "workflowservice/v1/service.proto"))
	requestResponseName := filepath.ToSlash(filepath.Join(stableRoot, "workflowservice/v1/request_response.proto"))
	enumName := filepath.ToSlash(filepath.Join(stableRoot, "enums/v1/workflow.proto"))

	var serviceFile, requestFile, enumFile *descriptorpb.FileDescriptorProto
	for _, file := range set.File {
		switch file.GetName() {
		case serviceName:
			serviceFile = file
		case requestResponseName:
			requestFile = file
		case enumName:
			enumFile = file
		}
	}
	if serviceFile == nil || requestFile == nil || enumFile == nil {
		return descriptorSnapshot{}, fmt.Errorf("descriptor set missing required files")
	}

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

	for _, msg := range requestFile.GetMessageType() {
		snapshot.WorkflowMessages[msg.GetName()] = descriptorMessageDef(msg)
		snapshot.WorkflowDescriptors[msg.GetName()] = proto.Clone(msg).(*descriptorpb.DescriptorProto)
	}
	snapshot.WorkflowPackage = requestFile.GetPackage()
	snapshot.WorkflowRequestFile = proto.Clone(requestFile).(*descriptorpb.FileDescriptorProto)

	for _, enum := range enumFile.GetEnumType() {
		values := make(map[string]int32, len(enum.GetValue()))
		for _, value := range enum.GetValue() {
			values[value.GetName()] = value.GetNumber()
		}
		snapshot.WorkflowEnums[enum.GetName()] = values
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
