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
// source: temporal/api/deployment/v1/message.proto

package deployment

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/common/v1"
	v11 "go.temporal.io/api/enums/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `Deployment` identifies a deployment of Temporal workers. The combination of deployment series
// name + build ID serves as the identifier. User can use `WorkerDeploymentOptions` in their worker
// programs to specify these values.
// [cleanup-wv-pre-release]
type Deployment struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Different versions of the same worker service/application are related together by having a
	// shared series name.
	// Out of all deployments of a series, one can be designated as the current deployment, which
	// receives new workflow executions and new tasks of workflows with
	// `VERSIONING_BEHAVIOR_AUTO_UPGRADE` versioning behavior.
	SeriesName string `protobuf:"bytes,1,opt,name=series_name,json=seriesName,proto3" json:"series_name,omitempty"`
	// Build ID changes with each version of the worker when the worker program code and/or config
	// changes.
	BuildId       string `protobuf:"bytes,2,opt,name=build_id,json=buildId,proto3" json:"build_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetSeriesName() string {
	if x != nil {
		return x.SeriesName
	}
	return ""
}

func (x *Deployment) GetBuildId() string {
	if x != nil {
		return x.BuildId
	}
	return ""
}

type WorkerDeploymentVersion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Uniquely identifies the worker deployment version.
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// The name of the deployment this version belongs too.
	DeploymentName string `protobuf:"bytes,2,opt,name=deployment_name,json=deploymentName,proto3" json:"deployment_name,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WorkerDeploymentVersion) Reset() {
	*x = WorkerDeploymentVersion{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerDeploymentVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeploymentVersion) ProtoMessage() {}

func (x *WorkerDeploymentVersion) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeploymentVersion.ProtoReflect.Descriptor instead.
func (*WorkerDeploymentVersion) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *WorkerDeploymentVersion) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *WorkerDeploymentVersion) GetDeploymentName() string {
	if x != nil {
		return x.DeploymentName
	}
	return ""
}

// `DeploymentInfo` holds information about a deployment. Deployment information is tracked
// automatically by server as soon as the first poll from that deployment reaches the server. There
// can be multiple task queue workers in a single deployment which are listed in this message.
// [cleanup-wv-pre-release]
type DeploymentInfo struct {
	state          protoimpl.MessageState          `protogen:"open.v1"`
	Deployment     *Deployment                     `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	CreateTime     *timestamppb.Timestamp          `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	TaskQueueInfos []*DeploymentInfo_TaskQueueInfo `protobuf:"bytes,3,rep,name=task_queue_infos,json=taskQueueInfos,proto3" json:"task_queue_infos,omitempty"`
	// A user-defined set of key-values. Can be updated as part of write operations to the
	// deployment, such as `SetCurrentDeployment`.
	Metadata map[string]*v1.Payload `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// If this deployment is the current deployment of its deployment series.
	IsCurrent     bool `protobuf:"varint,5,opt,name=is_current,json=isCurrent,proto3" json:"is_current,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploymentInfo) Reset() {
	*x = DeploymentInfo{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentInfo) ProtoMessage() {}

func (x *DeploymentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentInfo.ProtoReflect.Descriptor instead.
func (*DeploymentInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *DeploymentInfo) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *DeploymentInfo) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *DeploymentInfo) GetTaskQueueInfos() []*DeploymentInfo_TaskQueueInfo {
	if x != nil {
		return x.TaskQueueInfos
	}
	return nil
}

func (x *DeploymentInfo) GetMetadata() map[string]*v1.Payload {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DeploymentInfo) GetIsCurrent() bool {
	if x != nil {
		return x.IsCurrent
	}
	return false
}

// Used as part of Deployment write APIs to update metadata attached to a deployment.
// [cleanup-wv-pre-release]
type UpdateDeploymentMetadata struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UpsertEntries map[string]*v1.Payload `protobuf:"bytes,1,rep,name=upsert_entries,json=upsertEntries,proto3" json:"upsert_entries,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// List of keys to remove from the metadata.
	RemoveEntries []string `protobuf:"bytes,2,rep,name=remove_entries,json=removeEntries,proto3" json:"remove_entries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateDeploymentMetadata) Reset() {
	*x = UpdateDeploymentMetadata{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateDeploymentMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateDeploymentMetadata) ProtoMessage() {}

