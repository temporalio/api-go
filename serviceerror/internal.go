package serviceerror

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type (
	// Internal represents internal error.
	Internal struct {
		Message string
		st      *status.Status
	}
)

// NewInternal returns new Internal error.
func NewInternal(message string) error {
	return &Internal{
		Message: message,
	}
}

// NewInternalf returns new Internal error with formatted message.
func NewInternalf(format string, args ...any) error {
	return &Internal{
		Message: fmt.Sprintf(format, args...),
	}
}

// Error returns string message.
func (e *Internal) Error() string {
	return e.Message
}

func (e *Internal) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	return status.New(codes.Internal, e.Message)
}

func newInternal(st *status.Status) error {
	return &Internal{
		Message: st.Message(),
		st:      st,
	}
}

// IsInternal returns whether any error in the provided error's chain is an
// Internal error.
func IsInternal(err error) bool {
	var serr *Internal
	return errors.As(err, &serr)
}
