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

// This doesn't actually test the temporalproto package but we don't want a top-level
// test package.
package temporalproto_test

import (
	"reflect"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	namespacepb "go.temporal.io/api/namespace/v1"
	taskqueuepb "go.temporal.io/api/taskqueue/v1"
	"go.temporal.io/api/workflowservice/v1"
)

// 0x8f01 is invalid UTF-8
const invalidUTF8 = "\n\x8f\x01\n\x0ejunk\x12data"

// Let's just be sure
func TestInvalidUTF_Sample_IsActuallyInvalid(t *testing.T) {
	if utf8.ValidString(invalidUTF8) {
		t.Fatalf("Invalid UTF8 sample is actually valid")
	}
}

func TestProto_AllowsInvalidUTF8_InStrings(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input proto.Message
	}{{
		name: "Invalid UTF-8 in map value",
		input: &namespacepb.NamespaceInfo{
			Data: map[string]string{
				"valid UTF8": invalidUTF8,
			},
		},
	}, {
		name: "Invalid UTF-8 in map key",
		input: &namespacepb.NamespaceInfo{
			Data: map[string]string{
				invalidUTF8: "valid utf8",
			},
		},
	}, {
		name: "Invalid UTF-8 in struct field",
		input: &namespacepb.NamespaceInfo{
			Name: invalidUTF8,
		},
	}, {
		name: "Invalid UTF-8 in repeated string",
		input: &taskqueuepb.CompatibleVersionSet{
			BuildIds: []string{
				"valid utf-8",
				invalidUTF8,
			},
		},
	}, {
		name: "Invalid UTF-8 in nested message",
		input: &workflowservice.DescribeNamespaceResponse{
			NamespaceInfo: &namespacepb.NamespaceInfo{
				Name: invalidUTF8,
				Data: map[string]string{
					invalidUTF8: invalidUTF8,
				},
			},
		},
	}} {
		t.Run(tc.name, func(t *testing.T) {
			require := require.New(t)
			out := reflect.New(reflect.TypeOf(tc.input).Elem()).Interface().(proto.Message)
			bs, err := proto.Marshal(tc.input)
			require.NoError(err, "unable to marshal proto containing invalid UTF-8")
			require.NoError(proto.Unmarshal(bs, out), "unable to unmarshal proto containing invalid UTF-8")
			require.True(proto.Equal(tc.input, out), "unmarshaled proto does not match input")
		})
	}
}
