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
// source: temporal/api/namespace/v1/message.proto

package namespace

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/enums/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NamespaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	State       v1.NamespaceState `protobuf:"varint,2,opt,name=state,proto3,enum=temporal.api.enums.v1.NamespaceState" json:"state,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	OwnerEmail  string            `protobuf:"bytes,4,opt,name=owner_email,json=ownerEmail,proto3" json:"owner_email,omitempty"`
	// A key-value map for any customized purpose.
	Data map[string][]byte `protobuf:"bytes,5,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Id   string            `protobuf:"bytes,6,opt,name=id,proto3" json:"id,omitempty"`
	// All capabilities the namespace supports.
	Capabilities *NamespaceInfo_Capabilities `protobuf:"bytes,7,opt,name=capabilities,proto3" json:"capabilities,omitempty"`
	// Whether scheduled workflows are supported on this namespace. This is only needed
	// temporarily while the feature is experimental, so we can give it a high tag.
	SupportsSchedules bool `protobuf:"varint,100,opt,name=supports_schedules,json=supportsSchedules,proto3" json:"supports_schedules,omitempty"`
}

func (x *NamespaceInfo) Reset() {
	*x = NamespaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceInfo) ProtoMessage() {}

func (x *NamespaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceInfo.ProtoReflect.Descriptor instead.
func (*NamespaceInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_namespace_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *NamespaceInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NamespaceInfo) GetState() v1.NamespaceState {
	if x != nil {
		return x.State
	}
	return v1.NamespaceState(0)
}

func (x *NamespaceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *NamespaceInfo) GetOwnerEmail() string {
	if x != nil {
		return x.OwnerEmail
	}
	return ""
}

func (x *NamespaceInfo) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *NamespaceInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *NamespaceInfo) GetCapabilities() *NamespaceInfo_Capabilities {
	if x != nil {
		return x.Capabilities
	}
	return nil
}

func (x *NamespaceInfo) GetSupportsSchedules() bool {
	if x != nil {
		return x.SupportsSchedules
	}
	return false
}

type NamespaceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WorkflowExecutionRetentionTtl *durationpb.Duration `protobuf:"bytes,1,opt,name=workflow_execution_retention_ttl,json=workflowExecutionRetentionTtl,proto3" json:"workflow_execution_retention_ttl,omitempty"`
	BadBinaries                   *BadBinaries         `protobuf:"bytes,2,opt,name=bad_binaries,json=badBinaries,proto3" json:"bad_binaries,omitempty"`
	// If unspecified (ARCHIVAL_STATE_UNSPECIFIED) then default server configuration is used.
	HistoryArchivalState v1.ArchivalState `protobuf:"varint,3,opt,name=history_archival_state,json=historyArchivalState,proto3,enum=temporal.api.enums.v1.ArchivalState" json:"history_archival_state,omitempty"`
	HistoryArchivalUri   string           `protobuf:"bytes,4,opt,name=history_archival_uri,json=historyArchivalUri,proto3" json:"history_archival_uri,omitempty"`
	// If unspecified (ARCHIVAL_STATE_UNSPECIFIED) then default server configuration is used.
	VisibilityArchivalState v1.ArchivalState `protobuf:"varint,5,opt,name=visibility_archival_state,json=visibilityArchivalState,proto3,enum=temporal.api.enums.v1.ArchivalState" json:"visibility_archival_state,omitempty"`
	VisibilityArchivalUri   string           `protobuf:"bytes,6,opt,name=visibility_archival_uri,json=visibilityArchivalUri,proto3" json:"visibility_archival_uri,omitempty"`
	// Map from field name to alias.
	CustomSearchAttributeAliases map[string]string `protobuf:"bytes,7,rep,name=custom_search_attribute_aliases,json=customSearchAttributeAliases,proto3" json:"custom_search_attribute_aliases,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *NamespaceConfig) Reset() {
	*x = NamespaceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceConfig) ProtoMessage() {}

func (x *NamespaceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceConfig.ProtoReflect.Descriptor instead.
func (*NamespaceConfig) Descriptor() ([]byte, []int) {
	return file_temporal_api_namespace_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *NamespaceConfig) GetWorkflowExecutionRetentionTtl() *durationpb.Duration {
	if x != nil {
		return x.WorkflowExecutionRetentionTtl
	}
	return nil
}

func (x *NamespaceConfig) GetBadBinaries() *BadBinaries {
	if x != nil {
		return x.BadBinaries
	}
	return nil
}

func (x *NamespaceConfig) GetHistoryArchivalState() v1.ArchivalState {
	if x != nil {
		return x.HistoryArchivalState
	}
	return v1.ArchivalState(0)
}

func (x *NamespaceConfig) GetHistoryArchivalUri() string {
	if x != nil {
		return x.HistoryArchivalUri
	}
	return ""
}

func (x *NamespaceConfig) GetVisibilityArchivalState() v1.ArchivalState {
	if x != nil {
		return x.VisibilityArchivalState
	}
	return v1.ArchivalState(0)
}

func (x *NamespaceConfig) GetVisibilityArchivalUri() string {
	if x != nil {
		return x.VisibilityArchivalUri
	}
	return ""
}

func (x *NamespaceConfig) GetCustomSearchAttributeAliases() map[string]string {
	if x != nil {
		return x.CustomSearchAttributeAliases
	}
	return nil
}

type BadBinaries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Binaries map[string]*BadBinaryInfo `protobuf:"bytes,1,rep,name=binaries,proto3" json:"binaries,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *BadBinaries) Reset() {
	*x = BadBinaries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadBinaries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadBinaries) ProtoMessage() {}

