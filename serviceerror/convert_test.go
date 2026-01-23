package serviceerror_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	rpc "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/durationpb"

	enumspb "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/errordetails/v1"
	"go.temporal.io/api/serviceerror"
)

func TestFromStatus_NotFound(t *testing.T) {
	var err error
	st1 := status.New(codes.NotFound, "Not found.")
	err1 := serviceerror.FromStatus(st1)
	require.IsType(t, &serviceerror.NotFound{}, err1)
	require.Equal(t, codes.NotFound, err1.(*serviceerror.NotFound).Status().Code())
	require.Equal(t, "Not found.", err1.(*serviceerror.NotFound).Message)

	st2 := status.New(codes.NotFound, "Not found.")
	st2, err = st1.WithDetails(&errordetails.NamespaceNotFoundFailure{
		Namespace: "test-ns",
	})
	require.NoError(t, err)
	err2 := serviceerror.FromStatus(st2)
	require.IsType(t, &serviceerror.NamespaceNotFound{}, err2)
	require.Equal(t, codes.NotFound, err2.(*serviceerror.NamespaceNotFound).Status().Code())
	require.Equal(t, "Not found.", err2.(*serviceerror.NamespaceNotFound).Message)
	require.Equal(t, "test-ns", err2.(*serviceerror.NamespaceNotFound).Namespace)
}

func TestFromStatus_UnknownErrorDetails(t *testing.T) {
	st1 := status.FromProto(&rpc.Status{
		Code:    int32(codes.NotFound),
		Message: "Not found.",
	})

	err1 := serviceerror.FromStatus(st1)
	require.IsType(t, &serviceerror.NotFound{}, err1)
	require.Equal(t, codes.NotFound, err1.(*serviceerror.NotFound).Status().Code())
	require.Equal(t, "Not found.", err1.(*serviceerror.NotFound).Message)
}

func TestToStatus_UnknownErrorDetails(t *testing.T) {
	anyd, err := anypb.New(durationpb.New(time.Duration(time.Second)))
	if err != nil {
		t.Fatalf("Failed to create any out of duration: %s", err)
	}
	err1 := status.ErrorProto(&rpc.Status{
		Code:    int32(codes.NotFound),
		Message: "Not found.",
		Details: []*anypb.Any{anyd},
	})

	st1 := serviceerror.ToStatus(err1)
	require.Equal(t, codes.NotFound, st1.Code())
	require.Equal(t, "Not found.", st1.Message())
	require.Len(t, st1.Details(), 1)
	require.Equal(t, "type.googleapis.com/google.protobuf.Duration", st1.Proto().Details[0].TypeUrl)
}

func TestToStatus_NotServiceError(t *testing.T) {
	err1 := errors.New("some error")
	st1 := serviceerror.ToStatus(err1)
	require.Equal(t, codes.Unknown, st1.Code())
	require.Equal(t, "some error", st1.Message())
	require.Len(t, st1.Details(), 0)
}

