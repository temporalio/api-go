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
// source: temporal/api/enums/v1/task_queue.proto

package enums

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type TaskQueueKind int32

const (
	TaskQueueKind_TASK_QUEUE_KIND_UNSPECIFIED TaskQueueKind = 0
	// Tasks from a normal workflow task queue always include complete workflow history
	//
	// The task queue specified by the user is always a normal task queue. There can be as many
	// workers as desired for a single normal task queue. All those workers may pick up tasks from
	// that queue.
	TaskQueueKind_TASK_QUEUE_KIND_NORMAL TaskQueueKind = 1
	// A sticky queue only includes new history since the last workflow task, and they are
	// per-worker.
	//
	// Sticky queues are created dynamically by each worker during their start up. They only exist
	// for the lifetime of the worker process. Tasks in a sticky task queue are only available to
	// the worker that created the sticky queue.
	//
	// Sticky queues are only for workflow tasks. There are no sticky task queues for activities.
	TaskQueueKind_TASK_QUEUE_KIND_STICKY TaskQueueKind = 2
)

// Enum value maps for TaskQueueKind.
var (
	TaskQueueKind_name = map[int32]string{
		0: "TASK_QUEUE_KIND_UNSPECIFIED",
		1: "TASK_QUEUE_KIND_NORMAL",
		2: "TASK_QUEUE_KIND_STICKY",
	}
	TaskQueueKind_value = map[string]int32{
		"TASK_QUEUE_KIND_UNSPECIFIED": 0,
		"TASK_QUEUE_KIND_NORMAL":      1,
		"TASK_QUEUE_KIND_STICKY":      2,
	}
)

func (x TaskQueueKind) Enum() *TaskQueueKind {
	p := new(TaskQueueKind)
	*p = x
	return p
}

func (x TaskQueueKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskQueueKind) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[0].Descriptor()
}

func (TaskQueueKind) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[0]
}

func (x TaskQueueKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskQueueKind.Descriptor instead.
func (TaskQueueKind) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{0}
}

type TaskQueueType int32

const (
	TaskQueueType_TASK_QUEUE_TYPE_UNSPECIFIED TaskQueueType = 0
	// Workflow type of task queue.
	TaskQueueType_TASK_QUEUE_TYPE_WORKFLOW TaskQueueType = 1
	// Activity type of task queue.
	TaskQueueType_TASK_QUEUE_TYPE_ACTIVITY TaskQueueType = 2
	// Task queue type for dispatching Nexus requests.
	TaskQueueType_TASK_QUEUE_TYPE_NEXUS TaskQueueType = 3
)

// Enum value maps for TaskQueueType.
var (
	TaskQueueType_name = map[int32]string{
		0: "TASK_QUEUE_TYPE_UNSPECIFIED",
		1: "TASK_QUEUE_TYPE_WORKFLOW",
		2: "TASK_QUEUE_TYPE_ACTIVITY",
		3: "TASK_QUEUE_TYPE_NEXUS",
	}
	TaskQueueType_value = map[string]int32{
		"TASK_QUEUE_TYPE_UNSPECIFIED": 0,
		"TASK_QUEUE_TYPE_WORKFLOW":    1,
		"TASK_QUEUE_TYPE_ACTIVITY":    2,
		"TASK_QUEUE_TYPE_NEXUS":       3,
	}
)

func (x TaskQueueType) Enum() *TaskQueueType {
	p := new(TaskQueueType)
	*p = x
	return p
}

func (x TaskQueueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskQueueType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[1].Descriptor()
}

func (TaskQueueType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[1]
}

func (x TaskQueueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskQueueType.Descriptor instead.
func (TaskQueueType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{1}
}

// Specifies which category of tasks may reach a worker on a versioned task queue.
// Used both in a reachability query and its response.
// Deprecated.
type TaskReachability int32

const (
	TaskReachability_TASK_REACHABILITY_UNSPECIFIED TaskReachability = 0
	// There's a possiblity for a worker to receive new workflow tasks. Workers should *not* be retired.
	TaskReachability_TASK_REACHABILITY_NEW_WORKFLOWS TaskReachability = 1
	// There's a possiblity for a worker to receive existing workflow and activity tasks from existing workflows. Workers
	// should *not* be retired.
	// This enum value does not distinguish between open and closed workflows.
	TaskReachability_TASK_REACHABILITY_EXISTING_WORKFLOWS TaskReachability = 2
	// There's a possiblity for a worker to receive existing workflow and activity tasks from open workflows. Workers
	// should *not* be retired.
	TaskReachability_TASK_REACHABILITY_OPEN_WORKFLOWS TaskReachability = 3
	// There's a possiblity for a worker to receive existing workflow tasks from closed workflows. Workers may be
	// retired dependending on application requirements. For example, if there's no need to query closed workflows.
	TaskReachability_TASK_REACHABILITY_CLOSED_WORKFLOWS TaskReachability = 4
)

// Enum value maps for TaskReachability.
var (
	TaskReachability_name = map[int32]string{
		0: "TASK_REACHABILITY_UNSPECIFIED",
		1: "TASK_REACHABILITY_NEW_WORKFLOWS",
		2: "TASK_REACHABILITY_EXISTING_WORKFLOWS",
		3: "TASK_REACHABILITY_OPEN_WORKFLOWS",
		4: "TASK_REACHABILITY_CLOSED_WORKFLOWS",
	}
	TaskReachability_value = map[string]int32{
		"TASK_REACHABILITY_UNSPECIFIED":        0,
		"TASK_REACHABILITY_NEW_WORKFLOWS":      1,
		"TASK_REACHABILITY_EXISTING_WORKFLOWS": 2,
		"TASK_REACHABILITY_OPEN_WORKFLOWS":     3,
		"TASK_REACHABILITY_CLOSED_WORKFLOWS":   4,
	}
)

func (x TaskReachability) Enum() *TaskReachability {
	p := new(TaskReachability)
	*p = x
	return p
}

func (x TaskReachability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskReachability) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[2].Descriptor()
}

