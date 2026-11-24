package main

import (
	"slices"
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

// TestContainsMessageRecursive checks that a recursive search can find a payload through a previously visited type.
func TestContainsMessageRecursive(t *testing.T) {
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
			{Name: proto.String("A"), Field: []*descriptorpb.FieldDescriptorProto{field("b", ".recursive.B", 1), field("payload", ".recursive.Payload", 2)}},
			{Name: proto.String("B"), Field: []*descriptorpb.FieldDescriptorProto{field("a", ".recursive.A", 1)}},
			{Name: proto.String("Payload")},
		},
	}, nil)
	if err != nil {
		t.Fatal(err)
	}
	memo := map[protoreflect.FullName]bool{}
	for _, name := range []protoreflect.Name{"A", "B"} {
		if !containsMessage(file.Messages().ByName(name), []string{"recursive.Payload"}, memo) {
			t.Fatalf("%s should reach Payload through the recursive graph", name)
		}
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
