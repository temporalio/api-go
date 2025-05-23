// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/enums/v1/failed_cause.proto

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

// Workflow tasks can fail for various reasons. Note that some of these reasons can only originate
// from the server, and some of them can only originate from the SDK/worker.
type WorkflowTaskFailedCause int32

const (
	WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED WorkflowTaskFailedCause = 0
	// Between starting and completing the workflow task (with a workflow completion command), some
	// new command (like a signal) was processed into workflow history. The outstanding task will be
	// failed with this reason, and a worker must pick up a new task.
	WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND                                         WorkflowTaskFailedCause = 1
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES                          WorkflowTaskFailedCause = 2
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES                    WorkflowTaskFailedCause = 3
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES                                WorkflowTaskFailedCause = 4
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES                               WorkflowTaskFailedCause = 5
	WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES                              WorkflowTaskFailedCause = 6
	WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES                WorkflowTaskFailedCause = 7
	WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES                    WorkflowTaskFailedCause = 8
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES                  WorkflowTaskFailedCause = 9
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 10
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES                            WorkflowTaskFailedCause = 11
	WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID                                  WorkflowTaskFailedCause = 12
	// The worker wishes to fail the task and have the next one be generated on a normal, not sticky
	// queue. Generally workers should prefer to use the explicit `ResetStickyTaskQueue` RPC call.
	WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE                  WorkflowTaskFailedCause = 13
	WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE        WorkflowTaskFailedCause = 14
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 15
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES     WorkflowTaskFailedCause = 16
	WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND                      WorkflowTaskFailedCause = 17
	WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND                   WorkflowTaskFailedCause = 18
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE                    WorkflowTaskFailedCause = 19
	WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW                           WorkflowTaskFailedCause = 20
	WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY                               WorkflowTaskFailedCause = 21
	WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID           WorkflowTaskFailedCause = 22
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES                    WorkflowTaskFailedCause = 23
	// The worker encountered a mismatch while replaying history between what was expected, and
	// what the workflow code actually did.
	WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR                   WorkflowTaskFailedCause = 24
	WORKFLOW_TASK_FAILED_CAUSE_BAD_MODIFY_WORKFLOW_PROPERTIES_ATTRIBUTES WorkflowTaskFailedCause = 25
	// We send the below error codes to users when their requests would violate a size constraint
	// of their workflow. We do this to ensure that the state of their workflow does not become too
	// large because that can cause severe performance degradation. You can modify the thresholds for
	// each of these errors within your dynamic config.
	//
	// Spawning a new child workflow would cause this workflow to exceed its limit of pending child
	// workflows.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_CHILD_WORKFLOWS_LIMIT_EXCEEDED WorkflowTaskFailedCause = 26
	// Starting a new activity would cause this workflow to exceed its limit of pending activities
	// that we track.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_ACTIVITIES_LIMIT_EXCEEDED WorkflowTaskFailedCause = 27
	// A workflow has a buffer of signals that have not yet reached their destination. We return this
	// error when sending a new signal would exceed the capacity of this buffer.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_SIGNALS_LIMIT_EXCEEDED WorkflowTaskFailedCause = 28
	// Similarly, we have a buffer of pending requests to cancel other workflows. We return this error
	// when our capacity for pending cancel requests is already reached.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_REQUEST_CANCEL_LIMIT_EXCEEDED WorkflowTaskFailedCause = 29
	// Workflow execution update message (update.Acceptance, update.Rejection, or update.Response)
	// has wrong format, or missing required fields.
	WORKFLOW_TASK_FAILED_CAUSE_BAD_UPDATE_WORKFLOW_EXECUTION_MESSAGE WorkflowTaskFailedCause = 30
	// Similar to WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND, but for updates.
	WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_UPDATE WorkflowTaskFailedCause = 31
	// A workflow task completed with an invalid ScheduleNexusOperation command.
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_NEXUS_OPERATION_ATTRIBUTES WorkflowTaskFailedCause = 32
	// A workflow task completed requesting to schedule a Nexus Operation exceeding the server configured limit.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_NEXUS_OPERATIONS_LIMIT_EXCEEDED WorkflowTaskFailedCause = 33
	// A workflow task completed with an invalid RequestCancelNexusOperation command.
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_NEXUS_OPERATION_ATTRIBUTES WorkflowTaskFailedCause = 34
	// A workflow task completed requesting a feature that's disabled on the server (either system wide or - typically -
	// for the workflow's namespace).
	// Check the workflow task failure message for more information.
	WORKFLOW_TASK_FAILED_CAUSE_FEATURE_DISABLED WorkflowTaskFailedCause = 35
	// A workflow task failed because a grpc message was too large.
	WORKFLOW_TASK_FAILED_CAUSE_GRPC_MESSAGE_TOO_LARGE WorkflowTaskFailedCause = 36
)

