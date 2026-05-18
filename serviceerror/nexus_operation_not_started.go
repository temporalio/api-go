package serviceerror

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// NexusOperationNotStarted represents a nexus operation execution already started error.
	NexusOperationNotStarted struct {
		Message        string
		StartRequestId string
		st             *status.Status
	}
)

// NewNexusOperationNotStarted returns new NexusOperationNotStarted error.
func NewNexusOperationNotStarted(message, startRequestId string) error {
	return &NexusOperationNotStarted{
		Message:        message,
		StartRequestId: startRequestId,
	}
}

// NewNexusOperationNotStartedf returns new NexusOperationNotStarted error with formatted message.
func NewNexusOperationNotStartedf(startRequestId, runId, format string, args ...any) error {
	return &NexusOperationNotStarted{
		Message:        fmt.Sprintf(format, args...),
		StartRequestId: startRequestId,
	}
}

// Error returns string message.
func (e *NexusOperationNotStarted) Error() string {
	return e.Message
}

func (e *NexusOperationNotStarted) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.AlreadyExists, e.Message)
	st, _ = st.WithDetails(
		&errordetails.NexusOperationNotStartedFailure{
			StartRequestId: e.StartRequestId,
		},
	)
	return st
}

func newNexusOperationNotStarted(st *status.Status, errDetails *errordetails.NexusOperationNotStartedFailure) error {
	return &NexusOperationNotStarted{
		Message:        st.Message(),
		StartRequestId: errDetails.GetStartRequestId(),
		st:             st,
	}
}