func (TaskReachability) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[2]
}

func (x TaskReachability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskReachability.Descriptor instead.
func (TaskReachability) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{2}
}

// Specifies which category of tasks may reach a versioned worker of a certain Build ID.
//
// Task Reachability is eventually consistent; there may be a delay (up to few minutes) until it
// converges to the most accurate value but it is designed in a way to take the more conservative
// side until it converges. For example REACHABLE is more conservative than CLOSED_WORKFLOWS_ONLY.
//
// Note: future activities who inherit their workflow's Build ID but not its Task Queue will not be
// accounted for reachability as server cannot know if they'll happen as they do not use
// assignment rules of their Task Queue. Same goes for Child Workflows or Continue-As-New Workflows
// who inherit the parent/previous workflow's Build ID but not its Task Queue. In those cases, make
// sure to query reachability for the parent/previous workflow's Task Queue as well.
type BuildIdTaskReachability int32

const (
	// Task reachability is not reported
	BuildIdTaskReachability_BUILD_ID_TASK_REACHABILITY_UNSPECIFIED BuildIdTaskReachability = 0
	// Build ID may be used by new workflows or activities (base on versioning rules), or there MAY
	// be open workflows or backlogged activities assigned to it.
	BuildIdTaskReachability_BUILD_ID_TASK_REACHABILITY_REACHABLE BuildIdTaskReachability = 1
	// Build ID does not have open workflows and is not reachable by new workflows,
	// but MAY have closed workflows within the namespace retention period.
	// Not applicable to activity-only task queues.
	BuildIdTaskReachability_BUILD_ID_TASK_REACHABILITY_CLOSED_WORKFLOWS_ONLY BuildIdTaskReachability = 2
	// Build ID is not used for new executions, nor it has been used by any existing execution
	// within the retention period.
	BuildIdTaskReachability_BUILD_ID_TASK_REACHABILITY_UNREACHABLE BuildIdTaskReachability = 3
)

// Enum value maps for BuildIdTaskReachability.
var (
	BuildIdTaskReachability_name = map[int32]string{
		0: "BUILD_ID_TASK_REACHABILITY_UNSPECIFIED",
		1: "BUILD_ID_TASK_REACHABILITY_REACHABLE",
		2: "BUILD_ID_TASK_REACHABILITY_CLOSED_WORKFLOWS_ONLY",
		3: "BUILD_ID_TASK_REACHABILITY_UNREACHABLE",
	}
	BuildIdTaskReachability_value = map[string]int32{
		"BUILD_ID_TASK_REACHABILITY_UNSPECIFIED":           0,
		"BUILD_ID_TASK_REACHABILITY_REACHABLE":             1,
		"BUILD_ID_TASK_REACHABILITY_CLOSED_WORKFLOWS_ONLY": 2,
		"BUILD_ID_TASK_REACHABILITY_UNREACHABLE":           3,
	}
)

func (x BuildIdTaskReachability) Enum() *BuildIdTaskReachability {
	p := new(BuildIdTaskReachability)
	*p = x
	return p
}

func (x BuildIdTaskReachability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BuildIdTaskReachability) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[3].Descriptor()
}

func (BuildIdTaskReachability) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[3]
}

func (x BuildIdTaskReachability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BuildIdTaskReachability.Descriptor instead.
func (BuildIdTaskReachability) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{3}
}

type DescribeTaskQueueMode int32

const (
	// Unspecified means legacy behavior.
	DescribeTaskQueueMode_DESCRIBE_TASK_QUEUE_MODE_UNSPECIFIED DescribeTaskQueueMode = 0
	// Enhanced mode reports aggregated results for all partitions, supports Build IDs, and reports richer info.
	DescribeTaskQueueMode_DESCRIBE_TASK_QUEUE_MODE_ENHANCED DescribeTaskQueueMode = 1
)