// Enum value maps for WorkflowTaskFailedCause.
var (
	WorkflowTaskFailedCause_name = map[int32]string{
		0:  "WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED",
		1:  "WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND",
		2:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES",
		3:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES",
		4:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES",
		5:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES",
		6:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES",
		7:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES",
		8:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES",
		9:  "WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES",
		10: "WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES",
		11: "WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES",
		12: "WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID",
		13: "WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE",
		14: "WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE",
		15: "WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES",
		16: "WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES",
		17: "WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND",
		18: "WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND",
		19: "WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE",
		20: "WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW",
		21: "WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY",
		22: "WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID",
		23: "WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES",
		24: "WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR",
		25: "WORKFLOW_TASK_FAILED_CAUSE_BAD_MODIFY_WORKFLOW_PROPERTIES_ATTRIBUTES",
		26: "WORKFLOW_TASK_FAILED_CAUSE_PENDING_CHILD_WORKFLOWS_LIMIT_EXCEEDED",
		27: "WORKFLOW_TASK_FAILED_CAUSE_PENDING_ACTIVITIES_LIMIT_EXCEEDED",
		28: "WORKFLOW_TASK_FAILED_CAUSE_PENDING_SIGNALS_LIMIT_EXCEEDED",
		29: "WORKFLOW_TASK_FAILED_CAUSE_PENDING_REQUEST_CANCEL_LIMIT_EXCEEDED",
		30: "WORKFLOW_TASK_FAILED_CAUSE_BAD_UPDATE_WORKFLOW_EXECUTION_MESSAGE",
		31: "WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_UPDATE",
		32: "WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_NEXUS_OPERATION_ATTRIBUTES",
		33: "WORKFLOW_TASK_FAILED_CAUSE_PENDING_NEXUS_OPERATIONS_LIMIT_EXCEEDED",
		34: "WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_NEXUS_OPERATION_ATTRIBUTES",
		35: "WORKFLOW_TASK_FAILED_CAUSE_FEATURE_DISABLED",
		36: "WORKFLOW_TASK_FAILED_CAUSE_GRPC_MESSAGE_TOO_LARGE",
	}
	WorkflowTaskFailedCause_value = map[string]int32{
		"WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED":                                               0,
		"WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND":                                         1,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES":                          2,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES":                    3,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES":                                4,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES":                               5,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES":                              6,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES":                7,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES":                    8,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES":                  9,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES": 10,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES":                            11,
		"WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID":                                  12,
		"WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE":                                   13,
		"WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE":                         14,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES":                  15,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES":                      16,
		"WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND":                                       17,
		"WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND":                                    18,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE":                                     19,
		"WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW":                                            20,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY":                                                21,
		"WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID":                            22,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES":                                     23,
		"WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR":                                   24,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_MODIFY_WORKFLOW_PROPERTIES_ATTRIBUTES":                 25,
		"WORKFLOW_TASK_FAILED_CAUSE_PENDING_CHILD_WORKFLOWS_LIMIT_EXCEEDED":                    26,
		"WORKFLOW_TASK_FAILED_CAUSE_PENDING_ACTIVITIES_LIMIT_EXCEEDED":                         27,
		"WORKFLOW_TASK_FAILED_CAUSE_PENDING_SIGNALS_LIMIT_EXCEEDED":                            28,
		"WORKFLOW_TASK_FAILED_CAUSE_PENDING_REQUEST_CANCEL_LIMIT_EXCEEDED":                     29,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_UPDATE_WORKFLOW_EXECUTION_MESSAGE":                     30,
		"WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_UPDATE":                                          31,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_NEXUS_OPERATION_ATTRIBUTES":                   32,
		"WORKFLOW_TASK_FAILED_CAUSE_PENDING_NEXUS_OPERATIONS_LIMIT_EXCEEDED":                   33,
		"WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_NEXUS_OPERATION_ATTRIBUTES":             34,
		"WORKFLOW_TASK_FAILED_CAUSE_FEATURE_DISABLED":                                          35,
		"WORKFLOW_TASK_FAILED_CAUSE_GRPC_MESSAGE_TOO_LARGE":                                    36,
	}
)

