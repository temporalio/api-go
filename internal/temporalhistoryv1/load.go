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

// !!! This file is copied from internal/temporalhistoryv1 to history/v1.
// !!! DO NOT EDIT at history/v1/load.go.

package history

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type fixSpec struct {
	TypeName protoreflect.FullName
	Enums    map[string]string
	// Paths that correspond to other fixable types
	Messages map[string]protoreflect.FullName
}

type fixRegistry map[protoreflect.FullName]fixSpec

// Add a type to the registry, returning whether or not it already existed
func (f fixRegistry) Add(t protoreflect.FullName) (fixSpec, bool) {
	if _, exists := f[t]; exists {
		return fixSpec{}, true
	}

	spec := fixSpec{
		TypeName: t,
		Enums:    make(map[string]string),
		Messages: make(map[string]protoreflect.FullName),
	}
	f[t] = spec
	return spec, false
}

func (f fixRegistry) Lookup(spec fixSpec, name string) (fixSpec, bool) {
	pType, found := spec.Messages[name]
	if !found {
		return fixSpec{}, false
	}
	tSpec, found := f[pType]
	return tSpec, found
}

// Not initialized until prepareFixes.Do(...) has been called
var failureFullName string
var registry fixRegistry
var historySpec fixSpec
var prepareFixes sync.Once

// Recursively walk the protoreflect descriptors for the provided object,
// collecting all enum paths for later fixing. We keep a global map of fixes
// per proto type
func findFixes(md protoreflect.MessageDescriptor, spec fixSpec) {
	for i := 0; i < md.Fields().Len(); i++ {
		cd := md.Fields().Get(i)
		switch cd.Kind() {
		case protoreflect.EnumKind:
			ed := cd.Enum()
			spec.Enums[cd.JSONName()] = strcase.ToScreamingSnake(string(ed.Name()))
		case protoreflect.MessageKind:
			spec.Messages[cd.JSONName()] = cd.FullName()
			mSpec, exists := registry.Add(cd.FullName())
			if exists {
				continue
			}
			findFixes(cd.Message(), mSpec)
		}
	}
}

func fixupString(val string, name string, spec fixSpec) string {
	if typePrefix, ok := spec.Enums[name]; ok {
		newVal := strcase.ToScreamingSnake(val)
		if !strings.HasPrefix(newVal, typePrefix) {
			return fmt.Sprintf("%s_%s", typePrefix, newVal)
		}
	}
	return val
}

func fixupMsg(msg map[string]interface{}, spec fixSpec) {
	for k, v := range msg {
		switch vv := v.(type) {
		case string:
			msg[k] = fixupString(vv, k, spec)
		case map[string]interface{}:
			spec, found := registry.Lookup(spec, k)
			if !found {
				continue
			}
			fixupMsg(vv, spec)
		case []interface{}:
			spec, found := registry.Lookup(spec, k)
			if !found {
				continue
			}
			for i := 0; i < len(vv); i++ {
				switch vvv := vv[i].(type) {
				case string:
					vv[i] = fixupString(vvv, k, spec)
				case map[string]interface{}:
					fixupMsg(vvv, spec)
				}
			}
		}
	}
}

// LoadFromJSON deserializes history from a reader of JSON bytes. This does
// not close the reader if it is closeable. This function is compatible both
// with the "correct" SCREAMING_SNAKE enums of protojson as well as the old
// PascalCase enums of Temporal's older history exports.
func LoadFromJSON(r io.Reader) (*History, error) {
	prepareFixes.Do(func() {
		var hist History
		registry = make(fixRegistry)
		desc := hist.ProtoReflect().Type().Descriptor()
		historySpec, _ = registry.Add(desc.FullName())
		findFixes(desc, historySpec)

		// Iteratively prune empty specs from the registry.
		removed := make(map[protoreflect.FullName]struct{})
		changing := true
		for changing {
			changing = false
			for typ, spec := range registry {
				for k, mt := range spec.Messages {
					if _, rmed := removed[mt]; rmed {
						delete(spec.Messages, k)
					}
				}
				if len(spec.Enums) == 0 && len(spec.Messages) == 0 {
					changing = true
					removed[typ] = struct{}{}
					delete(registry, typ)
				}
			}
		}
	})

	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	msg := make(map[string]interface{})
	if err := json.Unmarshal(bs, &msg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal history: %w", err)
	}

	fixupMsg(msg, historySpec)

	out, err := json.Marshal(msg)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fixed history: %w", err)
	}

	opts := protojson.UnmarshalOptions{
		// Ignore unknown fields because if the histroy was generated with a different version of the proto
		// fields may have been added/removed.
		DiscardUnknown: true,
	}

	var hist History
	if err := opts.Unmarshal(out, &hist); err != nil {
		return nil, fmt.Errorf("failed to fix history: %w", err)
	}

	return &hist, nil
}
