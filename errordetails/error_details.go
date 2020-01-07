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

package errordetails

import (
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// Generate these 3 helper funcs for every failure from error_details.proto:
// - Build*Status
// - Is*Failure
// - Get*Failure

func NewDomainNotActiveStatus(msg, domainName, currentCluster, activeCluster string) *status.Status {
	st := status.New(codes.InvalidArgument, msg)
	st, _  = st.WithDetails(
		&DomainNotActiveFailure{
			DomainName:     domainName,
			CurrentCluster: currentCluster,
			ActiveCluster:  activeCluster,
		},
	)

	return st
}

func IsDomainNotActiveFailure(st *status.Status) bool {
	_, ok := GetDomainNotActiveFailure(st)
	return ok
}

func GetDomainNotActiveFailure(st *status.Status) (*DomainNotActiveFailure, bool) {
	if st == nil || st.Code() != codes.InvalidArgument{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*DomainNotActiveFailure)
	return failure, ok
}

func NewWorkflowExecutionAlreadyStartedStatus(msg, startRequestId, runId string) *status.Status {
	st := status.New(codes.AlreadyExists, msg)
	st, _  = st.WithDetails(
		&WorkflowExecutionAlreadyStartedFailure{
			StartRequestId: startRequestId,
			RunId:          runId,
		},
	)

	return st
}

func IsWorkflowExecutionAlreadyStartedFailure(st *status.Status) bool {
	_, ok := GetWorkflowExecutionAlreadyStartedFailure(st)
	return ok
}

func GetWorkflowExecutionAlreadyStartedFailure(st *status.Status) (*WorkflowExecutionAlreadyStartedFailure, bool) {
	if st == nil || st.Code() != codes.AlreadyExists{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*WorkflowExecutionAlreadyStartedFailure)
	return failure, ok
}

func NewClientVersionNotSupportedStatus(msg, featureVersion, clientImpl, supportedVersions string) *status.Status {
	st := status.New(codes.FailedPrecondition, msg)
	st, _  = st.WithDetails(
		&ClientVersionNotSupportedFailure{
			FeatureVersion:    featureVersion,
			ClientImpl:        clientImpl,
			SupportedVersions: supportedVersions,
		},
	)

	return st
}

func IsClientVersionNotSupportedFailure(st *status.Status) bool {
	_, ok := GetClientVersionNotSupportedFailure(st)
	return ok
}

func GetClientVersionNotSupportedFailure(st *status.Status) (*ClientVersionNotSupportedFailure, bool) {
	if st == nil || st.Code() != codes.AlreadyExists{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*ClientVersionNotSupportedFailure)
	return failure, ok
}

func getFirstDetail(st *status.Status) interface{}{
	details := st.Details()
	if len(details) > 0 {
		return details[0]
	}

	return nil
}