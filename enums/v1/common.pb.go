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
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: temporal/api/enums/v1/common.proto

package enums

import (
	reflect "reflect"
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

type EncodingType int32

const (
	EncodingType_ENCODING_TYPE_UNSPECIFIED EncodingType = 0
	EncodingType_ENCODING_TYPE_PROTO3      EncodingType = 1
	EncodingType_ENCODING_TYPE_JSON        EncodingType = 2
)

// Enum value maps for EncodingType.
var (
	EncodingType_name = map[int32]string{
		0: "ENCODING_TYPE_UNSPECIFIED",
		1: "ENCODING_TYPE_PROTO3",
		2: "ENCODING_TYPE_JSON",
	}
	EncodingType_value = map[string]int32{
		"ENCODING_TYPE_UNSPECIFIED": 0,
		"ENCODING_TYPE_PROTO3":      1,
		"ENCODING_TYPE_JSON":        2,
	}
)

func (x EncodingType) Enum() *EncodingType {
	p := new(EncodingType)
	*p = x
	return p
}

func (x EncodingType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EncodingType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[0].Descriptor()
}

func (EncodingType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[0]
}

func (x EncodingType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EncodingType.Descriptor instead.
func (EncodingType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{0}
}

type IndexedValueType int32

const (
	IndexedValueType_INDEXED_VALUE_TYPE_UNSPECIFIED  IndexedValueType = 0
	IndexedValueType_INDEXED_VALUE_TYPE_TEXT         IndexedValueType = 1
	IndexedValueType_INDEXED_VALUE_TYPE_KEYWORD      IndexedValueType = 2
	IndexedValueType_INDEXED_VALUE_TYPE_INT          IndexedValueType = 3
	IndexedValueType_INDEXED_VALUE_TYPE_DOUBLE       IndexedValueType = 4
	IndexedValueType_INDEXED_VALUE_TYPE_BOOL         IndexedValueType = 5
	IndexedValueType_INDEXED_VALUE_TYPE_DATETIME     IndexedValueType = 6
	IndexedValueType_INDEXED_VALUE_TYPE_KEYWORD_LIST IndexedValueType = 7
)

// Enum value maps for IndexedValueType.
var (
	IndexedValueType_name = map[int32]string{
		0: "INDEXED_VALUE_TYPE_UNSPECIFIED",
		1: "INDEXED_VALUE_TYPE_TEXT",
		2: "INDEXED_VALUE_TYPE_KEYWORD",
		3: "INDEXED_VALUE_TYPE_INT",
		4: "INDEXED_VALUE_TYPE_DOUBLE",
		5: "INDEXED_VALUE_TYPE_BOOL",
		6: "INDEXED_VALUE_TYPE_DATETIME",
		7: "INDEXED_VALUE_TYPE_KEYWORD_LIST",
	}
	IndexedValueType_value = map[string]int32{
		"INDEXED_VALUE_TYPE_UNSPECIFIED":  0,
		"INDEXED_VALUE_TYPE_TEXT":         1,
		"INDEXED_VALUE_TYPE_KEYWORD":      2,
		"INDEXED_VALUE_TYPE_INT":          3,
		"INDEXED_VALUE_TYPE_DOUBLE":       4,
		"INDEXED_VALUE_TYPE_BOOL":         5,
		"INDEXED_VALUE_TYPE_DATETIME":     6,
		"INDEXED_VALUE_TYPE_KEYWORD_LIST": 7,
	}
)

func (x IndexedValueType) Enum() *IndexedValueType {
	p := new(IndexedValueType)
	*p = x
	return p
}

func (x IndexedValueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IndexedValueType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[1].Descriptor()
}

func (IndexedValueType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[1]
}

func (x IndexedValueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IndexedValueType.Descriptor instead.
func (IndexedValueType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{1}
}

type Severity int32

const (
	Severity_SEVERITY_UNSPECIFIED Severity = 0
	Severity_SEVERITY_HIGH        Severity = 1
	Severity_SEVERITY_MEDIUM      Severity = 2
	Severity_SEVERITY_LOW         Severity = 3
)

// Enum value maps for Severity.
var (
	Severity_name = map[int32]string{
		0: "SEVERITY_UNSPECIFIED",
		1: "SEVERITY_HIGH",
		2: "SEVERITY_MEDIUM",
		3: "SEVERITY_LOW",
	}
	Severity_value = map[string]int32{
		"SEVERITY_UNSPECIFIED": 0,
		"SEVERITY_HIGH":        1,
		"SEVERITY_MEDIUM":      2,
		"SEVERITY_LOW":         3,
	}
)

func (x Severity) Enum() *Severity {
	p := new(Severity)
	*p = x
	return p
}

func (x Severity) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Severity) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_common_proto_enumTypes[2].Descriptor()
}

func (Severity) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_common_proto_enumTypes[2]
}

