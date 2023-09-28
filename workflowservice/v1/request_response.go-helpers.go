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

package workflowservice

import "google.golang.org/protobuf/proto"

func (val *RegisterNamespaceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RegisterNamespaceRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RegisterNamespaceRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RegisterNamespaceRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RegisterNamespaceRequest
	switch t := that.(type) {
	case *RegisterNamespaceRequest:
		that1 = t
	case RegisterNamespaceRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RegisterNamespaceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RegisterNamespaceResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RegisterNamespaceResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RegisterNamespaceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RegisterNamespaceResponse
	switch t := that.(type) {
	case *RegisterNamespaceResponse:
		that1 = t
	case RegisterNamespaceResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListNamespacesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListNamespacesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListNamespacesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListNamespacesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListNamespacesRequest
	switch t := that.(type) {
	case *ListNamespacesRequest:
		that1 = t
	case ListNamespacesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListNamespacesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListNamespacesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListNamespacesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListNamespacesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListNamespacesResponse
	switch t := that.(type) {
	case *ListNamespacesResponse:
		that1 = t
	case ListNamespacesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeNamespaceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeNamespaceRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeNamespaceRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeNamespaceRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeNamespaceRequest
	switch t := that.(type) {
	case *DescribeNamespaceRequest:
		that1 = t
	case DescribeNamespaceRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeNamespaceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeNamespaceResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeNamespaceResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeNamespaceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeNamespaceResponse
	switch t := that.(type) {
	case *DescribeNamespaceResponse:
		that1 = t
	case DescribeNamespaceResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateNamespaceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateNamespaceRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateNamespaceRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateNamespaceRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateNamespaceRequest
	switch t := that.(type) {
	case *UpdateNamespaceRequest:
		that1 = t
	case UpdateNamespaceRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateNamespaceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateNamespaceResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateNamespaceResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateNamespaceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateNamespaceResponse
	switch t := that.(type) {
	case *UpdateNamespaceResponse:
		that1 = t
	case UpdateNamespaceResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DeprecateNamespaceRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DeprecateNamespaceRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DeprecateNamespaceRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeprecateNamespaceRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeprecateNamespaceRequest
	switch t := that.(type) {
	case *DeprecateNamespaceRequest:
		that1 = t
	case DeprecateNamespaceRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DeprecateNamespaceResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DeprecateNamespaceResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DeprecateNamespaceResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeprecateNamespaceResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeprecateNamespaceResponse
	switch t := that.(type) {
	case *DeprecateNamespaceResponse:
		that1 = t
	case DeprecateNamespaceResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StartWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StartWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StartWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartWorkflowExecutionRequest
	switch t := that.(type) {
	case *StartWorkflowExecutionRequest:
		that1 = t
	case StartWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StartWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StartWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StartWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartWorkflowExecutionResponse
	switch t := that.(type) {
	case *StartWorkflowExecutionResponse:
		that1 = t
	case StartWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkflowExecutionHistoryRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkflowExecutionHistoryRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkflowExecutionHistoryRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryRequest
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryRequest:
		that1 = t
	case GetWorkflowExecutionHistoryRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkflowExecutionHistoryResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkflowExecutionHistoryResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkflowExecutionHistoryResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryResponse
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryResponse:
		that1 = t
	case GetWorkflowExecutionHistoryResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkflowExecutionHistoryReverseRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkflowExecutionHistoryReverseRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkflowExecutionHistoryReverseRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryReverseRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryReverseRequest
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryReverseRequest:
		that1 = t
	case GetWorkflowExecutionHistoryReverseRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkflowExecutionHistoryReverseResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkflowExecutionHistoryReverseResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkflowExecutionHistoryReverseResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkflowExecutionHistoryReverseResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkflowExecutionHistoryReverseResponse
	switch t := that.(type) {
	case *GetWorkflowExecutionHistoryReverseResponse:
		that1 = t
	case GetWorkflowExecutionHistoryReverseResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PollWorkflowTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PollWorkflowTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PollWorkflowTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowTaskQueueRequest
	switch t := that.(type) {
	case *PollWorkflowTaskQueueRequest:
		that1 = t
	case PollWorkflowTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PollWorkflowTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PollWorkflowTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PollWorkflowTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowTaskQueueResponse
	switch t := that.(type) {
	case *PollWorkflowTaskQueueResponse:
		that1 = t
	case PollWorkflowTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondWorkflowTaskCompletedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondWorkflowTaskCompletedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondWorkflowTaskCompletedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskCompletedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskCompletedRequest
	switch t := that.(type) {
	case *RespondWorkflowTaskCompletedRequest:
		that1 = t
	case RespondWorkflowTaskCompletedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondWorkflowTaskCompletedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondWorkflowTaskCompletedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondWorkflowTaskCompletedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskCompletedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskCompletedResponse
	switch t := that.(type) {
	case *RespondWorkflowTaskCompletedResponse:
		that1 = t
	case RespondWorkflowTaskCompletedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondWorkflowTaskFailedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondWorkflowTaskFailedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondWorkflowTaskFailedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskFailedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskFailedRequest
	switch t := that.(type) {
	case *RespondWorkflowTaskFailedRequest:
		that1 = t
	case RespondWorkflowTaskFailedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondWorkflowTaskFailedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondWorkflowTaskFailedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondWorkflowTaskFailedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondWorkflowTaskFailedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondWorkflowTaskFailedResponse
	switch t := that.(type) {
	case *RespondWorkflowTaskFailedResponse:
		that1 = t
	case RespondWorkflowTaskFailedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PollActivityTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PollActivityTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PollActivityTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollActivityTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollActivityTaskQueueRequest
	switch t := that.(type) {
	case *PollActivityTaskQueueRequest:
		that1 = t
	case PollActivityTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PollActivityTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PollActivityTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PollActivityTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollActivityTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollActivityTaskQueueResponse
	switch t := that.(type) {
	case *PollActivityTaskQueueResponse:
		that1 = t
	case PollActivityTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RecordActivityTaskHeartbeatRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RecordActivityTaskHeartbeatRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RecordActivityTaskHeartbeatRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskHeartbeatRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskHeartbeatRequest
	switch t := that.(type) {
	case *RecordActivityTaskHeartbeatRequest:
		that1 = t
	case RecordActivityTaskHeartbeatRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RecordActivityTaskHeartbeatResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RecordActivityTaskHeartbeatResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RecordActivityTaskHeartbeatResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskHeartbeatResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskHeartbeatResponse
	switch t := that.(type) {
	case *RecordActivityTaskHeartbeatResponse:
		that1 = t
	case RecordActivityTaskHeartbeatResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RecordActivityTaskHeartbeatByIdRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RecordActivityTaskHeartbeatByIdRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RecordActivityTaskHeartbeatByIdRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskHeartbeatByIdRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskHeartbeatByIdRequest
	switch t := that.(type) {
	case *RecordActivityTaskHeartbeatByIdRequest:
		that1 = t
	case RecordActivityTaskHeartbeatByIdRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RecordActivityTaskHeartbeatByIdResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RecordActivityTaskHeartbeatByIdResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RecordActivityTaskHeartbeatByIdResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordActivityTaskHeartbeatByIdResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordActivityTaskHeartbeatByIdResponse
	switch t := that.(type) {
	case *RecordActivityTaskHeartbeatByIdResponse:
		that1 = t
	case RecordActivityTaskHeartbeatByIdResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCompletedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCompletedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCompletedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCompletedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCompletedRequest
	switch t := that.(type) {
	case *RespondActivityTaskCompletedRequest:
		that1 = t
	case RespondActivityTaskCompletedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCompletedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCompletedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCompletedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCompletedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCompletedResponse
	switch t := that.(type) {
	case *RespondActivityTaskCompletedResponse:
		that1 = t
	case RespondActivityTaskCompletedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCompletedByIdRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCompletedByIdRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCompletedByIdRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCompletedByIdRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCompletedByIdRequest
	switch t := that.(type) {
	case *RespondActivityTaskCompletedByIdRequest:
		that1 = t
	case RespondActivityTaskCompletedByIdRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCompletedByIdResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCompletedByIdResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCompletedByIdResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCompletedByIdResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCompletedByIdResponse
	switch t := that.(type) {
	case *RespondActivityTaskCompletedByIdResponse:
		that1 = t
	case RespondActivityTaskCompletedByIdResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskFailedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskFailedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskFailedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskFailedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskFailedRequest
	switch t := that.(type) {
	case *RespondActivityTaskFailedRequest:
		that1 = t
	case RespondActivityTaskFailedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskFailedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskFailedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskFailedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskFailedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskFailedResponse
	switch t := that.(type) {
	case *RespondActivityTaskFailedResponse:
		that1 = t
	case RespondActivityTaskFailedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskFailedByIdRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskFailedByIdRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskFailedByIdRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskFailedByIdRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskFailedByIdRequest
	switch t := that.(type) {
	case *RespondActivityTaskFailedByIdRequest:
		that1 = t
	case RespondActivityTaskFailedByIdRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskFailedByIdResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskFailedByIdResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskFailedByIdResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskFailedByIdResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskFailedByIdResponse
	switch t := that.(type) {
	case *RespondActivityTaskFailedByIdResponse:
		that1 = t
	case RespondActivityTaskFailedByIdResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCanceledRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCanceledRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCanceledRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCanceledRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCanceledRequest
	switch t := that.(type) {
	case *RespondActivityTaskCanceledRequest:
		that1 = t
	case RespondActivityTaskCanceledRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCanceledResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCanceledResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCanceledResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCanceledResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCanceledResponse
	switch t := that.(type) {
	case *RespondActivityTaskCanceledResponse:
		that1 = t
	case RespondActivityTaskCanceledResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCanceledByIdRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCanceledByIdRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCanceledByIdRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCanceledByIdRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCanceledByIdRequest
	switch t := that.(type) {
	case *RespondActivityTaskCanceledByIdRequest:
		that1 = t
	case RespondActivityTaskCanceledByIdRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondActivityTaskCanceledByIdResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondActivityTaskCanceledByIdResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondActivityTaskCanceledByIdResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondActivityTaskCanceledByIdResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondActivityTaskCanceledByIdResponse
	switch t := that.(type) {
	case *RespondActivityTaskCanceledByIdResponse:
		that1 = t
	case RespondActivityTaskCanceledByIdResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RequestCancelWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RequestCancelWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RequestCancelWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelWorkflowExecutionRequest
	switch t := that.(type) {
	case *RequestCancelWorkflowExecutionRequest:
		that1 = t
	case RequestCancelWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RequestCancelWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RequestCancelWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RequestCancelWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelWorkflowExecutionResponse
	switch t := that.(type) {
	case *RequestCancelWorkflowExecutionResponse:
		that1 = t
	case RequestCancelWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SignalWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SignalWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two SignalWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWorkflowExecutionRequest
	switch t := that.(type) {
	case *SignalWorkflowExecutionRequest:
		that1 = t
	case SignalWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SignalWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SignalWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two SignalWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWorkflowExecutionResponse
	switch t := that.(type) {
	case *SignalWorkflowExecutionResponse:
		that1 = t
	case SignalWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SignalWithStartWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SignalWithStartWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two SignalWithStartWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWithStartWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWithStartWorkflowExecutionRequest
	switch t := that.(type) {
	case *SignalWithStartWorkflowExecutionRequest:
		that1 = t
	case SignalWithStartWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SignalWithStartWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SignalWithStartWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two SignalWithStartWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalWithStartWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalWithStartWorkflowExecutionResponse
	switch t := that.(type) {
	case *SignalWithStartWorkflowExecutionResponse:
		that1 = t
	case SignalWithStartWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ResetWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ResetWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ResetWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetWorkflowExecutionRequest
	switch t := that.(type) {
	case *ResetWorkflowExecutionRequest:
		that1 = t
	case ResetWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ResetWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ResetWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ResetWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetWorkflowExecutionResponse
	switch t := that.(type) {
	case *ResetWorkflowExecutionResponse:
		that1 = t
	case ResetWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *TerminateWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *TerminateWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two TerminateWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TerminateWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TerminateWorkflowExecutionRequest
	switch t := that.(type) {
	case *TerminateWorkflowExecutionRequest:
		that1 = t
	case TerminateWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *TerminateWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *TerminateWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two TerminateWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TerminateWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TerminateWorkflowExecutionResponse
	switch t := that.(type) {
	case *TerminateWorkflowExecutionResponse:
		that1 = t
	case TerminateWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DeleteWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DeleteWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DeleteWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowExecutionRequest
	switch t := that.(type) {
	case *DeleteWorkflowExecutionRequest:
		that1 = t
	case DeleteWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DeleteWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DeleteWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DeleteWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteWorkflowExecutionResponse
	switch t := that.(type) {
	case *DeleteWorkflowExecutionResponse:
		that1 = t
	case DeleteWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListOpenWorkflowExecutionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListOpenWorkflowExecutionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListOpenWorkflowExecutionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListOpenWorkflowExecutionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListOpenWorkflowExecutionsRequest
	switch t := that.(type) {
	case *ListOpenWorkflowExecutionsRequest:
		that1 = t
	case ListOpenWorkflowExecutionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListOpenWorkflowExecutionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListOpenWorkflowExecutionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListOpenWorkflowExecutionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListOpenWorkflowExecutionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListOpenWorkflowExecutionsResponse
	switch t := that.(type) {
	case *ListOpenWorkflowExecutionsResponse:
		that1 = t
	case ListOpenWorkflowExecutionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListClosedWorkflowExecutionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListClosedWorkflowExecutionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListClosedWorkflowExecutionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClosedWorkflowExecutionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClosedWorkflowExecutionsRequest
	switch t := that.(type) {
	case *ListClosedWorkflowExecutionsRequest:
		that1 = t
	case ListClosedWorkflowExecutionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListClosedWorkflowExecutionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListClosedWorkflowExecutionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListClosedWorkflowExecutionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListClosedWorkflowExecutionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListClosedWorkflowExecutionsResponse
	switch t := that.(type) {
	case *ListClosedWorkflowExecutionsResponse:
		that1 = t
	case ListClosedWorkflowExecutionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListWorkflowExecutionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListWorkflowExecutionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListWorkflowExecutionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListWorkflowExecutionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListWorkflowExecutionsRequest
	switch t := that.(type) {
	case *ListWorkflowExecutionsRequest:
		that1 = t
	case ListWorkflowExecutionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListWorkflowExecutionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListWorkflowExecutionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListWorkflowExecutionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListWorkflowExecutionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListWorkflowExecutionsResponse
	switch t := that.(type) {
	case *ListWorkflowExecutionsResponse:
		that1 = t
	case ListWorkflowExecutionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListArchivedWorkflowExecutionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListArchivedWorkflowExecutionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListArchivedWorkflowExecutionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListArchivedWorkflowExecutionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListArchivedWorkflowExecutionsRequest
	switch t := that.(type) {
	case *ListArchivedWorkflowExecutionsRequest:
		that1 = t
	case ListArchivedWorkflowExecutionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListArchivedWorkflowExecutionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListArchivedWorkflowExecutionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListArchivedWorkflowExecutionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListArchivedWorkflowExecutionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListArchivedWorkflowExecutionsResponse
	switch t := that.(type) {
	case *ListArchivedWorkflowExecutionsResponse:
		that1 = t
	case ListArchivedWorkflowExecutionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScanWorkflowExecutionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScanWorkflowExecutionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScanWorkflowExecutionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScanWorkflowExecutionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScanWorkflowExecutionsRequest
	switch t := that.(type) {
	case *ScanWorkflowExecutionsRequest:
		that1 = t
	case ScanWorkflowExecutionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScanWorkflowExecutionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScanWorkflowExecutionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScanWorkflowExecutionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScanWorkflowExecutionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScanWorkflowExecutionsResponse
	switch t := that.(type) {
	case *ScanWorkflowExecutionsResponse:
		that1 = t
	case ScanWorkflowExecutionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CountWorkflowExecutionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CountWorkflowExecutionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two CountWorkflowExecutionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CountWorkflowExecutionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CountWorkflowExecutionsRequest
	switch t := that.(type) {
	case *CountWorkflowExecutionsRequest:
		that1 = t
	case CountWorkflowExecutionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CountWorkflowExecutionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CountWorkflowExecutionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two CountWorkflowExecutionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CountWorkflowExecutionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CountWorkflowExecutionsResponse
	switch t := that.(type) {
	case *CountWorkflowExecutionsResponse:
		that1 = t
	case CountWorkflowExecutionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetSearchAttributesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetSearchAttributesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetSearchAttributesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetSearchAttributesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetSearchAttributesRequest
	switch t := that.(type) {
	case *GetSearchAttributesRequest:
		that1 = t
	case GetSearchAttributesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetSearchAttributesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetSearchAttributesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetSearchAttributesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetSearchAttributesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetSearchAttributesResponse
	switch t := that.(type) {
	case *GetSearchAttributesResponse:
		that1 = t
	case GetSearchAttributesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondQueryTaskCompletedRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondQueryTaskCompletedRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondQueryTaskCompletedRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondQueryTaskCompletedRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondQueryTaskCompletedRequest
	switch t := that.(type) {
	case *RespondQueryTaskCompletedRequest:
		that1 = t
	case RespondQueryTaskCompletedRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RespondQueryTaskCompletedResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RespondQueryTaskCompletedResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two RespondQueryTaskCompletedResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RespondQueryTaskCompletedResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RespondQueryTaskCompletedResponse
	switch t := that.(type) {
	case *RespondQueryTaskCompletedResponse:
		that1 = t
	case RespondQueryTaskCompletedResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ResetStickyTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ResetStickyTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ResetStickyTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetStickyTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetStickyTaskQueueRequest
	switch t := that.(type) {
	case *ResetStickyTaskQueueRequest:
		that1 = t
	case ResetStickyTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ResetStickyTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ResetStickyTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ResetStickyTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetStickyTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetStickyTaskQueueResponse
	switch t := that.(type) {
	case *ResetStickyTaskQueueResponse:
		that1 = t
	case ResetStickyTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *QueryWorkflowRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *QueryWorkflowRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two QueryWorkflowRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryWorkflowRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryWorkflowRequest
	switch t := that.(type) {
	case *QueryWorkflowRequest:
		that1 = t
	case QueryWorkflowRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *QueryWorkflowResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *QueryWorkflowResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two QueryWorkflowResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryWorkflowResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryWorkflowResponse
	switch t := that.(type) {
	case *QueryWorkflowResponse:
		that1 = t
	case QueryWorkflowResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeWorkflowExecutionRequest
	switch t := that.(type) {
	case *DescribeWorkflowExecutionRequest:
		that1 = t
	case DescribeWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeWorkflowExecutionResponse
	switch t := that.(type) {
	case *DescribeWorkflowExecutionResponse:
		that1 = t
	case DescribeWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeTaskQueueRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeTaskQueueRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeTaskQueueRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeTaskQueueRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeTaskQueueRequest
	switch t := that.(type) {
	case *DescribeTaskQueueRequest:
		that1 = t
	case DescribeTaskQueueRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeTaskQueueResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeTaskQueueResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeTaskQueueResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeTaskQueueResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeTaskQueueResponse
	switch t := that.(type) {
	case *DescribeTaskQueueResponse:
		that1 = t
	case DescribeTaskQueueResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetClusterInfoRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetClusterInfoRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetClusterInfoRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetClusterInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetClusterInfoRequest
	switch t := that.(type) {
	case *GetClusterInfoRequest:
		that1 = t
	case GetClusterInfoRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetClusterInfoResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetClusterInfoResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetClusterInfoResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetClusterInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetClusterInfoResponse
	switch t := that.(type) {
	case *GetClusterInfoResponse:
		that1 = t
	case GetClusterInfoResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetSystemInfoRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetSystemInfoRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetSystemInfoRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetSystemInfoRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetSystemInfoRequest
	switch t := that.(type) {
	case *GetSystemInfoRequest:
		that1 = t
	case GetSystemInfoRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetSystemInfoResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetSystemInfoResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetSystemInfoResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetSystemInfoResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetSystemInfoResponse
	switch t := that.(type) {
	case *GetSystemInfoResponse:
		that1 = t
	case GetSystemInfoResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListTaskQueuePartitionsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListTaskQueuePartitionsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListTaskQueuePartitionsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListTaskQueuePartitionsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListTaskQueuePartitionsRequest
	switch t := that.(type) {
	case *ListTaskQueuePartitionsRequest:
		that1 = t
	case ListTaskQueuePartitionsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListTaskQueuePartitionsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListTaskQueuePartitionsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListTaskQueuePartitionsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListTaskQueuePartitionsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListTaskQueuePartitionsResponse
	switch t := that.(type) {
	case *ListTaskQueuePartitionsResponse:
		that1 = t
	case ListTaskQueuePartitionsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CreateScheduleRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CreateScheduleRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two CreateScheduleRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CreateScheduleRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CreateScheduleRequest
	switch t := that.(type) {
	case *CreateScheduleRequest:
		that1 = t
	case CreateScheduleRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CreateScheduleResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CreateScheduleResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two CreateScheduleResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CreateScheduleResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CreateScheduleResponse
	switch t := that.(type) {
	case *CreateScheduleResponse:
		that1 = t
	case CreateScheduleResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeScheduleRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeScheduleRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeScheduleRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeScheduleRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeScheduleRequest
	switch t := that.(type) {
	case *DescribeScheduleRequest:
		that1 = t
	case DescribeScheduleRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeScheduleResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeScheduleResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeScheduleResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeScheduleResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeScheduleResponse
	switch t := that.(type) {
	case *DescribeScheduleResponse:
		that1 = t
	case DescribeScheduleResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateScheduleRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateScheduleRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateScheduleRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateScheduleRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateScheduleRequest
	switch t := that.(type) {
	case *UpdateScheduleRequest:
		that1 = t
	case UpdateScheduleRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateScheduleResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateScheduleResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateScheduleResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateScheduleResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateScheduleResponse
	switch t := that.(type) {
	case *UpdateScheduleResponse:
		that1 = t
	case UpdateScheduleResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PatchScheduleRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PatchScheduleRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PatchScheduleRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PatchScheduleRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PatchScheduleRequest
	switch t := that.(type) {
	case *PatchScheduleRequest:
		that1 = t
	case PatchScheduleRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PatchScheduleResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PatchScheduleResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PatchScheduleResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PatchScheduleResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PatchScheduleResponse
	switch t := that.(type) {
	case *PatchScheduleResponse:
		that1 = t
	case PatchScheduleResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListScheduleMatchingTimesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListScheduleMatchingTimesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListScheduleMatchingTimesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListScheduleMatchingTimesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListScheduleMatchingTimesRequest
	switch t := that.(type) {
	case *ListScheduleMatchingTimesRequest:
		that1 = t
	case ListScheduleMatchingTimesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListScheduleMatchingTimesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListScheduleMatchingTimesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListScheduleMatchingTimesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListScheduleMatchingTimesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListScheduleMatchingTimesResponse
	switch t := that.(type) {
	case *ListScheduleMatchingTimesResponse:
		that1 = t
	case ListScheduleMatchingTimesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DeleteScheduleRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DeleteScheduleRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DeleteScheduleRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteScheduleRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteScheduleRequest
	switch t := that.(type) {
	case *DeleteScheduleRequest:
		that1 = t
	case DeleteScheduleRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DeleteScheduleResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DeleteScheduleResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DeleteScheduleResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeleteScheduleResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeleteScheduleResponse
	switch t := that.(type) {
	case *DeleteScheduleResponse:
		that1 = t
	case DeleteScheduleResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListSchedulesRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListSchedulesRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListSchedulesRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListSchedulesRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListSchedulesRequest
	switch t := that.(type) {
	case *ListSchedulesRequest:
		that1 = t
	case ListSchedulesRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListSchedulesResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListSchedulesResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListSchedulesResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListSchedulesResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListSchedulesResponse
	switch t := that.(type) {
	case *ListSchedulesResponse:
		that1 = t
	case ListSchedulesResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateWorkerBuildIdCompatibilityRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateWorkerBuildIdCompatibilityRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateWorkerBuildIdCompatibilityRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkerBuildIdCompatibilityRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkerBuildIdCompatibilityRequest
	switch t := that.(type) {
	case *UpdateWorkerBuildIdCompatibilityRequest:
		that1 = t
	case UpdateWorkerBuildIdCompatibilityRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateWorkerBuildIdCompatibilityResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateWorkerBuildIdCompatibilityResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateWorkerBuildIdCompatibilityResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkerBuildIdCompatibilityResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkerBuildIdCompatibilityResponse
	switch t := that.(type) {
	case *UpdateWorkerBuildIdCompatibilityResponse:
		that1 = t
	case UpdateWorkerBuildIdCompatibilityResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkerBuildIdCompatibilityRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkerBuildIdCompatibilityRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkerBuildIdCompatibilityRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkerBuildIdCompatibilityRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkerBuildIdCompatibilityRequest
	switch t := that.(type) {
	case *GetWorkerBuildIdCompatibilityRequest:
		that1 = t
	case GetWorkerBuildIdCompatibilityRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkerBuildIdCompatibilityResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkerBuildIdCompatibilityResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkerBuildIdCompatibilityResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkerBuildIdCompatibilityResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkerBuildIdCompatibilityResponse
	switch t := that.(type) {
	case *GetWorkerBuildIdCompatibilityResponse:
		that1 = t
	case GetWorkerBuildIdCompatibilityResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkerTaskReachabilityRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkerTaskReachabilityRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkerTaskReachabilityRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkerTaskReachabilityRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkerTaskReachabilityRequest
	switch t := that.(type) {
	case *GetWorkerTaskReachabilityRequest:
		that1 = t
	case GetWorkerTaskReachabilityRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *GetWorkerTaskReachabilityResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *GetWorkerTaskReachabilityResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two GetWorkerTaskReachabilityResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *GetWorkerTaskReachabilityResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *GetWorkerTaskReachabilityResponse
	switch t := that.(type) {
	case *GetWorkerTaskReachabilityResponse:
		that1 = t
	case GetWorkerTaskReachabilityResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateWorkflowExecutionRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateWorkflowExecutionRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateWorkflowExecutionRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkflowExecutionRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkflowExecutionRequest
	switch t := that.(type) {
	case *UpdateWorkflowExecutionRequest:
		that1 = t
	case UpdateWorkflowExecutionRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpdateWorkflowExecutionResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpdateWorkflowExecutionResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two UpdateWorkflowExecutionResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateWorkflowExecutionResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateWorkflowExecutionResponse
	switch t := that.(type) {
	case *UpdateWorkflowExecutionResponse:
		that1 = t
	case UpdateWorkflowExecutionResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StartBatchOperationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StartBatchOperationRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StartBatchOperationRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartBatchOperationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartBatchOperationRequest
	switch t := that.(type) {
	case *StartBatchOperationRequest:
		that1 = t
	case StartBatchOperationRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StartBatchOperationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StartBatchOperationResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StartBatchOperationResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartBatchOperationResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartBatchOperationResponse
	switch t := that.(type) {
	case *StartBatchOperationResponse:
		that1 = t
	case StartBatchOperationResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StopBatchOperationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StopBatchOperationRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StopBatchOperationRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StopBatchOperationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StopBatchOperationRequest
	switch t := that.(type) {
	case *StopBatchOperationRequest:
		that1 = t
	case StopBatchOperationRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StopBatchOperationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StopBatchOperationResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StopBatchOperationResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StopBatchOperationResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StopBatchOperationResponse
	switch t := that.(type) {
	case *StopBatchOperationResponse:
		that1 = t
	case StopBatchOperationResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeBatchOperationRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeBatchOperationRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeBatchOperationRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeBatchOperationRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeBatchOperationRequest
	switch t := that.(type) {
	case *DescribeBatchOperationRequest:
		that1 = t
	case DescribeBatchOperationRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *DescribeBatchOperationResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *DescribeBatchOperationResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two DescribeBatchOperationResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DescribeBatchOperationResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DescribeBatchOperationResponse
	switch t := that.(type) {
	case *DescribeBatchOperationResponse:
		that1 = t
	case DescribeBatchOperationResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListBatchOperationsRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListBatchOperationsRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListBatchOperationsRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListBatchOperationsRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListBatchOperationsRequest
	switch t := that.(type) {
	case *ListBatchOperationsRequest:
		that1 = t
	case ListBatchOperationsRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ListBatchOperationsResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ListBatchOperationsResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ListBatchOperationsResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ListBatchOperationsResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ListBatchOperationsResponse
	switch t := that.(type) {
	case *ListBatchOperationsResponse:
		that1 = t
	case ListBatchOperationsResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PollWorkflowExecutionUpdateRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PollWorkflowExecutionUpdateRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PollWorkflowExecutionUpdateRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowExecutionUpdateRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowExecutionUpdateRequest
	switch t := that.(type) {
	case *PollWorkflowExecutionUpdateRequest:
		that1 = t
	case PollWorkflowExecutionUpdateRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PollWorkflowExecutionUpdateResponse) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PollWorkflowExecutionUpdateResponse) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PollWorkflowExecutionUpdateResponse values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollWorkflowExecutionUpdateResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollWorkflowExecutionUpdateResponse
	switch t := that.(type) {
	case *PollWorkflowExecutionUpdateResponse:
		that1 = t
	case PollWorkflowExecutionUpdateResponse:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