func (x WorkflowTaskFailedCause) Enum() *WorkflowTaskFailedCause {
	p := new(WorkflowTaskFailedCause)
	*p = x
	return p
}

func (x WorkflowTaskFailedCause) String() string {
	switch x {
	case WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED:
		return "Unspecified"
	case WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND:
		return "UnhandledCommand"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES:
		return "BadScheduleActivityAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES:
		return "BadRequestCancelActivityAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES:
		return "BadStartTimerAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES:

		// Deprecated: Use WorkflowTaskFailedCause.Descriptor instead.
		return "BadCancelTimerAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES:
		return "BadRecordMarkerAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES:
		return "BadCompleteWorkflowExecutionAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES:
		return "BadFailWorkflowExecutionAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES:
		return "BadCancelWorkflowExecutionAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES:
		return "BadRequestCancelExternalWorkflowExecutionAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES:
		return "BadContinueAsNewAttributes"

		// Enum value maps for StartChildWorkflowExecutionFailedCause.
	case WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID:
		return "StartTimerDuplicateId"
	case WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE:
		return "ResetStickyTaskQueue"
	case WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE:
		return "WorkflowWorkerUnhandledFailure"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES:
		return "BadSignalWorkflowExecutionAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES:
		return "BadStartChildExecutionAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND:
		return "ForceCloseCommand"
	case WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND:
		return "FailoverCloseCommand"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE:
		return "BadSignalInputSize"
	case WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW:
		return "ResetWorkflow"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY:
		return "BadBinary"
	case WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID:
		return "ScheduleActivityDuplicateId"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES:
		return "BadSearchAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR:
		return "NonDeterministicError"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_MODIFY_WORKFLOW_PROPERTIES_ATTRIBUTES:
		return "BadModifyWorkflowPropertiesAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_PENDING_CHILD_WORKFLOWS_LIMIT_EXCEEDED:
		return "PendingChildWorkflowsLimitExceeded"
	case WORKFLOW_TASK_FAILED_CAUSE_PENDING_ACTIVITIES_LIMIT_EXCEEDED:

		// Deprecated: Use StartChildWorkflowExecutionFailedCause.Descriptor instead.
		return "PendingActivitiesLimitExceeded"
	case WORKFLOW_TASK_FAILED_CAUSE_PENDING_SIGNALS_LIMIT_EXCEEDED:
		return "PendingSignalsLimitExceeded"
	case WORKFLOW_TASK_FAILED_CAUSE_PENDING_REQUEST_CANCEL_LIMIT_EXCEEDED:
		return "PendingRequestCancelLimitExceeded"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_UPDATE_WORKFLOW_EXECUTION_MESSAGE:
		return "BadUpdateWorkflowExecutionMessage"
	case WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_UPDATE:
		return "UnhandledUpdate"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_NEXUS_OPERATION_ATTRIBUTES:
		return "BadScheduleNexusOperationAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_PENDING_NEXUS_OPERATIONS_LIMIT_EXCEEDED:
		return "PendingNexusOperationsLimitExceeded"
	case WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_NEXUS_OPERATION_ATTRIBUTES:
		return "BadRequestCancelNexusOperationAttributes"
	case WORKFLOW_TASK_FAILED_CAUSE_FEATURE_DISABLED:

		// Enum value maps for CancelExternalWorkflowExecutionFailedCause.
		return "FeatureDisabled"
	case WORKFLOW_TASK_FAILED_CAUSE_GRPC_MESSAGE_TOO_LARGE:
		return "GrpcMessageTooLarge"
	default:
		return strconv.Itoa(int(x))
	}

}