func (x *BadBinaries) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadBinaries.ProtoReflect.Descriptor instead.
func (*BadBinaries) Descriptor() ([]byte, []int) {
	return file_temporal_api_namespace_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *BadBinaries) GetBinaries() map[string]*BadBinaryInfo {
	if x != nil {
		return x.Binaries
	}
	return nil
}

type BadBinaryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reason     string                 `protobuf:"bytes,1,opt,name=reason,proto3" json:"reason,omitempty"`
	Operator   string                 `protobuf:"bytes,2,opt,name=operator,proto3" json:"operator,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
}

func (x *BadBinaryInfo) Reset() {
	*x = BadBinaryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BadBinaryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BadBinaryInfo) ProtoMessage() {}

func (x *BadBinaryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BadBinaryInfo.ProtoReflect.Descriptor instead.
func (*BadBinaryInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_namespace_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *BadBinaryInfo) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *BadBinaryInfo) GetOperator() string {
	if x != nil {
		return x.Operator
	}
	return ""
}

func (x *BadBinaryInfo) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

type UpdateNamespaceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Description string `protobuf:"bytes,1,opt,name=description,proto3" json:"description,omitempty"`
	OwnerEmail  string `protobuf:"bytes,2,opt,name=owner_email,json=ownerEmail,proto3" json:"owner_email,omitempty"`
	// A key-value map for any customized purpose.
	// If data already exists on the namespace,
	// this will merge with the existing key values.
	Data map[string][]byte `protobuf:"bytes,3,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// New namespace state, server will reject if transition is not allowed.
	// Allowed transitions are:
	//
	//	Registered -> [ Deleted | Deprecated | Handover ]
	//	Handover -> [ Registered ]
	//
	// Default is NAMESPACE_STATE_UNSPECIFIED which is do not change state.
	State v1.NamespaceState `protobuf:"varint,4,opt,name=state,proto3,enum=temporal.api.enums.v1.NamespaceState" json:"state,omitempty"`
}

func (x *UpdateNamespaceInfo) Reset() {
	*x = UpdateNamespaceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateNamespaceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateNamespaceInfo) ProtoMessage() {}

func (x *UpdateNamespaceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateNamespaceInfo.ProtoReflect.Descriptor instead.
func (*UpdateNamespaceInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_namespace_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateNamespaceInfo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateNamespaceInfo) GetOwnerEmail() string {
	if x != nil {
		return x.OwnerEmail
	}
	return ""
}

func (x *UpdateNamespaceInfo) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UpdateNamespaceInfo) GetState() v1.NamespaceState {
	if x != nil {
		return x.State
	}
	return v1.NamespaceState(0)
}

type NamespaceFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// By default namespaces in NAMESPACE_STATE_DELETED state are not included.
	// Setting include_deleted to true will include deleted namespaces.
	// Note: Namespace is in NAMESPACE_STATE_DELETED state when it was deleted from the system but associated data is not deleted yet.
	IncludeDeleted bool `protobuf:"varint,1,opt,name=include_deleted,json=includeDeleted,proto3" json:"include_deleted,omitempty"`
}

func (x *NamespaceFilter) Reset() {
	*x = NamespaceFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceFilter) ProtoMessage() {}

func (x *NamespaceFilter) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceFilter.ProtoReflect.Descriptor instead.
func (*NamespaceFilter) Descriptor() ([]byte, []int) {
	return file_temporal_api_namespace_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *NamespaceFilter) GetIncludeDeleted() bool {
	if x != nil {
		return x.IncludeDeleted
	}
	return false
}

// Namespace capability details. Should contain what features are enabled in a namespace.
type NamespaceInfo_Capabilities struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// True if the namespace supports eager workflow start.
	EagerWorkflowStart bool `protobuf:"varint,1,opt,name=eager_workflow_start,json=eagerWorkflowStart,proto3" json:"eager_workflow_start,omitempty"`
	// True if the namespace supports sync update
	SyncUpdate bool `protobuf:"varint,2,opt,name=sync_update,json=syncUpdate,proto3" json:"sync_update,omitempty"`
	// True if the namespace supports async update
	AsyncUpdate bool `protobuf:"varint,3,opt,name=async_update,json=asyncUpdate,proto3" json:"async_update,omitempty"`
}

