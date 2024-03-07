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

package main_test

import (
	"go/format"
	"go/parser"
	"go/token"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	protogen "go.temporal.io/api/cmd/protogen"
)

const protoGeneratedMessageDescriptor = `package sdk

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ...

var File_temporal_api_sdk_v1_task_complete_metadata_proto protoreflect.FileDescriptor

var file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73,
	0x64, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x65, 0x74, 0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31, 0x22, 0xab, 0x01, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6f, 0x72, 0x65, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6c, 0x61, 0x67,
	0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x66,
	0x6c, 0x61, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x61, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x64, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x64, 0x6b,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x87, 0x01, 0x0a, 0x16, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64, 0x6b, 0x2e, 0x76, 0x31,
	0x42, 0x19, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x67,
	0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x64, 0x6b, 0xaa, 0x02, 0x15, 0x54,
	0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x53, 0x64,
	0x6b, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x18, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69,
	0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x64, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescOnce sync.Once
	file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescData = file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc
)

// ...
`
const afterUtf8Rewrite = `package sdk

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ...

var File_temporal_api_sdk_v1_task_complete_metadata_proto protoreflect.FileDescriptor

var file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc = []byte{
	0x0a, 0x30, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x64,
	0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x31, 0x22, 0xbb, 0x01, 0x0a, 0x1d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x0f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x75, 0x73, 0x65,
	0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x63, 0x6f,
	0x72, 0x65, 0x55, 0x73, 0x65, 0x64, 0x46, 0x6c, 0x61, 0x67, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12,
	0x2a, 0x0a, 0x0f, 0x6c, 0x61, 0x6e, 0x67, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x66, 0x6c, 0x61, 0x67,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0d, 0x52, 0x0d, 0x6c, 0x61, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x64,
	0x46, 0x6c, 0x61, 0x67, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x1d, 0x0a, 0x08, 0x73, 0x64, 0x6b,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x64, 0x6b, 0x4e,
	0x61, 0x6d, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x23, 0x0a, 0x0b, 0x73, 0x64, 0x6b, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x64, 0x6b,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x42, 0x87, 0x01, 0x0a, 0x16, 0x69,
	0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x64,
	0x6b, 0x2e, 0x76, 0x31, 0x42, 0x19, 0x54, 0x61, 0x73, 0x6b, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65,
	0x74, 0x65, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x1d, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x64, 0x6b, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x64, 0x6b, 0xaa, 0x02,
	0x15, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x53,
	0x64, 0x6b, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x18, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69,
	0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x53, 0x64, 0x6b, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescOnce sync.Once
	file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDescData = file_temporal_api_sdk_v1_task_complete_metadata_proto_rawDesc
)

// ...
`

func TestRewriteDisableUtf8(t *testing.T) {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "", protoGeneratedMessageDescriptor, parser.ParseComments)
	if err != nil {
		t.Errorf("Failed to parse code: %s", err)
		t.FailNow()
	}

	d := protogen.NewDisableUtf8Validation()
	d.Process(f)

	var b strings.Builder
	if err := format.Node(&b, fset, f); err != nil {
		t.Errorf("Failed to format AST: %s", err)
		t.FailNow()
	}
	require.Equal(t, afterUtf8Rewrite, b.String())
}