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
	unsafe "unsafe"

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
	BATCH_OPERATION_TYPE_UNSPECIFIED              BatchOperationType = 0
	BATCH_OPERATION_TYPE_TERMINATE                BatchOperationType = 1
	BATCH_OPERATION_TYPE_CANCEL                   BatchOperationType = 2
	BATCH_OPERATION_TYPE_SIGNAL                   BatchOperationType = 3
	BATCH_OPERATION_TYPE_DELETE                   BatchOperationType = 4
	BATCH_OPERATION_TYPE_RESET                    BatchOperationType = 5
	BATCH_OPERATION_TYPE_UPDATE_EXECUTION_OPTIONS BatchOperationType = 6
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
		6: "BATCH_OPERATION_TYPE_UPDATE_EXECUTION_OPTIONS",
	}
	BatchOperationType_value = map[string]int32{
		"BATCH_OPERATION_TYPE_UNSPECIFIED":              0,
		"BATCH_OPERATION_TYPE_TERMINATE":                1,
		"BATCH_OPERATION_TYPE_CANCEL":                   2,
		"BATCH_OPERATION_TYPE_SIGNAL":                   3,
		"BATCH_OPERATION_TYPE_DELETE":                   4,
		"BATCH_OPERATION_TYPE_RESET":                    5,
		"BATCH_OPERATION_TYPE_UPDATE_EXECUTION_OPTIONS": 6,
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
	case BATCH_OPERATION_TYPE_UPDATE_EXECUTION_OPTIONS:
		return "UpdateExecutionOptions"
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

const file_temporal_api_enums_v1_batch_operation_proto_rawDesc = "" +
	"\n" +
	"+temporal/api/enums/v1/batch_operation.proto\x12\x15temporal.api.enums.v1*\x94\x02\n" +
	"\x12BatchOperationType\x12$\n" +
	" BATCH_OPERATION_TYPE_UNSPECIFIED\x10\x00\x12\"\n" +
	"\x1eBATCH_OPERATION_TYPE_TERMINATE\x10\x01\x12\x1f\n" +
	"\x1bBATCH_OPERATION_TYPE_CANCEL\x10\x02\x12\x1f\n" +
	"\x1bBATCH_OPERATION_TYPE_SIGNAL\x10\x03\x12\x1f\n" +
	"\x1bBATCH_OPERATION_TYPE_DELETE\x10\x04\x12\x1e\n" +
	"\x1aBATCH_OPERATION_TYPE_RESET\x10\x05\x121\n" +
	"-BATCH_OPERATION_TYPE_UPDATE_EXECUTION_OPTIONS\x10\x06*\xa6\x01\n" +
	"\x13BatchOperationState\x12%\n" +
	"!BATCH_OPERATION_STATE_UNSPECIFIED\x10\x00\x12!\n" +
	"\x1dBATCH_OPERATION_STATE_RUNNING\x10\x01\x12#\n" +
	"\x1fBATCH_OPERATION_STATE_COMPLETED\x10\x02\x12 \n" +
	"\x1cBATCH_OPERATION_STATE_FAILED\x10\x03B\x8b\x01\n" +
	"\x18io.temporal.api.enums.v1B\x13BatchOperationProtoP\x01Z!go.temporal.io/api/enums/v1;enums\xaa\x02\x17Temporalio.Api.Enums.V1\xea\x02\x1aTemporalio::Api::Enums::V1b\x06proto3"

var (
	file_temporal_api_enums_v1_batch_operation_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_batch_operation_proto_rawDescData []byte
)

func file_temporal_api_enums_v1_batch_operation_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_batch_operation_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_batch_operation_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_batch_operation_proto_rawDesc), len(file_temporal_api_enums_v1_batch_operation_proto_rawDesc)))
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
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_batch_operation_proto_rawDesc), len(file_temporal_api_enums_v1_batch_operation_proto_rawDesc)),
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
	file_temporal_api_enums_v1_batch_operation_proto_goTypes = nil
	file_temporal_api_enums_v1_batch_operation_proto_depIdxs = nil
}
