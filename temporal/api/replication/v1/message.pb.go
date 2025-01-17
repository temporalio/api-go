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
// 	protoc-gen-go v1.36.4
// 	protoc        v5.26.1
// source: temporal/api/replication/v1/message.proto

package replication

import (
	v1 "go.temporal.io/api/enums/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ClusterReplicationConfig struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClusterName   string                 `protobuf:"bytes,1,opt,name=cluster_name,json=clusterName,proto3" json:"cluster_name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClusterReplicationConfig) Reset() {
	*x = ClusterReplicationConfig{}
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterReplicationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterReplicationConfig) ProtoMessage() {}

func (x *ClusterReplicationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterReplicationConfig.ProtoReflect.Descriptor instead.
func (*ClusterReplicationConfig) Descriptor() ([]byte, []int) {
	return file_temporal_api_replication_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *ClusterReplicationConfig) GetClusterName() string {
	if x != nil {
		return x.ClusterName
	}
	return ""
}

type NamespaceReplicationConfig struct {
	state             protoimpl.MessageState      `protogen:"open.v1"`
	ActiveClusterName string                      `protobuf:"bytes,1,opt,name=active_cluster_name,json=activeClusterName,proto3" json:"active_cluster_name,omitempty"`
	Clusters          []*ClusterReplicationConfig `protobuf:"bytes,2,rep,name=clusters,proto3" json:"clusters,omitempty"`
	State             v1.ReplicationState         `protobuf:"varint,3,opt,name=state,proto3,enum=temporal.api.enums.v1.ReplicationState" json:"state,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *NamespaceReplicationConfig) Reset() {
	*x = NamespaceReplicationConfig{}
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NamespaceReplicationConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NamespaceReplicationConfig) ProtoMessage() {}

func (x *NamespaceReplicationConfig) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NamespaceReplicationConfig.ProtoReflect.Descriptor instead.
func (*NamespaceReplicationConfig) Descriptor() ([]byte, []int) {
	return file_temporal_api_replication_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *NamespaceReplicationConfig) GetActiveClusterName() string {
	if x != nil {
		return x.ActiveClusterName
	}
	return ""
}

func (x *NamespaceReplicationConfig) GetClusters() []*ClusterReplicationConfig {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *NamespaceReplicationConfig) GetState() v1.ReplicationState {
	if x != nil {
		return x.State
	}
	return v1.ReplicationState(0)
}

// Represents a historical replication status of a Namespace
type FailoverStatus struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Timestamp when the Cluster switched to the following failover_version
	FailoverTime    *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=failover_time,json=failoverTime,proto3" json:"failover_time,omitempty"`
	FailoverVersion int64                  `protobuf:"varint,2,opt,name=failover_version,json=failoverVersion,proto3" json:"failover_version,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *FailoverStatus) Reset() {
	*x = FailoverStatus{}
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FailoverStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailoverStatus) ProtoMessage() {}

func (x *FailoverStatus) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_replication_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailoverStatus.ProtoReflect.Descriptor instead.
func (*FailoverStatus) Descriptor() ([]byte, []int) {
	return file_temporal_api_replication_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *FailoverStatus) GetFailoverTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FailoverTime
	}
	return nil
}

func (x *FailoverStatus) GetFailoverVersion() int64 {
	if x != nil {
		return x.FailoverVersion
	}
	return 0
}

var File_temporal_api_replication_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_replication_v1_message_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72,
	0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x25, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x3d, 0x0a, 0x18, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22,
	0xde, 0x01, 0x0a, 0x1a, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x70,
	0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x2e,
	0x0a, 0x13, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x51,
	0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x35, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x27, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x7c, 0x0a, 0x0e, 0x46, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0d, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x5f,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x66,
	0x61, 0x69, 0x6c, 0x6f, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0xa2,
	0x01, 0x0a, 0x1e, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x2d, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0xaa, 0x02, 0x1d, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x56, 0x31,
	0xea, 0x02, 0x20, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a,
	0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_temporal_api_replication_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_replication_v1_message_proto_rawDescData []byte
)

func file_temporal_api_replication_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_replication_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_replication_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_replication_v1_message_proto_rawDesc), len(file_temporal_api_replication_v1_message_proto_rawDesc)))
	})
	return file_temporal_api_replication_v1_message_proto_rawDescData
}

var file_temporal_api_replication_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_temporal_api_replication_v1_message_proto_goTypes = []any{
	(*ClusterReplicationConfig)(nil),   // 0: temporal.api.replication.v1.ClusterReplicationConfig
	(*NamespaceReplicationConfig)(nil), // 1: temporal.api.replication.v1.NamespaceReplicationConfig
	(*FailoverStatus)(nil),             // 2: temporal.api.replication.v1.FailoverStatus
	(v1.ReplicationState)(0),           // 3: temporal.api.enums.v1.ReplicationState
	(*timestamppb.Timestamp)(nil),      // 4: google.protobuf.Timestamp
}
var file_temporal_api_replication_v1_message_proto_depIdxs = []int32{
	0, // 0: temporal.api.replication.v1.NamespaceReplicationConfig.clusters:type_name -> temporal.api.replication.v1.ClusterReplicationConfig
	3, // 1: temporal.api.replication.v1.NamespaceReplicationConfig.state:type_name -> temporal.api.enums.v1.ReplicationState
	4, // 2: temporal.api.replication.v1.FailoverStatus.failover_time:type_name -> google.protobuf.Timestamp
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_temporal_api_replication_v1_message_proto_init() }
func file_temporal_api_replication_v1_message_proto_init() {
	if File_temporal_api_replication_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_replication_v1_message_proto_rawDesc), len(file_temporal_api_replication_v1_message_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_replication_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_replication_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_replication_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_replication_v1_message_proto = out.File
	file_temporal_api_replication_v1_message_proto_goTypes = nil
	file_temporal_api_replication_v1_message_proto_depIdxs = nil
}
