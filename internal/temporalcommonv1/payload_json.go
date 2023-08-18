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

package common

import (
	"bytes"
	"encoding/json"

	gogojsonpb "github.com/gogo/protobuf/jsonpb"
	jsonpb "go.temporal.io/api/internal/temporaljsonpb"
)

// !!! This file is copied from internal/temporalcommonv1 to common/v1.
// !!! DO NOT EDIT at common/v1/payload_json.go.

// Key on the marshaler metadata specifying whether shorthand is disabled.
//
// WARNING: This is internal API and should not be called externally.
const DisablePayloadShorthandMetadataKey = "__temporal_disable_payload_shorthand"

// MaybeMarshalJSONPB implements
// [go.temporal.io/api/internal/temporaljsonpb.JSONPBMaybeMarshaler.MaybeMarshalJSONPB].
//
// WARNING: This is internal API and should not be called externally.
func (p *Payloads) MaybeMarshalJSONPB(m *jsonpb.Marshaler, currIndent string) (handled bool, b []byte, err error) {
	// If this is nil, ignore
	if p == nil {
		return false, nil, nil
	}
	// If shorthand is disabled, ignore
	if disabled, _ := m.Metadata[DisablePayloadShorthandMetadataKey].(bool); disabled {
		return false, nil, nil
	}

	// We only support marshalling to shorthand if all payloads are handled or
	// there  are no payloads
	payloads := make([]interface{}, len(p.Payloads))
	for i, payload := range p.Payloads {
		// If any are not handled or there is an error, return
		if handled, payloads[i], err = payload.toJSONShorthand(); !handled || err != nil {
			return handled, nil, err
		}
	}
	// If we're indenting, use the current indent as prefix. Note, regardless of
	// m.EmitDefaults, we always use an explicit empty array here if there are no
	// values.
	if m.Indent == "" {
		b, err = json.Marshal(payloads)
	} else {
		b, err = json.MarshalIndent(payloads, currIndent, m.Indent)
	}
	return true, b, err
}

// MaybeUnmarshalJSONPB implements
// [go.temporal.io/api/internal/temporaljsonpb.JSONPBMaybeUnmarshaler.MaybeUnmarshalJSONPB].
//
// WARNING: This is internal API and should not be called externally.
func (p *Payloads) MaybeUnmarshalJSONPB(u *jsonpb.Unmarshaler, b []byte) (handled bool, err error) {
	// If this is nil, ignore (should never be)
	if p == nil {
		return false, nil
	}
	// If shorthand is disabled, ignore
	if disabled, _ := u.Metadata[DisablePayloadShorthandMetadataKey].(bool); disabled {
		return false, nil
	}
	// Try to deserialize into slice. If this fails, it is not shorthand and this
	// does not handle it. This means on invalid JSON, we let the proto JSON
	// handler fail instead of here.
	var payloadJSONs []json.RawMessage
	if json.Unmarshal(b, &payloadJSONs) != nil {
		return false, nil
	}
	// Convert each (some may be shorthand, some may not)
	p.Payloads = make([]*Payload, len(payloadJSONs))
	for i, payloadJSON := range payloadJSONs {
		p.Payloads[i] = &Payload{}
		p.Payloads[i].fromJSONMaybeShorthand(payloadJSON)
	}
	return true, nil
}

// MaybeMarshalJSONPB implements
// [go.temporal.io/api/internal/temporaljsonpb.JSONPBMaybeMarshaler.MaybeMarshalJSONPB].
//
// WARNING: This is internal API and should not be called externally.
func (p *Payload) MaybeMarshalJSONPB(m *jsonpb.Marshaler, currIndent string) (handled bool, b []byte, err error) {
	// If this is nil, ignore
	if p == nil {
		return false, nil, nil
	}
	// If shorthand is disabled, ignore
	if disabled, _ := m.Metadata[DisablePayloadShorthandMetadataKey].(bool); disabled {
		return false, nil, nil
	}
	handled, value, err := p.toJSONShorthand()
	if !handled || err != nil {
		return handled, nil, err
	}
	if m.Indent == "" {
		b, err = json.Marshal(value)
	} else {
		b, err = json.MarshalIndent(value, currIndent, m.Indent)
	}
	return true, b, err
}

