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
// source: temporal/api/cloud/region/v1/message.proto

package region

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

// The cloud provider that's hosting the region.
type Region_CloudProvider int32

const (
	Region_CLOUD_PROVIDER_UNSPECIFIED Region_CloudProvider = 0
	Region_CLOUD_PROVIDER_AWS         Region_CloudProvider = 1
	Region_CLOUD_PROVIDER_GCP         Region_CloudProvider = 2
)

// Enum value maps for Region_CloudProvider.
var (
	Region_CloudProvider_name = map[int32]string{
		0: "CLOUD_PROVIDER_UNSPECIFIED",
		1: "CLOUD_PROVIDER_AWS",
		2: "CLOUD_PROVIDER_GCP",
	}
	Region_CloudProvider_value = map[string]int32{
		"CLOUD_PROVIDER_UNSPECIFIED": 0,
		"CLOUD_PROVIDER_AWS":         1,
		"CLOUD_PROVIDER_GCP":         2,
	}
)

func (x Region_CloudProvider) Enum() *Region_CloudProvider {
	p := new(Region_CloudProvider)
	*p = x
	return p
}

func (x Region_CloudProvider) String() string {
	switch x {
	case Region_CLOUD_PROVIDER_UNSPECIFIED:
		return "RegionCloudProviderUnspecified"
	case Region_CLOUD_PROVIDER_AWS:
		return "RegionCloudProviderAws"
	case Region_CLOUD_PROVIDER_GCP:
		return "RegionCloudProviderGcp"
	default:
		return strconv.Itoa(int(x))
	}

}

func (Region_CloudProvider) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_cloud_region_v1_message_proto_enumTypes[0].Descriptor()
}

func (Region_CloudProvider) Type() protoreflect.EnumType {
	return &file_temporal_api_cloud_region_v1_message_proto_enumTypes[0]
}

func (x Region_CloudProvider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Region_CloudProvider.Descriptor instead.
func (Region_CloudProvider) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_cloud_region_v1_message_proto_rawDescGZIP(), []int{0, 0}
}

type Region struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the temporal cloud region.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The name of the cloud provider that's hosting the region.
	// Currently only "aws" is supported.
	// Deprecated: Not supported after 2024-10-01-00 api version. Use cloud_provider instead.
	// temporal:versioning:max_version=2024-10-01-00
	//
	// Deprecated: Marked as deprecated in temporal/api/cloud/region/v1/message.proto.
	CloudProviderDeprecated string `protobuf:"bytes,2,opt,name=cloud_provider_deprecated,json=cloudProviderDeprecated,proto3" json:"cloud_provider_deprecated,omitempty"`
	// The cloud provider that's hosting the region.
	// temporal:versioning:min_version=2024-10-01-00
	// temporal:enums:replaces=cloud_provider_deprecated
	CloudProvider Region_CloudProvider `protobuf:"varint,5,opt,name=cloud_provider,json=cloudProvider,proto3,enum=temporal.api.cloud.region.v1.Region_CloudProvider" json:"cloud_provider,omitempty"`
	// The region identifier as defined by the cloud provider.
	CloudProviderRegion string `protobuf:"bytes,3,opt,name=cloud_provider_region,json=cloudProviderRegion,proto3" json:"cloud_provider_region,omitempty"`
	// The human readable location of the region.
	Location      string `protobuf:"bytes,4,opt,name=location,proto3" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Region) Reset() {
	*x = Region{}
	mi := &file_temporal_api_cloud_region_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Region) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Region) ProtoMessage() {}

func (x *Region) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_region_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Region.ProtoReflect.Descriptor instead.
func (*Region) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_region_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Region) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

// Deprecated: Marked as deprecated in temporal/api/cloud/region/v1/message.proto.
func (x *Region) GetCloudProviderDeprecated() string {
	if x != nil {
		return x.CloudProviderDeprecated
	}
	return ""
}

func (x *Region) GetCloudProvider() Region_CloudProvider {
	if x != nil {
		return x.CloudProvider
	}
	return Region_CLOUD_PROVIDER_UNSPECIFIED
}

func (x *Region) GetCloudProviderRegion() string {
	if x != nil {
		return x.CloudProviderRegion
	}
	return ""
}

func (x *Region) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

var File_temporal_api_cloud_region_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_cloud_region_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0xe4, 0x02, 0x0a, 0x06, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3e, 0x0a, 0x19, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x02, 0x18, 0x01, 0x52, 0x17, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x44, 0x65, 0x70, 0x72, 0x65,
	0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x59, 0x0a, 0x0e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x32, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f,
	0x75, 0x64, 0x2e, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67,
	0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x0d, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x13, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x5f, 0x0a, 0x0d, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49,
	0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4c, 0x4f, 0x55, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49,
	0x44, 0x45, 0x52, 0x5f, 0x41, 0x57, 0x53, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x43, 0x4c, 0x4f,
	0x55, 0x44, 0x5f, 0x50, 0x52, 0x4f, 0x56, 0x49, 0x44, 0x45, 0x52, 0x5f, 0x47, 0x43, 0x50, 0x10,
	0x02, 0x42, 0xa2, 0x01, 0x0a, 0x1f, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0xaa, 0x02, 0x1e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x2e, 0x56,
	0x31, 0xea, 0x02, 0x22, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a,
	0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6c, 0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_cloud_region_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_region_v1_message_proto_rawDescData = file_temporal_api_cloud_region_v1_message_proto_rawDesc
)

func file_temporal_api_cloud_region_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_region_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_region_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_cloud_region_v1_message_proto_rawDescData)
	})
	return file_temporal_api_cloud_region_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_region_v1_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temporal_api_cloud_region_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_api_cloud_region_v1_message_proto_goTypes = []any{
	(Region_CloudProvider)(0), // 0: temporal.api.cloud.region.v1.Region.CloudProvider
	(*Region)(nil),            // 1: temporal.api.cloud.region.v1.Region
}
var file_temporal_api_cloud_region_v1_message_proto_depIdxs = []int32{
	0, // 0: temporal.api.cloud.region.v1.Region.cloud_provider:type_name -> temporal.api.cloud.region.v1.Region.CloudProvider
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_region_v1_message_proto_init() }
func file_temporal_api_cloud_region_v1_message_proto_init() {
	if File_temporal_api_cloud_region_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_cloud_region_v1_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_region_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_region_v1_message_proto_depIdxs,
		EnumInfos:         file_temporal_api_cloud_region_v1_message_proto_enumTypes,
		MessageInfos:      file_temporal_api_cloud_region_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_region_v1_message_proto = out.File
	file_temporal_api_cloud_region_v1_message_proto_rawDesc = nil
	file_temporal_api_cloud_region_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_region_v1_message_proto_depIdxs = nil
}
