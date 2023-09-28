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

package filter

import "google.golang.org/protobuf/proto"

func (val *WorkflowExecutionFilter) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *WorkflowExecutionFilter) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two WorkflowExecutionFilter values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionFilter
	switch t := that.(type) {
	case *WorkflowExecutionFilter:
		that1 = t
	case WorkflowExecutionFilter:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *WorkflowTypeFilter) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *WorkflowTypeFilter) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two WorkflowTypeFilter values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowTypeFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowTypeFilter
	switch t := that.(type) {
	case *WorkflowTypeFilter:
		that1 = t
	case WorkflowTypeFilter:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StartTimeFilter) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StartTimeFilter) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StartTimeFilter values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartTimeFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartTimeFilter
	switch t := that.(type) {
	case *StartTimeFilter:
		that1 = t
	case StartTimeFilter:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StatusFilter) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StatusFilter) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StatusFilter values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StatusFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StatusFilter
	switch t := that.(type) {
	case *StatusFilter:
		that1 = t
	case StatusFilter:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
