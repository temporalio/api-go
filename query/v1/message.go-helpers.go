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

package query

import "google.golang.org/protobuf/proto"

func (val *WorkflowQuery) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *WorkflowQuery) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *WorkflowQuery) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowQuery values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowQuery) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowQuery
	switch t := that.(type) {
	case *WorkflowQuery:
		that1 = t
	case WorkflowQuery:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *WorkflowQueryResult) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *WorkflowQueryResult) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *WorkflowQueryResult) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowQueryResult values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowQueryResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowQueryResult
	switch t := that.(type) {
	case *WorkflowQueryResult:
		that1 = t
	case WorkflowQueryResult:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *QueryRejected) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *QueryRejected) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *QueryRejected) Size() int {
	return proto.Size(val)
}

// Equal returns whether two QueryRejected values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryRejected) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryRejected
	switch t := that.(type) {
	case *QueryRejected:
		that1 = t
	case QueryRejected:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
