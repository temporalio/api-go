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

package version

import "google.golang.org/protobuf/proto"

func (val *ReleaseInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ReleaseInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ReleaseInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReleaseInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReleaseInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReleaseInfo
	switch t := that.(type) {
	case *ReleaseInfo:
		that1 = t
	case ReleaseInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *Alert) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *Alert) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *Alert) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Alert values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Alert) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Alert
	switch t := that.(type) {
	case *Alert:
		that1 = t
	case Alert:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *VersionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *VersionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *VersionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersionInfo
	switch t := that.(type) {
	case *VersionInfo:
		that1 = t
	case VersionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
