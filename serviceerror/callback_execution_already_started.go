package serviceerror

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

// CallbackExecutionAlreadyStarted represents the error arising when trying to start a callback already in progress.
type CallbackExecutionAlreadyStarted struct {
	Message        string
	StartRequestID string
	RunID          string
	CallbackID     string
	st             *status.Status
}

// NewCallbackExecutionAlreadyStarted returns new CallbackExecutionAlreadyStarted error.
func NewCallbackExecutionAlreadyStarted(message, startRequestID, runID, callbackID string) error {
	return &CallbackExecutionAlreadyStarted{
		Message:        message,
		StartRequestID: startRequestID,
		RunID:          runID,
		CallbackID:     callbackID,
	}
}

// NewCallbackExecutionAlreadyStartedf returns new CallbackExecutionAlreadyStarted error with formatted message.
func NewCallbackExecutionAlreadyStartedf(startRequestID, runID, callbackID, format string, args ...any) error {
	return &CallbackExecutionAlreadyStarted{
		Message:        fmt.Sprintf(format, args...),
		StartRequestID: startRequestID,
		RunID:          runID,
		CallbackID:     callbackID,
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
			StartRequestId: e.StartRequestID,
			RunId:          e.RunID,
			CallbackId:     e.CallbackID,
		},
	)
	return st
}

func newCallbackExecutionAlreadyStarted(st *status.Status, errDetails *errordetails.CallbackExecutionAlreadyStartedFailure) error {
	return &CallbackExecutionAlreadyStarted{
		Message:        st.Message(),
		StartRequestID: errDetails.GetStartRequestId(),
		RunID:          errDetails.GetRunId(),
		CallbackID:     errDetails.GetCallbackId(),
		st:             st,
	}
}
