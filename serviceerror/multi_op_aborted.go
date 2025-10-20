package serviceerror

import (
	"errors"
	"fmt"

	failurepb "go.temporal.io/api/failure/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MultiOperationAborted represents an aborted operation from a MultiOperationExecution.
type MultiOperationAborted struct {
	Message string
	st      *status.Status
}

// NewMultiOperationAborted returns MultiOperationAborted.
func NewMultiOperationAborted(message string) error {
	return &MultiOperationAborted{
		Message: message,
	}
}

// NewMultiOperationAbortedf returns MultiOperationAborted with formatted message.
func NewMultiOperationAbortedf(format string, args ...any) error {
	return &MultiOperationAborted{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns string message.
func (e MultiOperationAborted) Error() string {
	return e.Message
}

func (e MultiOperationAborted) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.Aborted, e.Error())
	st, _ = st.WithDetails(&failurepb.MultiOperationExecutionAborted{})
	return st
}

func newMultiOperationAborted(st *status.Status) error {
	return &MultiOperationAborted{
		Message: st.Message(),
		st:      st,
	}
}

// IsMultiOperationAborted returns whether any error in the provided error's chain is a
// MultiOperationAborted error.
func IsMultiOperationAborted(err error) bool {
	var serr *MultiOperationAborted
	return errors.As(err, &serr)
}
