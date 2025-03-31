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
// source: temporal/api/enums/v1/namespace.proto

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

type NamespaceState int32

const (
	NAMESPACE_STATE_UNSPECIFIED NamespaceState = 0
	NAMESPACE_STATE_REGISTERED  NamespaceState = 1
	NAMESPACE_STATE_DEPRECATED  NamespaceState = 2
	NAMESPACE_STATE_DELETED     NamespaceState = 3
)

// Enum value maps for NamespaceState.
var (
	NamespaceState_name = map[int32]string{
		0: "NAMESPACE_STATE_UNSPECIFIED",
		1: "NAMESPACE_STATE_REGISTERED",
		2: "NAMESPACE_STATE_DEPRECATED",
		3: "NAMESPACE_STATE_DELETED",
	}
	NamespaceState_value = map[string]int32{
		"NAMESPACE_STATE_UNSPECIFIED": 0,
		"NAMESPACE_STATE_REGISTERED":  1,
		"NAMESPACE_STATE_DEPRECATED":  2,
		"NAMESPACE_STATE_DELETED":     3,
	}
)

func (x NamespaceState) Enum() *NamespaceState {
	p := new(NamespaceState)
	*p = x
	return p
}

func (x NamespaceState) String() string {
	switch x {
	case NAMESPACE_STATE_UNSPECIFIED:
		return "Unspecified"
	case NAMESPACE_STATE_REGISTERED:
		return "Registered"
	case NAMESPACE_STATE_DEPRECATED:
		return "Deprecated"
	case NAMESPACE_STATE_DELETED:
		return "Deleted"
	default:
		return strconv.Itoa(int(x))
	}

}

func (NamespaceState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_namespace_proto_enumTypes[0].Descriptor()
}

func (NamespaceState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_namespace_proto_enumTypes[0]
}

