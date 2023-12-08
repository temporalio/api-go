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
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/filter/v1/message.proto

package filter

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/enums/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkflowExecutionFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowId string `protobuf:"bytes,1,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	RunId      string `protobuf:"bytes,2,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *WorkflowExecutionFilter) Reset() {
	*x = WorkflowExecutionFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_filter_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowExecutionFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecutionFilter) ProtoMessage() {}

func (x *WorkflowExecutionFilter) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_filter_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecutionFilter.ProtoReflect.Descriptor instead.
func (*WorkflowExecutionFilter) Descriptor() ([]byte, []int) {
	return file_temporal_api_filter_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowExecutionFilter) GetWorkflowId() string {
	if x != nil {
		return x.WorkflowId
	}
	return ""
}

func (x *WorkflowExecutionFilter) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type WorkflowTypeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *WorkflowTypeFilter) Reset() {
	*x = WorkflowTypeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_filter_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkflowTypeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowTypeFilter) ProtoMessage() {}

func (x *WorkflowTypeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_filter_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowTypeFilter.ProtoReflect.Descriptor instead.
func (*WorkflowTypeFilter) Descriptor() ([]byte, []int) {
	return file_temporal_api_filter_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *WorkflowTypeFilter) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type StartTimeFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EarliestTime *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=earliest_time,json=earliestTime,proto3" json:"earliest_time,omitempty"`
	LatestTime   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=latest_time,json=latestTime,proto3" json:"latest_time,omitempty"`
}

func (x *StartTimeFilter) Reset() {
	*x = StartTimeFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_filter_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartTimeFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartTimeFilter) ProtoMessage() {}

func (x *StartTimeFilter) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_filter_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartTimeFilter.ProtoReflect.Descriptor instead.
func (*StartTimeFilter) Descriptor() ([]byte, []int) {
	return file_temporal_api_filter_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *StartTimeFilter) GetEarliestTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EarliestTime
	}
	return nil
}

func (x *StartTimeFilter) GetLatestTime() *timestamppb.Timestamp {
	if x != nil {
		return x.LatestTime
	}
	return nil
}

type StatusFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status v1.WorkflowExecutionStatus `protobuf:"varint,1,opt,name=status,proto3,enum=temporal.api.enums.v1.WorkflowExecutionStatus" json:"status,omitempty"`
}

func (x *StatusFilter) Reset() {
	*x = StatusFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_filter_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusFilter) ProtoMessage() {}

func (x *StatusFilter) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_filter_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusFilter.ProtoReflect.Descriptor instead.
func (*StatusFilter) Descriptor() ([]byte, []int) {
	return file_temporal_api_filter_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *StatusFilter) GetStatus() v1.WorkflowExecutionStatus {
	if x != nil {
		return x.Status
	}
	return v1.WorkflowExecutionStatus(0)
}

var File_temporal_api_filter_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_filter_v1_message_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x16, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e,
	0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x51, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x12, 0x1f, 0x0a, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49,
	0x64, 0x12, 0x15, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x22, 0x28, 0x0a, 0x12, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x79, 0x70, 0x65, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x0f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x3f, 0x0a, 0x0d, 0x65, 0x61, 0x72, 0x6c, 0x69, 0x65,
	0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x65, 0x61, 0x72, 0x6c, 0x69,
	0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74,
	0x54, 0x69, 0x6d, 0x65, 0x22, 0x56, 0x0a, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x89, 0x01, 0x0a,
	0x19, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x6f, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x66,
	0x69, 0x6c, 0x74, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0xaa,
	0x02, 0x18, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69,
	0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1b, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_filter_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_filter_v1_message_proto_rawDescData = file_temporal_api_filter_v1_message_proto_rawDesc
)

func file_temporal_api_filter_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_filter_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_filter_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_filter_v1_message_proto_rawDescData)
	})
	return file_temporal_api_filter_v1_message_proto_rawDescData
}

var file_temporal_api_filter_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_temporal_api_filter_v1_message_proto_goTypes = []interface{}{
	(*WorkflowExecutionFilter)(nil), // 0: temporal.api.filter.v1.WorkflowExecutionFilter
	(*WorkflowTypeFilter)(nil),      // 1: temporal.api.filter.v1.WorkflowTypeFilter
	(*StartTimeFilter)(nil),         // 2: temporal.api.filter.v1.StartTimeFilter
	(*StatusFilter)(nil),            // 3: temporal.api.filter.v1.StatusFilter
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
	(v1.WorkflowExecutionStatus)(0), // 5: temporal.api.enums.v1.WorkflowExecutionStatus
}
var file_temporal_api_filter_v1_message_proto_depIdxs = []int32{
	4, // 0: temporal.api.filter.v1.StartTimeFilter.earliest_time:type_name -> google.protobuf.Timestamp
	4, // 1: temporal.api.filter.v1.StartTimeFilter.latest_time:type_name -> google.protobuf.Timestamp
	5, // 2: temporal.api.filter.v1.StatusFilter.status:type_name -> temporal.api.enums.v1.WorkflowExecutionStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_temporal_api_filter_v1_message_proto_init() }
func file_temporal_api_filter_v1_message_proto_init() {
	if File_temporal_api_filter_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_filter_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowExecutionFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_api_filter_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkflowTypeFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_api_filter_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartTimeFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_temporal_api_filter_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusFilter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_filter_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_filter_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_filter_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_filter_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_filter_v1_message_proto = out.File
	file_temporal_api_filter_v1_message_proto_rawDesc = nil
	file_temporal_api_filter_v1_message_proto_goTypes = nil
	file_temporal_api_filter_v1_message_proto_depIdxs = nil
}
