package serviceerror

import "github.com/gogo/status"

type(
	ServiceError interface{
		error
		status() *status.Status
	}
)
