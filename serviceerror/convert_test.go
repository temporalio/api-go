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

package serviceerror_test

import (
	"errors"
	"testing"

	"github.com/gogo/googleapis/google/rpc"
	"github.com/gogo/protobuf/types"
	"github.com/gogo/status"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"

	"go.temporal.io/api/errordetails/v1"
	"go.temporal.io/api/serviceerror"
)

func TestFromStatus_NotFound(t *testing.T) {
	var err error
	st1 := status.New(codes.NotFound, "Not found.")
	err1 := serviceerror.FromStatus(st1)
	require.IsType(t, &serviceerror.NotFound{}, err1)
	require.Equal(t, codes.NotFound, err1.(*serviceerror.NotFound).Status().Code())
	require.Equal(t, "Not found.", err1.(*serviceerror.NotFound).Message)

	st2 := status.New(codes.NotFound, "Not found.")
	st2, err = st1.WithDetails(&errordetails.NamespaceNotFoundFailure{
		Namespace: "test-ns",
	})
	require.NoError(t, err)
	err2 := serviceerror.FromStatus(st2)
	require.IsType(t, &serviceerror.NamespaceNotFound{}, err2)
	require.Equal(t, codes.NotFound, err2.(*serviceerror.NamespaceNotFound).Status().Code())
	require.Equal(t, "Not found.", err2.(*serviceerror.NamespaceNotFound).Message)
	require.Equal(t, "test-ns", err2.(*serviceerror.NamespaceNotFound).Namespace)
}

func TestFromStatus_UnknownErrorDetails(t *testing.T) {
	st1 := status.FromProto(&rpc.Status{
		Code:    int32(codes.NotFound),
		Message: "Not found.",
		Details: []*types.Any{{TypeUrl: "type.googleapis.com/some.unknown.Type"}},
	})

	err1 := serviceerror.FromStatus(st1)
	require.IsType(t, &serviceerror.NotFound{}, err1)
	require.Equal(t, codes.NotFound, err1.(*serviceerror.NotFound).Status().Code())
	require.Equal(t, "Not found.", err1.(*serviceerror.NotFound).Message)
}

func TestToStatus_UnknownErrorDetails(t *testing.T) {
	err1 := status.ErrorProto(&rpc.Status{
		Code:    int32(codes.NotFound),
		Message: "Not found.",
		Details: []*types.Any{{TypeUrl: "type.googleapis.com/some.unknown.Type"}},
	})

	st1 := serviceerror.ToStatus(err1)
	require.Equal(t, codes.NotFound, st1.Code())
	require.Equal(t, "Not found.", st1.Message())
	require.Len(t, st1.Details(), 1)
	require.Equal(t, "type.googleapis.com/some.unknown.Type", st1.Proto().Details[0].TypeUrl)
}

func TestToStatus_NotServiceError(t *testing.T) {
	err1 := errors.New("some error")
	st1 := serviceerror.ToStatus(err1)
	require.Equal(t, codes.Unknown, st1.Code())
	require.Equal(t, "some error", st1.Message())
	require.Len(t, st1.Details(), 0)
}
