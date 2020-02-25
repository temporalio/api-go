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
	"errors"

	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"go.temporal.io/temporal-proto/failure"
)
// ToStatus converts service error to gogo gRPC status.
// If error is not a service error it returns status with code Unknown.
func ToStatus(err error) *status.Status {
	if err == nil{
		return status.New(codes.OK, "")
	}

	if svcerr, ok := err.(ServiceError); ok {
		return svcerr.status()
	}

	// If err is gogo Status or gRPC status return it (this should never happen though).
	// Otherwise returns codes.Unknown with message from err.Error()
	// This should never happen and this check is a safety net.
	return status.Convert(err)
}

// FromStatus converts gogo gRPC status to service error.
func FromStatus(st *status.Status) error {
	if st == nil || st.Code() == codes.OK {
		return nil
	}

	// Simple case. Code to serviceerror is one to one mapping and there is no failure.
	switch st.Code() {
	case codes.Internal:
		return newInternal(st)
	case codes.DataLoss:
		return newDataLoss(st)
	case codes.NotFound:
		return newNotFound(st)
	case codes.ResourceExhausted:
		return newResourceExhausted(st)
	case codes.PermissionDenied:
		return newPermissionDenied(st)
	case codes.DeadlineExceeded:
		return newDeadlineExceeded(st)
	case codes.Unavailable:
		return newUnavailable(st)
	case codes.Unknown:
		// Unwrap error message from unknown error.
		return errors.New(st.Message())
	// Unsupported codes.
	case codes.Canceled,
		codes.OutOfRange,
		codes.Unimplemented,
		codes.Unauthenticated:
		// Use standard gRPC error representation for unsupported codes.
		return st.Err()
	}

	// Extract failure once to optimize performance.
	f := extractFailure(st)
	switch st.Code() {
	case codes.InvalidArgument:
		if f == nil {
			return newInvalidArgument(st)
		}
		switch f := f.(type) {
		case *failure.QueryFailed:
			return newQueryFailed(st)
		case *failure.CurrentBranchChanged:
			return newCurrentBranchChanged(st, f)
		}
	case codes.AlreadyExists:
		switch f := f.(type) {
		case *failure.DomainAlreadyExists:
			return newDomainAlreadyExists(st)
		case *failure.WorkflowExecutionAlreadyStarted:
			return newWorkflowExecutionAlreadyStarted(st, f)
		case *failure.CancellationAlreadyRequested:
			return newCancellationAlreadyRequested(st)
		case *failure.EventAlreadyStarted:
			return newEventAlreadyStarted(st)
		}
	case codes.FailedPrecondition:
		switch f := f.(type) {
		case *failure.DomainNotActive:
			return newDomainNotActive(st, f)
		case *failure.ClientVersionNotSupported:
			return newClientVersionNotSupported(st, f)
		}
	case codes.Aborted:
		switch f := f.(type) {
		case *failure.ShardOwnershipLost:
			return newShardOwnershipLost(st, f)
		case *failure.RetryTask:
			return newRetryTask(st, f)
		case *failure.RetryTaskV2:
			return newRetryTaskV2(st, f)
		}
	}

	// Code should have failure but it didn't (or failure is of an wrong type).
	// Use standard gRPC error representation here also.
	return st.Err()
}

func extractFailure(st *status.Status) interface{} {
	details := st.Details()
	if len(details) > 0 {
		return details[0]
	}

	return nil
}
