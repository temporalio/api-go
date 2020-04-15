// The MIT License (MIT)
//
// Copyright (c) 2020 Temporal Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package serviceerror

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"

	"go.temporal.io/temporal-proto/failure"
	"go.temporal.io/temporal-proto/primitives"
)

type (
	// RetryTask represents retry task error.
	RetryTask struct {
		Message     string
		NamespaceId primitives.UUID
		WorkflowId  string
		RunId       primitives.UUID
		NextEventId int64
		st          *status.Status
	}
)

// NewRetryTask returns new RetryTask error.
func NewRetryTask(message string, namespaceId primitives.UUID, workflowId string, runId primitives.UUID, nextEventId int64) *RetryTask {
	return &RetryTask{
		Message:     message,
		NamespaceId: namespaceId,
		WorkflowId:  workflowId,
		RunId:       runId,
		NextEventId: nextEventId,
	}
}

// Error returns string message.
func (e *RetryTask) Error() string {
	return e.Message
}

func (e *RetryTask) status() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.Aborted, e.Message)
	st, _ = st.WithDetails(
		&failure.RetryTask{
			NamespaceId:    e.NamespaceId,
			WorkflowId:  e.WorkflowId,
			RunId:       e.RunId,
			NextEventId: e.NextEventId,
		},
	)
	return st
}

func newRetryTask(st *status.Status, failure *failure.RetryTask) *RetryTask {
	return &RetryTask{
		Message:     st.Message(),
		NamespaceId:    failure.NamespaceId,
		WorkflowId:  failure.WorkflowId,
		RunId:       failure.RunId,
		NextEventId: failure.NextEventId,
		st:          st,
	}
}
