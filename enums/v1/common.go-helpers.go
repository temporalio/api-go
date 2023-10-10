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
	EncodingType_shortNameValue = map[string]int32{
		"Unspecified": 0,
		"Proto3":      1,
		"Json":        2,
	}
)

// EncodingTypeFromString parses a EncodingType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to EncodingType
func EncodingTypeFromString(s string) (EncodingType, error) {
	if v, ok := EncodingType_value[s]; ok {
		return EncodingType(v), nil
	} else if v, ok := EncodingType_shortNameValue[s]; ok {
		return EncodingType(v), nil
	}
	return EncodingType(0), fmt.Errorf("Invalid value for EncodingType: %s", s)
}

var (
	IndexedValueType_shortNameValue = map[string]int32{
		"Unspecified": 0,
		"Text":        1,
		"Keyword":     2,
		"Int":         3,
		"Double":      4,
		"Bool":        5,
		"Datetime":    6,
		"KeywordList": 7,
	}
)

// IndexedValueTypeFromString parses a IndexedValueType value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to IndexedValueType
func IndexedValueTypeFromString(s string) (IndexedValueType, error) {
	if v, ok := IndexedValueType_value[s]; ok {
		return IndexedValueType(v), nil
	} else if v, ok := IndexedValueType_shortNameValue[s]; ok {
		return IndexedValueType(v), nil
	}
	return IndexedValueType(0), fmt.Errorf("Invalid value for IndexedValueType: %s", s)
}

var (
	Severity_shortNameValue = map[string]int32{
		"Unspecified": 0,
		"High":        1,
		"Medium":      2,
		"Low":         3,
	}
)

// SeverityFromString parses a Severity value from  either the protojson
// canonical SCREAMING_CASE enum or the traditional temporal PascalCase enum to Severity
func SeverityFromString(s string) (Severity, error) {
	if v, ok := Severity_value[s]; ok {
		return Severity(v), nil
	} else if v, ok := Severity_shortNameValue[s]; ok {
		return Severity(v), nil
	}
	return Severity(0), fmt.Errorf("Invalid value for Severity: %s", s)
}
