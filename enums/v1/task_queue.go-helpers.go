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
	TaskQueueKind_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Normal":      1,
		"Sticky":      2,
	}
	TaskQueueKind_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Normal",
		2: "Sticky",
	}
)

// TaskQueueKindFromString parses a TaskQueueKind value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to TaskQueueKind
func TaskQueueKindFromString(s string) (TaskQueueKind, error) {
	if v, ok := TaskQueueKind_value[s]; ok {
		return TaskQueueKind(v), nil
	} else if v, ok := TaskQueueKind_shorthandValue[s]; ok {
		return TaskQueueKind(v), nil
	}
	return TaskQueueKind(0), fmt.Errorf("%s is not a valid TaskQueueKind", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	TaskQueueKind(-1).Shorthand() // will return "", false
func (e TaskQueueKind) Shorthand() (string, bool) {
	if s, ok := TaskQueueKind_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	TaskQueueType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Workflow":    1,
		"Activity":    2,
	}
	TaskQueueType_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Workflow",
		2: "Activity",
	}
)

// TaskQueueTypeFromString parses a TaskQueueType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to TaskQueueType
func TaskQueueTypeFromString(s string) (TaskQueueType, error) {
	if v, ok := TaskQueueType_value[s]; ok {
		return TaskQueueType(v), nil
	} else if v, ok := TaskQueueType_shorthandValue[s]; ok {
		return TaskQueueType(v), nil
	}
	return TaskQueueType(0), fmt.Errorf("%s is not a valid TaskQueueType", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	TaskQueueType(-1).Shorthand() // will return "", false
func (e TaskQueueType) Shorthand() (string, bool) {
	if s, ok := TaskQueueType_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	TaskReachability_shorthandValue = map[string]int32{
		"Unspecified":       0,
		"NewWorkflows":      1,
		"ExistingWorkflows": 2,
		"OpenWorkflows":     3,
		"ClosedWorkflows":   4,
	}
	TaskReachability_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "NewWorkflows",
		2: "ExistingWorkflows",
		3: "OpenWorkflows",
		4: "ClosedWorkflows",
	}
)

// TaskReachabilityFromString parses a TaskReachability value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to TaskReachability
func TaskReachabilityFromString(s string) (TaskReachability, error) {
	if v, ok := TaskReachability_value[s]; ok {
		return TaskReachability(v), nil
	} else if v, ok := TaskReachability_shorthandValue[s]; ok {
		return TaskReachability(v), nil
	}
	return TaskReachability(0), fmt.Errorf("%s is not a valid TaskReachability", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	TaskReachability(-1).Shorthand() // will return "", false
func (e TaskReachability) Shorthand() (string, bool) {
	if s, ok := TaskReachability_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}