// Enum value maps for DescribeTaskQueueMode.
var (
	DescribeTaskQueueMode_name = map[int32]string{
		0: "DESCRIBE_TASK_QUEUE_MODE_UNSPECIFIED",
		1: "DESCRIBE_TASK_QUEUE_MODE_ENHANCED",
	}
	DescribeTaskQueueMode_value = map[string]int32{
		"DESCRIBE_TASK_QUEUE_MODE_UNSPECIFIED": 0,
		"DESCRIBE_TASK_QUEUE_MODE_ENHANCED":    1,
	}
)

func (x DescribeTaskQueueMode) Enum() *DescribeTaskQueueMode {
	p := new(DescribeTaskQueueMode)
	*p = x
	return p
}

func (x DescribeTaskQueueMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DescribeTaskQueueMode) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[4].Descriptor()
}

func (DescribeTaskQueueMode) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[4]
}

func (x DescribeTaskQueueMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DescribeTaskQueueMode.Descriptor instead.
func (DescribeTaskQueueMode) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{4}
}

var File_temporal_api_enums_v1_task_queue_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_task_queue_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0x68, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x53, 0x54, 0x49, 0x43, 0x4b, 0x59, 0x10, 0x02, 0x2a, 0x87, 0x01, 0x0a, 0x0d, 0x54, 0x61,
	0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18,
	0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43,
	0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15, 0x54, 0x41, 0x53, 0x4b,
	0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x45, 0x58, 0x55,
	0x53, 0x10, 0x03, 0x2a, 0xd2, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x61, 0x63,
	0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x1d, 0x54, 0x41, 0x53, 0x4b,
	0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59,
	0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x10, 0x01,
	0x12, 0x28, 0x0a, 0x24, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x49, 0x4e, 0x47, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f,
	0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x10, 0x03,
	0x12, 0x26, 0x0a, 0x22, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x5f, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x10, 0x04, 0x2a, 0xd1, 0x01, 0x0a, 0x17, 0x42, 0x75, 0x69,
	0x6c, 0x64, 0x49, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x69,
	0x6c, 0x69, 0x74, 0x79, 0x12, 0x2a, 0x0a, 0x26, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x49, 0x44,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49,
	0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x28, 0x0a, 0x24, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x52,
	0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x34, 0x0a, 0x30, 0x42, 0x55,
	0x49, 0x4c, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43,
	0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x5f,
	0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02,
	0x12, 0x2a, 0x0a, 0x26, 0x42, 0x55, 0x49, 0x4c, 0x44, 0x5f, 0x49, 0x44, 0x5f, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55,
	0x4e, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x2a, 0x68, 0x0a, 0x15,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75,
	0x65, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x28, 0x0a, 0x24, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42,
	0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x44,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x25, 0x0a, 0x21, 0x44, 0x45, 0x53, 0x43, 0x52, 0x49, 0x42, 0x45, 0x5f, 0x54, 0x41, 0x53, 0x4b,
	0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x45, 0x4e, 0x48, 0x41,
	0x4e, 0x43, 0x45, 0x44, 0x10, 0x01, 0x42, 0x86, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_temporal_api_enums_v1_task_queue_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_task_queue_proto_rawDescData []byte
)

func file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_task_queue_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_task_queue_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_task_queue_proto_rawDesc), len(file_temporal_api_enums_v1_task_queue_proto_rawDesc)))
	})
	return file_temporal_api_enums_v1_task_queue_proto_rawDescData
}

var file_temporal_api_enums_v1_task_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_temporal_api_enums_v1_task_queue_proto_goTypes = []any{
	(TaskQueueKind)(0),           // 0: temporal.api.enums.v1.TaskQueueKind
	(TaskQueueType)(0),           // 1: temporal.api.enums.v1.TaskQueueType
	(TaskReachability)(0),        // 2: temporal.api.enums.v1.TaskReachability
	(BuildIdTaskReachability)(0), // 3: temporal.api.enums.v1.BuildIdTaskReachability
	(DescribeTaskQueueMode)(0),   // 4: temporal.api.enums.v1.DescribeTaskQueueMode
}
var file_temporal_api_enums_v1_task_queue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_task_queue_proto_init() }
func file_temporal_api_enums_v1_task_queue_proto_init() {
	if File_temporal_api_enums_v1_task_queue_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_task_queue_proto_rawDesc), len(file_temporal_api_enums_v1_task_queue_proto_rawDesc)),
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_task_queue_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_task_queue_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_task_queue_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_task_queue_proto = out.File
	file_temporal_api_enums_v1_task_queue_proto_goTypes = nil
	file_temporal_api_enums_v1_task_queue_proto_depIdxs = nil
}
