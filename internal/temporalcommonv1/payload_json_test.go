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
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/temporalproto"
)

var tests = []struct {
	name          string
	pb            proto.Message
	longformJSON  string
	shorthandJSON string
}{{
	name:          "json/plain",
	longformJSON:  `{"metadata":{"encoding":"anNvbi9wcm90b2J1Zg==","messageType":"dGVtcG9yYWwuYXBpLmNvbW1vbi52MS5Xb3JrZmxvd1R5cGU="},"data":"eyJuYW1lIjoid29ya2Zsb3ctbmFtZSJ9"}`,
	shorthandJSON: `{"_protoMessageType": "temporal.api.common.v1.WorkflowType", "name": "workflow-name"}`,
	pb: &common.Payload{
		Metadata: map[string][]byte{
			"encoding":    []byte("json/protobuf"),
			"messageType": []byte("temporal.api.common.v1.WorkflowType"),
		},
		Data: []byte(`{"name":"workflow-name"}`),
	},
}, {
	name:          "binary/null",
	longformJSON:  `{"metadata":{"encoding":"YmluYXJ5L251bGw="}}`,
	shorthandJSON: `{"metadata":{"encoding":"YmluYXJ5L251bGw="}}`,
	pb: &common.Payload{
		Metadata: map[string][]byte{
			"encoding": []byte("binary/null"),
		},
	},
}, {
	name: "memo with varying fields",
	longformJSON: `{"fields":{
                      "some-string":{"metadata":{"encoding":"anNvbi9wbGFpbg=="},"data":"InN0cmluZyI="},
                      "some-array":{"metadata":{"encoding":"anNvbi9wbGFpbg=="},"data":"WyJmb28iLDEyMyxmYWxzZV0="}}}`,
	shorthandJSON: `{"fields": {
                      "some-string": "string",
			          "some-array": ["foo", 123, false]}}`,
	pb: &common.Memo{
		Fields: map[string]*common.Payload{
			"some-string": {
				Metadata: map[string][]byte{
					"encoding": []byte("json/plain"),
				},
				Data: []byte(`"string"`),
			},
			"some-array": {
				Metadata: map[string][]byte{
					"encoding": []byte("json/plain"),
				},
				// NOTE: we don't include spurious spaces when re-encoding
				Data: []byte(`["foo",123,false]`),
			},
		},
	},
}, {
	name:          "json/plain with empty object",
	longformJSON:  `{"metadata":{"encoding":"anNvbi9wbGFpbg=="},"data":"eyJncmVldGluZyI6e319"}`,
	shorthandJSON: `{"greeting": {}}`,
	pb: &common.Payload{
		Metadata: map[string][]byte{
			"encoding": []byte("json/plain"),
		},
		Data: []byte(`{"greeting":{}}`),
	},
}, {
	name:          "json/plain with nested object",
	longformJSON:  `{"metadata":{"encoding":"anNvbi9wbGFpbg=="},"data":"eyJncmVldGluZyI6eyJuYW1lIjp7fX19"}`,
	shorthandJSON: `{"greeting": {"name": {}}}`,
	pb: &common.Payload{
		Metadata: map[string][]byte{
			"encoding": []byte("json/plain"),
		},
		Data: []byte(`{"greeting":{"name":{}}}`),
	},
}, {
	name:          "empty payloads",
	longformJSON:  `{}`,
	shorthandJSON: `[]`,
	pb:            &common.Payloads{},
}, {
	name:          "empty payloads with non-nil slice",
	longformJSON:  `{}`,
	shorthandJSON: `[]`,
	pb:            &common.Payloads{Payloads: []*common.Payload{}},
}, {
	name:          "payloads with two items",
	longformJSON:  `{"payloads":[{"data":"InN0cmluZyB2YWx1ZSI=","metadata":{"encoding":"anNvbi9wbGFpbg=="}},{"data":"MzI0Mw==","metadata":{"encoding":"anNvbi9wbGFpbg=="}}]}`,
	shorthandJSON: `["string value", 3243]`,
	pb: &common.Payloads{Payloads: []*common.Payload{
		&common.Payload{
			Metadata: map[string][]byte{"encoding": []byte("json/plain")},
			Data:     []byte(`"string value"`),
		},
		&common.Payload{
			Metadata: map[string][]byte{"encoding": []byte("json/plain")},
			Data:     []byte(`3243`),
		},
	}},
}}

func TestMaybeMarshal_ShorthandEnabled(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := temporalproto.CustomJSONMarshalOptions{
				Metadata: map[string]interface{}{
					common.EnablePayloadShorthandMetadataKey: true,
				},
			}
			got, err := opts.Marshal(tt.pb)
			require.NoError(t, err)
			t.Logf("Marshalled to %s", string(got))
			require.JSONEq(t, tt.shorthandJSON, string(got))
		})
	}
}

func TestMaybeMarshal_ShorthandDisabled(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var opts temporalproto.CustomJSONMarshalOptions
			got, err := opts.Marshal(tt.pb)
			require.NoError(t, err)
			t.Logf("Marshalled to %s", string(got))
			require.JSONEq(t, tt.longformJSON, string(got))
		})
	}
}

func TestMaybeUnmarshal_Shorthand(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out = tt.pb.ProtoReflect().New().Interface()
			opts := temporalproto.CustomJSONUnmarshalOptions{
				Metadata: map[string]interface{}{
					common.EnablePayloadShorthandMetadataKey: true,
				},
			}
			err := opts.Unmarshal([]byte(tt.shorthandJSON), out)
			require.NoError(t, err)
			require.True(t, proto.Equal(tt.pb, out))
		})
	}
}

func TestMaybeUnmarshal_Longform(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out = tt.pb.ProtoReflect().New().Interface()
			var opts temporalproto.CustomJSONUnmarshalOptions
			err := opts.Unmarshal([]byte(tt.longformJSON), out)
			require.NoError(t, err)
			require.True(t, proto.Equal(tt.pb, out))
		})
	}
}

func TestMaybeUnmarshal_Payloads_AcceptsNull(t *testing.T) {
	var out common.Payloads
	opts := temporalproto.CustomJSONUnmarshalOptions{
		Metadata: map[string]interface{}{
			common.EnablePayloadShorthandMetadataKey: true,
		},
	}
	err := opts.Unmarshal([]byte("null"), &out)
	require.NoError(t, err)
	require.Equal(t, 0, len(out.Payloads))
}

func TestMaybeMarshal_Payloads_Unhandled(t *testing.T) {
	opts := temporalproto.CustomJSONMarshalOptions{
		Metadata: map[string]interface{}{
			common.EnablePayloadShorthandMetadataKey: true,
		},
	}
	p := common.Payloads{Payloads: []*common.Payload{
		&common.Payload{
			Metadata: map[string][]byte{
				"encoding": []byte("json/plain"),
			},
			Data: []byte(`"string"`),
		},
		&common.Payload{ // this one can't be handled by shorthand because of extra metadata
			Metadata: map[string][]byte{
				"encoding":            []byte("json/plain"),
				"some other metadata": []byte("23"),
			},
			Data: []byte(`"string"`),
		},
	}}
	out, err := opts.Marshal(&p)
	require.NoError(t, err)
	var i any
	require.NoError(t, json.Unmarshal(out, &i), "must unmarshal as valid json")
	require.Equal(t, '{', rune(out[0]), "should encode as long-form, not shorthand")
}
