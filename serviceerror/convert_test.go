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
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	rpc "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"

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
	})

	err1 := serviceerror.FromStatus(st1)
	require.IsType(t, &serviceerror.NotFound{}, err1)
	require.Equal(t, codes.NotFound, err1.(*serviceerror.NotFound).Status().Code())
	require.Equal(t, "Not found.", err1.(*serviceerror.NotFound).Message)
}

func TestToStatus_UnknownErrorDetails(t *testing.T) {
	anyd, err := anypb.New(durationpb.New(time.Duration(time.Second)))
	if err != nil {
		t.Fatalf("Failed to create any out of duration: %s", err)
	}
	err1 := status.ErrorProto(&rpc.Status{
		Code:    int32(codes.NotFound),
		Message: "Not found.",
		Details: []*anypb.Any{anyd},
	})

	st1 := serviceerror.ToStatus(err1)
	require.Equal(t, codes.NotFound, st1.Code())
	require.Equal(t, "Not found.", st1.Message())
	require.Len(t, st1.Details(), 1)
	require.Equal(t, "type.googleapis.com/google.protobuf.Duration", st1.Proto().Details[0].TypeUrl)
}

func TestToStatus_NotServiceError(t *testing.T) {
	err1 := errors.New("some error")
	st1 := serviceerror.ToStatus(err1)
	require.Equal(t, codes.Unknown, st1.Code())
	require.Equal(t, "some error", st1.Message())
	require.Len(t, st1.Details(), 0)
}

func TestFromWrapped(t *testing.T) {
	err := &serviceerror.PermissionDenied{
		Message: "x is not allowed",
		Reason:  "arbitrary reason",
	}
	wrapped := fmt.Errorf("wrapped error: %w", err)
	s := serviceerror.ToStatus(wrapped)
	require.Equal(t, codes.PermissionDenied, s.Code())
	require.Equal(t, "wrapped error: x is not allowed", s.Message())
	require.True(t, proto.Equal(
		&errordetails.PermissionDeniedFailure{Reason: "arbitrary reason"},
		s.Details()[0].(*errordetails.PermissionDeniedFailure)))
}
