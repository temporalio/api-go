package errors

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

)

// FromStatus converts gRPC status to go error.
func FromStatus(st *status.Status) error {
	if st == nil || st.Code() == codes.OK {
		return nil
	}

	//switch st.Code() {
	//case codes.Canceled:
	//case codes.Unknown:
	//case codes.InvalidArgument:
	//case codes.DeadlineExceeded:
	//case codes.NotFound:
	//	err, _ := notFoundErrorFromStatus(st)
	//	return err
	//case codes.AlreadyExists:
	//case codes.PermissionDenied:
	//case codes.ResourceExhausted:
	//case codes.FailedPrecondition:
	//case codes.Aborted:
	//	if err, ok := shardOwnershipLostErrorFromStatus(st); ok {
	//		return err
	//	}
	//case codes.OutOfRange:
	//case codes.Unimplemented:
	//case codes.Internal:
	//case codes.Unavailable:
	//case codes.DataLoss:
	//case codes.Unauthenticated:
	//}

	return st.Err()
}
