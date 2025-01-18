// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.2
// source: internal/protojson/testprotos/test/weak2/test_weak.proto

package weak2

import (
	reflect "reflect"
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

type WeakImportMessage2 struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	A             *int32                 `protobuf:"varint,1,req,name=a" json:"a,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WeakImportMessage2) Reset() {
	*x = WeakImportMessage2{}
	mi := &file_internal_protojson_testprotos_test_weak2_test_weak_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WeakImportMessage2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WeakImportMessage2) ProtoMessage() {}

func (x *WeakImportMessage2) ProtoReflect() protoreflect.Message {
	mi := &file_internal_protojson_testprotos_test_weak2_test_weak_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WeakImportMessage2.ProtoReflect.Descriptor instead.
func (*WeakImportMessage2) Descriptor() ([]byte, []int) {
	return file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDescGZIP(), []int{0}
}

func (x *WeakImportMessage2) GetA() int32 {
	if x != nil && x.A != nil {
		return *x.A
	}
	return 0
}

var File_internal_protojson_testprotos_test_weak2_test_weak_proto protoreflect.FileDescriptor

var file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDesc = string([]byte{
	0x0a, 0x38, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x6a, 0x73, 0x6f, 0x6e, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x74, 0x65, 0x73, 0x74, 0x2f, 0x77, 0x65, 0x61, 0x6b, 0x32, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f,
	0x77, 0x65, 0x61, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x67, 0x6f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x77,
	0x65, 0x61, 0x6b, 0x22, 0x22, 0x0a, 0x12, 0x57, 0x65, 0x61, 0x6b, 0x49, 0x6d, 0x70, 0x6f, 0x72,
	0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x01, 0x61, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x6f, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6a, 0x73, 0x6f, 0x6e,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x74, 0x65, 0x73, 0x74,
	0x2f, 0x77, 0x65, 0x61, 0x6b, 0x32,
})

var (
	file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDescOnce sync.Once
	file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDescData []byte
)

func file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDescGZIP() []byte {
	file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDescOnce.Do(func() {
		file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDesc), len(file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDesc)))
	})
	return file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDescData
}

var file_internal_protojson_testprotos_test_weak2_test_weak_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_protojson_testprotos_test_weak2_test_weak_proto_goTypes = []any{
	(*WeakImportMessage2)(nil), // 0: goproto.proto.test.weak.WeakImportMessage2
}
var file_internal_protojson_testprotos_test_weak2_test_weak_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_protojson_testprotos_test_weak2_test_weak_proto_init() }
func file_internal_protojson_testprotos_test_weak2_test_weak_proto_init() {
	if File_internal_protojson_testprotos_test_weak2_test_weak_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDesc), len(file_internal_protojson_testprotos_test_weak2_test_weak_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_protojson_testprotos_test_weak2_test_weak_proto_goTypes,
		DependencyIndexes: file_internal_protojson_testprotos_test_weak2_test_weak_proto_depIdxs,
		MessageInfos:      file_internal_protojson_testprotos_test_weak2_test_weak_proto_msgTypes,
	}.Build()
	File_internal_protojson_testprotos_test_weak2_test_weak_proto = out.File
	file_internal_protojson_testprotos_test_weak2_test_weak_proto_goTypes = nil
	file_internal_protojson_testprotos_test_weak2_test_weak_proto_depIdxs = nil
}
