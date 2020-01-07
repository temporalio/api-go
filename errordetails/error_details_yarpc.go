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
	"go.uber.org/yarpc/encoding/protobuf"
	"go.uber.org/yarpc/yarpcerrors"
)

func NewDomainNotActiveErrorYARPC(msg, domainName, currentCluster, activeCluster string) error {
	return protobuf.NewError(yarpcerrors.CodeInvalidArgument, msg, protobuf.WithErrorDetails(
		&DomainNotActiveFailure{
			DomainName:     domainName,
			CurrentCluster: currentCluster,
			ActiveCluster:  activeCluster,
		},
	))
}

func IsDomainNotActiveFailureYARPC(st *yarpcerrors.Status) bool {
	_, ok := GetDomainNotActiveFailureYARPC(st)
	return ok
}

func GetDomainNotActiveFailureYARPC(err error) (*DomainNotActiveFailure, bool) {
	st := yarpcerrors.FromError(err)
	if st == nil || st.Code() != yarpcerrors.CodeInvalidArgument{
		return nil, false
	}

	failure, ok := getFirstDetailYARPC(err).(*DomainNotActiveFailure)
	return failure, ok
}

func NewWorkflowExecutionAlreadyStartedErrorYARPC(msg, startRequestId, runId string) error {
	return protobuf.NewError(yarpcerrors.CodeAlreadyExists, msg, protobuf.WithErrorDetails(
		&WorkflowExecutionAlreadyStartedFailure{
			StartRequestId: startRequestId,
			RunId:          runId,
		},
	))
}

func IsWorkflowExecutionAlreadyStartedFailureYARPC(st *yarpcerrors.Status) bool {
	_, ok := GetWorkflowExecutionAlreadyStartedFailureYARPC(st)
	return ok
}

func GetWorkflowExecutionAlreadyStartedFailureYARPC(err error) (*WorkflowExecutionAlreadyStartedFailure, bool) {
	st := yarpcerrors.FromError(err)
	if st == nil || st.Code() != yarpcerrors.CodeAlreadyExists{
		return nil, false
	}

	failure, ok := getFirstDetailYARPC(err).(*WorkflowExecutionAlreadyStartedFailure)
	return failure, ok
}

func NewClientVersionNotSupportedErrorYARPC(msg, featureVersion, clientImpl, supportedVersions string) error {
	return protobuf.NewError(yarpcerrors.CodeFailedPrecondition, msg, protobuf.WithErrorDetails(
		&ClientVersionNotSupportedFailure{
			FeatureVersion:    featureVersion,
			ClientImpl:        clientImpl,
			SupportedVersions: supportedVersions,
		},
	))
}

func IsClientVersionNotSupportedErrorYARPC(st *yarpcerrors.Status) bool {
	_, ok := GetClientVersionNotSupportedFailureYARPC(st)
	return ok
}

func GetClientVersionNotSupportedFailureYARPC(err error) (*ClientVersionNotSupportedFailure, bool) {
	st := yarpcerrors.FromError(err)
	if st == nil || st.Code() != yarpcerrors.CodeFailedPrecondition{
		return nil, false
	}

	failure, ok := getFirstDetailYARPC(err).(*ClientVersionNotSupportedFailure)
	return failure, ok
}

func getFirstDetailYARPC(err error) interface{}{
	details := protobuf.GetErrorDetails(err)
	if len(details) > 0 {
		return details[0]
	}

	return nil
}