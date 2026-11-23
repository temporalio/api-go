package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestAPINextProxy(t *testing.T) {
	root, err := filepath.Abs("../..")
	if err != nil {
		t.Fatal(err)
	}
	stage := t.TempDir()
	if err := os.CopyFS(stage, os.DirFS(filepath.Join(root, "api_next"))); err != nil {
		t.Fatal(err)
	}
	run := func(dir, name string, args ...string) {
		t.Helper()
		cmd := exec.Command(name, args...)
		cmd.Dir = dir
		if output, err := cmd.CombinedOutput(); err != nil {
			t.Fatalf("%s %v: %v\n%s", name, args, err, output)
		}
	}
	descriptorPath := filepath.Join(stage, "descriptor_set.pb")
	raw, err := os.ReadFile(descriptorPath)
	if err != nil {
		t.Fatalf("run make api-next before this test: %v", err)
	}
	var set descriptorpb.FileDescriptorSet
	if err := proto.Unmarshal(raw, &set); err != nil {
		t.Fatal(err)
	}
	field := func(name string, number int32, typ string) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{
			Name: proto.String(name), Number: proto.Int32(number),
			Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
			Type:  descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(), TypeName: proto.String(typ),
		}
	}
	feature := func(value string) []byte {
		return protowire.AppendString(protowire.AppendTag(nil, 98008, protowire.BytesType), value)
	}
	var input string
	for _, file := range set.File {
		if file.GetPackage() != "temporal.api.update.v1" {
			continue
		}
		input = file.GetName()
		subtype := &descriptorpb.DescriptorProto{
			Name: proto.String("DraftPayloadSubtype"), Options: &descriptorpb.MessageOptions{},
			Field: []*descriptorpb.FieldDescriptorProto{
				field("payload", 1, ".temporal.api.common.v1.Payload"),
				field("payloads", 2, ".temporal.api.common.v1.Payloads"),
				field("failure", 3, ".temporal.api.failure.v1.Failure"),
				field("cycle", 4, ".temporal.api.update.v1.DraftPayloadCycle"),
				field("nested", 5, ".temporal.api.update.v1.DraftPayloadSubtype.Nested"),
			},
			NestedType: []*descriptorpb.DescriptorProto{{
				Name: proto.String("Nested"), Field: []*descriptorpb.FieldDescriptorProto{field("payload", 1, ".temporal.api.common.v1.Payload")},
			}},
		}
		subtype.Options.ProtoReflect().SetUnknown(feature("message-feature"))
		file.MessageType = append(file.MessageType, subtype, &descriptorpb.DescriptorProto{
			Name: proto.String("DraftPayloadCycle"), Field: []*descriptorpb.FieldDescriptorProto{field("back", 1, ".temporal.api.update.v1.DraftPayloadSubtype")},
		})
		for _, message := range file.MessageType {
			if message.GetName() == "Request" {
				f := field("draft_payload", 100, ".temporal.api.update.v1.DraftPayloadSubtype")
				f.Options = &descriptorpb.FieldOptions{}
				f.Options.ProtoReflect().SetUnknown(feature("field-feature"))
				message.Field = append(message.Field, f)
			}
		}
	}
	if input == "" {
		t.Fatal("update descriptors missing")
	}
	raw, err = proto.Marshal(&set)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(descriptorPath, raw, 0600); err != nil {
		t.Fatal(err)
	}
	helper := filepath.Join(t.TempDir(), "protoc-gen-go-helpers")
	run(root, "go", "build", "-o", helper, "./cmd/protoc-gen-go-helpers")
	generated := t.TempDir()
	run(stage, "protoc", "--descriptor_set_in="+descriptorPath,
		"--plugin=protoc-gen-go-helpers="+helper, "--go_out=paths=source_relative:"+generated,
		"--go-helpers_out=paths=source_relative,draft_omission=true:"+generated, input)
	entries, err := os.ReadDir(filepath.Join(generated, filepath.Dir(input)))
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range entries {
		content, err := os.ReadFile(filepath.Join(generated, filepath.Dir(input), entry.Name()))
		if err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(stage, "update/v1", entry.Name()), content, 0600); err != nil {
			t.Fatal(err)
		}
	}
	regression, err := os.ReadFile("testdata/api_next_proxy_test.go")
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stage, "update/v1/proxy_regression_test.go"), regression, 0600); err != nil {
		t.Fatal(err)
	}
	generatorDir := filepath.Join(stage, "cmd/proxygenerator")
	run(generatorDir, "go", "run", "./")
	run(stage, "go", "test", "./update/v1", "-run", "TestDraftPayloadVisitors", "-count=1")
	run(generatorDir, "go", "run", "./", "-verifyOnly")
	for _, artifact := range []string{"interceptor.go", "service.go", "request_header.go"} {
		t.Run(artifact, func(t *testing.T) {
			path := filepath.Join(stage, "proxy", artifact)
			original, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, append(append([]byte{}, original...), []byte("\n// stale\n")...), 0600); err != nil {
				t.Fatal(err)
			}
			defer func() {
				if err := os.WriteFile(path, original, 0600); err != nil {
					t.Error(err)
				}
			}()
			cmd := exec.Command("go", "run", "./", "-verifyOnly")
			cmd.Dir = generatorDir
			if output, err := cmd.CombinedOutput(); err == nil || !strings.Contains(string(output), "generated file does not match existing file: ../../proxy/"+artifact) {
				t.Fatalf("stale %s verification: %v\n%s", artifact, err, output)
			}
		})
	}
}
