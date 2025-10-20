package serviceerror

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// WorkflowNotReady represents workflow state is not ready to handle the request error.
	WorkflowNotReady struct {
		Message string
		st      *status.Status
	}
)

// NewWorkflowNotReady returns new WorkflowNotReady
func NewWorkflowNotReady(message string) error {
	return &WorkflowNotReady{
		Message: message,
	}
}

// NewWorkflowNotReadyf returns new WorkflowNotReady error with formatted message.
func NewWorkflowNotReadyf(format string, args ...any) error {
	return &WorkflowNotReady{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns string message.
func (e *WorkflowNotReady) Error() string {
	return e.Message
}

func (e *WorkflowNotReady) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.FailedPrecondition, e.Message)
	st, _ = st.WithDetails(
		&errordetails.WorkflowNotReadyFailure{},
	)
	return st
}

func newWorkflowNotReady(st *status.Status) error {
	return &WorkflowNotReady{
		Message: st.Message(),
		st:      st,
	}
}

// IsWorkflowNotReady returns whether any error in the provided error's chain is a
// WorkflowNotReady error.
func IsWorkflowNotReady(err error) bool {
	var serr *WorkflowNotReady
	return errors.As(err, &serr)
}
