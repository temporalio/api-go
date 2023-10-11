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
	EventType_shorthandValue = map[string]int32{
		"Unspecified":                                     0,
		"WorkflowExecutionStarted":                        1,
		"WorkflowExecutionCompleted":                      2,
		"WorkflowExecutionFailed":                         3,
		"WorkflowExecutionTimedOut":                       4,
		"WorkflowTaskScheduled":                           5,
		"WorkflowTaskStarted":                             6,
		"WorkflowTaskCompleted":                           7,
		"WorkflowTaskTimedOut":                            8,
		"WorkflowTaskFailed":                              9,
		"ActivityTaskScheduled":                           10,
		"ActivityTaskStarted":                             11,
		"ActivityTaskCompleted":                           12,
		"ActivityTaskFailed":                              13,
		"ActivityTaskTimedOut":                            14,
		"ActivityTaskCancelRequested":                     15,
		"ActivityTaskCanceled":                            16,
		"TimerStarted":                                    17,
		"TimerFired":                                      18,
		"TimerCanceled":                                   19,
		"WorkflowExecutionCancelRequested":                20,
		"WorkflowExecutionCanceled":                       21,
		"RequestCancelExternalWorkflowExecutionInitiated": 22,
		"RequestCancelExternalWorkflowExecutionFailed":    23,
		"ExternalWorkflowExecutionCancelRequested":        24,
		"MarkerRecorded":                                  25,
		"WorkflowExecutionSignaled":                       26,
		"WorkflowExecutionTerminated":                     27,
		"WorkflowExecutionContinuedAsNew":                 28,
		"StartChildWorkflowExecutionInitiated":            29,
		"StartChildWorkflowExecutionFailed":               30,
		"ChildWorkflowExecutionStarted":                   31,
		"ChildWorkflowExecutionCompleted":                 32,
		"ChildWorkflowExecutionFailed":                    33,
		"ChildWorkflowExecutionCanceled":                  34,
		"ChildWorkflowExecutionTimedOut":                  35,
		"ChildWorkflowExecutionTerminated":                36,
		"SignalExternalWorkflowExecutionInitiated":        37,
		"SignalExternalWorkflowExecutionFailed":           38,
		"ExternalWorkflowExecutionSignaled":               39,
		"UpsertWorkflowSearchAttributes":                  40,
		"WorkflowExecutionUpdateAccepted":                 41,
		"WorkflowExecutionUpdateRejected":                 42,
		"WorkflowExecutionUpdateCompleted":                43,
		"WorkflowPropertiesModifiedExternally":            44,
		"ActivityPropertiesModifiedExternally":            45,
		"WorkflowPropertiesModified":                      46,
	}
	EventType_shorthandName = map[int32]string{
		0:  "Unspecified",
		1:  "WorkflowExecutionStarted",
		2:  "WorkflowExecutionCompleted",
		3:  "WorkflowExecutionFailed",
		4:  "WorkflowExecutionTimedOut",
		5:  "WorkflowTaskScheduled",
		6:  "WorkflowTaskStarted",
		7:  "WorkflowTaskCompleted",
		8:  "WorkflowTaskTimedOut",
		9:  "WorkflowTaskFailed",
		10: "ActivityTaskScheduled",
		11: "ActivityTaskStarted",
		12: "ActivityTaskCompleted",
		13: "ActivityTaskFailed",
		14: "ActivityTaskTimedOut",
		15: "ActivityTaskCancelRequested",
		16: "ActivityTaskCanceled",
		17: "TimerStarted",
		18: "TimerFired",
		19: "TimerCanceled",
		20: "WorkflowExecutionCancelRequested",
		21: "WorkflowExecutionCanceled",
		22: "RequestCancelExternalWorkflowExecutionInitiated",
		23: "RequestCancelExternalWorkflowExecutionFailed",
		24: "ExternalWorkflowExecutionCancelRequested",
		25: "MarkerRecorded",
		26: "WorkflowExecutionSignaled",
		27: "WorkflowExecutionTerminated",
		28: "WorkflowExecutionContinuedAsNew",
		29: "StartChildWorkflowExecutionInitiated",
		30: "StartChildWorkflowExecutionFailed",
		31: "ChildWorkflowExecutionStarted",
		32: "ChildWorkflowExecutionCompleted",
		33: "ChildWorkflowExecutionFailed",
		34: "ChildWorkflowExecutionCanceled",
		35: "ChildWorkflowExecutionTimedOut",
		36: "ChildWorkflowExecutionTerminated",
		37: "SignalExternalWorkflowExecutionInitiated",
		38: "SignalExternalWorkflowExecutionFailed",
		39: "ExternalWorkflowExecutionSignaled",
		40: "UpsertWorkflowSearchAttributes",
		41: "WorkflowExecutionUpdateAccepted",
		42: "WorkflowExecutionUpdateRejected",
		43: "WorkflowExecutionUpdateCompleted",
		44: "WorkflowPropertiesModifiedExternally",
		45: "ActivityPropertiesModifiedExternally",
		46: "WorkflowPropertiesModified",
	}
)

// EventTypeFromString parses a EventType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to EventType
func EventTypeFromString(s string) (EventType, error) {
	if v, ok := EventType_value[s]; ok {
		return EventType(v), nil
	} else if v, ok := EventType_shorthandValue[s]; ok {
		return EventType(v), nil
	}
	return EventType(0), fmt.Errorf("%s is not a valid EventType", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	EventType(-1).Shorthand() // will return "", false
func (e EventType) Shorthand() (string, bool) {
	if s, ok := EventType_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}