func TestMultiOperationExecution(t *testing.T) {
	t.Run("several errors", func(t *testing.T) {
		err := serviceerror.NewMultiOperationExecution(
			"MultiOperation could not be executed.",
			[]error{
				serviceerror.NewMultiOperationAborted("Operation was aborted."),
				serviceerror.NewInvalidArgument("invalid arg"),
				serviceerror.NewMultiOperationAborted("Operation was aborted."),
			})

		st := serviceerror.ToStatus(err)
		require.Equal(t, codes.InvalidArgument, st.Code())
		require.Equal(t, "MultiOperation could not be executed.", st.Message())
		require.Len(t, st.Details(), 1)

		failure := st.Details()[0].(*errordetails.MultiOperationExecutionFailure)
		st1 := failure.Statuses[0]
		require.Equal(t, int32(codes.Aborted), st1.GetCode())
		require.Equal(t, "Operation was aborted.", st1.GetMessage())
		st2 := failure.Statuses[1]
		require.Equal(t, int32(codes.InvalidArgument), st2.GetCode())
		require.Equal(t, "invalid arg", st2.GetMessage())
		st3 := failure.Statuses[2]
		require.Equal(t, int32(codes.Aborted), st3.GetCode())
		require.Equal(t, "Operation was aborted.", st3.GetMessage())

		errFromStatus := serviceerror.FromStatus(st)
		reconstructedStatus := serviceerror.ToStatus(errFromStatus)
		require.True(t, proto.Equal(st.Proto(), reconstructedStatus.Proto()))
	})

	t.Run("single multi operation aborted", func(t *testing.T) {
		err := serviceerror.NewMultiOperationExecution(
			"MultiOperation could not be executed.",
			[]error{
				serviceerror.NewMultiOperationAborted("Operation was aborted."),
			})

		st := serviceerror.ToStatus(err)
		require.Equal(t, codes.Aborted, st.Code())
		require.Equal(t, err.Error(), st.Message())
		require.Len(t, st.Details(), 1)
	})

	t.Run("no errors", func(t *testing.T) {
		err := serviceerror.NewMultiOperationExecution(
			"MultiOperation could not be executed.",
			[]error{})

		st := serviceerror.ToStatus(err)
		require.Equal(t, codes.Aborted, st.Code())
		require.Equal(t, err.Error(), st.Message())
		require.Len(t, st.Details(), 1)
		require.Empty(t, st.Details()[0].(*errordetails.MultiOperationExecutionFailure).Statuses)
	})
}

func TestMultiOperationExecution_PreservesResourceExhaustedDetails(t *testing.T) {
	resourceErr := &serviceerror.ResourceExhausted{
		Cause:   enumspb.RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT,
		Scope:   enumspb.RESOURCE_EXHAUSTED_SCOPE_NAMESPACE,
		Message: "rate limit exceeded",
	}
	multiErr := serviceerror.NewMultiOperationExecution(
		"multi-op failed",
		[]error{
			resourceErr,
			serviceerror.NewInvalidArgument("other error"),
		},
	)

	// Check round-trip conversion
	st := serviceerror.ToStatus(multiErr)
	reconstructed := serviceerror.FromStatus(st)

	multiOp, ok := reconstructed.(*serviceerror.MultiOperationExecution)
	require.True(t, ok)

	// Get the operation errors
	opErrors := multiOp.OperationErrors()
	require.Len(t, opErrors, 2)

	// Verify the first error is ResourceExhausted with preserved details
	reErr, ok := opErrors[0].(*serviceerror.ResourceExhausted)
	require.True(t, ok)
	require.Equal(t, enumspb.RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT, reErr.Cause)
	require.Equal(t, enumspb.RESOURCE_EXHAUSTED_SCOPE_NAMESPACE, reErr.Scope)
	require.Equal(t, "rate limit exceeded", reErr.Message)

	// Verify proto round-trip equality
	reconstructedStatus := serviceerror.ToStatus(reconstructed)
	require.True(t, proto.Equal(st.Proto(), reconstructedStatus.Proto()))
}

func TestMultiOperationAborted(t *testing.T) {
	err := serviceerror.NewMultiOperationAborted("Operation was aborted.")

	st := serviceerror.ToStatus(err)
	require.Equal(t, codes.Aborted, st.Code())
	require.Equal(t, err.Error(), st.Message())
	require.Len(t, st.Details(), 1)

	errFromStatus := serviceerror.FromStatus(st)
	require.Equal(t, err.Error(), errFromStatus.Error())

	reconstructedStatus := serviceerror.ToStatus(errFromStatus)
	require.True(t, proto.Equal(st.Proto(), reconstructedStatus.Proto()))
}

func TestFromWrapped(t *testing.T) {
	err := &serviceerror.PermissionDenied{
		Message: "x is not allowed",
		Reason:  "arbitrary reason",
	}
	wrapped := fmt.Errorf("wrapped error: %w", err)
	s := serviceerror.ToStatus(wrapped)
	require.Equal(t, codes.PermissionDenied, s.Code())
	require.Equal(t, "wrapped error: x is not allowed", s.Message())
	require.True(t, proto.Equal(
		&errordetails.PermissionDeniedFailure{Reason: "arbitrary reason"},
		s.Details()[0].(*errordetails.PermissionDeniedFailure)))
}