func (x *UpdateDeploymentMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateDeploymentMetadata.ProtoReflect.Descriptor instead.
func (*UpdateDeploymentMetadata) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateDeploymentMetadata) GetUpsertEntries() map[string]*v1.Payload {
	if x != nil {
		return x.UpsertEntries
	}
	return nil
}

func (x *UpdateDeploymentMetadata) GetRemoveEntries() []string {
	if x != nil {
		return x.RemoveEntries
	}
	return nil
}

// DeploymentListInfo is an abbreviated set of fields from DeploymentInfo that's returned in
// ListDeployments.
// [cleanup-wv-pre-release]
type DeploymentListInfo struct {
	state      protoimpl.MessageState `protogen:"open.v1"`
	Deployment *Deployment            `protobuf:"bytes,1,opt,name=deployment,proto3" json:"deployment,omitempty"`
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// If this deployment is the current deployment of its deployment series.
	IsCurrent     bool `protobuf:"varint,3,opt,name=is_current,json=isCurrent,proto3" json:"is_current,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeploymentListInfo) Reset() {
	*x = DeploymentListInfo{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentListInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentListInfo) ProtoMessage() {}

func (x *DeploymentListInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentListInfo.ProtoReflect.Descriptor instead.
func (*DeploymentListInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *DeploymentListInfo) GetDeployment() *Deployment {
	if x != nil {
		return x.Deployment
	}
	return nil
}

func (x *DeploymentListInfo) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *DeploymentListInfo) GetIsCurrent() bool {
	if x != nil {
		return x.IsCurrent
	}
	return false
}

// A Worker Deployment Version (Version, for short) represents all workers of the same
// code and config within a Deployment. Workers of the same Version are expected to
// behave exactly the same so when executions move between them there are no
// non-determinism issues.
// Worker Deployment Versions are created in Temporal server automatically when
// their first poller arrives to the server.
// Each Version has a Workflow Versioning Mode which is chosen by the app
// developer. (see WorkflowVersioningMode enum documentation)
type WorkerDeploymentVersionInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Identifies a Worker Deployment Version. Must be unique within the namespace.
	// Same ID cannot be used in multiple Deployments.
	// In some contexts, such as Reset-by-build-id, version might be used
	// as "Build ID".
	Version string `protobuf:"bytes,1,opt,name=version,proto3" json:"version,omitempty"`
	// Each Worker Version belongs to exactly one Deployment.
	DeploymentName string                 `protobuf:"bytes,2,opt,name=deployment_name,json=deploymentName,proto3" json:"deployment_name,omitempty"`
	CreateTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// All the Task Queues that have ever polled from this Deployment version.
	TaskQueueInfos []*WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo `protobuf:"bytes,4,rep,name=task_queue_infos,json=taskQueueInfos,proto3" json:"task_queue_infos,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *WorkerDeploymentVersionInfo) Reset() {
	*x = WorkerDeploymentVersionInfo{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerDeploymentVersionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeploymentVersionInfo) ProtoMessage() {}

func (x *WorkerDeploymentVersionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeploymentVersionInfo.ProtoReflect.Descriptor instead.
func (*WorkerDeploymentVersionInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *WorkerDeploymentVersionInfo) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *WorkerDeploymentVersionInfo) GetDeploymentName() string {
	if x != nil {
		return x.DeploymentName
	}
	return ""
}

func (x *WorkerDeploymentVersionInfo) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *WorkerDeploymentVersionInfo) GetTaskQueueInfos() []*WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo {
	if x != nil {
		return x.TaskQueueInfos
	}
	return nil
}

