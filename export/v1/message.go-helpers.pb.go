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

// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package export

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type WorkflowExecution to the protobuf v3 wire format
func (val *WorkflowExecution) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecution from the protobuf v3 wire format
func (val *WorkflowExecution) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecution) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecution values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecution) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecution
	switch t := that.(type) {
	case *WorkflowExecution:
		that1 = t
	case WorkflowExecution:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutions to the protobuf v3 wire format
func (val *WorkflowExecutions) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutions from the protobuf v3 wire format
func (val *WorkflowExecutions) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutions) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutions values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutions
	switch t := that.(type) {
	case *WorkflowExecutions:
		that1 = t
	case WorkflowExecutions:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}