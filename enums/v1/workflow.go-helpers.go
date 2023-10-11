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
	WorkflowIdReusePolicy_shorthandValue = map[string]int32{
		"Unspecified":              0,
		"AllowDuplicate":           1,
		"AllowDuplicateFailedOnly": 2,
		"RejectDuplicate":          3,
		"TerminateIfRunning":       4,
	}
	WorkflowIdReusePolicy_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "AllowDuplicate",
		2: "AllowDuplicateFailedOnly",
		3: "RejectDuplicate",
		4: "TerminateIfRunning",
	}
)

// WorkflowIdReusePolicyFromString parses a WorkflowIdReusePolicy value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to WorkflowIdReusePolicy
func WorkflowIdReusePolicyFromString(s string) (WorkflowIdReusePolicy, error) {
	if v, ok := WorkflowIdReusePolicy_value[s]; ok {
		return WorkflowIdReusePolicy(v), nil
	} else if v, ok := WorkflowIdReusePolicy_shorthandValue[s]; ok {
		return WorkflowIdReusePolicy(v), nil
	}
	return WorkflowIdReusePolicy(0), fmt.Errorf("%s is not a valid WorkflowIdReusePolicy", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	WorkflowIdReusePolicy(-1).Shorthand() // will return "", false
func (e WorkflowIdReusePolicy) Shorthand() (string, bool) {
	if s, ok := WorkflowIdReusePolicy_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	ParentClosePolicy_shorthandValue = map[string]int32{
		"Unspecified":   0,
		"Terminate":     1,
		"Abandon":       2,
		"RequestCancel": 3,
	}
	ParentClosePolicy_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Terminate",
		2: "Abandon",
		3: "RequestCancel",
	}
)

// ParentClosePolicyFromString parses a ParentClosePolicy value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ParentClosePolicy
func ParentClosePolicyFromString(s string) (ParentClosePolicy, error) {
	if v, ok := ParentClosePolicy_value[s]; ok {
		return ParentClosePolicy(v), nil
	} else if v, ok := ParentClosePolicy_shorthandValue[s]; ok {
		return ParentClosePolicy(v), nil
	}
	return ParentClosePolicy(0), fmt.Errorf("%s is not a valid ParentClosePolicy", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	ParentClosePolicy(-1).Shorthand() // will return "", false
func (e ParentClosePolicy) Shorthand() (string, bool) {
	if s, ok := ParentClosePolicy_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	ContinueAsNewInitiator_shorthandValue = map[string]int32{
		"Unspecified":  0,
		"Workflow":     1,
		"Retry":        2,
		"CronSchedule": 3,
	}
	ContinueAsNewInitiator_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Workflow",
		2: "Retry",
		3: "CronSchedule",
	}
)

// ContinueAsNewInitiatorFromString parses a ContinueAsNewInitiator value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ContinueAsNewInitiator
func ContinueAsNewInitiatorFromString(s string) (ContinueAsNewInitiator, error) {
	if v, ok := ContinueAsNewInitiator_value[s]; ok {
		return ContinueAsNewInitiator(v), nil
	} else if v, ok := ContinueAsNewInitiator_shorthandValue[s]; ok {
		return ContinueAsNewInitiator(v), nil
	}
	return ContinueAsNewInitiator(0), fmt.Errorf("%s is not a valid ContinueAsNewInitiator", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	ContinueAsNewInitiator(-1).Shorthand() // will return "", false
func (e ContinueAsNewInitiator) Shorthand() (string, bool) {
	if s, ok := ContinueAsNewInitiator_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	WorkflowExecutionStatus_shorthandValue = map[string]int32{
		"Unspecified":    0,
		"Running":        1,
		"Completed":      2,
		"Failed":         3,
		"Canceled":       4,
		"Terminated":     5,
		"ContinuedAsNew": 6,
		"TimedOut":       7,
	}
	WorkflowExecutionStatus_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Running",
		2: "Completed",
		3: "Failed",
		4: "Canceled",
		5: "Terminated",
		6: "ContinuedAsNew",
		7: "TimedOut",
	}
)

// WorkflowExecutionStatusFromString parses a WorkflowExecutionStatus value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to WorkflowExecutionStatus
func WorkflowExecutionStatusFromString(s string) (WorkflowExecutionStatus, error) {
	if v, ok := WorkflowExecutionStatus_value[s]; ok {
		return WorkflowExecutionStatus(v), nil
	} else if v, ok := WorkflowExecutionStatus_shorthandValue[s]; ok {
		return WorkflowExecutionStatus(v), nil
	}
	return WorkflowExecutionStatus(0), fmt.Errorf("%s is not a valid WorkflowExecutionStatus", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	WorkflowExecutionStatus(-1).Shorthand() // will return "", false
func (e WorkflowExecutionStatus) Shorthand() (string, bool) {
	if s, ok := WorkflowExecutionStatus_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	PendingActivityState_shorthandValue = map[string]int32{
		"Unspecified":     0,
		"Scheduled":       1,
		"Started":         2,
		"CancelRequested": 3,
	}
	PendingActivityState_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Scheduled",
		2: "Started",
		3: "CancelRequested",
	}
)

// PendingActivityStateFromString parses a PendingActivityState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to PendingActivityState
func PendingActivityStateFromString(s string) (PendingActivityState, error) {
	if v, ok := PendingActivityState_value[s]; ok {
		return PendingActivityState(v), nil
	} else if v, ok := PendingActivityState_shorthandValue[s]; ok {
		return PendingActivityState(v), nil
	}
	return PendingActivityState(0), fmt.Errorf("%s is not a valid PendingActivityState", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	PendingActivityState(-1).Shorthand() // will return "", false
func (e PendingActivityState) Shorthand() (string, bool) {
	if s, ok := PendingActivityState_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	PendingWorkflowTaskState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Scheduled":   1,
		"Started":     2,
	}
	PendingWorkflowTaskState_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Scheduled",
		2: "Started",
	}
)

// PendingWorkflowTaskStateFromString parses a PendingWorkflowTaskState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to PendingWorkflowTaskState
func PendingWorkflowTaskStateFromString(s string) (PendingWorkflowTaskState, error) {
	if v, ok := PendingWorkflowTaskState_value[s]; ok {
		return PendingWorkflowTaskState(v), nil
	} else if v, ok := PendingWorkflowTaskState_shorthandValue[s]; ok {
		return PendingWorkflowTaskState(v), nil
	}
	return PendingWorkflowTaskState(0), fmt.Errorf("%s is not a valid PendingWorkflowTaskState", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	PendingWorkflowTaskState(-1).Shorthand() // will return "", false
func (e PendingWorkflowTaskState) Shorthand() (string, bool) {
	if s, ok := PendingWorkflowTaskState_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	HistoryEventFilterType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"AllEvent":    1,
		"CloseEvent":  2,
	}
	HistoryEventFilterType_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "AllEvent",
		2: "CloseEvent",
	}
)

// HistoryEventFilterTypeFromString parses a HistoryEventFilterType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to HistoryEventFilterType
func HistoryEventFilterTypeFromString(s string) (HistoryEventFilterType, error) {
	if v, ok := HistoryEventFilterType_value[s]; ok {
		return HistoryEventFilterType(v), nil
	} else if v, ok := HistoryEventFilterType_shorthandValue[s]; ok {
		return HistoryEventFilterType(v), nil
	}
	return HistoryEventFilterType(0), fmt.Errorf("%s is not a valid HistoryEventFilterType", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	HistoryEventFilterType(-1).Shorthand() // will return "", false
func (e HistoryEventFilterType) Shorthand() (string, bool) {
	if s, ok := HistoryEventFilterType_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	RetryState_shorthandValue = map[string]int32{
		"Unspecified":            0,
		"InProgress":             1,
		"NonRetryableFailure":    2,
		"Timeout":                3,
		"MaximumAttemptsReached": 4,
		"RetryPolicyNotSet":      5,
		"InternalServerError":    6,
		"CancelRequested":        7,
	}
	RetryState_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "InProgress",
		2: "NonRetryableFailure",
		3: "Timeout",
		4: "MaximumAttemptsReached",
		5: "RetryPolicyNotSet",
		6: "InternalServerError",
		7: "CancelRequested",
	}
)

// RetryStateFromString parses a RetryState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to RetryState
func RetryStateFromString(s string) (RetryState, error) {
	if v, ok := RetryState_value[s]; ok {
		return RetryState(v), nil
	} else if v, ok := RetryState_shorthandValue[s]; ok {
		return RetryState(v), nil
	}
	return RetryState(0), fmt.Errorf("%s is not a valid RetryState", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	RetryState(-1).Shorthand() // will return "", false
func (e RetryState) Shorthand() (string, bool) {
	if s, ok := RetryState_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	TimeoutType_shorthandValue = map[string]int32{
		"Unspecified":     0,
		"StartToClose":    1,
		"ScheduleToStart": 2,
		"ScheduleToClose": 3,
		"Heartbeat":       4,
	}
	TimeoutType_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "StartToClose",
		2: "ScheduleToStart",
		3: "ScheduleToClose",
		4: "Heartbeat",
	}
)

// TimeoutTypeFromString parses a TimeoutType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to TimeoutType
func TimeoutTypeFromString(s string) (TimeoutType, error) {
	if v, ok := TimeoutType_value[s]; ok {
		return TimeoutType(v), nil
	} else if v, ok := TimeoutType_shorthandValue[s]; ok {
		return TimeoutType(v), nil
	}
	return TimeoutType(0), fmt.Errorf("%s is not a valid TimeoutType", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	TimeoutType(-1).Shorthand() // will return "", false
func (e TimeoutType) Shorthand() (string, bool) {
	if s, ok := TimeoutType_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}
