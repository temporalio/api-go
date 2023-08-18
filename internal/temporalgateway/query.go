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

// This code taken from the bottom of
// https://github.com/grpc-ecosystem/grpc-gateway/blob/v1.16.0/runtime/query.go.

package temporalgateway

import (
	"reflect"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

var (
	convFromType = map[reflect.Kind]reflect.Value{
		reflect.String:  reflect.ValueOf(runtime.String),
		reflect.Bool:    reflect.ValueOf(runtime.Bool),
		reflect.Float64: reflect.ValueOf(runtime.Float64),
		reflect.Float32: reflect.ValueOf(runtime.Float32),
		reflect.Int64:   reflect.ValueOf(runtime.Int64),
		reflect.Int32:   reflect.ValueOf(runtime.Int32),
		reflect.Uint64:  reflect.ValueOf(runtime.Uint64),
		reflect.Uint32:  reflect.ValueOf(runtime.Uint32),
		reflect.Slice:   reflect.ValueOf(runtime.Bytes),
	}
)
