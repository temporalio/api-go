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

package namespace

import "google.golang.org/protobuf/proto"

func (val *NamespaceInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NamespaceInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceInfo
	switch t := that.(type) {
	case *NamespaceInfo:
		that1 = t
	case NamespaceInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NamespaceConfig) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceConfig) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NamespaceConfig values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceConfig
	switch t := that.(type) {
	case *NamespaceConfig:
		that1 = t
	case NamespaceConfig:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *BadBinaries) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *BadBinaries) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two BadBinaries values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BadBinaries) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BadBinaries
	switch t := that.(type) {
	case *BadBinaries:
		that1 = t
	case BadBinaries:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *BadBinaryInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *BadBinaryInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two BadBinaryInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BadBinaryInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BadBinaryInfo
	switch t := that.(type) {
	case *BadBinaryInfo:
		that1 = t
	case BadBinaryInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateNamespaceInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateNamespaceInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateNamespaceInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateNamespaceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateNamespaceInfo
	switch t := that.(type) {
	case *UpdateNamespaceInfo:
		that1 = t
	case UpdateNamespaceInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NamespaceFilter) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceFilter) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NamespaceFilter values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceFilter
	switch t := that.(type) {
	case *NamespaceFilter:
		that1 = t
	case NamespaceFilter:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
