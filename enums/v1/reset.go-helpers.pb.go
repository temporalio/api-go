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
	ResetReapplyType_shorthandValue = map[string]int32{
		"Unspecified": 0,
		"Signal":      1,
		"None":        2,
	}
)

// ResetReapplyTypeFromString parses a ResetReapplyType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ResetReapplyType
func ResetReapplyTypeFromString(s string) (ResetReapplyType, error) {
	if v, ok := ResetReapplyType_value[s]; ok {
		return ResetReapplyType(v), nil
	} else if v, ok := ResetReapplyType_shorthandValue[s]; ok {
		return ResetReapplyType(v), nil
	}
	return ResetReapplyType(0), fmt.Errorf("%s is not a valid ResetReapplyType", s)
}

var (
	ResetType_shorthandValue = map[string]int32{
		"Unspecified":       0,
		"FirstWorkflowTask": 1,
		"LastWorkflowTask":  2,
	}
)

// ResetTypeFromString parses a ResetType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to ResetType
func ResetTypeFromString(s string) (ResetType, error) {
	if v, ok := ResetType_value[s]; ok {
		return ResetType(v), nil
	} else if v, ok := ResetType_shorthandValue[s]; ok {
		return ResetType(v), nil
	}
	return ResetType(0), fmt.Errorf("%s is not a valid ResetType", s)
}
