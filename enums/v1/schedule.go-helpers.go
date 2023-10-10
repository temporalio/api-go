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
	ScheduleOverlapPolicy_shortNameValue = map[string]int32{
		"Unspecified":    0,
		"Skip":           1,
		"BufferOne":      2,
		"BufferAll":      3,
		"CancelOther":    4,
		"TerminateOther": 5,
		"AllowAll":       6,
	}
)

// ScheduleOverlapPolicyFromString parses a ScheduleOverlapPolicy value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ScheduleOverlapPolicy
func ScheduleOverlapPolicyFromString(s string) (ScheduleOverlapPolicy, error) {
	if v, ok := ScheduleOverlapPolicy_value[s]; ok {
		return ScheduleOverlapPolicy(v), nil
	} else if v, ok := ScheduleOverlapPolicy_shortNameValue[s]; ok {
		return ScheduleOverlapPolicy(v), nil
	}
	return ScheduleOverlapPolicy(0), fmt.Errorf("Invalid value for ScheduleOverlapPolicy: %s", s)
}
