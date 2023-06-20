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

package proxy

import (
	"encoding/json"
	"io"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/internal/temporalgateway"
	jsonpb "go.temporal.io/api/internal/temporaljsonpb"
)

// JSONPBMarshaler is a protobuf JSON marshaler that supports Temporal-specific
// features. This is mostly equivalent to
// [github.com/gogo/protobuf/jsonpb.Marshaler].
//
// One feature is "shorthand payloads". During marshal when shorthand payloads
// are enabled (which is the default), JSON payloads are represented as their
// actual data instead of the protobuf default which would be a base64'd data
// field and base64'd metadata fields. For JSON proto payloads, the same occurs
// but a special field in the object of "_protoMessageType" is present with the
// qualified protobuf message name.
type JSONPBMarshaler struct{ underlying jsonpb.Marshaler }

// JSONPBMarshalerOptions is used for [NewJSONPBMarshaler]. Most of the options
// are copied from [github.com/gogo/protobuf/jsonpb.Marshaler].
type JSONPBMarshalerOptions struct {
	// Whether to render enum values as integers, as opposed to string values.
	EnumsAsInts bool

	// Whether to render fields with zero values.
	EmitDefaults bool

	// A string to indent each level by. The presence of this field will
	// also cause a space to appear between the field separator and
	// value, and for newlines to be appear between fields and array
	// elements.
	Indent string

	// Whether to use the original (.proto) name for fields.
	OrigName bool

	// A custom URL resolver to use when marshaling Any messages to JSON.
	// If unset, the default resolution strategy is to extract the
	// fully-qualified type name from the type URL and pass that to
	// proto.MessageType(string).
	AnyResolver gogojsonpb.AnyResolver

	// If true, this will never marshal to shorthand payloads. See
	// [JSONPBMarshaler] for more detail.
	DisablePayloadShorthand bool
}

// NewJSONPBMarshaler creates a marshaler that supports Temporal-specific
// features. See [JSONPBMarshaler] for more detail.
func NewJSONPBMarshaler(options JSONPBMarshalerOptions) (*JSONPBMarshaler, error) {
	ret := &JSONPBMarshaler{}
	ret.underlying.EnumsAsInts = options.EnumsAsInts
	ret.underlying.EmitDefaults = options.EmitDefaults
	ret.underlying.Indent = options.Indent
	ret.underlying.OrigName = options.OrigName
	ret.underlying.AnyResolver = options.AnyResolver
	if options.DisablePayloadShorthand {
		ret.underlying.Metadata = map[string]interface{}{
			common.DisablePayloadShorthandMetadataKey: true,
		}
	}
	return ret, nil
}

// Marshal is the Temporal-specific equivalent of
// [github.com/gogo/protobuf/jsonpb.Marshaler.Marshal].
func (j *JSONPBMarshaler) Marshal(out io.Writer, pb proto.Message) error {
	return j.underlying.Marshal(out, pb)
}

// Marshal is the Temporal-specific equivalent of
// [github.com/gogo/protobuf/jsonpb.Marshaler.MarshalToString].
func (j *JSONPBMarshaler) MarshalToString(pb proto.Message) (string, error) {
	return j.underlying.MarshalToString(pb)
}

// JSONPBUnmarshaler is a protobuf JSON unmarshaler that supports
// Temporal-specific features. This is mostly equivalent to
// [github.com/gogo/protobuf/jsonpb.Unmarshaler].
//
// One feature is "shorthand payloads". During unmarshal when a JSON is
// encountered that cannot be converted to a traditional protobuf JSON payload
// with metadata and data, it is assumed to be "shorthand". This means the JSON
// itself is the payload and it is turned into a payload with the proper
// metadata set. If the JSON is an object with a "_protoMessageType" field, it
// is assumed to be a proto JSON payload with that field containing the
// qualified message name.
type JSONPBUnmarshaler struct{ underlying jsonpb.Unmarshaler }

// JSONPBUnmarshalerOptions is used for [NewJSONPBUnmarshaler]. Most of the
// options are copied from [github.com/gogo/protobuf/jsonpb.Unmarshaler].
type JSONPBUnmarshalerOptions struct {
	// Whether to allow messages to contain unknown fields, as opposed to
	// failing to unmarshal.
	AllowUnknownFields bool

	// A custom URL resolver to use when unmarshaling Any messages from JSON.
	// If unset, the default resolution strategy is to extract the
	// fully-qualified type name from the type URL and pass that to
	// proto.MessageType(string).
	AnyResolver gogojsonpb.AnyResolver

	// If true, this will never unmarshal from shorthand payloads. See
	// [JSONPBUnmarshaler] for more detail.
	DisablePayloadShorthand bool
}

func NewJSONPBUnmarshaler(options JSONPBUnmarshalerOptions) (*JSONPBUnmarshaler, error) {
	ret := &JSONPBUnmarshaler{}
	ret.underlying.AllowUnknownFields = options.AllowUnknownFields
	ret.underlying.AnyResolver = options.AnyResolver
	if options.DisablePayloadShorthand {
		ret.underlying.Metadata = map[string]interface{}{
			common.DisablePayloadShorthandMetadataKey: true,
		}
	}
	return ret, nil
}

// Unmarshal is the Temporal-specific equivalent of
// [github.com/gogo/protobuf/jsonpb.Unmarshaler.Unmarshal].
func (j *JSONPBUnmarshaler) Unmarshal(r io.Reader, pb proto.Message) error {
	return j.underlying.Unmarshal(r, pb)
}

// UnmarshalNext is the Temporal-specific equivalent of
// [github.com/gogo/protobuf/jsonpb.Unmarshaler.UnmarshalNext].
func (j *JSONPBUnmarshaler) UnmarshalNext(dec *json.Decoder, pb proto.Message) error {
	return j.underlying.UnmarshalNext(dec, pb)
}

// NewGRPCGatewayJSONPBMarshaler creates a new gRPC gateway marshaler for the
// given marshaler/unmarshaler pair.
func NewGRPCGatewayJSONPBMarshaler(marshaler *JSONPBMarshaler, unmarshaler *JSONPBUnmarshaler) runtime.Marshaler {
	return &temporalgateway.JSONPb{
		Marshaler:   marshaler.underlying,
		Unmarshaler: unmarshaler.underlying,
	}
}
