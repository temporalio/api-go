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
// source: temporal/api/batch/v1/message.proto

package batch

import (
	reflect "reflect"
	sync "sync"

	v11 "go.temporal.io/api/common/v1"
	v1 "go.temporal.io/api/enums/v1"
	v12 "go.temporal.io/api/workflow/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	fieldmaskpb "google.golang.org/protobuf/types/known/fieldmaskpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BatchOperationInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Batch job ID
	JobId string `protobuf:"bytes,1,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	// Batch operation state
	State v1.BatchOperationState `protobuf:"varint,2,opt,name=state,proto3,enum=temporal.api.enums.v1.BatchOperationState" json:"state,omitempty"`
	// Batch operation start time
	StartTime *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	// Batch operation close time
	CloseTime     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=close_time,json=closeTime,proto3" json:"close_time,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchOperationInfo) Reset() {
	*x = BatchOperationInfo{}
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchOperationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationInfo) ProtoMessage() {}

func (x *BatchOperationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationInfo.ProtoReflect.Descriptor instead.
func (*BatchOperationInfo) Descriptor() ([]byte, []int) {
	return file_temporal_api_batch_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *BatchOperationInfo) GetJobId() string {
	if x != nil {
		return x.JobId
	}
	return ""
}

func (x *BatchOperationInfo) GetState() v1.BatchOperationState {
	if x != nil {
		return x.State
	}
	return v1.BatchOperationState(0)
}

func (x *BatchOperationInfo) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *BatchOperationInfo) GetCloseTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CloseTime
	}
	return nil
}

// BatchOperationTermination sends terminate requests to batch workflows.
// Keep the parameter in sync with temporal.api.workflowservice.v1.TerminateWorkflowExecutionRequest.
// Ignore first_execution_run_id because this is used for single workflow operation.
type BatchOperationTermination struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Serialized value(s) to provide to the termination event
	Details *v11.Payloads `protobuf:"bytes,1,opt,name=details,proto3" json:"details,omitempty"`
	// The identity of the worker/client
	Identity      string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchOperationTermination) Reset() {
	*x = BatchOperationTermination{}
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchOperationTermination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationTermination) ProtoMessage() {}

func (x *BatchOperationTermination) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationTermination.ProtoReflect.Descriptor instead.
func (*BatchOperationTermination) Descriptor() ([]byte, []int) {
	return file_temporal_api_batch_v1_message_proto_rawDescGZIP(), []int{1}
}

func (x *BatchOperationTermination) GetDetails() *v11.Payloads {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *BatchOperationTermination) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

// BatchOperationSignal sends signals to batch workflows.
// Keep the parameter in sync with temporal.api.workflowservice.v1.SignalWorkflowExecutionRequest.
type BatchOperationSignal struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The workflow author-defined name of the signal to send to the workflow
	Signal string `protobuf:"bytes,1,opt,name=signal,proto3" json:"signal,omitempty"`
	// Serialized value(s) to provide with the signal
	Input *v11.Payloads `protobuf:"bytes,2,opt,name=input,proto3" json:"input,omitempty"`
	// Headers that are passed with the signal to the processing workflow.
	// These can include things like auth or tracing tokens.
	Header *v11.Header `protobuf:"bytes,3,opt,name=header,proto3" json:"header,omitempty"`
	// The identity of the worker/client
	Identity      string `protobuf:"bytes,4,opt,name=identity,proto3" json:"identity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchOperationSignal) Reset() {
	*x = BatchOperationSignal{}
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchOperationSignal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationSignal) ProtoMessage() {}

func (x *BatchOperationSignal) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationSignal.ProtoReflect.Descriptor instead.
func (*BatchOperationSignal) Descriptor() ([]byte, []int) {
	return file_temporal_api_batch_v1_message_proto_rawDescGZIP(), []int{2}
}

func (x *BatchOperationSignal) GetSignal() string {
	if x != nil {
		return x.Signal
	}
	return ""
}

