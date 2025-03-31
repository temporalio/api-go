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
// source: temporal/api/enums/v1/task_queue.proto

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

type TaskQueueKind int32

const (
	TASK_QUEUE_KIND_UNSPECIFIED TaskQueueKind = 0
	// Tasks from a normal workflow task queue always include complete workflow history
	//
	// The task queue specified by the user is always a normal task queue. There can be as many
	// workers as desired for a single normal task queue. All those workers may pick up tasks from
	// that queue.
	TASK_QUEUE_KIND_NORMAL TaskQueueKind = 1
	// A sticky queue only includes new history since the last workflow task, and they are
	// per-worker.
	//
	// Sticky queues are created dynamically by each worker during their start up. They only exist
	// for the lifetime of the worker process. Tasks in a sticky task queue are only available to
	// the worker that created the sticky queue.
	//
	// Sticky queues are only for workflow tasks. There are no sticky task queues for activities.
	TASK_QUEUE_KIND_STICKY TaskQueueKind = 2
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
	switch x {
	case TASK_QUEUE_KIND_UNSPECIFIED:
		return "Unspecified"
	case TASK_QUEUE_KIND_NORMAL:
		return "Normal"
	case TASK_QUEUE_KIND_STICKY:
		return "Sticky"
	default:
		return strconv.Itoa(int(x))
	}

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
	TASK_QUEUE_TYPE_UNSPECIFIED TaskQueueType = 0
	// Workflow type of task queue.
	TASK_QUEUE_TYPE_WORKFLOW TaskQueueType = 1
	// Activity type of task queue.
	TASK_QUEUE_TYPE_ACTIVITY TaskQueueType = 2
	// Task queue type for dispatching Nexus requests.
	TASK_QUEUE_TYPE_NEXUS TaskQueueType = 3
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
	switch x {
	case TASK_QUEUE_TYPE_UNSPECIFIED:
		return "Unspecified"
	case TASK_QUEUE_TYPE_WORKFLOW:
		return "Workflow"
	case TASK_QUEUE_TYPE_ACTIVITY:
		return "Activity"
	case TASK_QUEUE_TYPE_NEXUS:
		return "Nexus"
	default:
		return strconv.Itoa(int(x))
	}

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
	TASK_REACHABILITY_UNSPECIFIED TaskReachability = 0
	// There's a possiblity for a worker to receive new workflow tasks. Workers should *not* be retired.
	TASK_REACHABILITY_NEW_WORKFLOWS TaskReachability = 1
	// There's a possiblity for a worker to receive existing workflow and activity tasks from existing workflows. Workers
	// should *not* be retired.
	// This enum value does not distinguish between open and closed workflows.
	TASK_REACHABILITY_EXISTING_WORKFLOWS TaskReachability = 2
	// There's a possiblity for a worker to receive existing workflow and activity tasks from open workflows. Workers
	// should *not* be retired.
	TASK_REACHABILITY_OPEN_WORKFLOWS TaskReachability = 3
	// There's a possiblity for a worker to receive existing workflow tasks from closed workflows. Workers may be
	// retired dependending on application requirements. For example, if there's no need to query closed workflows.
	TASK_REACHABILITY_CLOSED_WORKFLOWS TaskReachability = 4
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
	switch x {
	case TASK_REACHABILITY_UNSPECIFIED:
		return "Unspecified"
	case TASK_REACHABILITY_NEW_WORKFLOWS:
		return "NewWorkflows"
	case TASK_REACHABILITY_EXISTING_WORKFLOWS:
		return "ExistingWorkflows"
	case TASK_REACHABILITY_OPEN_WORKFLOWS:
		return "OpenWorkflows"
	case TASK_REACHABILITY_CLOSED_WORKFLOWS:
		return "ClosedWorkflows"
	default:
		return strconv.Itoa(int(x))
	}

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
	BUILD_ID_TASK_REACHABILITY_UNSPECIFIED BuildIdTaskReachability = 0
	// Build ID may be used by new workflows or activities (base on versioning rules), or there MAY
	// be open workflows or backlogged activities assigned to it.
	BUILD_ID_TASK_REACHABILITY_REACHABLE BuildIdTaskReachability = 1
	// Build ID does not have open workflows and is not reachable by new workflows,
	// but MAY have closed workflows within the namespace retention period.
	// Not applicable to activity-only task queues.
	BUILD_ID_TASK_REACHABILITY_CLOSED_WORKFLOWS_ONLY BuildIdTaskReachability = 2
	// Build ID is not used for new executions, nor it has been used by any existing execution
	// within the retention period.
	BUILD_ID_TASK_REACHABILITY_UNREACHABLE BuildIdTaskReachability = 3
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
	switch x {
	case BUILD_ID_TASK_REACHABILITY_UNSPECIFIED:
		return "Unspecified"
	case BUILD_ID_TASK_REACHABILITY_REACHABLE:
		return "Reachable"
	case BUILD_ID_TASK_REACHABILITY_CLOSED_WORKFLOWS_ONLY:
		return "ClosedWorkflowsOnly"
	case BUILD_ID_TASK_REACHABILITY_UNREACHABLE:
		return "Unreachable"
	default:
		return strconv.Itoa(int(x))
	}

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
	DESCRIBE_TASK_QUEUE_MODE_UNSPECIFIED DescribeTaskQueueMode = 0
	// Enhanced mode reports aggregated results for all partitions, supports Build IDs, and reports richer info.
	DESCRIBE_TASK_QUEUE_MODE_ENHANCED DescribeTaskQueueMode = 1
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
	switch x {
	case DESCRIBE_TASK_QUEUE_MODE_UNSPECIFIED:
		return "Unspecified"
	case DESCRIBE_TASK_QUEUE_MODE_ENHANCED:
		return "Enhanced"
	default:
		return strconv.Itoa(int(x))
	}

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

const file_temporal_api_enums_v1_task_queue_proto_rawDesc = "" +
	"\n" +
	"&temporal/api/enums/v1/task_queue.proto\x12\x15temporal.api.enums.v1*h\n" +
	"\rTaskQueueKind\x12\x1f\n" +
	"\x1bTASK_QUEUE_KIND_UNSPECIFIED\x10\x00\x12\x1a\n" +
	"\x16TASK_QUEUE_KIND_NORMAL\x10\x01\x12\x1a\n" +
	"\x16TASK_QUEUE_KIND_STICKY\x10\x02*\x87\x01\n" +
	"\rTaskQueueType\x12\x1f\n" +
	"\x1bTASK_QUEUE_TYPE_UNSPECIFIED\x10\x00\x12\x1c\n" +
	"\x18TASK_QUEUE_TYPE_WORKFLOW\x10\x01\x12\x1c\n" +
	"\x18TASK_QUEUE_TYPE_ACTIVITY\x10\x02\x12\x19\n" +
	"\x15TASK_QUEUE_TYPE_NEXUS\x10\x03*\xd2\x01\n" +
	"\x10TaskReachability\x12!\n" +
	"\x1dTASK_REACHABILITY_UNSPECIFIED\x10\x00\x12#\n" +
	"\x1fTASK_REACHABILITY_NEW_WORKFLOWS\x10\x01\x12(\n" +
	"$TASK_REACHABILITY_EXISTING_WORKFLOWS\x10\x02\x12$\n" +
	" TASK_REACHABILITY_OPEN_WORKFLOWS\x10\x03\x12&\n" +
	"\"TASK_REACHABILITY_CLOSED_WORKFLOWS\x10\x04*\xd1\x01\n" +
	"\x17BuildIdTaskReachability\x12*\n" +
	"&BUILD_ID_TASK_REACHABILITY_UNSPECIFIED\x10\x00\x12(\n" +
	"$BUILD_ID_TASK_REACHABILITY_REACHABLE\x10\x01\x124\n" +
	"0BUILD_ID_TASK_REACHABILITY_CLOSED_WORKFLOWS_ONLY\x10\x02\x12*\n" +
	"&BUILD_ID_TASK_REACHABILITY_UNREACHABLE\x10\x03*h\n" +
	"\x15DescribeTaskQueueMode\x12(\n" +
	"$DESCRIBE_TASK_QUEUE_MODE_UNSPECIFIED\x10\x00\x12%\n" +
	"!DESCRIBE_TASK_QUEUE_MODE_ENHANCED\x10\x01B\x86\x01\n" +
	"\x18io.temporal.api.enums.v1B\x0eTaskQueueProtoP\x01Z!go.temporal.io/api/enums/v1;enums\xaa\x02\x17Temporalio.Api.Enums.V1\xea\x02\x1aTemporalio::Api::Enums::V1b\x06proto3"

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
