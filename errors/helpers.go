package errors

import "github.com/gogo/status"

func getFirstDetail(st *status.Status) interface{} {
	details := st.Details()
	if len(details) > 0 {
		return details[0]
	}

	return nil
}
