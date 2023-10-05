// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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

package duration

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

// Value safely converts a protobuf duration to a go duration
func Value(d *durationpb.Duration) time.Duration {
	if d == nil {
		return 0
	}
	return d.AsDuration()
}

// MinProto returns the proto-compatible minimum duration
func MinProto(d1 *durationpb.Duration, d2 *durationpb.Duration) *durationpb.Duration {
	d1v, d2v := Value(d1), Value(d2)
	if d1v > d2v {
		return durationpb.New(d2v)
	}
	return durationpb.New(d1v)
}

// Proto returns a proto-compatible duration from a duration value
func Proto(d time.Duration) *durationpb.Duration {
	return durationpb.New(d)
}

// Ptr returns a pointer to a duration object
func Ptr(d time.Duration) *time.Duration {
	return &d
}