func (x *BatchOperationSignal) GetInput() *v11.Payloads {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *BatchOperationSignal) GetHeader() *v11.Header {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *BatchOperationSignal) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

// BatchOperationCancellation sends cancel requests to batch workflows.
// Keep the parameter in sync with temporal.api.workflowservice.v1.RequestCancelWorkflowExecutionRequest.
// Ignore first_execution_run_id because this is used for single workflow operation.
type BatchOperationCancellation struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The identity of the worker/client
	Identity      string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchOperationCancellation) Reset() {
	*x = BatchOperationCancellation{}
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchOperationCancellation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationCancellation) ProtoMessage() {}

func (x *BatchOperationCancellation) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationCancellation.ProtoReflect.Descriptor instead.
func (*BatchOperationCancellation) Descriptor() ([]byte, []int) {
	return file_temporal_api_batch_v1_message_proto_rawDescGZIP(), []int{3}
}

func (x *BatchOperationCancellation) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

// BatchOperationDeletion sends deletion requests to batch workflows.
// Keep the parameter in sync with temporal.api.workflowservice.v1.DeleteWorkflowExecutionRequest.
type BatchOperationDeletion struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The identity of the worker/client
	Identity      string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchOperationDeletion) Reset() {
	*x = BatchOperationDeletion{}
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchOperationDeletion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationDeletion) ProtoMessage() {}

func (x *BatchOperationDeletion) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationDeletion.ProtoReflect.Descriptor instead.
func (*BatchOperationDeletion) Descriptor() ([]byte, []int) {
	return file_temporal_api_batch_v1_message_proto_rawDescGZIP(), []int{4}
}

func (x *BatchOperationDeletion) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

// BatchOperationReset sends reset requests to batch workflows.
// Keep the parameter in sync with temporal.api.workflowservice.v1.ResetWorkflowExecutionRequest.
type BatchOperationReset struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The identity of the worker/client.
	Identity string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	// Describes what to reset to and how. If set, `reset_type` and `reset_reapply_type` are ignored.
	Options *v11.ResetOptions `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
	// Reset type (deprecated, use `options`).
	ResetType v1.ResetType `protobuf:"varint,1,opt,name=reset_type,json=resetType,proto3,enum=temporal.api.enums.v1.ResetType" json:"reset_type,omitempty"`
	// History event reapply options (deprecated, use `options`).
	ResetReapplyType v1.ResetReapplyType `protobuf:"varint,2,opt,name=reset_reapply_type,json=resetReapplyType,proto3,enum=temporal.api.enums.v1.ResetReapplyType" json:"reset_reapply_type,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *BatchOperationReset) Reset() {
	*x = BatchOperationReset{}
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchOperationReset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationReset) ProtoMessage() {}

func (x *BatchOperationReset) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationReset.ProtoReflect.Descriptor instead.
func (*BatchOperationReset) Descriptor() ([]byte, []int) {
	return file_temporal_api_batch_v1_message_proto_rawDescGZIP(), []int{5}
}

func (x *BatchOperationReset) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *BatchOperationReset) GetOptions() *v11.ResetOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *BatchOperationReset) GetResetType() v1.ResetType {
	if x != nil {
		return x.ResetType
	}
	return v1.ResetType(0)
}

func (x *BatchOperationReset) GetResetReapplyType() v1.ResetReapplyType {
	if x != nil {
		return x.ResetReapplyType
	}
	return v1.ResetReapplyType(0)
}