func (WorkflowTaskFailedCause) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_failed_cause_proto_enumTypes[0].Descriptor()
}

func (WorkflowTaskFailedCause) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_failed_cause_proto_enumTypes[0]
}

func (x WorkflowTaskFailedCause) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func (WorkflowTaskFailedCause) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_failed_cause_proto_rawDescGZIP(), []int{0}
}

type StartChildWorkflowExecutionFailedCause int32

const (
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED             StartChildWorkflowExecutionFailedCause = 0
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS StartChildWorkflowExecutionFailedCause = 1
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND     StartChildWorkflowExecutionFailedCause = 2
)

var (
	StartChildWorkflowExecutionFailedCause_name = map[int32]string{
		0: "START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED",
		1: "START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS",
		2: "START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND",
	}
	StartChildWorkflowExecutionFailedCause_value = map[string]int32{
		"START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED":             0,
		"START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS": 1,
		"START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND":     2,
	}
)

func (x StartChildWorkflowExecutionFailedCause) Enum() *StartChildWorkflowExecutionFailedCause {
	p := new(StartChildWorkflowExecutionFailedCause)
	*p = x
	return p
}

func (x StartChildWorkflowExecutionFailedCause) String() string {
	switch x {
	case START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED:
		return "Unspecified"
	case START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS:
		return "WorkflowAlreadyExists"
	case START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND:
		return "NamespaceNotFound"
	default:
		return strconv.Itoa(int(x))
	}

}

func (StartChildWorkflowExecutionFailedCause) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_failed_cause_proto_enumTypes[1].Descriptor()
}

func (StartChildWorkflowExecutionFailedCause) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_failed_cause_proto_enumTypes[1]
}

func (x StartChildWorkflowExecutionFailedCause) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func (StartChildWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_failed_cause_proto_rawDescGZIP(), []int{1}
}

type CancelExternalWorkflowExecutionFailedCause int32

const (
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           CancelExternalWorkflowExecutionFailedCause = 0
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND CancelExternalWorkflowExecutionFailedCause = 1
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   CancelExternalWorkflowExecutionFailedCause = 2
)

var (
	CancelExternalWorkflowExecutionFailedCause_name = map[int32]string{
		0: "CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED",
		1: "CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND",
		2: "CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND",
	}
	CancelExternalWorkflowExecutionFailedCause_value = map[string]int32{
		"CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED":                           0,
		"CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND": 1,
		"CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND":                   2,
	}
)

func (x CancelExternalWorkflowExecutionFailedCause) Enum() *CancelExternalWorkflowExecutionFailedCause {
	p := new(CancelExternalWorkflowExecutionFailedCause)
	*p = x
	return p
}

func (x CancelExternalWorkflowExecutionFailedCause) String() string {
	switch x {
	case CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED:
		return "Unspecified"
	case CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND:
		return "ExternalWorkflowExecutionNotFound"
	case CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND:
		return "NamespaceNotFound"
	default:
		return strconv.Itoa(int(x))
	}

}

func (CancelExternalWorkflowExecutionFailedCause) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_failed_cause_proto_enumTypes[2].Descriptor()
}

func (CancelExternalWorkflowExecutionFailedCause) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_failed_cause_proto_enumTypes[2]
}

func (x CancelExternalWorkflowExecutionFailedCause) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CancelExternalWorkflowExecutionFailedCause.Descriptor instead.
func (CancelExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_failed_cause_proto_rawDescGZIP(), []int{2}
}

type SignalExternalWorkflowExecutionFailedCause int32

const (
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           SignalExternalWorkflowExecutionFailedCause = 0
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND SignalExternalWorkflowExecutionFailedCause = 1
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   SignalExternalWorkflowExecutionFailedCause = 2
	// Signal count limit is per workflow and controlled by server dynamic config "history.maximumSignalsPerExecution"
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_SIGNAL_COUNT_LIMIT_EXCEEDED SignalExternalWorkflowExecutionFailedCause = 3
)

