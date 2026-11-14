package serviceerror

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// WorkflowTaskCompletionBufferLost represents a lost buffer during processing of a
	// paginated workflow task completion request. Clients should restart pagination by
	// resending all pages from page 0 with the same task token.
	WorkflowTaskCompletionBufferLost struct {
		Message string
		st      *status.Status
	}
)

// NewWorkflowTaskCompletionBufferLost returns a new WorkflowTaskCompletionBufferLost error.
func NewWorkflowTaskCompletionBufferLost(message string) error {
	return &WorkflowTaskCompletionBufferLost{
		Message: message,
	}
}

// NewWorkflowTaskCompletionBufferLostf returns a new WorkflowTaskCompletionBufferLost error with a formatted message.
func NewWorkflowTaskCompletionBufferLostf(format string, args ...any) error {
	return &WorkflowTaskCompletionBufferLost{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns string message.
func (e *WorkflowTaskCompletionBufferLost) Error() string {
	return e.Message
}

func (e *WorkflowTaskCompletionBufferLost) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.Aborted, e.Message)
	st, _ = st.WithDetails(
		&errordetails.WorkflowTaskCompletionBufferLostFailure{},
	)
	return st
}

func newWorkflowTaskCompletionBufferLost(st *status.Status) error {
	return &WorkflowTaskCompletionBufferLost{
		Message: st.Message(),
		st:      st,
	}
}
