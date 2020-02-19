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

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/temporal-proto/errordetails"
)

// FromStatus converts gRPC status to service error.
func FromStatus(st *status.Status) error {
	if st == nil || st.Code() == codes.OK {
		return nil
	}

	// Simple case. Code to error is one to one mapping and there is no failure.
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
	case codes.Unknown:
		// Unwrap error message from unknown error.
		return errors.New(st.Message())
	// Unsupported codes.
	case codes.Canceled,
		codes.OutOfRange,
		codes.Unimplemented,
		codes.Unavailable,
		codes.Unauthenticated:
		// Use standard gRPC error representation for unsupported codes.
		return st.Err()
	}

	// Extract failure once to optimize performance.
	failure := extractFailure(st)
	switch st.Code() {
	case codes.InvalidArgument:
		if failure == nil {
			return newInvalidArgument(st)
		}
		switch f := failure.(type) {
		case *errordetails.QueryFailedFailure:
			return newQueryFailed(st)
		case *errordetails.CurrentBranchChangedFailure:
			return newCurrentBranchChanged(st, f)
		}
	case codes.AlreadyExists:
		switch f := failure.(type) {
		case *errordetails.DomainAlreadyExistsFailure:
			return newDomainAlreadyExists(st)
		case *errordetails.WorkflowExecutionAlreadyStartedFailure:
			return newWorkflowExecutionAlreadyStarted(st, f)
		case *errordetails.CancellationAlreadyRequestedFailure:
			return newCancellationAlreadyRequested(st)
		case *errordetails.EventAlreadyStartedFailure:
			return newEventAlreadyStarted(st)
		}
	case codes.FailedPrecondition:
		switch f := failure.(type) {
		case *errordetails.DomainNotActiveFailure:
			return newDomainNotActive(st, f)
		case *errordetails.ClientVersionNotSupportedFailure:
			return newClientVersionNotSupported(st, f)
		}
	case codes.Aborted:
		switch f := failure.(type) {
		case *errordetails.ShardOwnershipLostFailure:
			return newShardOwnershipLost(st, f)
		case *errordetails.RetryTaskFailure:
			return newRetryTask(st, f)
		case *errordetails.RetryTaskV2Failure:
			return newRetryTaskV2(st, f)
		}
	}

	// Code suppose to have failure but it didn't (or failure has wrong type).
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