// Enum value maps for SignalExternalWorkflowExecutionFailedCause.
var (
	SignalExternalWorkflowExecutionFailedCause_name = map[int32]string{
		0: "SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED",
		1: "SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND",
		2: "SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND",
		3: "SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_SIGNAL_COUNT_LIMIT_EXCEEDED",
	}
	SignalExternalWorkflowExecutionFailedCause_value = map[string]int32{
		"SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED":                           0,
		"SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND": 1,
		"SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND":                   2,
		"SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_SIGNAL_COUNT_LIMIT_EXCEEDED":           3,
	}
)

func (x SignalExternalWorkflowExecutionFailedCause) Enum() *SignalExternalWorkflowExecutionFailedCause {
	p := new(SignalExternalWorkflowExecutionFailedCause)
	*p = x
	return p
}

func (x SignalExternalWorkflowExecutionFailedCause) String() string {
	switch x {
	case SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED:
		return "Unspecified"
	case SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND:
		return "ExternalWorkflowExecutionNotFound"
	case SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND:
		return "NamespaceNotFound"
	case SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_SIGNAL_COUNT_LIMIT_EXCEEDED:
		return "SignalCountLimitExceeded"
	default:
		return strconv.Itoa(int(x))
	}

}

func (SignalExternalWorkflowExecutionFailedCause) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_failed_cause_proto_enumTypes[3].Descriptor()
}

func (SignalExternalWorkflowExecutionFailedCause) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_failed_cause_proto_enumTypes[3]
}

func (x SignalExternalWorkflowExecutionFailedCause) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SignalExternalWorkflowExecutionFailedCause.Descriptor instead.
func (SignalExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_failed_cause_proto_rawDescGZIP(), []int{3}
}

type ResourceExhaustedCause int32

const (
	RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED ResourceExhaustedCause = 0
	// Caller exceeds request per second limit.
	RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT ResourceExhaustedCause = 1
	// Caller exceeds max concurrent request limit.
	RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT ResourceExhaustedCause = 2
	// System overloaded.
	RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED ResourceExhaustedCause = 3
	// Namespace exceeds persistence rate limit.
	RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT ResourceExhaustedCause = 4
	// Workflow is busy
	RESOURCE_EXHAUSTED_CAUSE_BUSY_WORKFLOW ResourceExhaustedCause = 5
	// Caller exceeds action per second limit.
	RESOURCE_EXHAUSTED_CAUSE_APS_LIMIT ResourceExhaustedCause = 6
	// Persistence storage limit exceeded.
	RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_STORAGE_LIMIT ResourceExhaustedCause = 7
	// Circuit breaker is open/half-open.
	RESOURCE_EXHAUSTED_CAUSE_CIRCUIT_BREAKER_OPEN ResourceExhaustedCause = 8
	// Namespace exceeds operations rate limit.
	RESOURCE_EXHAUSTED_CAUSE_OPS_LIMIT ResourceExhaustedCause = 9
)

// Enum value maps for ResourceExhaustedCause.
var (
	ResourceExhaustedCause_name = map[int32]string{
		0: "RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED",
		1: "RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT",
		2: "RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT",
		3: "RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED",
		4: "RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT",
		5: "RESOURCE_EXHAUSTED_CAUSE_BUSY_WORKFLOW",
		6: "RESOURCE_EXHAUSTED_CAUSE_APS_LIMIT",
		7: "RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_STORAGE_LIMIT",
		8: "RESOURCE_EXHAUSTED_CAUSE_CIRCUIT_BREAKER_OPEN",
		9: "RESOURCE_EXHAUSTED_CAUSE_OPS_LIMIT",
	}
	ResourceExhaustedCause_value = map[string]int32{
		"RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED":               0,
		"RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT":                 1,
		"RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT":          2,
		"RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED":         3,
		"RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT":         4,
		"RESOURCE_EXHAUSTED_CAUSE_BUSY_WORKFLOW":             5,
		"RESOURCE_EXHAUSTED_CAUSE_APS_LIMIT":                 6,
		"RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_STORAGE_LIMIT": 7,
		"RESOURCE_EXHAUSTED_CAUSE_CIRCUIT_BREAKER_OPEN":      8,
		"RESOURCE_EXHAUSTED_CAUSE_OPS_LIMIT":                 9,
	}
)

