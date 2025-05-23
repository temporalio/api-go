// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/export/v1/message.proto

package export

import (
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"

	v1 "go.temporal.io/api/history/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkflowExecution struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	History       *v1.History            `protobuf:"bytes,1,opt,name=history,proto3" json:"history,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowExecution) Reset() {
	*x = WorkflowExecution{}
	mi := &file_temporal_api_export_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowExecution) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecution) ProtoMessage() {}

func (x *WorkflowExecution) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_export_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecution.ProtoReflect.Descriptor instead.
func (*WorkflowExecution) Descriptor() ([]byte, []int) {
	return file_temporal_api_export_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowExecution) GetHistory() *v1.History {
	if x != nil {
		return x.History
	}
	return nil
}

// WorkflowExecutions is used by the Cloud Export feature to deserialize
// the exported file. It encapsulates a collection of workflow execution information.
type WorkflowExecutions struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*WorkflowExecution   `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowExecutions) Reset() {
	*x = WorkflowExecutions{}
	mi := &file_temporal_api_export_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowExecutions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecutions) ProtoMessage() {}

func (x *WorkflowExecutions) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_export_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecutions.ProtoReflect.Descriptor instead.
func (*WorkflowExecutions) Descriptor() ([]byte, []int) {
	return file_temporal_api_export_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowExecutions) GetItems() []*WorkflowExecution {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_temporal_api_export_v1_message_proto protoreflect.FileDescriptor

const file_temporal_api_export_v1_message_proto_rawDesc = "" +
	"\n" +
	"$temporal/api/export/v1/message.proto\x12\x16temporal.api.export.v1\x1a%temporal/api/history/v1/message.proto\"O\n" +
	"\x11WorkflowExecution\x12:\n" +
	"\ahistory\x18\x01 \x01(\v2 .temporal.api.history.v1.HistoryR\ahistory\"U\n" +
	"\x12WorkflowExecutions\x12?\n" +
	"\x05items\x18\x01 \x03(\v2).temporal.api.export.v1.WorkflowExecutionR\x05itemsB\x89\x01\n" +
	"\x19io.temporal.api.export.v1B\fMessageProtoP\x01Z#go.temporal.io/api/export/v1;export\xaa\x02\x18Temporalio.Api.Export.V1\xea\x02\x1bTemporalio::Api::Export::V1b\x06proto3"

var (
	file_temporal_api_export_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_export_v1_message_proto_rawDescData []byte
)

func file_temporal_api_export_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_export_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_export_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_export_v1_message_proto_rawDesc), len(file_temporal_api_export_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_export_v1_message_proto_rawDescData
}

var file_temporal_api_export_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_temporal_api_export_v1_message_proto_goTypes = []any{
	(*WorkflowExecution)(nil),  // 0: temporal.api.export.v1.WorkflowExecution
	(*WorkflowExecutions)(nil), // 1: temporal.api.export.v1.WorkflowExecutions
	(*v1.History)(nil),         // 2: temporal.api.history.v1.History
}
var file_temporal_api_export_v1_message_proto_depIdxs = []int32{
	2, // 0: temporal.api.export.v1.WorkflowExecution.history:type_name -> temporal.api.history.v1.History
	0, // 1: temporal.api.export.v1.WorkflowExecutions.items:type_name -> temporal.api.export.v1.WorkflowExecution
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_temporal_api_export_v1_message_proto_init() }
func file_temporal_api_export_v1_message_proto_init() {
	if File_temporal_api_export_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_export_v1_message_proto_rawDesc), len(file_temporal_api_export_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_export_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_export_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_export_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_export_v1_message_proto = out.File
	file_temporal_api_export_v1_message_proto_goTypes = nil
	file_temporal_api_export_v1_message_proto_depIdxs = nil
}
