package errors

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type (
	// InvalidArgument represents invalid argument error.
	InvalidArgument struct {
		Message string
		st *status.Status
	}
)

// NewInvalidArgument returns new InvalidArgument error.
func NewInvalidArgument(message string) *InvalidArgument {
	return &InvalidArgument{
		Message: message,
	}
}

// Error returns string message.
func (e *InvalidArgument) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *InvalidArgument) GRPCStatus() *status.Status {
	if e.st != nil{
		return e.st
	}

	return status.New(codes.InvalidArgument, e.Message)
}

func invalidArgument(st *status.Status) (*InvalidArgument, bool) {
	if st == nil || st.Code() != codes.InvalidArgument {
		return nil, false
	}

	return &InvalidArgument{
		Message: st.Message(),
		st: st,
	}, true
}
