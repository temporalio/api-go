package errors

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type (
	// NotFound represents not found error.
	NotFound struct {
		Message string
		st *status.Status
	}
)

// NewNotFound returns new NotFound error.
func NewNotFound(message string) *NotFound {
	return &NotFound{
		Message: message,
	}
}

// Error returns string message.
func (e *NotFound) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *NotFound) GRPCStatus() *status.Status {
	if e.st != nil{
		return e.st
	}

	return status.New(codes.NotFound, e.Message)
}

func notFound(st *status.Status) (*NotFound, bool) {
	if st == nil || st.Code() != codes.NotFound {
		return nil, false
	}

	return &NotFound{
		Message: st.Message(),
		st: st,
	}, true
}