func (x ResourceExhaustedCause) Enum() *ResourceExhaustedCause {
	p := new(ResourceExhaustedCause)
	*p = x
	return p
}

func (x ResourceExhaustedCause) String() string {
	switch x {
	case RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED:
		return "Unspecified"
	case RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT:
		return "RpsLimit"
	case RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT:
		return "ConcurrentLimit"
	case RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED:
		return "SystemOverloaded"
	case RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT:
		return "PersistenceLimit"
	case RESOURCE_EXHAUSTED_CAUSE_BUSY_WORKFLOW:
		return "BusyWorkflow"
	case RESOURCE_EXHAUSTED_CAUSE_APS_LIMIT:
		return "ApsLimit"
	case RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_STORAGE_LIMIT:

		// Deprecated: Use ResourceExhaustedCause.Descriptor instead.
		return "PersistenceStorageLimit"
	case RESOURCE_EXHAUSTED_CAUSE_CIRCUIT_BREAKER_OPEN:
		return "CircuitBreakerOpen"
	case RESOURCE_EXHAUSTED_CAUSE_OPS_LIMIT:
		return "OpsLimit"
	default:
		return strconv.Itoa(int(x))
	}

}

func (ResourceExhaustedCause) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_failed_cause_proto_enumTypes[4].Descriptor()
}

func (ResourceExhaustedCause) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_failed_cause_proto_enumTypes[4]
}

func (x ResourceExhaustedCause) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

func (ResourceExhaustedCause) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_failed_cause_proto_rawDescGZIP(), []int{4}
}

type ResourceExhaustedScope int32

const (
	RESOURCE_EXHAUSTED_SCOPE_UNSPECIFIED ResourceExhaustedScope = 0
	// Exhausted resource is a system-level resource.
	RESOURCE_EXHAUSTED_SCOPE_NAMESPACE ResourceExhaustedScope = 1
	// Exhausted resource is a namespace-level resource.
	RESOURCE_EXHAUSTED_SCOPE_SYSTEM ResourceExhaustedScope = 2
)

// Enum value maps for ResourceExhaustedScope.
var (
	ResourceExhaustedScope_name = map[int32]string{
		0: "RESOURCE_EXHAUSTED_SCOPE_UNSPECIFIED",
		1: "RESOURCE_EXHAUSTED_SCOPE_NAMESPACE",
		2: "RESOURCE_EXHAUSTED_SCOPE_SYSTEM",
	}
	ResourceExhaustedScope_value = map[string]int32{
		"RESOURCE_EXHAUSTED_SCOPE_UNSPECIFIED": 0,
		"RESOURCE_EXHAUSTED_SCOPE_NAMESPACE":   1,
		"RESOURCE_EXHAUSTED_SCOPE_SYSTEM":      2,
	}
)

func (x ResourceExhaustedScope) Enum() *ResourceExhaustedScope {
	p := new(ResourceExhaustedScope)
	*p = x
	return p
}

func (x ResourceExhaustedScope) String() string {
	switch x {
	case RESOURCE_EXHAUSTED_SCOPE_UNSPECIFIED:
		return "Unspecified"
	case RESOURCE_EXHAUSTED_SCOPE_NAMESPACE:
		return "Namespace"
	case RESOURCE_EXHAUSTED_SCOPE_SYSTEM:
		return "System"
	default:
		return strconv.Itoa(int(x))
	}

}

func (ResourceExhaustedScope) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_failed_cause_proto_enumTypes[5].Descriptor()
}

func (ResourceExhaustedScope) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_failed_cause_proto_enumTypes[5]
}

func (x ResourceExhaustedScope) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceExhaustedScope.Descriptor instead.
func (ResourceExhaustedScope) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_failed_cause_proto_rawDescGZIP(), []int{5}
}