func (x *NamespaceInfo_Capabilities) Reset() {
	*x = NamespaceInfo_Capabilities{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NamespaceInfo_Capabilities) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceInfo_Capabilities) ProtoMessage() {}

func (x *NamespaceInfo_Capabilities) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_namespace_v1_message_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceInfo_Capabilities.ProtoReflect.Descriptor instead.
func (*NamespaceInfo_Capabilities) Descriptor() ([]byte, []int) {
	return file_temporal_api_namespace_v1_message_proto_rawDescGZIP(), []int{0, 1}
}

func (x *NamespaceInfo_Capabilities) GetEagerWorkflowStart() bool {
	if x != nil {
		return x.EagerWorkflowStart
	}
	return false
}

func (x *NamespaceInfo_Capabilities) GetSyncUpdate() bool {
	if x != nil {
		return x.SyncUpdate
	}
	return false
}

func (x *NamespaceInfo_Capabilities) GetAsyncUpdate() bool {
	if x != nil {
		return x.AsyncUpdate
	}
	return false
}

var File_temporal_api_namespace_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_namespace_v1_message_proto_rawDesc = []byte{
	0x0a, 0x27, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x04, 0x0a,
	0x0d, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x46, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x32, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x59, 0x0a, 0x0c, 0x63, 0x61,
	0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x61, 0x70, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x0c, 0x63, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x0a, 0x12, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x18, 0x64, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x84, 0x01,
	0x0a, 0x0c, 0x43, 0x61, 0x70, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x30,
	0x0a, 0x14, 0x65, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x65, 0x61,
	0x67, 0x65, 0x72, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x61, 0x73, 0x79, 0x6e, 0x63, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x22, 0xcf, 0x05, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x62, 0x0a, 0x20, 0x77, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x74, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x1d, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x74, 0x6c, 0x12, 0x49, 0x0a, 0x0c,
	0x62, 0x61, 0x64, 0x5f, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42,
	0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x52, 0x0b, 0x62, 0x61, 0x64, 0x42,
	0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x5a, 0x0a, 0x16, 0x68, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x14, 0x68,
	0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x61,
	0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x75, 0x72, 0x69, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x12, 0x68, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x61, 0x6c, 0x55, 0x72, 0x69, 0x12, 0x60, 0x0a, 0x19, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c,
	0x69, 0x74, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x73, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x17,
	0x76, 0x69, 0x73, 0x69, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a, 0x17, 0x76, 0x69, 0x73, 0x69, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x5f, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x5f, 0x75,
	0x72, 0x69, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x15, 0x76, 0x69, 0x73, 0x69, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x41, 0x72, 0x63, 0x68, 0x69, 0x76, 0x61, 0x6c, 0x55, 0x72, 0x69, 0x12,
	0x93, 0x01, 0x0a, 0x1f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x5f, 0x73, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x5f, 0x61, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x5f, 0x61, 0x6c, 0x69, 0x61,
	0x73, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x4c, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x41, 0x6c, 0x69, 0x61, 0x73,
	0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x1c, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x65, 0x73, 0x1a, 0x4f, 0x0a, 0x21, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x41, 0x6c,
	0x69, 0x61, 0x73, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc6, 0x01, 0x0a, 0x0b, 0x42, 0x61, 0x64, 0x42, 0x69,
	0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x12, 0x50, 0x0a, 0x08, 0x62, 0x69, 0x6e, 0x61, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73,
	0x2e, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08,
	0x62, 0x69, 0x6e, 0x61, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x65, 0x0a, 0x0d, 0x42, 0x69, 0x6e, 0x61,
	0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3e, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x80, 0x01, 0x0a, 0x0d, 0x42, 0x61, 0x64, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6f, 0x70, 0x65,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x22, 0x9c, 0x02, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x4c, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x25, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x1a, 0x37, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x3a, 0x0a, 0x0f, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x12, 0x27, 0x0a, 0x0f, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x42, 0x98, 0x01,
	0x0a, 0x1c, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x29,
	0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0xaa, 0x02, 0x1b, 0x54, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x4e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_namespace_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_namespace_v1_message_proto_rawDescData = file_temporal_api_namespace_v1_message_proto_rawDesc
)

func file_temporal_api_namespace_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_namespace_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_namespace_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_namespace_v1_message_proto_rawDescData)
	})
	return file_temporal_api_namespace_v1_message_proto_rawDescData
}

