package errors

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

type (
	// DataLoss represents data loss error.
	DataLoss struct {
		Message string
		st *status.Status
	}
)

// NewDataLoss returns new DataLoss error.
func NewDataLoss(message string) *DataLoss {
	return &DataLoss{
		Message: message,
	}
}

// Error returns string message.
func (e *DataLoss) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *DataLoss) GRPCStatus() *status.Status {
	if e.st != nil{
		return e.st
	}

	return status.New(codes.DataLoss, e.Message)
}

func dataLoss(st *status.Status) (*DataLoss, bool) {
	if st == nil || st.Code() != codes.DataLoss {
		return nil, false
	}

	return &DataLoss{
		Message: st.Message(),
		st: st,
	}, true
}
