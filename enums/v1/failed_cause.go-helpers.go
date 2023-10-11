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
	WorkflowTaskFailedCause_shorthandValue = map[string]int32{
		"Unspecified":                                         0,
		"UnhandledCommand":                                    1,
		"BadScheduleActivityAttributes":                       2,
		"BadRequestCancelActivityAttributes":                  3,
		"BadStartTimerAttributes":                             4,
		"BadCancelTimerAttributes":                            5,
		"BadRecordMarkerAttributes":                           6,
		"BadCompleteWorkflowExecutionAttributes":              7,
		"BadFailWorkflowExecutionAttributes":                  8,
		"BadCancelWorkflowExecutionAttributes":                9,
		"BadRequestCancelExternalWorkflowExecutionAttributes": 10,
		"BadContinueAsNewAttributes":                          11,
		"StartTimerDuplicateId":                               12,
		"ResetStickyTaskQueue":                                13,
		"WorkflowWorkerUnhandledFailure":                      14,
		"BadSignalWorkflowExecutionAttributes":                15,
		"BadStartChildExecutionAttributes":                    16,
		"ForceCloseCommand":                                   17,
		"FailoverCloseCommand":                                18,
		"BadSignalInputSize":                                  19,
		"ResetWorkflow":                                       20,
		"BadBinary":                                           21,
		"ScheduleActivityDuplicateId":                         22,
		"BadSearchAttributes":                                 23,
		"NonDeterministicError":                               24,
		"BadModifyWorkflowPropertiesAttributes":               25,
		"PendingChildWorkflowsLimitExceeded":                  26,
		"PendingActivitiesLimitExceeded":                      27,
		"PendingSignalsLimitExceeded":                         28,
		"PendingRequestCancelLimitExceeded":                   29,
		"BadUpdateWorkflowExecutionMessage":                   30,
		"UnhandledUpdate":                                     31,
	}
	WorkflowTaskFailedCause_shorthandName = map[int32]string{
		0:  "Unspecified",
		1:  "UnhandledCommand",
		2:  "BadScheduleActivityAttributes",
		3:  "BadRequestCancelActivityAttributes",
		4:  "BadStartTimerAttributes",
		5:  "BadCancelTimerAttributes",
		6:  "BadRecordMarkerAttributes",
		7:  "BadCompleteWorkflowExecutionAttributes",
		8:  "BadFailWorkflowExecutionAttributes",
		9:  "BadCancelWorkflowExecutionAttributes",
		10: "BadRequestCancelExternalWorkflowExecutionAttributes",
		11: "BadContinueAsNewAttributes",
		12: "StartTimerDuplicateId",
		13: "ResetStickyTaskQueue",
		14: "WorkflowWorkerUnhandledFailure",
		15: "BadSignalWorkflowExecutionAttributes",
		16: "BadStartChildExecutionAttributes",
		17: "ForceCloseCommand",
		18: "FailoverCloseCommand",
		19: "BadSignalInputSize",
		20: "ResetWorkflow",
		21: "BadBinary",
		22: "ScheduleActivityDuplicateId",
		23: "BadSearchAttributes",
		24: "NonDeterministicError",
		25: "BadModifyWorkflowPropertiesAttributes",
		26: "PendingChildWorkflowsLimitExceeded",
		27: "PendingActivitiesLimitExceeded",
		28: "PendingSignalsLimitExceeded",
		29: "PendingRequestCancelLimitExceeded",
		30: "BadUpdateWorkflowExecutionMessage",
		31: "UnhandledUpdate",
	}
)

// WorkflowTaskFailedCauseFromString parses a WorkflowTaskFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to WorkflowTaskFailedCause
func WorkflowTaskFailedCauseFromString(s string) (WorkflowTaskFailedCause, error) {
	if v, ok := WorkflowTaskFailedCause_value[s]; ok {
		return WorkflowTaskFailedCause(v), nil
	} else if v, ok := WorkflowTaskFailedCause_shorthandValue[s]; ok {
		return WorkflowTaskFailedCause(v), nil
	}
	return WorkflowTaskFailedCause(0), fmt.Errorf("%s is not a valid WorkflowTaskFailedCause", s)
}

var (
	StartChildWorkflowExecutionFailedCause_shorthandValue = map[string]int32{
		"Unspecified":           0,
		"WorkflowAlreadyExists": 1,
		"NamespaceNotFound":     2,
	}
	StartChildWorkflowExecutionFailedCause_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "WorkflowAlreadyExists",
		2: "NamespaceNotFound",
	}
)