var file_temporal_api_namespace_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_temporal_api_namespace_v1_message_proto_goTypes = []interface{}{
	(*NamespaceInfo)(nil),              // 0: temporal.api.namespace.v1.NamespaceInfo
	(*NamespaceConfig)(nil),            // 1: temporal.api.namespace.v1.NamespaceConfig
	(*BadBinaries)(nil),                // 2: temporal.api.namespace.v1.BadBinaries
	(*BadBinaryInfo)(nil),              // 3: temporal.api.namespace.v1.BadBinaryInfo
	(*UpdateNamespaceInfo)(nil),        // 4: temporal.api.namespace.v1.UpdateNamespaceInfo
	(*NamespaceFilter)(nil),            // 5: temporal.api.namespace.v1.NamespaceFilter
	nil,                                // 6: temporal.api.namespace.v1.NamespaceInfo.DataEntry
	(*NamespaceInfo_Capabilities)(nil), // 7: temporal.api.namespace.v1.NamespaceInfo.Capabilities
	nil,                                // 8: temporal.api.namespace.v1.NamespaceConfig.CustomSearchAttributeAliasesEntry
	nil,                                // 9: temporal.api.namespace.v1.BadBinaries.BinariesEntry
	nil,                                // 10: temporal.api.namespace.v1.UpdateNamespaceInfo.DataEntry
	(v1.NamespaceState)(0),             // 11: temporal.api.enums.v1.NamespaceState
	(*durationpb.Duration)(nil),        // 12: google.protobuf.Duration
	(v1.ArchivalState)(0),              // 13: temporal.api.enums.v1.ArchivalState
	(*timestamppb.Timestamp)(nil),      // 14: google.protobuf.Timestamp
}
var file_temporal_api_namespace_v1_message_proto_depIdxs = []int32{
	11, // 0: temporal.api.namespace.v1.NamespaceInfo.state:type_name -> temporal.api.enums.v1.NamespaceState
	6,  // 1: temporal.api.namespace.v1.NamespaceInfo.data:type_name -> temporal.api.namespace.v1.NamespaceInfo.DataEntry
	7,  // 2: temporal.api.namespace.v1.NamespaceInfo.capabilities:type_name -> temporal.api.namespace.v1.NamespaceInfo.Capabilities
	12, // 3: temporal.api.namespace.v1.NamespaceConfig.workflow_execution_retention_ttl:type_name -> google.protobuf.Duration
	2,  // 4: temporal.api.namespace.v1.NamespaceConfig.bad_binaries:type_name -> temporal.api.namespace.v1.BadBinaries
	13, // 5: temporal.api.namespace.v1.NamespaceConfig.history_archival_state:type_name -> temporal.api.enums.v1.ArchivalState
	13, // 6: temporal.api.namespace.v1.NamespaceConfig.visibility_archival_state:type_name -> temporal.api.enums.v1.ArchivalState
	8,  // 7: temporal.api.namespace.v1.NamespaceConfig.custom_search_attribute_aliases:type_name -> temporal.api.namespace.v1.NamespaceConfig.CustomSearchAttributeAliasesEntry
	9,  // 8: temporal.api.namespace.v1.BadBinaries.binaries:type_name -> temporal.api.namespace.v1.BadBinaries.BinariesEntry
	14, // 9: temporal.api.namespace.v1.BadBinaryInfo.create_time:type_name -> google.protobuf.Timestamp
	10, // 10: temporal.api.namespace.v1.UpdateNamespaceInfo.data:type_name -> temporal.api.namespace.v1.UpdateNamespaceInfo.DataEntry
	11, // 11: temporal.api.namespace.v1.UpdateNamespaceInfo.state:type_name -> temporal.api.enums.v1.NamespaceState
	3,  // 12: temporal.api.namespace.v1.BadBinaries.BinariesEntry.value:type_name -> temporal.api.namespace.v1.BadBinaryInfo
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_temporal_api_namespace_v1_message_proto_init() }
func file_temporal_api_namespace_v1_message_proto_init() {
	if File_temporal_api_namespace_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_namespace_v1_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceInfo); i {
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
		file_temporal_api_namespace_v1_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceConfig); i {
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
		file_temporal_api_namespace_v1_message_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadBinaries); i {
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
		file_temporal_api_namespace_v1_message_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BadBinaryInfo); i {
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
		file_temporal_api_namespace_v1_message_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateNamespaceInfo); i {
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
		file_temporal_api_namespace_v1_message_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceFilter); i {
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
		file_temporal_api_namespace_v1_message_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NamespaceInfo_Capabilities); i {
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
			RawDescriptor: file_temporal_api_namespace_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_namespace_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_namespace_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_namespace_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_namespace_v1_message_proto = out.File
	file_temporal_api_namespace_v1_message_proto_rawDesc = nil
	file_temporal_api_namespace_v1_message_proto_goTypes = nil
	file_temporal_api_namespace_v1_message_proto_depIdxs = nil
}
