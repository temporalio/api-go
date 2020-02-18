package serviceerror

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"go.temporal.io/temporal-proto/errordetails"
)

type (
	// ShardOwnershipLost represents shard ownership lost error.
	ShardOwnershipLost struct {
		Message string
		Owner   string
		st      *status.Status
	}
)

// NewShardOwnershipLost returns new ShardOwnershipLost error.
func NewShardOwnershipLost(message, owner string) *ShardOwnershipLost {
	return &ShardOwnershipLost{
		Message: message,
		Owner:   owner,
	}
}

// Error returns string message.
func (e *ShardOwnershipLost) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *ShardOwnershipLost) GRPCStatus() *status.Status {
	if e.st != nil{
		return e.st
	}

	st := status.New(codes.Aborted, e.Message)
	st, _ = st.WithDetails(
		&errordetails.ShardOwnershipLostFailure{
			Owner: e.Owner,
		},
	)
	return st
}

func shardOwnershipLost(st *status.Status) (*ShardOwnershipLost, bool) {
	if st == nil || st.Code() != codes.Aborted {
		return nil, false
	}

	if failure, ok := getFailure(st).(*errordetails.ShardOwnershipLostFailure); ok {
		return &ShardOwnershipLost{
			Message: st.Message(),
			Owner:   failure.Owner,
			st:      st,
		}, true
	}

	return nil, false
}
