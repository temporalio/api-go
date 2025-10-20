package serviceerror

import (
	"errors"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/api/errordetails/v1"
)

type (
	// NamespaceNotFound represents namespace not found error.
	NamespaceNotFound struct {
		Message   string
		Namespace string
		st        *status.Status
	}
)

// NewNamespaceNotFound returns new NamespaceNotFound error.
func NewNamespaceNotFound(namespace string) error {
	return &NamespaceNotFound{
		Message: fmt.Sprintf(
			"Namespace %s is not found.",
			namespace,
		),
		Namespace: namespace,
	}
}

// Error returns string message.
func (e *NamespaceNotFound) Error() string {
	return e.Message
}

func (e *NamespaceNotFound) Status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.NotFound, e.Message)
	st, _ = st.WithDetails(
		&errordetails.NamespaceNotFoundFailure{
			Namespace: e.Namespace,
		},
	)
	return st
}

func newNamespaceNotFound(st *status.Status, errDetails *errordetails.NamespaceNotFoundFailure) error {
	return &NamespaceNotFound{
		Message:   st.Message(),
		Namespace: errDetails.GetNamespace(),
		st:        st,
	}
}

// IsNamespaceNotFound returns whether any error in the provided error's chain is a
// NamespaceNotFound error.
func IsNamespaceNotFound(err error) bool {
	var serr *NamespaceNotFound
	return errors.As(err, &serr)
}