// BatchOperationUpdateWorkflowExecutionOptions sends UpdateWorkflowExecutionOptions requests to batch workflows.
// Keep the parameters in sync with temporal.api.workflowservice.v1.UpdateWorkflowExecutionOptionsRequest.
type BatchOperationUpdateWorkflowExecutionOptions struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// The identity of the worker/client.
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	// Workflow Execution options. Partial updates are accepted and controlled by update_mask.
	WorkflowExecutionOptions *v12.WorkflowExecutionOptions `protobuf:"bytes,2,opt,name=workflow_execution_options,json=workflowExecutionOptions,proto3" json:"workflow_execution_options,omitempty"`
	// Controls which fields from `workflow_execution_options` will be applied.
	// To unset a field, set it to null and use the update mask to indicate that it should be mutated.
	UpdateMask    *fieldmaskpb.FieldMask `protobuf:"bytes,3,opt,name=update_mask,json=updateMask,proto3" json:"update_mask,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *BatchOperationUpdateWorkflowExecutionOptions) Reset() {
	*x = BatchOperationUpdateWorkflowExecutionOptions{}
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BatchOperationUpdateWorkflowExecutionOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchOperationUpdateWorkflowExecutionOptions) ProtoMessage() {}

func (x *BatchOperationUpdateWorkflowExecutionOptions) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_batch_v1_message_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BatchOperationUpdateWorkflowExecutionOptions.ProtoReflect.Descriptor instead.
func (*BatchOperationUpdateWorkflowExecutionOptions) Descriptor() ([]byte, []int) {
	return file_temporal_api_batch_v1_message_proto_rawDescGZIP(), []int{6}
}

func (x *BatchOperationUpdateWorkflowExecutionOptions) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *BatchOperationUpdateWorkflowExecutionOptions) GetWorkflowExecutionOptions() *v12.WorkflowExecutionOptions {
	if x != nil {
		return x.WorkflowExecutionOptions
	}
	return nil
}

func (x *BatchOperationUpdateWorkflowExecutionOptions) GetUpdateMask() *fieldmaskpb.FieldMask {
	if x != nil {
		return x.UpdateMask
	}
	return nil
}

var File_temporal_api_batch_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_batch_v1_message_proto_rawDesc = []byte{
	0x0a, 0x23, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2b, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x74,
	0x63, 0x68, 0x5f, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x21, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe3, 0x01,
	0x0a, 0x12, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x0a, 0x06, 0x6a, 0x6f, 0x62, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6a, 0x6f, 0x62, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x39, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x6c, 0x6f, 0x73,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0x73, 0x0a, 0x19, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x3a, 0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x73, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xba, 0x01, 0x0a, 0x14, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x6c, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x12, 0x36, 0x0a, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75,
	0x74, 0x12, 0x36, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x38, 0x0a, 0x1a, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x34, 0x0a, 0x16, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x89, 0x02, 0x0a, 0x13, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f,
	0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x12, 0x1a, 0x0a,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3e, 0x0a, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3f, 0x0a, 0x0a, 0x72, 0x65, 0x73,
	0x65, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75,
	0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x09, 0x72, 0x65, 0x73, 0x65, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x55, 0x0a, 0x12, 0x72, 0x65,
	0x73, 0x65, 0x74, 0x5f, 0x72, 0x65, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70, 0x65, 0x52,
	0x10, 0x72, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x61, 0x70, 0x70, 0x6c, 0x79, 0x54, 0x79, 0x70,
	0x65, 0x22, 0xf9, 0x01, 0x0a, 0x2c, 0x42, 0x61, 0x74, 0x63, 0x68, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x70,
	0x0a, 0x1a, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x32, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x18, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x12, 0x3b, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6d, 0x61, 0x73, 0x6b, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4d, 0x61, 0x73,
	0x6b, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x73, 0x6b, 0x42, 0x84, 0x01,
	0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x62,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x61, 0x74, 0x63, 0x68, 0xaa, 0x02, 0x17,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x42,
	0x61, 0x74, 0x63, 0x68, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_batch_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_batch_v1_message_proto_rawDescData = file_temporal_api_batch_v1_message_proto_rawDesc
)

func file_temporal_api_batch_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_batch_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_batch_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_batch_v1_message_proto_rawDescData)
	})
	return file_temporal_api_batch_v1_message_proto_rawDescData
}

var file_temporal_api_batch_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_temporal_api_batch_v1_message_proto_goTypes = []any{
	(*BatchOperationInfo)(nil),                           // 0: temporal.api.batch.v1.BatchOperationInfo
	(*BatchOperationTermination)(nil),                    // 1: temporal.api.batch.v1.BatchOperationTermination
	(*BatchOperationSignal)(nil),                         // 2: temporal.api.batch.v1.BatchOperationSignal
	(*BatchOperationCancellation)(nil),                   // 3: temporal.api.batch.v1.BatchOperationCancellation
	(*BatchOperationDeletion)(nil),                       // 4: temporal.api.batch.v1.BatchOperationDeletion
	(*BatchOperationReset)(nil),                          // 5: temporal.api.batch.v1.BatchOperationReset
	(*BatchOperationUpdateWorkflowExecutionOptions)(nil), // 6: temporal.api.batch.v1.BatchOperationUpdateWorkflowExecutionOptions
	(v1.BatchOperationState)(0),                          // 7: temporal.api.enums.v1.BatchOperationState
	(*timestamppb.Timestamp)(nil),                        // 8: google.protobuf.Timestamp
	(*v11.Payloads)(nil),                                 // 9: temporal.api.common.v1.Payloads
	(*v11.Header)(nil),                                   // 10: temporal.api.common.v1.Header
	(*v11.ResetOptions)(nil),                             // 11: temporal.api.common.v1.ResetOptions
	(v1.ResetType)(0),                                    // 12: temporal.api.enums.v1.ResetType
	(v1.ResetReapplyType)(0),                             // 13: temporal.api.enums.v1.ResetReapplyType
	(*v12.WorkflowExecutionOptions)(nil),                 // 14: temporal.api.workflow.v1.WorkflowExecutionOptions
	(*fieldmaskpb.FieldMask)(nil),                        // 15: google.protobuf.FieldMask
}
var file_temporal_api_batch_v1_message_proto_depIdxs = []int32{
	7,  // 0: temporal.api.batch.v1.BatchOperationInfo.state:type_name -> temporal.api.enums.v1.BatchOperationState
	8,  // 1: temporal.api.batch.v1.BatchOperationInfo.start_time:type_name -> google.protobuf.Timestamp
	8,  // 2: temporal.api.batch.v1.BatchOperationInfo.close_time:type_name -> google.protobuf.Timestamp
	9,  // 3: temporal.api.batch.v1.BatchOperationTermination.details:type_name -> temporal.api.common.v1.Payloads
	9,  // 4: temporal.api.batch.v1.BatchOperationSignal.input:type_name -> temporal.api.common.v1.Payloads
	10, // 5: temporal.api.batch.v1.BatchOperationSignal.header:type_name -> temporal.api.common.v1.Header
	11, // 6: temporal.api.batch.v1.BatchOperationReset.options:type_name -> temporal.api.common.v1.ResetOptions
	12, // 7: temporal.api.batch.v1.BatchOperationReset.reset_type:type_name -> temporal.api.enums.v1.ResetType
	13, // 8: temporal.api.batch.v1.BatchOperationReset.reset_reapply_type:type_name -> temporal.api.enums.v1.ResetReapplyType
	14, // 9: temporal.api.batch.v1.BatchOperationUpdateWorkflowExecutionOptions.workflow_execution_options:type_name -> temporal.api.workflow.v1.WorkflowExecutionOptions
	15, // 10: temporal.api.batch.v1.BatchOperationUpdateWorkflowExecutionOptions.update_mask:type_name -> google.protobuf.FieldMask
	11, // [11:11] is the sub-list for method output_type
	11, // [11:11] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_temporal_api_batch_v1_message_proto_init() }
func file_temporal_api_batch_v1_message_proto_init() {
	if File_temporal_api_batch_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_batch_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_batch_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_batch_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_batch_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_batch_v1_message_proto = out.File
	file_temporal_api_batch_v1_message_proto_rawDesc = nil
	file_temporal_api_batch_v1_message_proto_goTypes = nil
	file_temporal_api_batch_v1_message_proto_depIdxs = nil
}
