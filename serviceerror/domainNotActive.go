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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"go.temporal.io/temporal-proto/errordetails"
)

type (
	// DomainNotActive represents domain not active error.
	DomainNotActive struct {
		Message        string
		DomainName     string
		CurrentCluster string
		ActiveCluster  string
		st             *status.Status
	}
)

// NewDomainNotActive returns new DomainNotActive error.
func NewDomainNotActive(message, domainName, currentCluster, activeCluster string) *DomainNotActive {
	return &DomainNotActive{
		Message:        message,
		DomainName:     domainName,
		CurrentCluster: currentCluster,
		ActiveCluster:  activeCluster,
	}
}

// Error returns string message.
func (e *DomainNotActive) Error() string {
	return e.Message
}

// GRPCStatus returns corresponding gRPC status.Status.
func (e *DomainNotActive) GRPCStatus() *status.Status {
	if e.st != nil {
		return e.st
	}

	st := status.New(codes.FailedPrecondition, e.Message)
	st, _ = st.WithDetails(
		&errordetails.DomainNotActiveFailure{
			DomainName:     e.DomainName,
			CurrentCluster: e.CurrentCluster,
			ActiveCluster:  e.ActiveCluster,
		},
	)
	return st
}

func newDomainNotActive(st *status.Status, failure *errordetails.DomainNotActiveFailure) *DomainNotActive {
	return &DomainNotActive{
		Message:        st.Message(),
		DomainName:     failure.DomainName,
		CurrentCluster: failure.CurrentCluster,
		ActiveCluster:  failure.ActiveCluster,
		st:             st,
	}
}
