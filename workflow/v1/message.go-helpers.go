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

package workflow

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type WorkflowExecutionInfo to the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionInfo from the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionInfo
	switch t := that.(type) {
	case *WorkflowExecutionInfo:
		that1 = t
	case WorkflowExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionConfig to the protobuf v3 wire format
func (val *WorkflowExecutionConfig) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionConfig from the protobuf v3 wire format
func (val *WorkflowExecutionConfig) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionConfig) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionConfig values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionConfig
	switch t := that.(type) {
	case *WorkflowExecutionConfig:
		that1 = t
	case WorkflowExecutionConfig:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingActivityInfo to the protobuf v3 wire format
func (val *PendingActivityInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingActivityInfo from the protobuf v3 wire format
func (val *PendingActivityInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingActivityInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingActivityInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingActivityInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingActivityInfo
	switch t := that.(type) {
	case *PendingActivityInfo:
		that1 = t
	case PendingActivityInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingChildExecutionInfo to the protobuf v3 wire format
func (val *PendingChildExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingChildExecutionInfo from the protobuf v3 wire format
func (val *PendingChildExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingChildExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingChildExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingChildExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingChildExecutionInfo
	switch t := that.(type) {
	case *PendingChildExecutionInfo:
		that1 = t
	case PendingChildExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingWorkflowTaskInfo to the protobuf v3 wire format
func (val *PendingWorkflowTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingWorkflowTaskInfo from the protobuf v3 wire format
func (val *PendingWorkflowTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingWorkflowTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingWorkflowTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingWorkflowTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingWorkflowTaskInfo
	switch t := that.(type) {
	case *PendingWorkflowTaskInfo:
		that1 = t
	case PendingWorkflowTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetPoints to the protobuf v3 wire format
func (val *ResetPoints) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetPoints from the protobuf v3 wire format
func (val *ResetPoints) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetPoints) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetPoints values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetPoints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetPoints
	switch t := that.(type) {
	case *ResetPoints:
		that1 = t
	case ResetPoints:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetPointInfo to the protobuf v3 wire format
func (val *ResetPointInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetPointInfo from the protobuf v3 wire format
func (val *ResetPointInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetPointInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetPointInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetPointInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetPointInfo
	switch t := that.(type) {
	case *ResetPointInfo:
		that1 = t
	case ResetPointInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NewWorkflowExecutionInfo to the protobuf v3 wire format
func (val *NewWorkflowExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NewWorkflowExecutionInfo from the protobuf v3 wire format
func (val *NewWorkflowExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NewWorkflowExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NewWorkflowExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NewWorkflowExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NewWorkflowExecutionInfo
	switch t := that.(type) {
	case *NewWorkflowExecutionInfo:
		that1 = t
	case NewWorkflowExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}