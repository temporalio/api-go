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

package common_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/temporalproto"
)

var tests = []struct {
	name string
	pb   proto.Message
	json string
}{{
	name: "binary/plain",
	pb: &common.Payload{
		Metadata: map[string][]byte{
			"encoding": []byte("binary/plain"),
		},
		Data: []byte("bytes")},
	json: `{"metadata": {"encoding": "YmluYXJ5L3BsYWlu"}, "data": "Ynl0ZXM="}`,
}, {
	name: "json/plain",
	json: `{"_protoMessageType": "temporal.api.common.v1.WorkflowType", "name": "workflow-name"}`,
	pb: &common.Payload{
		Metadata: map[string][]byte{
			"encoding":    []byte("json/protobuf"),
			"messageType": []byte("temporal.api.common.v1.WorkflowType"),
		},
		Data: []byte(`{"name":"workflow-name"}`),
	},
}, {
	name: "memo with varying fields",
	json: `{"fields": {
                      "some-string": "string",
			          "some-array": ["foo", 123, false]}}`,
	pb: &common.Memo{
		Fields: map[string]*common.Payload{
			"some-string": &common.Payload{
				Metadata: map[string][]byte{
					"encoding": []byte("json/plain"),
				},
				Data: []byte(`"string"`),
			},
			"some-array": &common.Payload{
				Metadata: map[string][]byte{
					"encoding": []byte("json/plain"),
				},
				// NOTE: we don't include spurious spaces when re-encoding
				Data: []byte(`["foo",123,false]`),
			},
		},
	},
}}

func TestMaybeMarshal(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts temporalproto.CustomJSONMarshalOptions
			got, err := opts.Marshal(tt.pb)
			require.NoError(t, err)
			require.JSONEq(t, tt.json, string(got), tt.pb)
		})
	}
}

func TestMaybeUnmarshal(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out = tt.pb.ProtoReflect().New().Interface()
			var opts temporalproto.CustomJSONUnmarshalOptions
			err := opts.Unmarshal([]byte(tt.json), out)
			require.NoError(t, err)
			if !proto.Equal(tt.pb, out) {
				t.Errorf("protos mismatched\n%#v\n%#v", tt.pb, out)
			}
		})
	}
}
