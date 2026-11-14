package main

import (
	"testing"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func TestContainsMessageRecursive(t *testing.T) {
	field := func(name, typ string, number int32) *descriptorpb.FieldDescriptorProto {
		return &descriptorpb.FieldDescriptorProto{
			Name: proto.String(name), Number: proto.Int32(number), TypeName: proto.String(typ),
			Type: descriptorpb.FieldDescriptorProto_TYPE_MESSAGE.Enum(), Label: descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum(),
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
