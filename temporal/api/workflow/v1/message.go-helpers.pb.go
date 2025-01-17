// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package workflow

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type WorkflowExecutionInfo to the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionInfo from the protobuf v3 wire format
func (val *WorkflowExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionInfo
	switch t := that.(type) {
	case *WorkflowExecutionInfo:
		that1 = t
	case WorkflowExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionExtendedInfo to the protobuf v3 wire format
func (val *WorkflowExecutionExtendedInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionExtendedInfo from the protobuf v3 wire format
func (val *WorkflowExecutionExtendedInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionExtendedInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionExtendedInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionExtendedInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionExtendedInfo
	switch t := that.(type) {
	case *WorkflowExecutionExtendedInfo:
		that1 = t
	case WorkflowExecutionExtendedInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionVersioningInfo to the protobuf v3 wire format
func (val *WorkflowExecutionVersioningInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionVersioningInfo from the protobuf v3 wire format
func (val *WorkflowExecutionVersioningInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionVersioningInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionVersioningInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionVersioningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionVersioningInfo
	switch t := that.(type) {
	case *WorkflowExecutionVersioningInfo:
		that1 = t
	case WorkflowExecutionVersioningInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeploymentTransition to the protobuf v3 wire format
func (val *DeploymentTransition) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeploymentTransition from the protobuf v3 wire format
func (val *DeploymentTransition) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeploymentTransition) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeploymentTransition values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeploymentTransition) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeploymentTransition
	switch t := that.(type) {
	case *DeploymentTransition:
		that1 = t
	case DeploymentTransition:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionConfig to the protobuf v3 wire format
func (val *WorkflowExecutionConfig) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionConfig from the protobuf v3 wire format
func (val *WorkflowExecutionConfig) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionConfig) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionConfig values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionConfig
	switch t := that.(type) {
	case *WorkflowExecutionConfig:
		that1 = t
	case WorkflowExecutionConfig:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingActivityInfo to the protobuf v3 wire format
func (val *PendingActivityInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingActivityInfo from the protobuf v3 wire format
func (val *PendingActivityInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingActivityInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingActivityInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingActivityInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingActivityInfo
	switch t := that.(type) {
	case *PendingActivityInfo:
		that1 = t
	case PendingActivityInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingChildExecutionInfo to the protobuf v3 wire format
func (val *PendingChildExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingChildExecutionInfo from the protobuf v3 wire format
func (val *PendingChildExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingChildExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingChildExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingChildExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingChildExecutionInfo
	switch t := that.(type) {
	case *PendingChildExecutionInfo:
		that1 = t
	case PendingChildExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingWorkflowTaskInfo to the protobuf v3 wire format
func (val *PendingWorkflowTaskInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingWorkflowTaskInfo from the protobuf v3 wire format
func (val *PendingWorkflowTaskInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingWorkflowTaskInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingWorkflowTaskInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingWorkflowTaskInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingWorkflowTaskInfo
	switch t := that.(type) {
	case *PendingWorkflowTaskInfo:
		that1 = t
	case PendingWorkflowTaskInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetPoints to the protobuf v3 wire format
func (val *ResetPoints) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetPoints from the protobuf v3 wire format
func (val *ResetPoints) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetPoints) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetPoints values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetPoints) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetPoints
	switch t := that.(type) {
	case *ResetPoints:
		that1 = t
	case ResetPoints:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type ResetPointInfo to the protobuf v3 wire format
func (val *ResetPointInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ResetPointInfo from the protobuf v3 wire format
func (val *ResetPointInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ResetPointInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ResetPointInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ResetPointInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ResetPointInfo
	switch t := that.(type) {
	case *ResetPointInfo:
		that1 = t
	case ResetPointInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NewWorkflowExecutionInfo to the protobuf v3 wire format
func (val *NewWorkflowExecutionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NewWorkflowExecutionInfo from the protobuf v3 wire format
func (val *NewWorkflowExecutionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NewWorkflowExecutionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NewWorkflowExecutionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NewWorkflowExecutionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NewWorkflowExecutionInfo
	switch t := that.(type) {
	case *NewWorkflowExecutionInfo:
		that1 = t
	case NewWorkflowExecutionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CallbackInfo to the protobuf v3 wire format
func (val *CallbackInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CallbackInfo from the protobuf v3 wire format
func (val *CallbackInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CallbackInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CallbackInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CallbackInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CallbackInfo
	switch t := that.(type) {
	case *CallbackInfo:
		that1 = t
	case CallbackInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PendingNexusOperationInfo to the protobuf v3 wire format
func (val *PendingNexusOperationInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PendingNexusOperationInfo from the protobuf v3 wire format
func (val *PendingNexusOperationInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PendingNexusOperationInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PendingNexusOperationInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PendingNexusOperationInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PendingNexusOperationInfo
	switch t := that.(type) {
	case *PendingNexusOperationInfo:
		that1 = t
	case PendingNexusOperationInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NexusOperationCancellationInfo to the protobuf v3 wire format
func (val *NexusOperationCancellationInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NexusOperationCancellationInfo from the protobuf v3 wire format
func (val *NexusOperationCancellationInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NexusOperationCancellationInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NexusOperationCancellationInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NexusOperationCancellationInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NexusOperationCancellationInfo
	switch t := that.(type) {
	case *NexusOperationCancellationInfo:
		that1 = t
	case NexusOperationCancellationInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowExecutionOptions to the protobuf v3 wire format
func (val *WorkflowExecutionOptions) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowExecutionOptions from the protobuf v3 wire format
func (val *WorkflowExecutionOptions) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowExecutionOptions) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowExecutionOptions values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowExecutionOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowExecutionOptions
	switch t := that.(type) {
	case *WorkflowExecutionOptions:
		that1 = t
	case WorkflowExecutionOptions:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VersioningOverride to the protobuf v3 wire format
func (val *VersioningOverride) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VersioningOverride from the protobuf v3 wire format
func (val *VersioningOverride) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VersioningOverride) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersioningOverride values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersioningOverride) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersioningOverride
	switch t := that.(type) {
	case *VersioningOverride:
		that1 = t
	case VersioningOverride:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
