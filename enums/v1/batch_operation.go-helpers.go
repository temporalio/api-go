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
	BatchOperationType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Terminate":   1,
		"Cancel":      2,
		"Signal":      3,
		"Delete":      4,
		"Reset":       5,
	}
	BatchOperationType_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Terminate",
		2: "Cancel",
		3: "Signal",
		4: "Delete",
		5: "Reset",
	}
)

// BatchOperationTypeFromString parses a BatchOperationType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to BatchOperationType
func BatchOperationTypeFromString(s string) (BatchOperationType, error) {
	if v, ok := BatchOperationType_value[s]; ok {
		return BatchOperationType(v), nil
	} else if v, ok := BatchOperationType_shorthandValue[s]; ok {
		return BatchOperationType(v), nil
	}
	return BatchOperationType(0), fmt.Errorf("%s is not a valid BatchOperationType", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	BatchOperationType(-1).Shorthand() // will return "", false
func (e BatchOperationType) Shorthand() (string, bool) {
	if s, ok := BatchOperationType_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}

var (
	BatchOperationState_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Running":     1,
		"Completed":   2,
		"Failed":      3,
	}
	BatchOperationState_shorthandName = map[int32]string{
		0: "Unspecified",
		1: "Running",
		2: "Completed",
		3: "Failed",
	}
)

// BatchOperationStateFromString parses a BatchOperationState value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to BatchOperationState
func BatchOperationStateFromString(s string) (BatchOperationState, error) {
	if v, ok := BatchOperationState_value[s]; ok {
		return BatchOperationState(v), nil
	} else if v, ok := BatchOperationState_shorthandValue[s]; ok {
		return BatchOperationState(v), nil
	}
	return BatchOperationState(0), fmt.Errorf("%s is not a valid BatchOperationState", s)
}

// Shorthand returns the shorthand temporal PascalCase variant of this enum's string representation.
// For example, CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED will return as "Unspecified".
// This also returns whether the value is valid to prevent bugs caused by invalid casts:
//
//	BatchOperationState(-1).Shorthand() // will return "", false
func (e BatchOperationState) Shorthand() (string, bool) {
	if s, ok := BatchOperationState_shorthandName[int32(e)]; ok {
		return s, true
	}
	return "", false
}
