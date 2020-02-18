// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serviceerror

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

)

// FromStatus converts gRPC status to go error.
func FromStatus(st *status.Status) error {
	if st == nil || st.Code() == codes.OK {
		return nil
	}

	//switch st.Code() {
	//case codes.Canceled:
	//case codes.Unknown:
	//case codes.InvalidArgument:
	//case codes.DeadlineExceeded:
	//case codes.NotFound:
	//	err, _ := notFoundErrorFromStatus(st)
	//	return err
	//case codes.AlreadyExists:
	//case codes.PermissionDenied:
	//case codes.ResourceExhausted:
	//case codes.FailedPrecondition:
	//case codes.Aborted:
	//	if err, ok := shardOwnershipLostErrorFromStatus(st); ok {
	//		return err
	//	}
	//case codes.OutOfRange:
	//case codes.Unimplemented:
	//case codes.Internal:
	//case codes.Unavailable:
	//case codes.DataLoss:
	//case codes.Unauthenticated:
	//}

	return st.Err()
}
