package main

import (
	"slices"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

// TestGatherMessagesContainingTargetsRecursive checks that cyclic references can lead to a payload.
func TestGatherMessagesContainingTargetsRecursive(t *testing.T) {
	field := func(name, typ string, number int32) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{
			Name:     proto.String(name),
			Number:   proto.Int32(number),
			TypeName: proto.String(typ),
			Type:     descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(),
			Label:    descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
		}
	}
	file, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name: proto.String("recursive.proto"), Package: proto.String("recursive"), Syntax: proto.String("proto3"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: proto.String("A"), Field: []*descriptorpb.FieldDescriptorProto{
				field("b", ".recursive.B", 1),
				field("payload", ".recursive.Payload", 2),
			}},
			{Name: proto.String("B"), Field: []*descriptorpb.FieldDescriptorProto{
				field("a", ".recursive.A", 1),
			}},
			{Name: proto.String("NoPayload"), Field: []*descriptorpb.FieldDescriptorProto{
				field("first", ".recursive.Leaf", 1),
				field("second", ".recursive.Leaf", 2),
			}},
			{Name: proto.String("Leaf")},
			{Name: proto.String("Payload")},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	var files protoregistry.Files
	if err := files.RegisterFile(file); err != nil {
		t.Fatal(err)
	}
	matches, err := gatherMessagesContainingTargets(&files, []string{"recursive.Payload"})
	if err != nil {
		t.Fatal(err)
	}

	var names []protoreflect.Name
	for _, match := range matches {
		names = append(names, match.Name())
	}
	if want := []protoreflect.Name{"A", "B"}; !slices.Equal(names, want) {
		t.Fatalf("messages containing a payload = %v, want %v", names, want)
	}
}

// TestGatherMatchesKeepsRecursivePayloadGetter checks that the generator retains a recursive getter that can lead to a payload.
func TestGatherMatchesKeepsRecursivePayloadGetter(t *testing.T) {
	payloads, err := lookupTypes("go.temporal.io/api/cmd/proxygenerator/testdata/recursive", []string{"Payload"})
	if err != nil {
		t.Fatal(err)
	}

	file, err := protodesc.NewFile(&descriptorpb.FileDescriptorProto{
		Name:    proto.String("recursive.proto"),
		Package: proto.String("temporal.api.cmd.proxygenerator.testdata.recursive"),
		MessageType: []*descriptorpb.DescriptorProto{
			{Name: proto.String("Node")},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}

	records, err := gatherMatchesToTypeRecords([]protoreflect.MessageDescriptor{file.Messages().ByName("Node")}, payloads, payloads)
	if err != nil {
		t.Fatal(err)
	}

	record := records["*recursive.Node"]
	if record == nil || !slices.Contains(record.Methods, "GetChild") {
		t.Fatal("Node.GetChild should be visited because the child can contain a Payload")
	}
}
