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
// - New*Status
// - Is*Status
// - Get*Failure

// NewDomainNotActiveStatus returns new status with DomainNotActiveFailure in details.
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

// IsDomainNotActiveStatus checks if status has DomainNotActiveFailure in details.
func IsDomainNotActiveStatus(st *status.Status) bool {
	_, ok := GetDomainNotActiveFailure(st)
	return ok
}

// GetDomainNotActiveFailure returns DomainNotActiveFailure from status details.
func GetDomainNotActiveFailure(st *status.Status) (*DomainNotActiveFailure, bool) {
	if st == nil || st.Code() != codes.InvalidArgument{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*DomainNotActiveFailure)
	return failure, ok
}

// NewWorkflowExecutionAlreadyStartedStatus returns new status with WorkflowExecutionAlreadyStartedFailure in details.
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

// IsWorkflowExecutionAlreadyStartedStatus checks if status has WorkflowExecutionAlreadyStartedFailure in details.
func IsWorkflowExecutionAlreadyStartedStatus(st *status.Status) bool {
	_, ok := GetWorkflowExecutionAlreadyStartedFailure(st)
	return ok
}

// GetWorkflowExecutionAlreadyStartedFailure returns WorkflowExecutionAlreadyStartedFailure from status details.
func GetWorkflowExecutionAlreadyStartedFailure(st *status.Status) (*WorkflowExecutionAlreadyStartedFailure, bool) {
	if st == nil || st.Code() != codes.AlreadyExists{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*WorkflowExecutionAlreadyStartedFailure)
	return failure, ok
}

// NewClientVersionNotSupportedStatus returns new status with ClientVersionNotSupportedFailure in details.
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

// IsClientVersionNotSupportedStatus checks if status has ClientVersionNotSupportedFailure in details.
func IsClientVersionNotSupportedStatus(st *status.Status) bool {
	_, ok := GetClientVersionNotSupportedFailure(st)
	return ok
}

// GetClientVersionNotSupportedFailure returns ClientVersionNotSupportedFailure from status details.
func GetClientVersionNotSupportedFailure(st *status.Status) (*ClientVersionNotSupportedFailure, bool) {
	if st == nil || st.Code() != codes.FailedPrecondition{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*ClientVersionNotSupportedFailure)
	return failure, ok
}

// NewShardOwnershipLostStatus returns new status with ShardOwnershipLostFailure in details.
func NewShardOwnershipLostStatus(msg, owner string) *status.Status {
	st := status.New(codes.Aborted, msg)
	st, _  = st.WithDetails(
		&ShardOwnershipLostFailure{
			Owner:    owner,
		},
	)
	return st
}

// IsShardOwnershipLostStatus checks if status has ShardOwnershipLostFailure in details.
func IsShardOwnershipLostStatus(st *status.Status) bool {
	_, ok := GetShardOwnershipLostFailure(st)
	return ok
}

// GetShardOwnershipLostFailure returns ShardOwnershipLostFailure from status details.
func GetShardOwnershipLostFailure(st *status.Status) (*ShardOwnershipLostFailure, bool) {
	if st == nil || st.Code() != codes.Aborted{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*ShardOwnershipLostFailure)
	return failure, ok
}

// NewRetryTaskStatus returns new status with RetryTaskFailure in details.
func NewRetryTaskStatus(msg, domainId, workflowId, RunId string, nextEventId int64) *status.Status {
	st := status.New(codes.Aborted, msg)
	st, _  = st.WithDetails(
		&RetryTaskFailure{
			DomainId:    domainId,
			WorkflowId:  workflowId,
			RunId:       RunId,
			NextEventId: nextEventId,
		},
	)
	return st
}

// IsRetryTaskStatus checks if status has RetryTaskFailure in details.
func IsRetryTaskStatus(st *status.Status) bool {
	_, ok := GetRetryTaskFailure(st)
	return ok
}

// GetRetryTaskFailure returns RetryTaskFailure from status details.
func GetRetryTaskFailure(st *status.Status) (*RetryTaskFailure, bool) {
	if st == nil || st.Code() != codes.Aborted{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*RetryTaskFailure)
	return failure, ok
}

// NewRetryTaskV2Status returns new status with RetryTaskV2Failure in details.
func NewRetryTaskV2Status(msg, domainId, workflowId, RunId string, startEventId, startEventVersion, endEventId, endEventVersion int64) *status.Status {
	st := status.New(codes.Aborted, msg)
	st, _  = st.WithDetails(
		&RetryTaskV2Failure{
			DomainId:          domainId,
			WorkflowId:        workflowId,
			RunId:             RunId,
			StartEventId:      startEventId,
			StartEventVersion: startEventVersion,
			EndEventId:        endEventId,
			EndEventVersion:   endEventVersion,
		},
	)
	return st
}

// IsRetryTaskV2Status checks if status has RetryTaskV2Failure in details.
func IsRetryTaskV2Status(st *status.Status) bool {
	_, ok := GetRetryTaskV2Failure(st)
	return ok
}

// GetRetryTaskV2Failure returns RetryTaskV2Failure from status details.
func GetRetryTaskV2Failure(st *status.Status) (*RetryTaskV2Failure, bool) {
	if st == nil || st.Code() != codes.Aborted{
		return nil, false
	}

	failure, ok := getFirstDetail(st).(*RetryTaskV2Failure)
	return failure, ok
}

func getFirstDetail(st *status.Status) interface{}{
	details := st.Details()
	if len(details) > 0 {
		return details[0]
	}

	return nil
}