type DeploymentInfo_TaskQueueInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	Name  string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type  v11.TaskQueueType      `protobuf:"varint,2,opt,name=type,proto3,enum=temporal.api.enums.v1.TaskQueueType" json:"type,omitempty"`
	// When server saw the first poller for this task queue in this deployment.
	FirstPollerTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=first_poller_time,json=firstPollerTime,proto3" json:"first_poller_time,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *DeploymentInfo_TaskQueueInfo) Reset() {
	*x = DeploymentInfo_TaskQueueInfo{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeploymentInfo_TaskQueueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeploymentInfo_TaskQueueInfo) ProtoMessage() {}

func (x *DeploymentInfo_TaskQueueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeploymentInfo_TaskQueueInfo.ProtoReflect.Descriptor instead.
func (*DeploymentInfo_TaskQueueInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{2, 1}
}

func (x *DeploymentInfo_TaskQueueInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *DeploymentInfo_TaskQueueInfo) GetType() v11.TaskQueueType {
	if x != nil {
		return x.Type
	}
	return v11.TaskQueueType(0)
}

func (x *DeploymentInfo_TaskQueueInfo) GetFirstPollerTime() *timestamppb.Timestamp {
	if x != nil {
		return x.FirstPollerTime
	}
	return nil
}

type WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type          v11.TaskQueueType      `protobuf:"varint,2,opt,name=type,proto3,enum=temporal.api.enums.v1.TaskQueueType" json:"type,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo) Reset() {
	*x = WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo{}
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo) ProtoMessage() {}

func (x *WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_deployment_v1_message_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo.ProtoReflect.Descriptor instead.
func (*WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_deployment_v1_message_proto_rawDescGZIP(), []int{5, 0}
}

func (x *WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo) GetType() v11.TaskQueueType {
	if x != nil {
		return x.Type
	}
	return v11.TaskQueueType(0)
}

var File_temporal_api_deployment_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_deployment_v1_message_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0a, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65,
	0x72, 0x69, 0x65, 0x73, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x1d, 0x0a, 0x08, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x22, 0x64, 0x0a, 0x17, 0x57, 0x6f, 0x72,
	0x6b, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12,
	0x2b, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x22, 0x9c, 0x05, 0x0a, 0x0e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x0a,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x66, 0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f,
	0x69, 0x6e, 0x66, 0x6f, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f,
	0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x58, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x08, 0x6d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x1a, 0x64, 0x0a, 0x0d, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42, 0x02, 0x68, 0x00,
	0x12, 0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xb1, 0x01, 0x0a,
	0x0d, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x42, 0x02, 0x68, 0x00, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4a,
	0x0a, 0x11, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x70, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x52, 0x0f, 0x66, 0x69, 0x72, 0x73, 0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x22, 0xa4, 0x02, 0x0a, 0x18, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x72, 0x0a, 0x0e, 0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x5f, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x47, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0d,
	0x75, 0x70, 0x73, 0x65, 0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x02, 0x68, 0x00,
	0x12, 0x29, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x69, 0x65, 0x73, 0x42, 0x02, 0x68, 0x00, 0x1a, 0x69, 0x0a, 0x12, 0x55, 0x70, 0x73,
	0x65, 0x72, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x42,
	0x02, 0x68, 0x00, 0x12, 0x39, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc4,
	0x01, 0x0a, 0x12, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4a, 0x0a, 0x0a, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x0a, 0x64,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3f, 0x0a, 0x0b,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x69, 0x73,
	0x5f, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69,
	0x73, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x02, 0x68, 0x00, 0x22, 0x9b, 0x03, 0x0a, 0x1b,
	0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c, 0x0a, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2b, 0x0a, 0x0f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e,
	0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x3f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x7e,
	0x0a, 0x10, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x50, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x65, 0x72, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x65, 0x72, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x74, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x73, 0x42, 0x02, 0x68, 0x00, 0x1a, 0x70, 0x0a, 0x18, 0x57, 0x6f, 0x72, 0x6b, 0x65,
	0x72, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3c, 0x0a, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65,
	0x42, 0x02, 0x68, 0x00, 0x42, 0x9d, 0x01, 0x0a, 0x1d, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2b, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0xaa, 0x02,
	0x1c, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1f, 0x54, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x44,
	0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_deployment_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_deployment_v1_message_proto_rawDescData = file_temporal_api_deployment_v1_message_proto_rawDesc
)

func file_temporal_api_deployment_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_deployment_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_deployment_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_deployment_v1_message_proto_rawDescData)
	})
	return file_temporal_api_deployment_v1_message_proto_rawDescData
}

var file_temporal_api_deployment_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_temporal_api_deployment_v1_message_proto_goTypes = []any{
	(*Deployment)(nil),                   // 0: temporal.api.deployment.v1.Deployment
	(*WorkerDeploymentVersion)(nil),      // 1: temporal.api.deployment.v1.WorkerDeploymentVersion
	(*DeploymentInfo)(nil),               // 2: temporal.api.deployment.v1.DeploymentInfo
	(*UpdateDeploymentMetadata)(nil),     // 3: temporal.api.deployment.v1.UpdateDeploymentMetadata
	(*DeploymentListInfo)(nil),           // 4: temporal.api.deployment.v1.DeploymentListInfo
	(*WorkerDeploymentVersionInfo)(nil),  // 5: temporal.api.deployment.v1.WorkerDeploymentVersionInfo
	nil,                                  // 6: temporal.api.deployment.v1.DeploymentInfo.MetadataEntry
	(*DeploymentInfo_TaskQueueInfo)(nil), // 7: temporal.api.deployment.v1.DeploymentInfo.TaskQueueInfo
	nil,                                  // 8: temporal.api.deployment.v1.UpdateDeploymentMetadata.UpsertEntriesEntry
	(*WorkerDeploymentVersionInfo_WorkerBuildTaskQueueInfo)(nil), // 9: temporal.api.deployment.v1.WorkerDeploymentVersionInfo.WorkerBuildTaskQueueInfo
	(*timestamppb.Timestamp)(nil),                                // 10: google.protobuf.Timestamp
	(*v1.Payload)(nil),                                           // 11: temporal.api.common.v1.Payload
	(v11.TaskQueueType)(0),                                       // 12: temporal.api.enums.v1.TaskQueueType
}
var file_temporal_api_deployment_v1_message_proto_depIdxs = []int32{
	0,  // 0: temporal.api.deployment.v1.DeploymentInfo.deployment:type_name -> temporal.api.deployment.v1.Deployment
	10, // 1: temporal.api.deployment.v1.DeploymentInfo.create_time:type_name -> google.protobuf.Timestamp
	7,  // 2: temporal.api.deployment.v1.DeploymentInfo.task_queue_infos:type_name -> temporal.api.deployment.v1.DeploymentInfo.TaskQueueInfo
	6,  // 3: temporal.api.deployment.v1.DeploymentInfo.metadata:type_name -> temporal.api.deployment.v1.DeploymentInfo.MetadataEntry
	8,  // 4: temporal.api.deployment.v1.UpdateDeploymentMetadata.upsert_entries:type_name -> temporal.api.deployment.v1.UpdateDeploymentMetadata.UpsertEntriesEntry
	0,  // 5: temporal.api.deployment.v1.DeploymentListInfo.deployment:type_name -> temporal.api.deployment.v1.Deployment
	10, // 6: temporal.api.deployment.v1.DeploymentListInfo.create_time:type_name -> google.protobuf.Timestamp
	10, // 7: temporal.api.deployment.v1.WorkerDeploymentVersionInfo.create_time:type_name -> google.protobuf.Timestamp
	9,  // 8: temporal.api.deployment.v1.WorkerDeploymentVersionInfo.task_queue_infos:type_name -> temporal.api.deployment.v1.WorkerDeploymentVersionInfo.WorkerBuildTaskQueueInfo
	11, // 9: temporal.api.deployment.v1.DeploymentInfo.MetadataEntry.value:type_name -> temporal.api.common.v1.Payload
	12, // 10: temporal.api.deployment.v1.DeploymentInfo.TaskQueueInfo.type:type_name -> temporal.api.enums.v1.TaskQueueType
	10, // 11: temporal.api.deployment.v1.DeploymentInfo.TaskQueueInfo.first_poller_time:type_name -> google.protobuf.Timestamp
	11, // 12: temporal.api.deployment.v1.UpdateDeploymentMetadata.UpsertEntriesEntry.value:type_name -> temporal.api.common.v1.Payload
	12, // 13: temporal.api.deployment.v1.WorkerDeploymentVersionInfo.WorkerBuildTaskQueueInfo.type:type_name -> temporal.api.enums.v1.TaskQueueType
	14, // [14:14] is the sub-list for method output_type
	14, // [14:14] is the sub-list for method input_type
	14, // [14:14] is the sub-list for extension type_name
	14, // [14:14] is the sub-list for extension extendee
	0,  // [0:14] is the sub-list for field type_name
}

func init() { file_temporal_api_deployment_v1_message_proto_init() }
func file_temporal_api_deployment_v1_message_proto_init() {
	if File_temporal_api_deployment_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_deployment_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_deployment_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_deployment_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_deployment_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_deployment_v1_message_proto = out.File
	file_temporal_api_deployment_v1_message_proto_rawDesc = nil
	file_temporal_api_deployment_v1_message_proto_goTypes = nil
	file_temporal_api_deployment_v1_message_proto_depIdxs = nil
}
