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
package nexus

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type Failure to the protobuf v3 wire format
func (val *Failure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Failure from the protobuf v3 wire format
func (val *Failure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
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

// Marshal an object of type HandlerError to the protobuf v3 wire format
func (val *HandlerError) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type HandlerError from the protobuf v3 wire format
func (val *HandlerError) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *HandlerError) Size() int {
	return proto.Size(val)
}

// Equal returns whether two HandlerError values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *HandlerError) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *HandlerError
	switch t := that.(type) {
	case *HandlerError:
		that1 = t
	case HandlerError:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UnsuccessfulOperationError to the protobuf v3 wire format
func (val *UnsuccessfulOperationError) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UnsuccessfulOperationError from the protobuf v3 wire format
func (val *UnsuccessfulOperationError) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UnsuccessfulOperationError) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UnsuccessfulOperationError values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UnsuccessfulOperationError) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UnsuccessfulOperationError
	switch t := that.(type) {
	case *UnsuccessfulOperationError:
		that1 = t
	case UnsuccessfulOperationError:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StartOperationRequest to the protobuf v3 wire format
func (val *StartOperationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StartOperationRequest from the protobuf v3 wire format
func (val *StartOperationRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StartOperationRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartOperationRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartOperationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartOperationRequest
	switch t := that.(type) {
	case *StartOperationRequest:
		that1 = t
	case StartOperationRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CancelOperationRequest to the protobuf v3 wire format
func (val *CancelOperationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CancelOperationRequest from the protobuf v3 wire format
func (val *CancelOperationRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CancelOperationRequest) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelOperationRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelOperationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelOperationRequest
	switch t := that.(type) {
	case *CancelOperationRequest:
		that1 = t
	case CancelOperationRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Request to the protobuf v3 wire format
func (val *Request) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Request from the protobuf v3 wire format
func (val *Request) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Request) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Request values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Request) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Request
	switch t := that.(type) {
	case *Request:
		that1 = t
	case Request:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StartOperationResponse to the protobuf v3 wire format
func (val *StartOperationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StartOperationResponse from the protobuf v3 wire format
func (val *StartOperationResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StartOperationResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartOperationResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartOperationResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartOperationResponse
	switch t := that.(type) {
	case *StartOperationResponse:
		that1 = t
	case StartOperationResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CancelOperationResponse to the protobuf v3 wire format
func (val *CancelOperationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CancelOperationResponse from the protobuf v3 wire format
func (val *CancelOperationResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CancelOperationResponse) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelOperationResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelOperationResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelOperationResponse
	switch t := that.(type) {
	case *CancelOperationResponse:
		that1 = t
	case CancelOperationResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Response to the protobuf v3 wire format
func (val *Response) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Response from the protobuf v3 wire format
func (val *Response) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Response) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Response values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Response) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Response
	switch t := that.(type) {
	case *Response:
		that1 = t
	case Response:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type IncomingService to the protobuf v3 wire format
func (val *IncomingService) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type IncomingService from the protobuf v3 wire format
func (val *IncomingService) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *IncomingService) Size() int {
	return proto.Size(val)
}

// Equal returns whether two IncomingService values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IncomingService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IncomingService
	switch t := that.(type) {
	case *IncomingService:
		that1 = t
	case IncomingService:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type IncomingServiceSpec to the protobuf v3 wire format
func (val *IncomingServiceSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type IncomingServiceSpec from the protobuf v3 wire format
func (val *IncomingServiceSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *IncomingServiceSpec) Size() int {
	return proto.Size(val)
}

// Equal returns whether two IncomingServiceSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IncomingServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IncomingServiceSpec
	switch t := that.(type) {
	case *IncomingServiceSpec:
		that1 = t
	case IncomingServiceSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type OutgoingService to the protobuf v3 wire format
func (val *OutgoingService) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type OutgoingService from the protobuf v3 wire format
func (val *OutgoingService) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *OutgoingService) Size() int {
	return proto.Size(val)
}

// Equal returns whether two OutgoingService values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *OutgoingService) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *OutgoingService
	switch t := that.(type) {
	case *OutgoingService:
		that1 = t
	case OutgoingService:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type OutgoingServiceSpec to the protobuf v3 wire format
func (val *OutgoingServiceSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type OutgoingServiceSpec from the protobuf v3 wire format
func (val *OutgoingServiceSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *OutgoingServiceSpec) Size() int {
	return proto.Size(val)
}

// Equal returns whether two OutgoingServiceSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *OutgoingServiceSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *OutgoingServiceSpec
	switch t := that.(type) {
	case *OutgoingServiceSpec:
		that1 = t
	case OutgoingServiceSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
