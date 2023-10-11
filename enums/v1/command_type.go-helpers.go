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
	CommandType_shorthandValue = map[string]int32{
		"Unspecified":                            0,
		"ScheduleActivityTask":                   1,
		"RequestCancelActivityTask":              2,
		"StartTimer":                             3,
		"CompleteWorkflowExecution":              4,
		"FailWorkflowExecution":                  5,
		"CancelTimer":                            6,
		"CancelWorkflowExecution":                7,
		"RequestCancelExternalWorkflowExecution": 8,
		"RecordMarker":                           9,
		"ContinueAsNewWorkflowExecution":         10,
		"StartChildWorkflowExecution":            11,
		"SignalExternalWorkflowExecution":        12,
		"UpsertWorkflowSearchAttributes":         13,
		"ProtocolMessage":                        14,
		"ModifyWorkflowProperties":               16,
	}
	CommandType_shorthandName = map[int32]string{
		0:  "Unspecified",
		1:  "ScheduleActivityTask",
		2:  "RequestCancelActivityTask",
		3:  "StartTimer",
		4:  "CompleteWorkflowExecution",
		5:  "FailWorkflowExecution",
		6:  "CancelTimer",
		7:  "CancelWorkflowExecution",
		8:  "RequestCancelExternalWorkflowExecution",
		9:  "RecordMarker",
		10: "ContinueAsNewWorkflowExecution",
		11: "StartChildWorkflowExecution",
		12: "SignalExternalWorkflowExecution",
		13: "UpsertWorkflowSearchAttributes",
		14: "ProtocolMessage",
		16: "ModifyWorkflowProperties",
	}
)

// CommandTypeFromString parses a CommandType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to CommandType
func CommandTypeFromString(s string) (CommandType, error) {
	if v, ok := CommandType_value[s]; ok {
		return CommandType(v), nil
	} else if v, ok := CommandType_shorthandValue[s]; ok {
		return CommandType(v), nil
	}
	return CommandType(0), fmt.Errorf("%s is not a valid CommandType", s)
}