var File_temporal_api_enums_v1_failed_cause_proto protoreflect.FileDescriptor

const file_temporal_api_enums_v1_failed_cause_proto_rawDesc = "" +
	"\n" +
	"(temporal/api/enums/v1/failed_cause.proto\x12\x15temporal.api.enums.v1*\xa5\x12\n" +
	"\x17WorkflowTaskFailedCause\x12*\n" +
	"&WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED\x10\x00\x120\n" +
	",WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND\x10\x01\x12?\n" +
	";WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES\x10\x02\x12E\n" +
	"AWORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES\x10\x03\x129\n" +
	"5WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES\x10\x04\x12:\n" +
	"6WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES\x10\x05\x12;\n" +
	"7WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES\x10\x06\x12I\n" +
	"EWORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES\x10\a\x12E\n" +
	"AWORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES\x10\b\x12G\n" +
	"CWORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES\x10\t\x12X\n" +
	"TWORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES\x10\n" +
	"\x12=\n" +
	"9WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES\x10\v\x127\n" +
	"3WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID\x10\f\x126\n" +
	"2WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE\x10\r\x12@\n" +
	"<WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE\x10\x0e\x12G\n" +
	"CWORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES\x10\x0f\x12C\n" +
	"?WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES\x10\x10\x122\n" +
	".WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND\x10\x11\x125\n" +
	"1WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND\x10\x12\x124\n" +
	"0WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE\x10\x13\x12-\n" +
	")WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW\x10\x14\x12)\n" +
	"%WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY\x10\x15\x12=\n" +
	"9WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID\x10\x16\x124\n" +
	"0WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES\x10\x17\x126\n" +
	"2WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR\x10\x18\x12H\n" +
	"DWORKFLOW_TASK_FAILED_CAUSE_BAD_MODIFY_WORKFLOW_PROPERTIES_ATTRIBUTES\x10\x19\x12E\n" +
	"AWORKFLOW_TASK_FAILED_CAUSE_PENDING_CHILD_WORKFLOWS_LIMIT_EXCEEDED\x10\x1a\x12@\n" +
	"<WORKFLOW_TASK_FAILED_CAUSE_PENDING_ACTIVITIES_LIMIT_EXCEEDED\x10\x1b\x12=\n" +
	"9WORKFLOW_TASK_FAILED_CAUSE_PENDING_SIGNALS_LIMIT_EXCEEDED\x10\x1c\x12D\n" +
	"@WORKFLOW_TASK_FAILED_CAUSE_PENDING_REQUEST_CANCEL_LIMIT_EXCEEDED\x10\x1d\x12D\n" +
	"@WORKFLOW_TASK_FAILED_CAUSE_BAD_UPDATE_WORKFLOW_EXECUTION_MESSAGE\x10\x1e\x12/\n" +
	"+WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_UPDATE\x10\x1f\x12F\n" +
	"BWORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_NEXUS_OPERATION_ATTRIBUTES\x10 \x12F\n" +
	"BWORKFLOW_TASK_FAILED_CAUSE_PENDING_NEXUS_OPERATIONS_LIMIT_EXCEEDED\x10!\x12L\n" +
	"HWORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_NEXUS_OPERATION_ATTRIBUTES\x10\"\x12/\n" +
	"+WORKFLOW_TASK_FAILED_CAUSE_FEATURE_DISABLED\x10#\x125\n" +
	"1WORKFLOW_TASK_FAILED_CAUSE_GRPC_MESSAGE_TOO_LARGE\x10$*\xf3\x01\n" +
	"&StartChildWorkflowExecutionFailedCause\x12;\n" +
	"7START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED\x10\x00\x12G\n" +
	"CSTART_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS\x10\x01\x12C\n" +
	"?START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND\x10\x02*\x91\x02\n" +
	"*CancelExternalWorkflowExecutionFailedCause\x12?\n" +
	";CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED\x10\x00\x12Y\n" +
	"UCANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND\x10\x01\x12G\n" +
	"CCANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND\x10\x02*\xe2\x02\n" +
	"*SignalExternalWorkflowExecutionFailedCause\x12?\n" +
	";SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED\x10\x00\x12Y\n" +
	"USIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND\x10\x01\x12G\n" +
	"CSIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND\x10\x02\x12O\n" +
	"KSIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_SIGNAL_COUNT_LIMIT_EXCEEDED\x10\x03*\xe0\x03\n" +
	"\x16ResourceExhaustedCause\x12(\n" +
	"$RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED\x10\x00\x12&\n" +
	"\"RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT\x10\x01\x12-\n" +
	")RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT\x10\x02\x12.\n" +
	"*RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED\x10\x03\x12.\n" +
	"*RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT\x10\x04\x12*\n" +
	"&RESOURCE_EXHAUSTED_CAUSE_BUSY_WORKFLOW\x10\x05\x12&\n" +
	"\"RESOURCE_EXHAUSTED_CAUSE_APS_LIMIT\x10\x06\x126\n" +
	"2RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_STORAGE_LIMIT\x10\a\x121\n" +
	"-RESOURCE_EXHAUSTED_CAUSE_CIRCUIT_BREAKER_OPEN\x10\b\x12&\n" +
	"\"RESOURCE_EXHAUSTED_CAUSE_OPS_LIMIT\x10\t*\x8f\x01\n" +
	"\x16ResourceExhaustedScope\x12(\n" +
	"$RESOURCE_EXHAUSTED_SCOPE_UNSPECIFIED\x10\x00\x12&\n" +
	"\"RESOURCE_EXHAUSTED_SCOPE_NAMESPACE\x10\x01\x12#\n" +
	"\x1fRESOURCE_EXHAUSTED_SCOPE_SYSTEM\x10\x02B\x88\x01\n" +
	"\x18io.temporal.api.enums.v1B\x10FailedCauseProtoP\x01Z!go.temporal.io/api/enums/v1;enums\xaa\x02\x17Temporalio.Api.Enums.V1\xea\x02\x1aTemporalio::Api::Enums::V1b\x06proto3"

