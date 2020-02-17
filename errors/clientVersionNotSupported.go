package errors

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"go.temporal.io/temporal-proto/errordetails"
)

type (
	// ClientVersionNotSupported represents client version is not supported error.
	ClientVersionNotSupported struct {
		Message           string
		FeatureVersion    string
		ClientImpl        string
		SupportedVersions string
		st                *status.Status
	}
)

// NewClientVersionNotSupported returns new ClientVersionNotSupported error.
func NewClientVersionNotSupported(message, featureVersion, clientImpl, supportedVersions string) *ClientVersionNotSupported {
	return &ClientVersionNotSupported{
		Message:           message,
		FeatureVersion:    featureVersion,
		ClientImpl:        clientImpl,
		SupportedVersions: supportedVersions,
	}
}

// Error returns string message.
func (e *ClientVersionNotSupported) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *ClientVersionNotSupported) GRPCStatus() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.FailedPrecondition, e.Message)
	st, _ = st.WithDetails(
		&errordetails.ClientVersionNotSupportedFailure{
			FeatureVersion:    e.FeatureVersion,
			ClientImpl:        e.ClientImpl,
			SupportedVersions: e.SupportedVersions,
		},
	)
	return st
}

func clientVersionNotSupported(st *status.Status) (*ClientVersionNotSupported, bool) {
	if st == nil || st.Code() != codes.FailedPrecondition {
		return nil, false
	}

	if failure, ok := getFirstDetail(st).(*errordetails.ClientVersionNotSupportedFailure); ok {
		return &ClientVersionNotSupported{
			Message:           st.Message(),
			FeatureVersion:    failure.FeatureVersion,
			ClientImpl:        failure.ClientImpl,
			SupportedVersions: failure.SupportedVersions,
			st:                st,
		}, true
	}

	return nil, false
}