// StartChildWorkflowExecutionFailedCauseFromString parses a StartChildWorkflowExecutionFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to StartChildWorkflowExecutionFailedCause
func StartChildWorkflowExecutionFailedCauseFromString(s string) (StartChildWorkflowExecutionFailedCause, error) {
	if v, ok := StartChildWorkflowExecutionFailedCause_value[s]; ok {
		return StartChildWorkflowExecutionFailedCause(v), nil
	} else if v, ok := StartChildWorkflowExecutionFailedCause_shorthandValue[s]; ok {
		return StartChildWorkflowExecutionFailedCause(v), nil
	}
	return StartChildWorkflowExecutionFailedCause(0), fmt.Errorf("%s is not a valid StartChildWorkflowExecutionFailedCause", s)
}

var (
	CancelExternalWorkflowExecutionFailedCause_shorthandValue = map[string]int32{
		"Unspecified":                       0,
		"ExternalWorkflowExecutionNotFound": 1,
		"NamespaceNotFound":                 2,
	}
	CancelExternalWorkflowExecutionFailedCause_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "ExternalWorkflowExecutionNotFound",
		2: "NamespaceNotFound",
	}
)

// CancelExternalWorkflowExecutionFailedCauseFromString parses a CancelExternalWorkflowExecutionFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to CancelExternalWorkflowExecutionFailedCause
func CancelExternalWorkflowExecutionFailedCauseFromString(s string) (CancelExternalWorkflowExecutionFailedCause, error) {
	if v, ok := CancelExternalWorkflowExecutionFailedCause_value[s]; ok {
		return CancelExternalWorkflowExecutionFailedCause(v), nil
	} else if v, ok := CancelExternalWorkflowExecutionFailedCause_shorthandValue[s]; ok {
		return CancelExternalWorkflowExecutionFailedCause(v), nil
	}
	return CancelExternalWorkflowExecutionFailedCause(0), fmt.Errorf("%s is not a valid CancelExternalWorkflowExecutionFailedCause", s)
}

var (
	SignalExternalWorkflowExecutionFailedCause_shorthandValue = map[string]int32{
		"Unspecified":                       0,
		"ExternalWorkflowExecutionNotFound": 1,
		"NamespaceNotFound":                 2,
		"SignalCountLimitExceeded":          3,
	}
	SignalExternalWorkflowExecutionFailedCause_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "ExternalWorkflowExecutionNotFound",
		2: "NamespaceNotFound",
		3: "SignalCountLimitExceeded",
	}
)

// SignalExternalWorkflowExecutionFailedCauseFromString parses a SignalExternalWorkflowExecutionFailedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to SignalExternalWorkflowExecutionFailedCause
func SignalExternalWorkflowExecutionFailedCauseFromString(s string) (SignalExternalWorkflowExecutionFailedCause, error) {
	if v, ok := SignalExternalWorkflowExecutionFailedCause_value[s]; ok {
		return SignalExternalWorkflowExecutionFailedCause(v), nil
	} else if v, ok := SignalExternalWorkflowExecutionFailedCause_shorthandValue[s]; ok {
		return SignalExternalWorkflowExecutionFailedCause(v), nil
	}
	return SignalExternalWorkflowExecutionFailedCause(0), fmt.Errorf("%s is not a valid SignalExternalWorkflowExecutionFailedCause", s)
}

var (
	ResourceExhaustedCause_shorthandValue = map[string]int32{
		"Unspecified":      0,
		"RpsLimit":         1,
		"ConcurrentLimit":  2,
		"SystemOverloaded": 3,
		"PersistenceLimit": 4,
		"BusyWorkflow":     5,
		"ApsLimit":         6,
	}
	ResourceExhaustedCause_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "RpsLimit",
		2: "ConcurrentLimit",
		3: "SystemOverloaded",
		4: "PersistenceLimit",
		5: "BusyWorkflow",
		6: "ApsLimit",
	}
)

// ResourceExhaustedCauseFromString parses a ResourceExhaustedCause value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ResourceExhaustedCause
func ResourceExhaustedCauseFromString(s string) (ResourceExhaustedCause, error) {
	if v, ok := ResourceExhaustedCause_value[s]; ok {
		return ResourceExhaustedCause(v), nil
	} else if v, ok := ResourceExhaustedCause_shorthandValue[s]; ok {
		return ResourceExhaustedCause(v), nil
	}
	return ResourceExhaustedCause(0), fmt.Errorf("%s is not a valid ResourceExhaustedCause", s)
}