// MaybeUnmarshalJSONPB implements
// [go.temporal.io/api/internal/temporaljsonpb.JSONPBMaybeUnmarshaler.MaybeUnmarshalJSONPB].
//
// WARNING: This is internal API and should not be called externally.
func (p *Payload) MaybeUnmarshalJSONPB(u *jsonpb.Unmarshaler, b []byte) (handled bool, err error) {
	// If this is nil, ignore (should never be)
	if p == nil {
		return false, nil
	}
	// If shorthand is disabled, ignore
	if disabled, _ := u.Metadata[DisablePayloadShorthandMetadataKey].(bool); disabled {
		return false, nil
	}
	// Always considered handled, unmarshaler ignored (unknown fields always
	// disallowed for non-shorthand payloads at this time)
	p.fromJSONMaybeShorthand(b)
	return true, nil
}

func (p *Payload) toJSONShorthand() (handled bool, value interface{}, err error) {
	// Only support binary null, plain JSON and proto JSON
	switch string(p.Metadata["encoding"]) {
	case "binary/null":
		// Leave value as nil
		handled = true
	case "json/plain":
		// Must only have this single metadata
		if len(p.Metadata) != 1 {
			return false, nil, nil
		}
		// We unmarshal because we may have to indent. We let this error fail the
		// marshaller.
		handled = true
		err = json.Unmarshal(p.Data, &value)
	case "json/protobuf":
		// Must have the message type and no other metadata
		msgType := string(p.Metadata["messageType"])
		if msgType == "" || len(p.Metadata) != 2 {
			return false, nil, nil
		}
		// Since this is a proto object, this must unmarshal to a object. We let
		// this error fail the marshaller.
		var valueMap map[string]interface{}
		handled = true
		err = json.Unmarshal(p.Data, &valueMap)
		// Put the message type on the object
		if valueMap != nil {
			valueMap["_protoMessageType"] = msgType
		}
		value = valueMap
	}
	return
}

func (p *Payload) fromJSONMaybeShorthand(b []byte) {
	// We need to try to deserialize into the regular payload first. If it works
	// and there is metadata _and_ data actually present (or null with a null
	// metadata encoding), we assume it's a non-shorthand payload. If it fails
	// (which it will if not an object or there is an unknown field or if
	// 'metadata' is not string + base64 or if 'data' is not base64), we assume
	// shorthand. We are ok disallowing unknown fields for payloads here even if
	// the outer unmarshaler allows them.
	if gogojsonpb.Unmarshal(bytes.NewReader(b), p) == nil && len(p.Metadata) > 0 {
		// A raw payload must either have data or a binary/null encoding
		if len(p.Data) > 0 || string(p.Metadata["encoding"]) == "binary/null" {
			return
		}
	}

	// If it's "null", set no data and just metadata
	if string(b) == "null" {
		p.Data = nil
		p.Metadata = map[string][]byte{"encoding": []byte("binary/null")}
		return
	}

	// Now that we know it is shorthand, it might be a proto JSON with a message
	// type. If it does have the message type, we need to remove it and
	// re-serialize it to data. So the quickest way to check whether it has the
	// message type is to search for the key.
	p.Data = b
	p.Metadata = map[string][]byte{"encoding": []byte("json/plain")}
	if bytes.Contains(p.Data, []byte(`"_protoMessageType"`)) {
		// Try to unmarshal into map, extract and remove key, and re-serialize
		var valueMap map[string]interface{}
		if json.Unmarshal(p.Data, &valueMap) == nil {
			if msgType, _ := valueMap["_protoMessageType"].(string); msgType != "" {
				// Now we know it's a proto JSON, so remove the key and re-serialize
				delete(valueMap, "_protoMessageType")
				// This won't error. The resulting JSON payload data may not be exactly
				// what user passed in sans message type (e.g. user may have indented or
				// did not have same field order), but that is acceptable when going
				// from shorthand to non-shorthand.
				p.Data, _ = json.Marshal(valueMap)
				p.Metadata["encoding"] = []byte("json/protobuf")
				p.Metadata["messageType"] = []byte(msgType)
			}
		}
	}
}