var (
	file_temporal_api_enums_v1_failed_cause_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_failed_cause_proto_rawDescData []byte
)

func file_temporal_api_enums_v1_failed_cause_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_failed_cause_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_failed_cause_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_failed_cause_proto_rawDesc), len(file_temporal_api_enums_v1_failed_cause_proto_rawDesc)))
	})
	return file_temporal_api_enums_v1_failed_cause_proto_rawDescData
}

var file_temporal_api_enums_v1_failed_cause_proto_enumTypes = make([]protoimpl.EnumInfo, 6)
var file_temporal_api_enums_v1_failed_cause_proto_goTypes = []any{
	(WorkflowTaskFailedCause)(0),                    // 0: temporal.api.enums.v1.WorkflowTaskFailedCause
	(StartChildWorkflowExecutionFailedCause)(0),     // 1: temporal.api.enums.v1.StartChildWorkflowExecutionFailedCause
	(CancelExternalWorkflowExecutionFailedCause)(0), // 2: temporal.api.enums.v1.CancelExternalWorkflowExecutionFailedCause
	(SignalExternalWorkflowExecutionFailedCause)(0), // 3: temporal.api.enums.v1.SignalExternalWorkflowExecutionFailedCause
	(ResourceExhaustedCause)(0),                     // 4: temporal.api.enums.v1.ResourceExhaustedCause
	(ResourceExhaustedScope)(0),                     // 5: temporal.api.enums.v1.ResourceExhaustedScope
}
var file_temporal_api_enums_v1_failed_cause_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_failed_cause_proto_init() }
func file_temporal_api_enums_v1_failed_cause_proto_init() {
	if File_temporal_api_enums_v1_failed_cause_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_failed_cause_proto_rawDesc), len(file_temporal_api_enums_v1_failed_cause_proto_rawDesc)),
			NumEnums:      6,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_failed_cause_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_failed_cause_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_failed_cause_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_failed_cause_proto = out.File
	file_temporal_api_enums_v1_failed_cause_proto_goTypes = nil
	file_temporal_api_enums_v1_failed_cause_proto_depIdxs = nil
}
