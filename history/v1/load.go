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

const pathSep = "."

type enumFix struct {
	Path       string
	TypePrefix string
}

// Map of JSON path to SCREAMING_CASE type prefix for proto enums
// Not initialized until prepareFixes.Do(...) has been called
var historyPathsToFix = make(map[string]string)
var prepareFixes sync.Once

// Recursively walk the protoreflect descriptions for the History object
// collecting all enum paths for later fixing
func collectPaths(path []string, fd protoreflect.FieldDescriptor) {
	if fe := fd.Enum(); fe != nil {
		for i := 0; i < fe.Values().Len(); i++ {
			fullPath := strings.Join(append(path, fd.JSONName()), pathSep)
			historyPathsToFix[fullPath] = strcase.ToScreamingSnake(string(fe.Name()))
		}
	}

	if fm := fd.Message(); fm != nil {
		path = append(path, fd.JSONName())
		for i := 0; i < fm.Fields().Len(); i++ {
			cd := fm.Fields().Get(i)
			// Avoid blowing the stack when a proto may contain itself
			// I'm looking at you temporal.api.failure.v1.Failure...
			if cd.Name() == fd.Name() {
				continue
			}
			collectPaths(path, fm.Fields().Get(i))
		}

	}
}

func visitAndFix(val interface{}, path string, fixes map[string]string) {
	switch v := val.(type) {
	case []interface{}:
		for i := range v {
			if vv, ok := v[i].(map[string]interface{}); ok {
				visitAndFix(vv, path, fixes)
			}
		}
	case map[string]interface{}:
		for kk, vv := range v {
			fullPath := fmt.Sprintf("%s.%s", path, kk)
			switch vvv := vv.(type) {
			case string:
				if typePrefix, ok := fixes[fullPath]; ok {
					newVal := strcase.ToScreamingSnake(vvv)
					if !strings.HasPrefix(newVal, typePrefix) {
						newVal = fmt.Sprintf("%s_%s", typePrefix, newVal)
					}
					v[kk] = newVal
				}
			case []interface{}:
				visitAndFix(vvv, fullPath, fixes)
			case map[string]interface{}:
				visitAndFix(vvv, fullPath, fixes)
			}
		}
	}
}

type LoadOptions struct {
	LastEventID int64
}

type LoadOption = func(*LoadOptions)

func WithLastEventID(i int64) LoadOption {
	return func(opts *LoadOptions) {
		opts.LastEventID = i
	}
}

// LoadFromJSON deserializes history from a reader of JSON bytes. This does
// not close the reader if it is closeable. This function is compatible both
// with the "correct" SCREAMING_SNAKE enums of protojson as well as the old
// PascalCase enums of Temporal's older history exports.
func LoadFromJSON(r io.Reader, options ...LoadOption) (*History, error) {
	var loadOpts LoadOptions
	for i := range options {
		options[i](&loadOpts)
	}

	prepareFixes.Do(func() {
		var hist History

		desc := hist.ProtoReflect().Type().Descriptor()
		for i := 0; i < desc.Fields().Len(); i++ {
			collectPaths(nil, desc.Fields().Get(i))
		}
	})

	bs, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("failed to read input: %w", err)
	}

	blob := make(map[string]interface{})
	if err := json.Unmarshal(bs, &blob); err != nil {
		return nil, fmt.Errorf("failed to unmarshal history: %w", err)
	}

	for k, v := range blob {
		visitAndFix(v, k, historyPathsToFix)
	}

	out, err := json.Marshal(blob)
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

	// If there is a last event ID, slice the rest off
	if loadOpts.LastEventID > 0 {
		for i, event := range hist.Events {
			if event.EventId == loadOpts.LastEventID {
				// Inclusive
				hist.Events = hist.Events[:i+1]
				break
			}
		}
	}

	return &hist, nil
}
