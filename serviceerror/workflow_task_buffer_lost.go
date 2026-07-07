package serviceerror

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// WorkflowTaskBufferLost represents a lost buffer during processing of a pagination requestt.
	WorkflowTaskBufferLost struct {
		Message string
		st      *status.Status
	}
)

// NewWorkflowTaskBufferLost returns a new WorkflowTaskBufferLost error.
func NewWorkflowTaskBufferLost(message string) error {
	return &WorkflowTaskBufferLost{
		Message: message,
	}
}

// NewWorkflowTaskBufferLostf returns a new WorkflowTaskBufferLost error with a formatted message.
func NewWorkflowTaskBufferLostf(format string, args ...any) error {
	return &WorkflowTaskBufferLost{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns string message.
func (e *WorkflowTaskBufferLost) Error() string {
	return e.Message
}

func (e *WorkflowTaskBufferLost) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.Aborted, e.Message)
	st, _ = st.WithDetails(
		&errordetails.WorkflowTaskCompletionBufferLostFailure{},
	)
	return st
}

func newWorkflowTaskBufferLost(st *status.Status) error {
	return &WorkflowTaskBufferLost{
		Message: st.Message(),
		st:      st,
	}
}
