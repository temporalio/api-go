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

package test

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
	namespacepb "go.temporal.io/api/namespace/v1"
	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/protobuf/proto"
)

const serializedNSWithInvalidUTF8 = "CrUBKrIBChtfX3RlbXBvcmFsLXN0b3JhZ2UtbWV0YWRhdGESkgEKjwEKDmFwLW5vcnRoZWFzdC0xEn0IAVp5Cg5hcC1ub3J0aGVhc3QtMRIjbnMtc3RvcmFnZS1wcm9kLWFwLW5vcnRoZWFzdC0xLTM2ZTQaQmFybjphd3M6aWFtOjo1MTEwMjI4MjM1NTQ6cm9sZS9ucy1zdG9yYWdlLXByb2QtYXAtbm9ydGhlYXN0LTEtMzZlNA=="

var nsWithInvalidUTF8 = &workflowservice.DescribeNamespaceResponse{
	NamespaceInfo: &namespacepb.NamespaceInfo{
		Data: map[string]string{
			// 0x8f01 is invalid UTF-8
			"metadata": "\n\x8f\x01\n\x0ejunk\x12data",
		},
	},
}

func TestUnmarshalNamespaceInfo_AllowsInvalidUTF8(t *testing.T) {
	require := require.New(t)
	ns := &workflowservice.DescribeNamespaceResponse{}
	bs, err := base64.StdEncoding.DecodeString(serializedNSWithInvalidUTF8)
	require.NoError(err, "unable to decode serialized ns info")
	require.NoError(proto.Unmarshal(bs, ns), "unable to unmarshal namespace detail containing invalid UTF-8")
	require.True(proto.Equal(nsWithInvalidUTF8, ns), "unmarshaled namespace detail does not match expectation")
}

func TestMarshalNamespaceInfo_AllowsInvalidUTF8(t *testing.T) {
	require := require.New(t)
	bs, err := proto.Marshal(nsWithInvalidUTF8)
	require.NoError(err, "unable to marshal namespace detail containing invalid UTF-8")
	require.Equal(serializedNSWithInvalidUTF8, base64.StdEncoding.EncodeToString(bs), "marshaled namespace detail does not match expectation")
}
