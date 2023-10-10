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

package enums

import (
	"fmt"
)

var (
	WorkflowIdReusePolicy_shortNameValue = map[string]int32{
		"Unspecified":              0,
		"AllowDuplicate":           1,
		"AllowDuplicateFailedOnly": 2,
		"RejectDuplicate":          3,
		"TerminateIfRunning":       4,
	}
)

// WorkflowIdReusePolicyFromString parses a WorkflowIdReusePolicy value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to WorkflowIdReusePolicy
func WorkflowIdReusePolicyFromString(s string) (WorkflowIdReusePolicy, error) {
	if v, ok := WorkflowIdReusePolicy_value[s]; ok {
		return WorkflowIdReusePolicy(v), nil
	} else if v, ok := WorkflowIdReusePolicy_shortNameValue[s]; ok {
		return WorkflowIdReusePolicy(v), nil
	}
	return WorkflowIdReusePolicy(0), fmt.Errorf("Invalid value for WorkflowIdReusePolicy: %s", s)
}

var (
	ParentClosePolicy_shortNameValue = map[string]int32{
		"Unspecified":   0,
		"Terminate":     1,
		"Abandon":       2,
		"RequestCancel": 3,
	}
)

// ParentClosePolicyFromString parses a ParentClosePolicy value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ParentClosePolicy
func ParentClosePolicyFromString(s string) (ParentClosePolicy, error) {
	if v, ok := ParentClosePolicy_value[s]; ok {
		return ParentClosePolicy(v), nil
	} else if v, ok := ParentClosePolicy_shortNameValue[s]; ok {
		return ParentClosePolicy(v), nil
	}
	return ParentClosePolicy(0), fmt.Errorf("Invalid value for ParentClosePolicy: %s", s)
}

var (
	ContinueAsNewInitiator_shortNameValue = map[string]int32{
		"Unspecified":  0,
		"Workflow":     1,
		"Retry":        2,
		"CronSchedule": 3,
	}
)

// ContinueAsNewInitiatorFromString parses a ContinueAsNewInitiator value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ContinueAsNewInitiator
func ContinueAsNewInitiatorFromString(s string) (ContinueAsNewInitiator, error) {
	if v, ok := ContinueAsNewInitiator_value[s]; ok {
		return ContinueAsNewInitiator(v), nil
	} else if v, ok := ContinueAsNewInitiator_shortNameValue[s]; ok {
		return ContinueAsNewInitiator(v), nil
	}
	return ContinueAsNewInitiator(0), fmt.Errorf("Invalid value for ContinueAsNewInitiator: %s", s)
}

var (
	WorkflowExecutionStatus_shortNameValue = map[string]int32{
		"Unspecified":    0,
		"Running":        1,
		"Completed":      2,
		"Failed":         3,
		"Canceled":       4,
		"Terminated":     5,
		"ContinuedAsNew": 6,
		"TimedOut":       7,
	}
)

// WorkflowExecutionStatusFromString parses a WorkflowExecutionStatus value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to WorkflowExecutionStatus
func WorkflowExecutionStatusFromString(s string) (WorkflowExecutionStatus, error) {
	if v, ok := WorkflowExecutionStatus_value[s]; ok {
		return WorkflowExecutionStatus(v), nil
	} else if v, ok := WorkflowExecutionStatus_shortNameValue[s]; ok {
		return WorkflowExecutionStatus(v), nil
	}
	return WorkflowExecutionStatus(0), fmt.Errorf("Invalid value for WorkflowExecutionStatus: %s", s)
}

var (
	PendingActivityState_shortNameValue = map[string]int32{
		"Unspecified":     0,
		"Scheduled":       1,
		"Started":         2,
		"CancelRequested": 3,
	}
)

// PendingActivityStateFromString parses a PendingActivityState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to PendingActivityState
func PendingActivityStateFromString(s string) (PendingActivityState, error) {
	if v, ok := PendingActivityState_value[s]; ok {
		return PendingActivityState(v), nil
	} else if v, ok := PendingActivityState_shortNameValue[s]; ok {
		return PendingActivityState(v), nil
	}
	return PendingActivityState(0), fmt.Errorf("Invalid value for PendingActivityState: %s", s)
}

var (
	PendingWorkflowTaskState_shortNameValue = map[string]int32{
		"Unspecified": 0,
		"Scheduled":   1,
		"Started":     2,
	}
)

// PendingWorkflowTaskStateFromString parses a PendingWorkflowTaskState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to PendingWorkflowTaskState
func PendingWorkflowTaskStateFromString(s string) (PendingWorkflowTaskState, error) {
	if v, ok := PendingWorkflowTaskState_value[s]; ok {
		return PendingWorkflowTaskState(v), nil
	} else if v, ok := PendingWorkflowTaskState_shortNameValue[s]; ok {
		return PendingWorkflowTaskState(v), nil
	}
	return PendingWorkflowTaskState(0), fmt.Errorf("Invalid value for PendingWorkflowTaskState: %s", s)
}

var (
	HistoryEventFilterType_shortNameValue = map[string]int32{
		"Unspecified": 0,
		"AllEvent":    1,
		"CloseEvent":  2,
	}
)

// HistoryEventFilterTypeFromString parses a HistoryEventFilterType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to HistoryEventFilterType
func HistoryEventFilterTypeFromString(s string) (HistoryEventFilterType, error) {
	if v, ok := HistoryEventFilterType_value[s]; ok {
		return HistoryEventFilterType(v), nil
	} else if v, ok := HistoryEventFilterType_shortNameValue[s]; ok {
		return HistoryEventFilterType(v), nil
	}
	return HistoryEventFilterType(0), fmt.Errorf("Invalid value for HistoryEventFilterType: %s", s)
}

var (
	RetryState_shortNameValue = map[string]int32{
		"Unspecified":            0,
		"InProgress":             1,
		"NonRetryableFailure":    2,
		"Timeout":                3,
		"MaximumAttemptsReached": 4,
		"RetryPolicyNotSet":      5,
		"InternalServerError":    6,
		"CancelRequested":        7,
	}
)

// RetryStateFromString parses a RetryState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to RetryState
func RetryStateFromString(s string) (RetryState, error) {
	if v, ok := RetryState_value[s]; ok {
		return RetryState(v), nil
	} else if v, ok := RetryState_shortNameValue[s]; ok {
		return RetryState(v), nil
	}
	return RetryState(0), fmt.Errorf("Invalid value for RetryState: %s", s)
}

var (
	TimeoutType_shortNameValue = map[string]int32{
		"Unspecified":     0,
		"StartToClose":    1,
		"ScheduleToStart": 2,
		"ScheduleToClose": 3,
		"Heartbeat":       4,
	}
)

// TimeoutTypeFromString parses a TimeoutType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to TimeoutType
func TimeoutTypeFromString(s string) (TimeoutType, error) {
	if v, ok := TimeoutType_value[s]; ok {
		return TimeoutType(v), nil
	} else if v, ok := TimeoutType_shortNameValue[s]; ok {
		return TimeoutType(v), nil
	}
	return TimeoutType(0), fmt.Errorf("Invalid value for TimeoutType: %s", s)
}
