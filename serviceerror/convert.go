package serviceerror

import (
	"context"
	"errors"
	"unicode/utf8"

	spb "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
	"go.temporal.io/api/failure/v1"
)

const (
	maxMessageLength       = 4000
	truncatedMessageSuffix = "... <truncated>"
)

// ToStatus converts service error to gRPC Status.
// If error is not a service error it returns status with code Unknown.
func ToStatus(err error) (st *status.Status) {
	defer func() {
		if len(st.Message()) <= maxMessageLength {
			return
		}
		spb := st.Proto()
		spb.Message = truncateMessage(spb.Message)
		st = status.FromProto(spb)
	}()

	if err == nil {
		return status.New(codes.OK, "")
	}

	if svcerr, ok := err.(ServiceError); ok {
		return svcerr.Status()
	}

	errMsg := truncateMessage(err.Error())
	// err does not implement ServiceError directly, but check if it wraps it.
	// This path does more allocation so prefer to return a ServiceError directly if possible.
	var svcerr ServiceError
	if errors.As(err, &svcerr) {
		s := svcerr.Status().Proto()
		s.Message = errMsg // don't lose the wrapped message
		return status.FromProto(s)
	}

	// Special case for context.DeadlineExceeded and context.Canceled because they can happen in unpredictable places.
	if errors.Is(err, context.DeadlineExceeded) {
		return status.New(codes.DeadlineExceeded, errMsg)
	}
	if errors.Is(err, context.Canceled) {
		return status.New(codes.Canceled, errMsg)
	}

	// Internal logic of status.Convert is:
	//   - if err is already Status or gRPC Status, then just return it (this should never happen though).
	//   - otherwise returns codes.Unknown with message from err.Error() (this might happen if some generic go error reach to this point).
	return status.Convert(err)
}

// FromStatus converts gRPC Status to service error.
func FromStatus(st *status.Status) error {
	if st == nil || st.Code() == codes.OK {
		return nil
	}

	errDetails := extractErrorDetails(st)

	// Special case: MultiOperation error can have any status code.
	if err, ok := errDetails.(*errordetails.MultiOperationExecutionFailure); ok {
		errs := make([]error, len(err.Statuses))
		for i, opStatus := range err.Statuses {
			errs[i] = FromStatus(status.FromProto(&spb.Status{
				Code:    opStatus.Code,
				Message: opStatus.Message,
				Details: opStatus.Details,
			}))
		}
		return newMultiOperationExecution(st, errs)
	}

	// If there was an error during details extraction, for example unknown message type,
	// which can happen when new error details are added and getting read by old clients,
	// then errDetails will be of type `error` with corresponding error inside.
	// This error is ignored and `serviceerror` is built using `st.Code()` only.
	switch st.Code() {
	case codes.DataLoss:
		return newDataLoss(st)
	case codes.DeadlineExceeded:
		return newDeadlineExceeded(st)
	case codes.Canceled:
		return newCanceled(st)
	case codes.Unavailable:
		switch errDetails := errDetails.(type) {
		case *errordetails.NamespaceUnavailableFailure:
			return newNamespaceUnavailable(st, errDetails)
		default:
			return newUnavailable(st)
		}
	case codes.Unimplemented:
		return newUnimplemented(st)
	case codes.Unknown:
		// Unwrap error message from unknown error.
		return errors.New(st.Message())
	case codes.Aborted:
		switch errDetails.(type) {
		case *failure.MultiOperationExecutionAborted:
			return newMultiOperationAborted(st)
		default:
			return newAborted(st)
		}
	case codes.Internal:
		switch errDetails := errDetails.(type) {
		case *errordetails.SystemWorkflowFailure:
			return newSystemWorkflow(st, errDetails)
		default:
			return newInternal(st)
		}
	case codes.NotFound:
		switch errDetails := errDetails.(type) {
		case *errordetails.NotFoundFailure:
			return newNotFound(st, errDetails)
		case *errordetails.NamespaceNotFoundFailure:
			return newNamespaceNotFound(st, errDetails)
		default:
			return newNotFound(st, nil)
		}
	case codes.InvalidArgument:
		switch errDetails := errDetails.(type) {
		case *errordetails.QueryFailedFailure:
			return newQueryFailed(st, errDetails)
		default:
			return newInvalidArgument(st)
		}
	case codes.ResourceExhausted:
		switch errDetails := errDetails.(type) {
		case *errordetails.ResourceExhaustedFailure:
			return newResourceExhausted(st, errDetails)
		default:
			return newResourceExhausted(st, nil)
		}
	case codes.AlreadyExists:
		switch errDetails := errDetails.(type) {
		case *errordetails.NamespaceAlreadyExistsFailure:
			return newNamespaceAlreadyExists(st)
		case *errordetails.WorkflowExecutionAlreadyStartedFailure:
			return newWorkflowExecutionAlreadyStarted(st, errDetails)
		case *errordetails.ActivityExecutionAlreadyStartedFailure:
			return newActivityExecutionAlreadyStarted(st, errDetails)
		case *errordetails.CancellationAlreadyRequestedFailure:
			return newCancellationAlreadyRequested(st)
		default:
			return newAlreadyExists(st)
		}
	case codes.FailedPrecondition:
		switch errDetails := errDetails.(type) {
		case *errordetails.NamespaceNotActiveFailure:
			return newNamespaceNotActive(st, errDetails)
		case *errordetails.NamespaceInvalidStateFailure:
			return newNamespaceInvalidState(st, errDetails)
		case *errordetails.ClientVersionNotSupportedFailure:
			return newClientVersionNotSupported(st, errDetails)
		case *errordetails.ServerVersionNotSupportedFailure:
			return newServerVersionNotSupported(st, errDetails)
		case *errordetails.WorkflowNotReadyFailure:
			return newWorkflowNotReady(st)
		default:
			return newFailedPrecondition(st)
		}
	case codes.PermissionDenied:
		switch errDetails := errDetails.(type) {
		case *errordetails.PermissionDeniedFailure:
			return newPermissionDenied(st, errDetails)
		default:
			return newPermissionDenied(st, nil)
		}
	case codes.OutOfRange:
		switch errDetails := errDetails.(type) {
		case *errordetails.NewerBuildExistsFailure:
			return newNewerBuildExists(st, errDetails)
		default:
			// fall through to st.Err()
		}
	// Unsupported code:
	case codes.Unauthenticated:
		// fall through to st.Err()
	}

	// `st.Code()` has unknown value (should never happen).
	// Use standard gRPC error representation "rpc error: code = %s desc = %s".
	return st.Err()
}

func extractErrorDetails(st *status.Status) any {
	details := st.Details()
	if len(details) > 0 {
		return details[0]
	}

	return nil
}

// truncateUTF8 truncates s to no more than n _bytes_, and returns a valid utf-8 string as long
// as the input is a valid utf-8 string.
// Note that truncation pays attention to codepoints only! This may truncate in the middle of a
// grapheme cluster.
func truncateUTF8(s string, n int) string {
	if len(s) <= n {
		return s
	} else if n <= 0 {
		return ""
	}
	for n > 0 && !utf8.RuneStart(s[n]) {
		n--
	}
	return s[:n]
}

func truncateMessage(s string) string {
	if len(s) <= maxMessageLength {
		return s
	}
	return truncateUTF8(s, maxMessageLength-len(truncatedMessageSuffix)) + truncatedMessageSuffix
}
