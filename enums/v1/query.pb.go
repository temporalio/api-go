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
// source: temporal/api/enums/v1/query.proto

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

type QueryResultType int32

const (
	QUERY_RESULT_TYPE_UNSPECIFIED QueryResultType = 0
	QUERY_RESULT_TYPE_ANSWERED    QueryResultType = 1
	QUERY_RESULT_TYPE_FAILED      QueryResultType = 2
)

// Enum value maps for QueryResultType.
var (
	QueryResultType_name = map[int32]string{
		0: "QUERY_RESULT_TYPE_UNSPECIFIED",
		1: "QUERY_RESULT_TYPE_ANSWERED",
		2: "QUERY_RESULT_TYPE_FAILED",
	}
	QueryResultType_value = map[string]int32{
		"QUERY_RESULT_TYPE_UNSPECIFIED": 0,
		"QUERY_RESULT_TYPE_ANSWERED":    1,
		"QUERY_RESULT_TYPE_FAILED":      2,
	}
)

func (x QueryResultType) Enum() *QueryResultType {
	p := new(QueryResultType)
	*p = x
	return p
}

func (x QueryResultType) String() string {
	switch x {
	case QUERY_RESULT_TYPE_UNSPECIFIED:
		return "Unspecified"
	case QUERY_RESULT_TYPE_ANSWERED:
		return "Answered"
	case QUERY_RESULT_TYPE_FAILED:
		return "Failed"
	default:
		return strconv.Itoa(int(x))
	}

}

func (QueryResultType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_query_proto_enumTypes[0].Descriptor()
}

func (QueryResultType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_query_proto_enumTypes[0]
}

func (x QueryResultType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryResultType.Descriptor instead.
func (QueryResultType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_query_proto_rawDescGZIP(), []int{0}
}

type QueryRejectCondition int32

const (
	QUERY_REJECT_CONDITION_UNSPECIFIED QueryRejectCondition = 0
	// None indicates that query should not be rejected.
	QUERY_REJECT_CONDITION_NONE QueryRejectCondition = 1
	// NotOpen indicates that query should be rejected if workflow is not open.
	QUERY_REJECT_CONDITION_NOT_OPEN QueryRejectCondition = 2
	// NotCompletedCleanly indicates that query should be rejected if workflow did not complete cleanly.
	QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY QueryRejectCondition = 3
)

// Enum value maps for QueryRejectCondition.
var (
	QueryRejectCondition_name = map[int32]string{
		0: "QUERY_REJECT_CONDITION_UNSPECIFIED",
		1: "QUERY_REJECT_CONDITION_NONE",
		2: "QUERY_REJECT_CONDITION_NOT_OPEN",
		3: "QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY",
	}
	QueryRejectCondition_value = map[string]int32{
		"QUERY_REJECT_CONDITION_UNSPECIFIED":           0,
		"QUERY_REJECT_CONDITION_NONE":                  1,
		"QUERY_REJECT_CONDITION_NOT_OPEN":              2,
		"QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY": 3,
	}
)

func (x QueryRejectCondition) Enum() *QueryRejectCondition {
	p := new(QueryRejectCondition)
	*p = x
	return p
}

func (x QueryRejectCondition) String() string {
	switch x {
	case QUERY_REJECT_CONDITION_UNSPECIFIED:
		return "Unspecified"
	case QUERY_REJECT_CONDITION_NONE:
		return "None"
	case QUERY_REJECT_CONDITION_NOT_OPEN:
		return "NotOpen"
	case QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY:
		return "NotCompletedCleanly"
	default:
		return strconv.Itoa(int(x))
	}

}

func (QueryRejectCondition) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_query_proto_enumTypes[1].Descriptor()
}

func (QueryRejectCondition) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_query_proto_enumTypes[1]
}

func (x QueryRejectCondition) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use QueryRejectCondition.Descriptor instead.
func (QueryRejectCondition) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_query_proto_rawDescGZIP(), []int{1}
}

var File_temporal_api_enums_v1_query_proto protoreflect.FileDescriptor

const file_temporal_api_enums_v1_query_proto_rawDesc = "" +
	"\n" +
	"!temporal/api/enums/v1/query.proto\x12\x15temporal.api.enums.v1*r\n" +
	"\x0fQueryResultType\x12!\n" +
	"\x1dQUERY_RESULT_TYPE_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aQUERY_RESULT_TYPE_ANSWERED\x10\x01\x12\x1c\n" +
	"\x18QUERY_RESULT_TYPE_FAILED\x10\x02*\xb6\x01\n" +
	"\x14QueryRejectCondition\x12&\n" +
	"\"QUERY_REJECT_CONDITION_UNSPECIFIED\x10\x00\x12\x1f\n" +
	"\x1bQUERY_REJECT_CONDITION_NONE\x10\x01\x12#\n" +
	"\x1fQUERY_REJECT_CONDITION_NOT_OPEN\x10\x02\x120\n" +
	",QUERY_REJECT_CONDITION_NOT_COMPLETED_CLEANLY\x10\x03B\x82\x01\n" +
	"\x18io.temporal.api.enums.v1B\n" +
	"QueryProtoP\x01Z!go.temporal.io/api/enums/v1;enums\xaa\x02\x17Temporalio.Api.Enums.V1\xea\x02\x1aTemporalio::Api::Enums::V1b\x06proto3"

var (
	file_temporal_api_enums_v1_query_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_query_proto_rawDescData []byte
)

func file_temporal_api_enums_v1_query_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_query_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_query_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_query_proto_rawDesc), len(file_temporal_api_enums_v1_query_proto_rawDesc)))
	})
	return file_temporal_api_enums_v1_query_proto_rawDescData
}

var file_temporal_api_enums_v1_query_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_temporal_api_enums_v1_query_proto_goTypes = []any{
	(QueryResultType)(0),      // 0: temporal.api.enums.v1.QueryResultType
	(QueryRejectCondition)(0), // 1: temporal.api.enums.v1.QueryRejectCondition
}
var file_temporal_api_enums_v1_query_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_query_proto_init() }
func file_temporal_api_enums_v1_query_proto_init() {
	if File_temporal_api_enums_v1_query_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_query_proto_rawDesc), len(file_temporal_api_enums_v1_query_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_query_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_query_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_query_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_query_proto = out.File
	file_temporal_api_enums_v1_query_proto_goTypes = nil
	file_temporal_api_enums_v1_query_proto_depIdxs = nil
}
