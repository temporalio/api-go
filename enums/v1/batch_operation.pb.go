// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
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
// source: temporal/api/enums/v1/batch_operation.proto

package enums

import (
	reflect "reflect"
	"strconv"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatchOperationType int32

const (
	BATCH_OPERATION_TYPE_UNSPECIFIED BatchOperationType = 0
	BATCH_OPERATION_TYPE_TERMINATE   BatchOperationType = 1
	BATCH_OPERATION_TYPE_CANCEL      BatchOperationType = 2
	BATCH_OPERATION_TYPE_SIGNAL      BatchOperationType = 3
	BATCH_OPERATION_TYPE_DELETE      BatchOperationType = 4
	BATCH_OPERATION_TYPE_RESET       BatchOperationType = 5
)

// Enum value maps for BatchOperationType.
var (
	BatchOperationType_name = map[int32]string{
		0: "BATCH_OPERATION_TYPE_UNSPECIFIED",
		1: "BATCH_OPERATION_TYPE_TERMINATE",
		2: "BATCH_OPERATION_TYPE_CANCEL",
		3: "BATCH_OPERATION_TYPE_SIGNAL",
		4: "BATCH_OPERATION_TYPE_DELETE",
		5: "BATCH_OPERATION_TYPE_RESET",
	}
	BatchOperationType_value = map[string]int32{
		"BATCH_OPERATION_TYPE_UNSPECIFIED": 0,
		"BATCH_OPERATION_TYPE_TERMINATE":   1,
		"BATCH_OPERATION_TYPE_CANCEL":      2,
		"BATCH_OPERATION_TYPE_SIGNAL":      3,
		"BATCH_OPERATION_TYPE_DELETE":      4,
		"BATCH_OPERATION_TYPE_RESET":       5,
	}
)

func (x BatchOperationType) Enum() *BatchOperationType {
	p := new(BatchOperationType)
	*p = x
	return p
}

func (x BatchOperationType) String() string {
	switch x {
	case BATCH_OPERATION_TYPE_UNSPECIFIED:
		return "Unspecified"
	case BATCH_OPERATION_TYPE_TERMINATE:
		return "Terminate"
	case BATCH_OPERATION_TYPE_CANCEL:
		return "Cancel"
	case BATCH_OPERATION_TYPE_SIGNAL:
		return "Signal"
	case BATCH_OPERATION_TYPE_DELETE:
		return "Delete"
	case BATCH_OPERATION_TYPE_RESET:
		return "Reset"
	default:
		return strconv.Itoa(int(x))
	}

}

func (BatchOperationType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_batch_operation_proto_enumTypes[0].Descriptor()
}

func (BatchOperationType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_batch_operation_proto_enumTypes[0]
}

func (x BatchOperationType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatchOperationType.Descriptor instead.
func (BatchOperationType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_batch_operation_proto_rawDescGZIP(), []int{0}
}

type BatchOperationState int32

const (
	BATCH_OPERATION_STATE_UNSPECIFIED BatchOperationState = 0
	BATCH_OPERATION_STATE_RUNNING     BatchOperationState = 1
	BATCH_OPERATION_STATE_COMPLETED   BatchOperationState = 2
	BATCH_OPERATION_STATE_FAILED      BatchOperationState = 3
)

// Enum value maps for BatchOperationState.
var (
	BatchOperationState_name = map[int32]string{
		0: "BATCH_OPERATION_STATE_UNSPECIFIED",
		1: "BATCH_OPERATION_STATE_RUNNING",
		2: "BATCH_OPERATION_STATE_COMPLETED",
		3: "BATCH_OPERATION_STATE_FAILED",
	}
	BatchOperationState_value = map[string]int32{
		"BATCH_OPERATION_STATE_UNSPECIFIED": 0,
		"BATCH_OPERATION_STATE_RUNNING":     1,
		"BATCH_OPERATION_STATE_COMPLETED":   2,
		"BATCH_OPERATION_STATE_FAILED":      3,
	}
)

func (x BatchOperationState) Enum() *BatchOperationState {
	p := new(BatchOperationState)
	*p = x
	return p
}

func (x BatchOperationState) String() string {
	switch x {
	case BATCH_OPERATION_STATE_UNSPECIFIED:
		return "Unspecified"
	case BATCH_OPERATION_STATE_RUNNING:
		return "Running"
	case BATCH_OPERATION_STATE_COMPLETED:
		return "Completed"
	case BATCH_OPERATION_STATE_FAILED:
		return "Failed"
	default:
		return strconv.Itoa(int(x))
	}

}

func (BatchOperationState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_batch_operation_proto_enumTypes[1].Descriptor()
}

func (BatchOperationState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_batch_operation_proto_enumTypes[1]
}

func (x BatchOperationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BatchOperationState.Descriptor instead.
func (BatchOperationState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_batch_operation_proto_rawDescGZIP(), []int{1}
}

var File_temporal_api_enums_v1_batch_operation_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_batch_operation_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x6f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d,
	0x73, 0x2e, 0x76, 0x31, 0x2a, 0xe1, 0x01, 0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x20, 0x42,
	0x41, 0x54, 0x43, 0x48, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x22, 0x0a, 0x1e, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e,
	0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4f,
	0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x10, 0x02, 0x12, 0x1f, 0x0a, 0x1b, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f,
	0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53,
	0x49, 0x47, 0x4e, 0x41, 0x4c, 0x10, 0x03, 0x12, 0x1f, 0x0a, 0x1b, 0x42, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x44, 0x45, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x04, 0x12, 0x1e, 0x0a, 0x1a, 0x42, 0x41, 0x54, 0x43,
	0x48, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x52, 0x45, 0x53, 0x45, 0x54, 0x10, 0x05, 0x2a, 0xa6, 0x01, 0x0a, 0x13, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x25, 0x0a, 0x21, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54,
	0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x42, 0x41, 0x54, 0x43, 0x48,
	0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x23, 0x0a, 0x1f, 0x42, 0x41,
	0x54, 0x43, 0x48, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12,
	0x20, 0x0a, 0x1c, 0x42, 0x41, 0x54, 0x43, 0x48, 0x5f, 0x4f, 0x50, 0x45, 0x52, 0x41, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10,
	0x03, 0x42, 0x8b, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x13,
	0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_batch_operation_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_batch_operation_proto_rawDescData = file_temporal_api_enums_v1_batch_operation_proto_rawDesc
)

func file_temporal_api_enums_v1_batch_operation_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_batch_operation_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_batch_operation_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_batch_operation_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_batch_operation_proto_rawDescData
}

var file_temporal_api_enums_v1_batch_operation_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_temporal_api_enums_v1_batch_operation_proto_goTypes = []any{
	(BatchOperationType)(0),  // 0: temporal.api.enums.v1.BatchOperationType
	(BatchOperationState)(0), // 1: temporal.api.enums.v1.BatchOperationState
}
var file_temporal_api_enums_v1_batch_operation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_batch_operation_proto_init() }
func file_temporal_api_enums_v1_batch_operation_proto_init() {
	if File_temporal_api_enums_v1_batch_operation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_batch_operation_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_batch_operation_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_batch_operation_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_batch_operation_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_batch_operation_proto = out.File
	file_temporal_api_enums_v1_batch_operation_proto_rawDesc = nil
	file_temporal_api_enums_v1_batch_operation_proto_goTypes = nil
	file_temporal_api_enums_v1_batch_operation_proto_depIdxs = nil
}
