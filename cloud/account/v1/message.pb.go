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
// source: temporal/api/cloud/account/v1/message.proto

package account

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/cloud/resource/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetricsSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The ca cert(s) in PEM format that clients connecting to the metrics endpoint can use for authentication.
	// This must only be one value, but the CA can have a chain.
	AcceptedClientCa []byte `protobuf:"bytes,2,opt,name=accepted_client_ca,json=acceptedClientCa,proto3" json:"accepted_client_ca,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *MetricsSpec) Reset() {
	*x = MetricsSpec{}
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsSpec) ProtoMessage() {}

func (x *MetricsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsSpec.ProtoReflect.Descriptor instead.
func (*MetricsSpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_account_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *MetricsSpec) GetAcceptedClientCa() []byte {
	if x != nil {
		return x.AcceptedClientCa
	}
	return nil
}

type AccountSpec struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The metrics specification for this account.
	// If not specified, metrics will not be enabled.
	Metrics       *MetricsSpec `protobuf:"bytes,1,opt,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccountSpec) Reset() {
	*x = AccountSpec{}
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccountSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountSpec) ProtoMessage() {}

func (x *AccountSpec) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountSpec.ProtoReflect.Descriptor instead.
func (*AccountSpec) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_account_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *AccountSpec) GetMetrics() *MetricsSpec {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type Metrics struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The prometheus metrics endpoint uri.
	// This is only populated when the metrics is enabled in the metrics specification.
	Uri           string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_account_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *Metrics) GetUri() string {
	if x != nil {
		return x.Uri
	}
	return ""
}

type Account struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The id of the account.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// The account specification.
	Spec *AccountSpec `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	// The current version of the account specification.
	// The next update operation will have to include this version.
	ResourceVersion string `protobuf:"bytes,3,opt,name=resource_version,json=resourceVersion,proto3" json:"resource_version,omitempty"`
	// The current state of the account.
	State v1.ResourceState `protobuf:"varint,4,opt,name=state,proto3,enum=temporal.api.cloud.resource.v1.ResourceState" json:"state,omitempty"`
	// The id of the async operation that is updating the account, if any.
	AsyncOperationId string `protobuf:"bytes,5,opt,name=async_operation_id,json=asyncOperationId,proto3" json:"async_operation_id,omitempty"`
	// Information related to metrics.
	Metrics       *Metrics `protobuf:"bytes,6,opt,name=metrics,proto3" json:"metrics,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Account) Reset() {
	*x = Account{}
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_cloud_account_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_temporal_api_cloud_account_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *Account) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Account) GetSpec() *AccountSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

func (x *Account) GetResourceVersion() string {
	if x != nil {
		return x.ResourceVersion
	}
	return ""
}

func (x *Account) GetState() v1.ResourceState {
	if x != nil {
		return x.State
	}
	return v1.ResourceState(0)
}

func (x *Account) GetAsyncOperationId() string {
	if x != nil {
		return x.AsyncOperationId
	}
	return ""
}

func (x *Account) GetMetrics() *Metrics {
	if x != nil {
		return x.Metrics
	}
	return nil
}

var File_temporal_api_cloud_account_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_cloud_account_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c,
	0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x2c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3f, 0x0a, 0x0b, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x30, 0x0a, 0x12, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x5f, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x63, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x61,
	0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x43, 0x61, 0x42, 0x02,
	0x68, 0x00, 0x22, 0x57, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x48, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x53, 0x70, 0x65, 0x63, 0x52, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x42, 0x02, 0x68, 0x00, 0x22, 0x1f, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x12, 0x14, 0x0a, 0x03, 0x75, 0x72, 0x69, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x72, 0x69, 0x42, 0x02, 0x68, 0x00, 0x22, 0xd1, 0x02, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x42, 0x0a, 0x04, 0x73, 0x70, 0x65, 0x63, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x52, 0x04, 0x73,
	0x70, 0x65, 0x63, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2d, 0x0a, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x47, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x30, 0x0a, 0x12, 0x61, 0x73, 0x79,
	0x6e, 0x63, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x44, 0x0a, 0x07, 0x6d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x02, 0x68, 0x00, 0x42, 0xa7, 0x01, 0x0a,
	0x20, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31,
	0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x2b, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0xaa, 0x02, 0x1f, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x6c, 0x6f, 0x75, 0x64,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x23, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x6c,
	0x6f, 0x75, 0x64, 0x3a, 0x3a, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_cloud_account_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_cloud_account_v1_message_proto_rawDescData = file_temporal_api_cloud_account_v1_message_proto_rawDesc
)

func file_temporal_api_cloud_account_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_cloud_account_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_cloud_account_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_cloud_account_v1_message_proto_rawDescData)
	})
	return file_temporal_api_cloud_account_v1_message_proto_rawDescData
}

var file_temporal_api_cloud_account_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_temporal_api_cloud_account_v1_message_proto_goTypes = []any{
	(*MetricsSpec)(nil),   // 0: temporal.api.cloud.account.v1.MetricsSpec
	(*AccountSpec)(nil),   // 1: temporal.api.cloud.account.v1.AccountSpec
	(*Metrics)(nil),       // 2: temporal.api.cloud.account.v1.Metrics
	(*Account)(nil),       // 3: temporal.api.cloud.account.v1.Account
	(v1.ResourceState)(0), // 4: temporal.api.cloud.resource.v1.ResourceState
}
var file_temporal_api_cloud_account_v1_message_proto_depIdxs = []int32{
	0, // 0: temporal.api.cloud.account.v1.AccountSpec.metrics:type_name -> temporal.api.cloud.account.v1.MetricsSpec
	1, // 1: temporal.api.cloud.account.v1.Account.spec:type_name -> temporal.api.cloud.account.v1.AccountSpec
	4, // 2: temporal.api.cloud.account.v1.Account.state:type_name -> temporal.api.cloud.resource.v1.ResourceState
	2, // 3: temporal.api.cloud.account.v1.Account.metrics:type_name -> temporal.api.cloud.account.v1.Metrics
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_temporal_api_cloud_account_v1_message_proto_init() }
func file_temporal_api_cloud_account_v1_message_proto_init() {
	if File_temporal_api_cloud_account_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_cloud_account_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_cloud_account_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_cloud_account_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_cloud_account_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_cloud_account_v1_message_proto = out.File
	file_temporal_api_cloud_account_v1_message_proto_rawDesc = nil
	file_temporal_api_cloud_account_v1_message_proto_goTypes = nil
	file_temporal_api_cloud_account_v1_message_proto_depIdxs = nil
}
