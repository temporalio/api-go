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

package failure

import "google.golang.org/protobuf/proto"

func (val *ApplicationFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ApplicationFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ApplicationFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ApplicationFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ApplicationFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ApplicationFailureInfo
	switch t := that.(type) {
	case *ApplicationFailureInfo:
		that1 = t
	case ApplicationFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *TimeoutFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *TimeoutFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *TimeoutFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimeoutFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimeoutFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimeoutFailureInfo
	switch t := that.(type) {
	case *TimeoutFailureInfo:
		that1 = t
	case TimeoutFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CanceledFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CanceledFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *CanceledFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CanceledFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CanceledFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CanceledFailureInfo
	switch t := that.(type) {
	case *CanceledFailureInfo:
		that1 = t
	case CanceledFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *TerminatedFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *TerminatedFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *TerminatedFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TerminatedFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TerminatedFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TerminatedFailureInfo
	switch t := that.(type) {
	case *TerminatedFailureInfo:
		that1 = t
	case TerminatedFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ServerFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ServerFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ServerFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ServerFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ServerFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ServerFailureInfo
	switch t := that.(type) {
	case *ServerFailureInfo:
		that1 = t
	case ServerFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ResetWorkflowFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ResetWorkflowFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ResetWorkflowFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetWorkflowFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetWorkflowFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetWorkflowFailureInfo
	switch t := that.(type) {
	case *ResetWorkflowFailureInfo:
		that1 = t
	case ResetWorkflowFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ActivityFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ActivityFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ActivityFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ActivityFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ActivityFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ActivityFailureInfo
	switch t := that.(type) {
	case *ActivityFailureInfo:
		that1 = t
	case ActivityFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ChildWorkflowExecutionFailureInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ChildWorkflowExecutionFailureInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ChildWorkflowExecutionFailureInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ChildWorkflowExecutionFailureInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ChildWorkflowExecutionFailureInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ChildWorkflowExecutionFailureInfo
	switch t := that.(type) {
	case *ChildWorkflowExecutionFailureInfo:
		that1 = t
	case ChildWorkflowExecutionFailureInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *Failure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *Failure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *Failure) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Failure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Failure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Failure
	switch t := that.(type) {
	case *Failure:
		that1 = t
	case Failure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
