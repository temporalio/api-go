package serviceerror

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// CallbackExecutionAlreadyStarted represents the error arising when trying to start a callback already in progress.
	CallbackExecutionAlreadyStarted struct {
		Message        string
		StartRequestId string
		RunId          string
		st             *status.Status
	}
)

// NewCallbackExecutionAlreadyStarted returns new CallbackExecutionAlreadyStarted error.
func NewCallbackExecutionAlreadyStarted(message, startRequestId, runId string) error {
	return &CallbackExecutionAlreadyStarted{
		Message:        message,
		StartRequestId: startRequestId,
		RunId:          runId,
	}
}

// NewCallbackExecutionAlreadyStartedf returns new CallbackExecutionAlreadyStarted error with formatted message.
func NewCallbackExecutionAlreadyStartedf(startRequestId, runId, format string, args ...any) error {
	return &CallbackExecutionAlreadyStarted{
		Message:        fmt.Sprintf(format, args...),
		StartRequestId: startRequestId,
		RunId:          runId,
	}
}

// Error returns string message.
func (e *CallbackExecutionAlreadyStarted) Error() string {
	return e.Message
}

func (e *CallbackExecutionAlreadyStarted) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.AlreadyExists, e.Message)
	st, _ = st.WithDetails(
		&errordetails.CallbackExecutionAlreadyStartedFailure{
			StartRequestId: e.StartRequestId,
			RunId:          e.RunId,
		},
	)
	return st
}

func newCallbackExecutionAlreadyStarted(st *status.Status, errDetails *errordetails.CallbackExecutionAlreadyStartedFailure) error {
	return &CallbackExecutionAlreadyStarted{
		Message:        st.Message(),
		StartRequestId: errDetails.GetStartRequestId(),
		RunId:          errDetails.GetRunId(),
		st:             st,
	}
}