func (x Severity) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Severity.Descriptor instead.
func (Severity) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_common_proto_rawDescGZIP(), []int{2}
}

var File_temporal_api_enums_v1_common_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_common_proto_rawDesc = []byte{
	0x0a, 0x22, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x5f, 0x0a, 0x0c, 0x45,
	0x6e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x45,
	0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4e,
	0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x54,
	0x4f, 0x33, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x45, 0x4e, 0x43, 0x4f, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4a, 0x53, 0x4f, 0x4e, 0x10, 0x02, 0x2a, 0x91, 0x02, 0x0a,
	0x10, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x22, 0x0a, 0x1e, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c,
	0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44,
	0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x54, 0x45, 0x58, 0x54,
	0x10, 0x01, 0x12, 0x1e, 0x0a, 0x1a, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44,
	0x10, 0x02, 0x12, 0x1a, 0x0a, 0x16, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41,
	0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x1d,
	0x0a, 0x19, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x44, 0x4f, 0x55, 0x42, 0x4c, 0x45, 0x10, 0x04, 0x12, 0x1b, 0x0a,
	0x17, 0x49, 0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x42, 0x4f, 0x4f, 0x4c, 0x10, 0x05, 0x12, 0x1f, 0x0a, 0x1b, 0x49, 0x4e,
	0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x44, 0x41, 0x54, 0x45, 0x54, 0x49, 0x4d, 0x45, 0x10, 0x06, 0x12, 0x23, 0x0a, 0x1f, 0x49,
	0x4e, 0x44, 0x45, 0x58, 0x45, 0x44, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x4b, 0x45, 0x59, 0x57, 0x4f, 0x52, 0x44, 0x5f, 0x4c, 0x49, 0x53, 0x54, 0x10, 0x07,
	0x2a, 0x5e, 0x0a, 0x08, 0x53, 0x65, 0x76, 0x65, 0x72, 0x69, 0x74, 0x79, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49,
	0x54, 0x59, 0x5f, 0x48, 0x49, 0x47, 0x48, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x56,
	0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4d, 0x45, 0x44, 0x49, 0x55, 0x4d, 0x10, 0x02, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x45, 0x56, 0x45, 0x52, 0x49, 0x54, 0x59, 0x5f, 0x4c, 0x4f, 0x57, 0x10, 0x03,
	0x42, 0x83, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0b, 0x43,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa,
	0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69,
	0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75,
	0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_common_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_common_proto_rawDescData = file_temporal_api_enums_v1_common_proto_rawDesc
)

func file_temporal_api_enums_v1_common_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_common_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_common_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_common_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_common_proto_rawDescData
}

var file_temporal_api_enums_v1_common_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_temporal_api_enums_v1_common_proto_goTypes = []interface{}{
	(EncodingType)(0),     // 0: temporal.api.enums.v1.EncodingType
	(IndexedValueType)(0), // 1: temporal.api.enums.v1.IndexedValueType
	(Severity)(0),         // 2: temporal.api.enums.v1.Severity
}
var file_temporal_api_enums_v1_common_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_common_proto_init() }
func file_temporal_api_enums_v1_common_proto_init() {
	if File_temporal_api_enums_v1_common_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_common_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_common_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_common_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_common_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_common_proto = out.File
	file_temporal_api_enums_v1_common_proto_rawDesc = nil
	file_temporal_api_enums_v1_common_proto_goTypes = nil
	file_temporal_api_enums_v1_common_proto_depIdxs = nil
}
