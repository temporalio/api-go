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

package errordetails

import "google.golang.org/protobuf/proto"

func (val *NotFoundFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NotFoundFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NotFoundFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NotFoundFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NotFoundFailure
	switch t := that.(type) {
	case *NotFoundFailure:
		that1 = t
	case NotFoundFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *WorkflowExecutionAlreadyStartedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *WorkflowExecutionAlreadyStartedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two WorkflowExecutionAlreadyStartedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionAlreadyStartedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionAlreadyStartedFailure
	switch t := that.(type) {
	case *WorkflowExecutionAlreadyStartedFailure:
		that1 = t
	case WorkflowExecutionAlreadyStartedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NamespaceNotActiveFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceNotActiveFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NamespaceNotActiveFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceNotActiveFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceNotActiveFailure
	switch t := that.(type) {
	case *NamespaceNotActiveFailure:
		that1 = t
	case NamespaceNotActiveFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NamespaceInvalidStateFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceInvalidStateFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NamespaceInvalidStateFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceInvalidStateFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceInvalidStateFailure
	switch t := that.(type) {
	case *NamespaceInvalidStateFailure:
		that1 = t
	case NamespaceInvalidStateFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NamespaceNotFoundFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceNotFoundFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NamespaceNotFoundFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceNotFoundFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceNotFoundFailure
	switch t := that.(type) {
	case *NamespaceNotFoundFailure:
		that1 = t
	case NamespaceNotFoundFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NamespaceAlreadyExistsFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NamespaceAlreadyExistsFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NamespaceAlreadyExistsFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceAlreadyExistsFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceAlreadyExistsFailure
	switch t := that.(type) {
	case *NamespaceAlreadyExistsFailure:
		that1 = t
	case NamespaceAlreadyExistsFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ClientVersionNotSupportedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ClientVersionNotSupportedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ClientVersionNotSupportedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ClientVersionNotSupportedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ClientVersionNotSupportedFailure
	switch t := that.(type) {
	case *ClientVersionNotSupportedFailure:
		that1 = t
	case ClientVersionNotSupportedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ServerVersionNotSupportedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ServerVersionNotSupportedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ServerVersionNotSupportedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ServerVersionNotSupportedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ServerVersionNotSupportedFailure
	switch t := that.(type) {
	case *ServerVersionNotSupportedFailure:
		that1 = t
	case ServerVersionNotSupportedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CancellationAlreadyRequestedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CancellationAlreadyRequestedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two CancellationAlreadyRequestedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancellationAlreadyRequestedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancellationAlreadyRequestedFailure
	switch t := that.(type) {
	case *CancellationAlreadyRequestedFailure:
		that1 = t
	case CancellationAlreadyRequestedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *QueryFailedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *QueryFailedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two QueryFailedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *QueryFailedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *QueryFailedFailure
	switch t := that.(type) {
	case *QueryFailedFailure:
		that1 = t
	case QueryFailedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *PermissionDeniedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *PermissionDeniedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two PermissionDeniedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PermissionDeniedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PermissionDeniedFailure
	switch t := that.(type) {
	case *PermissionDeniedFailure:
		that1 = t
	case PermissionDeniedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ResourceExhaustedFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ResourceExhaustedFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ResourceExhaustedFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResourceExhaustedFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResourceExhaustedFailure
	switch t := that.(type) {
	case *ResourceExhaustedFailure:
		that1 = t
	case ResourceExhaustedFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SystemWorkflowFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SystemWorkflowFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two SystemWorkflowFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SystemWorkflowFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SystemWorkflowFailure
	switch t := that.(type) {
	case *SystemWorkflowFailure:
		that1 = t
	case SystemWorkflowFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *WorkflowNotReadyFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *WorkflowNotReadyFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two WorkflowNotReadyFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowNotReadyFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowNotReadyFailure
	switch t := that.(type) {
	case *WorkflowNotReadyFailure:
		that1 = t
	case WorkflowNotReadyFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *NewerBuildExistsFailure) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *NewerBuildExistsFailure) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two NewerBuildExistsFailure values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NewerBuildExistsFailure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NewerBuildExistsFailure
	switch t := that.(type) {
	case *NewerBuildExistsFailure:
		that1 = t
	case NewerBuildExistsFailure:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
