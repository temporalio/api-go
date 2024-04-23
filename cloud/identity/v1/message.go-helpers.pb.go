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
package identity

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type AccountAccess to the protobuf v3 wire format
func (val *AccountAccess) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type AccountAccess from the protobuf v3 wire format
func (val *AccountAccess) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *AccountAccess) Size() int {
	return proto.Size(val)
}

// Equal returns whether two AccountAccess values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *AccountAccess) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *AccountAccess
	switch t := that.(type) {
	case *AccountAccess:
		that1 = t
	case AccountAccess:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NamespaceAccess to the protobuf v3 wire format
func (val *NamespaceAccess) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NamespaceAccess from the protobuf v3 wire format
func (val *NamespaceAccess) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NamespaceAccess) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NamespaceAccess values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceAccess) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceAccess
	switch t := that.(type) {
	case *NamespaceAccess:
		that1 = t
	case NamespaceAccess:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Access to the protobuf v3 wire format
func (val *Access) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Access from the protobuf v3 wire format
func (val *Access) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Access) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Access values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Access) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Access
	switch t := that.(type) {
	case *Access:
		that1 = t
	case Access:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UserSpec to the protobuf v3 wire format
func (val *UserSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UserSpec from the protobuf v3 wire format
func (val *UserSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UserSpec) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UserSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UserSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UserSpec
	switch t := that.(type) {
	case *UserSpec:
		that1 = t
	case UserSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Invitation to the protobuf v3 wire format
func (val *Invitation) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Invitation from the protobuf v3 wire format
func (val *Invitation) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Invitation) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Invitation values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Invitation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Invitation
	switch t := that.(type) {
	case *Invitation:
		that1 = t
	case Invitation:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type User to the protobuf v3 wire format
func (val *User) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type User from the protobuf v3 wire format
func (val *User) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *User) Size() int {
	return proto.Size(val)
}

// Equal returns whether two User values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *User) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *User
	switch t := that.(type) {
	case *User:
		that1 = t
	case User:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
