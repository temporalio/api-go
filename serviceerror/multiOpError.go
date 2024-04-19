// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package serviceerror

import (
	"errors"

	"go.temporal.io/api/errordetails/v1"
	rpc "google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// MultiOperationExecutionError represents an error for a failed MultiOperationExecution.
type MultiOperationExecutionError struct {
	Message string
	errs    []error
	st      *status.Status
}

// NewMultiOperationExecutionError returns a new MultiOperationExecutionError.
func NewMultiOperationExecutionError(message string, errs []error) error {
	return &MultiOperationExecutionError{Message: message, errs: errs}
}

// Error returns string message.
func (e *MultiOperationExecutionError) Error() string {
	return e.Message
}

func (e *MultiOperationExecutionError) OperationErrors() []error {
	return e.errs
}

func (e *MultiOperationExecutionError) Status() *status.Status {
	var code *codes.Code
	failure := &errordetails.MultiOperationExecutionFailure{
		Statuses: []*rpc.Status{},
	}

	var abortedErr *MultiOperationAborted
	for _, err := range e.errs {
		st := ToStatus(err)
		// the first non-OK and non-Aborted code becomes the code for the entire Status
		if code == nil && st.Code() != codes.OK && !errors.As(err, &abortedErr) {
			c := st.Code()
			code = &c
		}
		failure.Statuses = append(failure.Statuses, st.Proto())
	}

	// this should never happen, but it's better to set it to `Aborted` than to panic
	if code == nil {
		c := codes.Aborted
		code = &c
	}

	st := status.New(*code, e.Error())
	st, _ = st.WithDetails(failure)
	return st
}

func newMultiOperationExecutionError(st *status.Status, errs []error) error {
	return &MultiOperationExecutionError{
		Message: st.Message(),
		errs:    errs,
		st:      st,
	}
}