func (x NamespaceState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NamespaceState.Descriptor instead.
func (NamespaceState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_namespace_proto_rawDescGZIP(), []int{0}
}

type ArchivalState int32

const (
	ARCHIVAL_STATE_UNSPECIFIED ArchivalState = 0
	ARCHIVAL_STATE_DISABLED    ArchivalState = 1
	ARCHIVAL_STATE_ENABLED     ArchivalState = 2
)

// Enum value maps for ArchivalState.
var (
	ArchivalState_name = map[int32]string{
		0: "ARCHIVAL_STATE_UNSPECIFIED",
		1: "ARCHIVAL_STATE_DISABLED",
		2: "ARCHIVAL_STATE_ENABLED",
	}
	ArchivalState_value = map[string]int32{
		"ARCHIVAL_STATE_UNSPECIFIED": 0,
		"ARCHIVAL_STATE_DISABLED":    1,
		"ARCHIVAL_STATE_ENABLED":     2,
	}
)

func (x ArchivalState) Enum() *ArchivalState {
	p := new(ArchivalState)
	*p = x
	return p
}

func (x ArchivalState) String() string {
	switch x {
	case ARCHIVAL_STATE_UNSPECIFIED:
		return "Unspecified"
	case ARCHIVAL_STATE_DISABLED:
		return "Disabled"
	case ARCHIVAL_STATE_ENABLED:
		return "Enabled"
	default:
		return strconv.Itoa(int(x))
	}

}

func (ArchivalState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_namespace_proto_enumTypes[1].Descriptor()
}

func (ArchivalState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_namespace_proto_enumTypes[1]
}

func (x ArchivalState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ArchivalState.Descriptor instead.
func (ArchivalState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_namespace_proto_rawDescGZIP(), []int{1}
}

type ReplicationState int32

const (
	REPLICATION_STATE_UNSPECIFIED ReplicationState = 0
	REPLICATION_STATE_NORMAL      ReplicationState = 1
	REPLICATION_STATE_HANDOVER    ReplicationState = 2
)

// Enum value maps for ReplicationState.
var (
	ReplicationState_name = map[int32]string{
		0: "REPLICATION_STATE_UNSPECIFIED",
		1: "REPLICATION_STATE_NORMAL",
		2: "REPLICATION_STATE_HANDOVER",
	}
	ReplicationState_value = map[string]int32{
		"REPLICATION_STATE_UNSPECIFIED": 0,
		"REPLICATION_STATE_NORMAL":      1,
		"REPLICATION_STATE_HANDOVER":    2,
	}
)

func (x ReplicationState) Enum() *ReplicationState {
	p := new(ReplicationState)
	*p = x
	return p
}

func (x ReplicationState) String() string {
	switch x {
	case REPLICATION_STATE_UNSPECIFIED:
		return "Unspecified"
	case REPLICATION_STATE_NORMAL:
		return "Normal"
	case REPLICATION_STATE_HANDOVER:
		return "Handover"
	default:
		return strconv.Itoa(int(x))
	}

}

func (ReplicationState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_namespace_proto_enumTypes[2].Descriptor()
}

func (ReplicationState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_namespace_proto_enumTypes[2]
}

func (x ReplicationState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReplicationState.Descriptor instead.
func (ReplicationState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_namespace_proto_rawDescGZIP(), []int{2}
}

var File_temporal_api_enums_v1_namespace_proto protoreflect.FileDescriptor

const file_temporal_api_enums_v1_namespace_proto_rawDesc = "" +
	"\n" +
	"%temporal/api/enums/v1/namespace.proto\x12\x15temporal.api.enums.v1*\x8e\x01\n" +
	"\x0eNamespaceState\x12\x1f\n" +
	"\x1bNAMESPACE_STATE_UNSPECIFIED\x10\x00\x12\x1e\n" +
	"\x1aNAMESPACE_STATE_REGISTERED\x10\x01\x12\x1e\n" +
	"\x1aNAMESPACE_STATE_DEPRECATED\x10\x02\x12\x1b\n" +
	"\x17NAMESPACE_STATE_DELETED\x10\x03*h\n" +
	"\rArchivalState\x12\x1e\n" +
	"\x1aARCHIVAL_STATE_UNSPECIFIED\x10\x00\x12\x1b\n" +
	"\x17ARCHIVAL_STATE_DISABLED\x10\x01\x12\x1a\n" +
	"\x16ARCHIVAL_STATE_ENABLED\x10\x02*s\n" +
	"\x10ReplicationState\x12!\n" +
	"\x1dREPLICATION_STATE_UNSPECIFIED\x10\x00\x12\x1c\n" +
	"\x18REPLICATION_STATE_NORMAL\x10\x01\x12\x1e\n" +
	"\x1aREPLICATION_STATE_HANDOVER\x10\x02B\x86\x01\n" +
	"\x18io.temporal.api.enums.v1B\x0eNamespaceProtoP\x01Z!go.temporal.io/api/enums/v1;enums\xaa\x02\x17Temporalio.Api.Enums.V1\xea\x02\x1aTemporalio::Api::Enums::V1b\x06proto3"

var (
	file_temporal_api_enums_v1_namespace_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_namespace_proto_rawDescData []byte
)

func file_temporal_api_enums_v1_namespace_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_namespace_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_namespace_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_namespace_proto_rawDesc), len(file_temporal_api_enums_v1_namespace_proto_rawDesc)))
	})
	return file_temporal_api_enums_v1_namespace_proto_rawDescData
}

var file_temporal_api_enums_v1_namespace_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_temporal_api_enums_v1_namespace_proto_goTypes = []any{
	(NamespaceState)(0),   // 0: temporal.api.enums.v1.NamespaceState
	(ArchivalState)(0),    // 1: temporal.api.enums.v1.ArchivalState
	(ReplicationState)(0), // 2: temporal.api.enums.v1.ReplicationState
}
var file_temporal_api_enums_v1_namespace_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_namespace_proto_init() }
func file_temporal_api_enums_v1_namespace_proto_init() {
	if File_temporal_api_enums_v1_namespace_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_namespace_proto_rawDesc), len(file_temporal_api_enums_v1_namespace_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_namespace_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_namespace_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_namespace_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_namespace_proto = out.File
	file_temporal_api_enums_v1_namespace_proto_goTypes = nil
	file_temporal_api_enums_v1_namespace_proto_depIdxs = nil
}
