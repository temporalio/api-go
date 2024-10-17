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
// source: temporal/api/activity/v1/message.proto

package activity

import (
	reflect "reflect"
	sync "sync"

	v11 "go.temporal.io/api/common/v1"
	v1 "go.temporal.io/api/taskqueue/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ActivityOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskQueue *v1.TaskQueue `protobuf:"bytes,1,opt,name=task_queue,json=taskQueue,proto3" json:"task_queue,omitempty"`
	// Indicates how long the caller is willing to wait for an activity completion. Limits how long
	// retries will be attempted. Either this or `start_to_close_timeout` must be specified.
	//
	// (-- api-linter: core::0140::prepositions=disabled
	//
	//	aip.dev/not-precedent: "to" is used to indicate interval. --)
	ScheduleToCloseTimeout *durationpb.Duration `protobuf:"bytes,2,opt,name=schedule_to_close_timeout,json=scheduleToCloseTimeout,proto3" json:"schedule_to_close_timeout,omitempty"`
	// Limits time an activity task can stay in a task queue before a worker picks it up. This
	// timeout is always non retryable, as all a retry would achieve is to put it back into the same
	// queue. Defaults to `schedule_to_close_timeout` or workflow execution timeout if not
	// specified.
	//
	// (-- api-linter: core::0140::prepositions=disabled
	//
	//	aip.dev/not-precedent: "to" is used to indicate interval. --)
	ScheduleToStartTimeout *durationpb.Duration `protobuf:"bytes,3,opt,name=schedule_to_start_timeout,json=scheduleToStartTimeout,proto3" json:"schedule_to_start_timeout,omitempty"`
	// Maximum time an activity is allowed to execute after being picked up by a worker. This
	// timeout is always retryable. Either this or `schedule_to_close_timeout` must be
	// specified.
	//
	// (-- api-linter: core::0140::prepositions=disabled
	//
	//	aip.dev/not-precedent: "to" is used to indicate interval. --)
	StartToCloseTimeout *durationpb.Duration `protobuf:"bytes,4,opt,name=start_to_close_timeout,json=startToCloseTimeout,proto3" json:"start_to_close_timeout,omitempty"`
	// Maximum permitted time between successful worker heartbeats.
	HeartbeatTimeout *durationpb.Duration `protobuf:"bytes,5,opt,name=heartbeat_timeout,json=heartbeatTimeout,proto3" json:"heartbeat_timeout,omitempty"`
	RetryPolicy      *v11.RetryPolicy     `protobuf:"bytes,6,opt,name=retry_policy,json=retryPolicy,proto3" json:"retry_policy,omitempty"`
}

func (x *ActivityOptions) Reset() {
	*x = ActivityOptions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_temporal_api_activity_v1_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityOptions) ProtoMessage() {}

func (x *ActivityOptions) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_activity_v1_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityOptions.ProtoReflect.Descriptor instead.
func (*ActivityOptions) Descriptor() ([]byte, []int) {
	return file_temporal_api_activity_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *ActivityOptions) GetTaskQueue() *v1.TaskQueue {
	if x != nil {
		return x.TaskQueue
	}
	return nil
}

func (x *ActivityOptions) GetScheduleToCloseTimeout() *durationpb.Duration {
	if x != nil {
		return x.ScheduleToCloseTimeout
	}
	return nil
}

func (x *ActivityOptions) GetScheduleToStartTimeout() *durationpb.Duration {
	if x != nil {
		return x.ScheduleToStartTimeout
	}
	return nil
}

func (x *ActivityOptions) GetStartToCloseTimeout() *durationpb.Duration {
	if x != nil {
		return x.StartToCloseTimeout
	}
	return nil
}

func (x *ActivityOptions) GetHeartbeatTimeout() *durationpb.Duration {
	if x != nil {
		return x.HeartbeatTimeout
	}
	return nil
}

func (x *ActivityOptions) GetRetryPolicy() *v11.RetryPolicy {
	if x != nil {
		return x.RetryPolicy
	}
	return nil
}

var File_temporal_api_activity_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_activity_v1_message_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31,
	0x1a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x76, 0x31, 0x2f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xfa, 0x03, 0x0a, 0x0f,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12,
	0x47, 0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x24, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61,
	0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65,
	0x75, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x58, 0x0a, 0x19, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x16, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x54, 0x6f, 0x43, 0x6c,
	0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x58, 0x0a,
	0x19, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x74, 0x61,
	0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x16, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x54, 0x6f, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x52, 0x0a, 0x16, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x6f,
	0x5f, 0x63, 0x6c, 0x6f, 0x73, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x6f, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x6f,
	0x75, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4a, 0x0a, 0x11, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x10, 0x68, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x02, 0x68, 0x00,
	0x12, 0x4a, 0x0a, 0x0c, 0x72, 0x65, 0x74, 0x72, 0x79, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x74, 0x72, 0x79,
	0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0x02, 0x68, 0x00, 0x42, 0x93, 0x01, 0x0a, 0x1b, 0x69,
	0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x27, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0xaa, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1d,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_activity_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_activity_v1_message_proto_rawDescData = file_temporal_api_activity_v1_message_proto_rawDesc
)

func file_temporal_api_activity_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_activity_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_activity_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_activity_v1_message_proto_rawDescData)
	})
	return file_temporal_api_activity_v1_message_proto_rawDescData
}

var file_temporal_api_activity_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_api_activity_v1_message_proto_goTypes = []any{
	(*ActivityOptions)(nil),     // 0: temporal.api.activity.v1.ActivityOptions
	(*v1.TaskQueue)(nil),        // 1: temporal.api.taskqueue.v1.TaskQueue
	(*durationpb.Duration)(nil), // 2: google.protobuf.Duration
	(*v11.RetryPolicy)(nil),     // 3: temporal.api.common.v1.RetryPolicy
}
var file_temporal_api_activity_v1_message_proto_depIdxs = []int32{
	1, // 0: temporal.api.activity.v1.ActivityOptions.task_queue:type_name -> temporal.api.taskqueue.v1.TaskQueue
	2, // 1: temporal.api.activity.v1.ActivityOptions.schedule_to_close_timeout:type_name -> google.protobuf.Duration
	2, // 2: temporal.api.activity.v1.ActivityOptions.schedule_to_start_timeout:type_name -> google.protobuf.Duration
	2, // 3: temporal.api.activity.v1.ActivityOptions.start_to_close_timeout:type_name -> google.protobuf.Duration
	2, // 4: temporal.api.activity.v1.ActivityOptions.heartbeat_timeout:type_name -> google.protobuf.Duration
	3, // 5: temporal.api.activity.v1.ActivityOptions.retry_policy:type_name -> temporal.api.common.v1.RetryPolicy
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_temporal_api_activity_v1_message_proto_init() }
func file_temporal_api_activity_v1_message_proto_init() {
	if File_temporal_api_activity_v1_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_temporal_api_activity_v1_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*ActivityOptions); i {
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
			RawDescriptor: file_temporal_api_activity_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_activity_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_activity_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_activity_v1_message_proto_msgTypes,
	}.Build()
	File_temporal_api_activity_v1_message_proto = out.File
	file_temporal_api_activity_v1_message_proto_rawDesc = nil
	file_temporal_api_activity_v1_message_proto_goTypes = nil
	file_temporal_api_activity_v1_message_proto_depIdxs = nil
}
