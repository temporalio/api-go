package main

import (
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"

	"go.temporal.io/api/protometa/v1"
)

func TestDraftRPCFieldReferences(t *testing.T) {
	tests := []struct {
		name   string
		rule   *annotations.HttpRule
		header string
		want   string
	}{
		{
			name:   "request header precedes HTTP",
			rule:   &annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/items/{a}"}},
			header: "{b}",
			want:   "request header references draft field b",
		},
		{
			name: "body precedes custom path",
			rule: &annotations.HttpRule{
				Pattern: &annotations.HttpRule_Custom{Custom: &annotations.CustomHttpPattern{Kind: "REPORT", Path: "/items/{a}"}},
				Body:    "b",
			},
			want: "http rule references draft field b",
		},
		{
			name: "additional binding precedes response body",
			rule: &annotations.HttpRule{
				ResponseBody:       "b",
				AdditionalBindings: []*annotations.HttpRule{{Pattern: &annotations.HttpRule_Get{Get: "/items/{a=items/*}"}}},
			},
			want: "http rule references draft field a",
		},
		{
			name:   "stable fields",
			rule:   &annotations.HttpRule{Pattern: &annotations.HttpRule_Get{Get: "/items/{stable}"}},
			header: "item-{stable}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := rpcFixtureRequest(tt.rule, tt.header)
			gen, err := (protogen.Options{}).New(request)
			require.NoError(t, err)
			_, err = draftPlan(gen)
			if tt.want == "" {
				require.NoError(t, err)
			} else {
				require.EqualError(t, err, "omission.Service.Call: "+tt.want)
			}
		})
	}
}

func TestDraftRPCOptionSchema(t *testing.T) {
	for _, tt := range []struct {
		name   string
		file   string
		option string
		change func(*descriptorpb.FileDescriptorProto)
	}{
		{
			name:   "singular request header",
			file:   protometa.File_temporal_api_protometa_v1_annotations_proto.Path(),
			option: requestHeaderOption,
			change: func(file *descriptorpb.FileDescriptorProto) {
				file.Extension[0].Label = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
			},
		},
		{
			name:   "renamed request header value",
			file:   protometa.File_temporal_api_protometa_v1_annotations_proto.Path(),
			option: requestHeaderOption,
			change: func(file *descriptorpb.FileDescriptorProto) {
				file.MessageType[0].Field[1].Name = proto.String("template")
			},
		},
		{
			name:   "repeated HTTP option",
			file:   annotations.File_google_api_annotations_proto.Path(),
			option: httpRuleOption,
			change: func(file *descriptorpb.FileDescriptorProto) {
				file.Extension[0].Label = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
			},
		},
		{
			name:   "renamed custom HTTP path",
			file:   annotations.File_google_api_http_proto.Path(),
			option: httpRuleOption,
			change: func(file *descriptorpb.FileDescriptorProto) {
				file.MessageType[2].Field[1].Name = proto.String("template")
			},
		},
		{
			name:   "additional HTTP bindings have a different message type",
			file:   annotations.File_google_api_http_proto.Path(),
			option: httpRuleOption,
			change: func(file *descriptorpb.FileDescriptorProto) {
				rule := file.MessageType[1]
				i := slices.IndexFunc(rule.Field, func(field *descriptorpb.FieldDescriptorProto) bool {
					return field.GetName() == "additional_bindings"
				})
				rule.Field[i].TypeName = proto.String(".google.api.CustomHttpPattern")
			},
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			rule := &annotations.HttpRule{Pattern: &annotations.HttpRule_Custom{
				Custom: &annotations.CustomHttpPattern{Kind: "REPORT", Path: "/items/{stable}"},
			}}
			request := rpcFixtureRequest(rule, "{stable}")
			i := slices.IndexFunc(request.ProtoFile, func(file *descriptorpb.FileDescriptorProto) bool {
				return file.GetName() == tt.file
			})
			require.NotEqual(t, -1, i)
			tt.change(request.ProtoFile[i])
			gen, err := (protogen.Options{}).New(request)
			require.NoError(t, err)
			_, err = draftPlan(gen)
			require.ErrorContains(t, err, tt.option+": malformed RPC option schema")
		})
	}
}

func rpcFixtureRequest(rule *annotations.HttpRule, header string) *pluginpb.CodeGeneratorRequest {
	request := fixtureRequest()
	request.ProtoFile = append(request.ProtoFile,
		protodesc.ToFileDescriptorProto(annotations.File_google_api_http_proto),
		protodesc.ToFileDescriptorProto(annotations.File_google_api_annotations_proto),
		protodesc.ToFileDescriptorProto(protometa.File_temporal_api_protometa_v1_annotations_proto),
	)
	options := &descriptorpb.MethodOptions{}
	proto.SetExtension(options, annotations.E_Http, rule)
	if header != "" {
		proto.SetExtension(options, protometa.E_RequestHeader, []*protometa.RequestHeaderAnnotation{{Header: "item", Value: header}})
	}
	request.ProtoFile = append(request.ProtoFile, &descriptorpb.FileDescriptorProto{
		Name:    proto.String("service.proto"),
		Package: proto.String("omission"),
		Syntax:  proto.String("proto3"),
		Dependency: []string{
			"child/message.proto",
			annotations.File_google_api_annotations_proto.Path(),
			protometa.File_temporal_api_protometa_v1_annotations_proto.Path(),
		},
		Options: &descriptorpb.FileOptions{GoPackage: proto.String("go.temporal.io/api/test/rpc")},
		Service: []*descriptorpb.ServiceDescriptorProto{{
			Name: proto.String("Service"),
			Method: []*descriptorpb.MethodDescriptorProto{{
				Name:       proto.String("Call"),
				InputType:  proto.String(".omission.child.Leaf"),
				OutputType: proto.String(".omission.child.Leaf"),
				Options:    options,
			}},
		}},
	})
	request.FileToGenerate = append(request.FileToGenerate, "service.proto")
	return request
}
