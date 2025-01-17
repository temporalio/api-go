// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.26.1
// source: temporal/api/sdk/v1/task_complete_metadata.proto

package sdk

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkflowTaskCompletedMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Internal flags used by the core SDK. SDKs using flags must comply with the following behavior:
	//
	// During replay:
	//   - If a flag is not recognized (value is too high or not defined), it must fail the workflow
	//     task.
	//   - If a flag is recognized, it is stored in a set of used flags for the run. Code checks for
	//     that flag during and after this WFT are allowed to assume that the flag is present.
	//   - If a code check for a flag does not find the flag in the set of used flags, it must take
	//     the branch corresponding to the absence of that flag.
	//
	// During non-replay execution of new WFTs:
	//   - The SDK is free to use all flags it knows about. It must record any newly-used (IE: not
	//     previously recorded) flags when completing the WFT.
	//
	// SDKs which are too old to even know about this field at all are considered to produce
	// undefined behavior if they replay workflows which used this mechanism.
	//
	// (-- api-linter: core::0141::forbidden-types=disabled
	//
	//	aip.dev/not-precedent: These really shouldn't have negative values. --)
	CoreUsedFlags []uint32 `protobuf:"varint,1,rep,packed,name=core_used_flags,json=coreUsedFlags,proto3" json:"core_used_flags,omitempty"`
	// Flags used by the SDK lang. No attempt is made to distinguish between different SDK languages
	// here as processing a workflow with a different language than the one which authored it is
	// already undefined behavior. See `core_used_patches` for more.
	//
	// (-- api-linter: core::0141::forbidden-types=disabled
	//
	//	aip.dev/not-precedent: These really shouldn't have negative values. --)
	LangUsedFlags []uint32 `protobuf:"varint,2,rep,packed,name=lang_used_flags,json=langUsedFlags,proto3" json:"lang_used_flags,omitempty"`
	// Name of the SDK that processed the task. This is usually something like "temporal-go" and is
	// usually the same as client-name gRPC header. This should only be set if its value changed
	// since the last time recorded on the workflow (or be set on the first task).
	//
	// (-- api-linter: core::0122::name-suffix=disabled
	//
	//	aip.dev/not-precedent: We're ok with a name suffix here. --)
	SdkName string `protobuf:"bytes,3,opt,name=sdk_name,json=sdkName,proto3" json:"sdk_name,omitempty"`
	// Version of the SDK that processed the task. This is usually something like "1.20.0" and is
	// usually the same as client-version gRPC header. This should only be set if its value changed
	// since the last time recorded on the workflow (or be set on the first task).
	SdkVersion    string `protobuf:"bytes,4,opt,name=sdk_version,json=sdkVersion,proto3" json:"sdk_version,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowTaskCompletedMetadata) Reset() {
	*x = WorkflowTaskCompletedMetadata{}
	mi := &file_temporal_api_sdk_v1_task_complete_metadata_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowTaskCompletedMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowTaskCompletedMetadata) ProtoMessage() {}

func (x *WorkflowTaskCompletedMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_sdk_v1_task_complete_metadata_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowTaskCompletedMetadata.ProtoReflect.Descriptor instead.
func (*WorkflowTaskCompletedMetadata) Descriptor() ([]byte, []int) {
	return file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowTaskCompletedMetadata) GetCoreUsedFlags() []uint32 {
	if x != nil {
		return x.CoreUsedFlags
	}
	return nil
}

func (x *WorkflowTaskCompletedMetadata) GetLangUsedFlags() []uint32 {
	if x != nil {
		return x.LangUsedFlags
	}
	return nil
}

func (x *WorkflowTaskCompletedMetadata) GetSdkName() string {
	if x != nil {
		return x.SdkName
	}
	return ""
}

func (x *WorkflowTaskCompletedMetadata) GetSdkVersion() string {
	if x != nil {
		return x.SdkVersion
	}
	return ""
}

var File_temporal_api_sdk_v1_task_complete_metadata_proto protoreflect.FileDescriptor

var file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc = string([]byte{
	0x0a, 0x30, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0xab, 0x01, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x65, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x61, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x64, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x64, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x87, 0x01, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31,
	0x42, 0x19, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x67,
	0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x64, 0x6b, 0xaa, 0x02, 0x15, 0x54,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x64,
	0x6b, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x18, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69,
	0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x64, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescOnce sync.Once
	file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescData []byte
)

func file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescGZIP() []byte {
	file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescOnce.Do(func() {
		file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc), len(file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc)))
	})
	return file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescData
}

var file_temporal_api_sdk_v1_task_complete_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_api_sdk_v1_task_complete_metadata_proto_goTypes = []any{
	(*WorkflowTaskCompletedMetadata)(nil), // 0: temporal.api.sdk.v1.WorkflowTaskCompletedMetadata
}
var file_temporal_api_sdk_v1_task_complete_metadata_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_sdk_v1_task_complete_metadata_proto_init() }
func file_temporal_api_sdk_v1_task_complete_metadata_proto_init() {
	if File_temporal_api_sdk_v1_task_complete_metadata_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc), len(file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_sdk_v1_task_complete_metadata_proto_goTypes,
		DependencyIndexes: file_temporal_api_sdk_v1_task_complete_metadata_proto_depIdxs,
		MessageInfos:      file_temporal_api_sdk_v1_task_complete_metadata_proto_msgTypes,
	}.Build()
	File_temporal_api_sdk_v1_task_complete_metadata_proto = out.File
	file_temporal_api_sdk_v1_task_complete_metadata_proto_goTypes = nil
	file_temporal_api_sdk_v1_task_complete_metadata_proto_depIdxs = nil
}
