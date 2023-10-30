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

package temporalproto

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// JSONUnmarshaller unmarshals proto structs from either the old temporal-style JSON with camelCase enums
// or the protojson canonical format (SCREAMING_SNAKE enums).
type JSONUnmarshaller struct {
	DiscardUnknown bool
}

// JSONDecoders decodes proto structs from either the old temporal-style JSON with camelCase enums
// or the protojson canonical format (SCREAMING_SNAKE enums).
type JSONDecoder struct {
	unmarshaller JSONUnmarshaller
	json         *json.Decoder
}

type fixSpec struct {
	TypeName protoreflect.FullName
	Enums    map[string]string
	// Paths that correspond to other fixable types
	Messages map[string]protoreflect.FullName
	Creator  uint64
	Ready    int32
	Wait     chan struct{}
}

type fixRegistry struct {
	sync.RWMutex
	nextToken uint64
	specs     map[protoreflect.FullName]*fixSpec
}

var failureFullName string
var registry = &fixRegistry{
	specs: make(map[protoreflect.FullName]*fixSpec),
}

func (f *fixSpec) Empty() bool {
	return len(f.Enums) == 0 && len(f.Messages) == 0
}

// Ensure that the specified proto type is fixable
func (f *fixRegistry) Ensure(desc protoreflect.MessageDescriptor) *fixSpec {
	if spec, ready := registry.Ready(desc.FullName()); ready {
		return spec
	}
	worker := atomic.AddUint64(&f.nextToken, 1)
	spec := registry.Add(worker, desc)
	return spec
}

func (f *fixRegistry) ensureWorker(worker uint64, desc protoreflect.MessageDescriptor) *fixSpec {
	if spec, ready := registry.Ready(desc.FullName()); ready {
		return spec
	}
	return registry.Add(worker, desc)
}

func (f *fixRegistry) Ready(t protoreflect.FullName) (*fixSpec, bool) {
	f.RLock()
	defer f.RUnlock()
	spec, exists := f.specs[t]
	if !exists {
		return nil, false
	}

	// We don't actually need an atomic read here as sloppiness is ok, but go requires it.
	if atomic.LoadInt32(&spec.Ready) == 1 {
		return spec, true
	}
	return nil, false
}

// Add a type to the registry, returning whether or not it already existed
func (f *fixRegistry) Add(worker uint64, desc protoreflect.MessageDescriptor) *fixSpec {
	f.Lock()
	if spec, exists := f.specs[desc.FullName()]; exists {
		f.Unlock()
		if spec.Creator == worker {
			// we can safely return the in-progress spec
			return spec
		}
		// wait for completion
		<-spec.Wait
		return spec
	}

	spec := &fixSpec{
		TypeName: desc.FullName(),
		Enums:    make(map[string]string),
		Messages: make(map[string]protoreflect.FullName),
		Creator:  worker,
		Wait:     make(chan struct{}),
	}
	f.specs[desc.FullName()] = spec
	f.Unlock()

	f.findFixes(worker, desc, spec)
	atomic.StoreInt32(&spec.Ready, 1)
	close(spec.Wait)
	return spec
}

func (f *fixRegistry) Lookup(spec *fixSpec, name string) (*fixSpec, bool) {
	f.RLock()
	defer f.RUnlock()

	pType, found := spec.Messages[name]
	if !found {
		return nil, false
	}
	tSpec, found := f.specs[pType]
	if found {
		if atomic.LoadInt32(&tSpec.Ready) != 1 {
			<-tSpec.Wait
		}
	}
	return tSpec, found
}

// Recursively walk the protoreflect descriptors for the provided object,
// collecting all enum paths for later fixing. We keep a global map of fixes
// per proto type
func (f *fixRegistry) findFixes(worker uint64, md protoreflect.MessageDescriptor, spec *fixSpec) {
	for i := 0; i < md.Fields().Len(); i++ {
		fd := md.Fields().Get(i)
		switch fd.Kind() {
		case protoreflect.EnumKind:
			ed := fd.Enum()
			spec.Enums[fd.JSONName()] = strcase.ToScreamingSnake(string(ed.Name()))
		case protoreflect.MessageKind:
			fmd := fd.Message()
			spec.Messages[fd.JSONName()] = fmd.FullName()
			f.ensureWorker(worker, fmd)
		}
	}
}

func fixupString(val string, name string, spec *fixSpec) string {
	if typePrefix, ok := spec.Enums[name]; ok {
		newVal := strcase.ToScreamingSnake(val)
		if !strings.HasPrefix(newVal, typePrefix) {
			return fmt.Sprintf("%s_%s", typePrefix, newVal)
		}
	}
	return val
}

func fixupMsg(msg map[string]interface{}, spec *fixSpec) {
	for k, v := range msg {
		switch vv := v.(type) {
		case string:
			msg[k] = fixupString(vv, k, spec)
		case map[string]interface{}:
			spec, found := registry.Lookup(spec, k)
			if !found || spec.Empty() {
				continue
			}
			fixupMsg(vv, spec)
		case []interface{}:
			ispec, found := registry.Lookup(spec, k)
			if !found || spec.Empty() {
				continue
			}
			for i := 0; i < len(vv); i++ {
				switch vvv := vv[i].(type) {
				case string:
					vv[i] = fixupString(vvv, k, ispec)
				case map[string]interface{}:
					fixupMsg(vvv, ispec)
				}
			}
		}
	}
}

func (j JSONUnmarshaller) Unmarshal(bs []byte, m proto.Message) error {
	spec := registry.Ensure(m.ProtoReflect().Descriptor())

	var msg map[string]any
	if err := json.Unmarshal(bs, &msg); err != nil {
		return err
	}

	fixupMsg(msg, spec)

	out, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal fixed history: %w", err)
	}

	opts := protojson.UnmarshalOptions{
		DiscardUnknown: j.DiscardUnknown,
	}

	return opts.Unmarshal(out, m)
}

// Unmarshal a single proto object from the provided slice of bytes. This function
// is compatible both with the "correct" SCREAMING_SNAKE enums of protojson as well
// as the PascalCase enums of earlier releases.
//
// This does not support decoding slices of proto objects. To do that use the Decoder in a loop.
func UnmarshalJSON(bs []byte, m proto.Message) error {
	return JSONUnmarshaller{}.Unmarshal(bs, m)
}

func NewJSONDecoder(r io.Reader, discardUnknown bool) *JSONDecoder {
	return &JSONDecoder{
		json: json.NewDecoder(r),
		unmarshaller: JSONUnmarshaller{
			DiscardUnknown: discardUnknown,
		},
	}
}

// Decode deserializes a protobuf message from a reader of JSON bytes. This does
// not close the reader if it is closeable. This function is compatible both
// with the "correct" SCREAMING_SNAKE enums of protojson as well as the old
// PascalCase enums of Temporal's older history exports.
func (dec *JSONDecoder) Decode(m proto.Message) error {
	var obj json.RawMessage
	if err := dec.json.Decode(&obj); err != nil {
		return err
	}

	return dec.unmarshaller.Unmarshal(obj, m)
}
