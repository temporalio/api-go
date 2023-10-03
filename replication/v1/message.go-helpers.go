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

package replication

import "google.golang.org/protobuf/proto"

func (val *ClusterReplicationConfig) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ClusterReplicationConfig) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ClusterReplicationConfig) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ClusterReplicationConfig values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ClusterReplicationConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ClusterReplicationConfig
	switch t := that.(type) {
	case *ClusterReplicationConfig:
		that1 = t
	case ClusterReplicationConfig:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NamespaceReplicationConfig) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceReplicationConfig) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *NamespaceReplicationConfig) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NamespaceReplicationConfig values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceReplicationConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceReplicationConfig
	switch t := that.(type) {
	case *NamespaceReplicationConfig:
		that1 = t
	case NamespaceReplicationConfig:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *FailoverStatus) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *FailoverStatus) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *FailoverStatus) Size() int {
	return proto.Size(val)
}

// Equal returns whether two FailoverStatus values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *FailoverStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *FailoverStatus
	switch t := that.(type) {
	case *FailoverStatus:
		that1 = t
	case FailoverStatus:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
