package serviceerror

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type (
	// ResourceExhausted represents resource exhausted error.
	ResourceExhausted struct {
		Message string
		st *status.Status
	}
)

// NewResourceExhausted returns new ResourceExhausted error.
func NewResourceExhausted(message string) *ResourceExhausted {
	return &ResourceExhausted{
		Message: message,
	}
}

// Error returns string message.
func (e *ResourceExhausted) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *ResourceExhausted) GRPCStatus() *status.Status {
	if e.st != nil{
		return e.st
	}

	return status.New(codes.ResourceExhausted, e.Message)
}

func resourceExhausted(st *status.Status) (*ResourceExhausted, bool) {
	if st == nil || st.Code() != codes.ResourceExhausted {
		return nil, false
	}

	return &ResourceExhausted{
		Message: st.Message(),
		st: st,
	}, true
}
