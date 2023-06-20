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

package proxy_test

import (
	"encoding/base64"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/enums/v1"
	"go.temporal.io/api/failure/v1"
	"go.temporal.io/api/proxy"
)

var marshaler, _ = proxy.NewJSONPBMarshaler(proxy.JSONPBMarshalerOptions{})
var unmarshaler, _ = proxy.NewJSONPBUnmarshaler(proxy.JSONPBUnmarshalerOptions{})

func TestShorthandPayloads(t *testing.T) {
	// Make a bunch of memo values to confirm payloads of different types
	var memo common.Memo
	binaryVal := base64.StdEncoding.EncodeToString([]byte("bytes"))
	assertProtoJSON(t, &memo, `{
		"fields": {
			"some-string": "string",
			"some-array": ["foo", 123, false],
			"some-bool": false,
			"some-float": 1.23,
			"some-int": 123,
			"some-null": null,
			"some-bytes": {
				"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("binary/plain"))+`" },
				"data": "`+binaryVal+`"
			},
			"some-proto": {
				"_protoMessageType": "temporal.api.common.v1.WorkflowType",
				"name": "workflow-name"
			},
			"some-raw-payload": {
				"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("my-encoding"))+`" },
				"data": "`+base64.StdEncoding.EncodeToString([]byte("raw-value"))+`"
			}
		}
	}`)
	// Confirm some payloads
	assertPayloadData(t, memo.Fields["some-string"], "json/plain", `"string"`)
	assertPayloadData(t, memo.Fields["some-float"], "json/plain", "1.23")
	assertPayloadData(t, memo.Fields["some-null"], "binary/null", "")
	assertPayloadData(t, memo.Fields["some-bytes"], "binary/plain", "bytes")
	assertPayloadJSONData(t, memo.Fields["some-proto"], "json/protobuf", `{"name":"workflow-name"}`)
	require.Equal(t, "temporal.api.common.v1.WorkflowType", string(memo.Fields["some-proto"].Metadata["messageType"]))
	assertPayloadData(t, memo.Fields["some-raw-payload"], "my-encoding", "raw-value")
}

func TestShorthandPayloadAmbiguity(t *testing.T) {
	// This demonstrates what happens if you give certain kinda-payload-like
	// forms
	var memo common.Memo
	assertProtoJSON(t, &memo, `{
		"fields": {
			"ok-payload": {
				"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("my-encoding"))+`" },
				"data": "`+base64.StdEncoding.EncodeToString([]byte("raw-value"))+`"
			},
			"like-payload-not-base64-data": {
				"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("my-encoding"))+`" },
				"data": "not@base64"
			},
			"like-payload-not-base64-metadata": {
				"metadata": { "some-meta": "some-value" },
				"data": "`+base64.StdEncoding.EncodeToString([]byte("raw-value"))+`"
			},
			"like-payload-no-data": {
				"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("my-encoding"))+`" }
			},
			"like-payload-no-metadata": {
				"data": "`+base64.StdEncoding.EncodeToString([]byte("raw-value"))+`"
			}
		}
	}`)
	assertPayloadData(t, memo.Fields["ok-payload"], "my-encoding", "raw-value")
	assertPayloadJSONData(t, memo.Fields["like-payload-not-base64-data"], "json/plain", `{
		"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("my-encoding"))+`" },
		"data": "not@base64"
	}`)
	assertPayloadJSONData(t, memo.Fields["like-payload-not-base64-metadata"], "json/plain", `{
		"metadata": { "some-meta": "some-value" },
		"data": "`+base64.StdEncoding.EncodeToString([]byte("raw-value"))+`"
	}`)
	assertPayloadJSONData(t, memo.Fields["like-payload-no-data"], "json/plain", `{
		"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("my-encoding"))+`" }
	}`)
	assertPayloadJSONData(t, memo.Fields["like-payload-not-base64-data"], "json/plain", `{
		"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("my-encoding"))+`" },
		"data": "not@base64"
	}`)
	assertPayloadJSONData(t, memo.Fields["like-payload-no-metadata"], "json/plain", `{
		"data": "`+base64.StdEncoding.EncodeToString([]byte("raw-value"))+`"
	}`)

	// A no-data raw payload is fine for binary/null, but we test separately
	// because it marshals back to shorthand
	memo.Reset()
	require.NoError(t, unmarshaler.Unmarshal(strings.NewReader(`{
		"fields": {
			"ok-null-payload": {
				"metadata": { "encoding": "`+base64.StdEncoding.EncodeToString([]byte("binary/null"))+`" }
			}
		}
	}`), &memo))
	assertPayloadData(t, memo.Fields["ok-null-payload"], "binary/null", "")
}

func TestEnumString(t *testing.T) {
	// Confirm the shorthand strings are used, not actual underscore strings
	var failureInfo failure.TimeoutFailureInfo
	assertProtoJSON(t, &failureInfo, `{
		"timeoutType": "ScheduleToStart",
		"lastHeartbeatDetails": ["some-string", { "some-obj-key": "some-obj-value" }]
	}`)
	require.Equal(t, enums.TIMEOUT_TYPE_SCHEDULE_TO_START, failureInfo.TimeoutType)
	// Might as well assert payloads while we're here
	require.Len(t, failureInfo.LastHeartbeatDetails.Payloads, 2)
	assertPayloadJSONData(t, failureInfo.LastHeartbeatDetails.Payloads[0],
		"json/plain", `"some-string"`)
	assertPayloadJSONData(t, failureInfo.LastHeartbeatDetails.Payloads[1],
		"json/plain", `{ "some-obj-key": "some-obj-value" }`)
}

func TestBadJSON(t *testing.T) {
	// Unknown field
	var failureInfo failure.TimeoutFailureInfo
	err := unmarshaler.Unmarshal(strings.NewReader(`{ "unknown-field": "some-value" }`), &failureInfo)
	require.ErrorContains(t, err, "unknown field")

	// Unknown enum value
	err = unmarshaler.Unmarshal(strings.NewReader(`{ "timeoutType": "Whatever" }`), &failureInfo)
	require.ErrorContains(t, err, "for enum")

	// Enum value using normal proto string
	err = unmarshaler.Unmarshal(strings.NewReader(`{ "timeoutType": "TIMEOUT_TYPE_SCHEDULE_TO_START" }`), &failureInfo)
	require.ErrorContains(t, err, "for enum")

	// Bad field type
	err = unmarshaler.Unmarshal(strings.NewReader(`{ "timeoutType": { } }`), &failureInfo)
	require.ErrorContains(t, err, "cannot unmarshal")
}

func assertProtoJSON(t *testing.T, p proto.Message, json string) {
	// Unmarshal
	require.NoError(t, unmarshaler.Unmarshal(strings.NewReader(json), p))

	// Marshal
	actual, err := marshaler.MarshalToString(p)
	require.NoError(t, err)

	// Compare JSON
	require.JSONEq(t, json, actual)
}

func assertPayloadData(t *testing.T, p *common.Payload, encoding, str string) {
	require.Equal(t, encoding, string(p.Metadata["encoding"]))
	require.Equal(t, str, string(p.Data))
}

func assertPayloadJSONData(t *testing.T, p *common.Payload, encoding, jsonStr string) {
	require.Equal(t, encoding, string(p.Metadata["encoding"]))
	require.JSONEq(t, jsonStr, string(p.Data))
}
