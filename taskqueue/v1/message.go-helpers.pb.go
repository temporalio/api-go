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

// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package taskqueue

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type TaskQueue to the protobuf v3 wire format
func (val *TaskQueue) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueue from the protobuf v3 wire format
func (val *TaskQueue) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueue) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueue values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueue) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueue
	switch t := that.(type) {
	case *TaskQueue:
		that1 = t
	case TaskQueue:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskQueueMetadata to the protobuf v3 wire format
func (val *TaskQueueMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueueMetadata from the protobuf v3 wire format
func (val *TaskQueueMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueueMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueueMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueueMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueueMetadata
	switch t := that.(type) {
	case *TaskQueueMetadata:
		that1 = t
	case TaskQueueMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskQueueVersionGroup to the protobuf v3 wire format
func (val *TaskQueueVersionGroup) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueueVersionGroup from the protobuf v3 wire format
func (val *TaskQueueVersionGroup) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueueVersionGroup) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueueVersionGroup values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueueVersionGroup) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueueVersionGroup
	switch t := that.(type) {
	case *TaskQueueVersionGroup:
		that1 = t
	case TaskQueueVersionGroup:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskQueueVersionInfo to the protobuf v3 wire format
func (val *TaskQueueVersionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueueVersionInfo from the protobuf v3 wire format
func (val *TaskQueueVersionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueueVersionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueueVersionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueueVersionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueueVersionInfo
	switch t := that.(type) {
	case *TaskQueueVersionInfo:
		that1 = t
	case TaskQueueVersionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskQueueTypeInfo to the protobuf v3 wire format
func (val *TaskQueueTypeInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueueTypeInfo from the protobuf v3 wire format
func (val *TaskQueueTypeInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueueTypeInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueueTypeInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueueTypeInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueueTypeInfo
	switch t := that.(type) {
	case *TaskQueueTypeInfo:
		that1 = t
	case TaskQueueTypeInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BacklogInfo to the protobuf v3 wire format
func (val *BacklogInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BacklogInfo from the protobuf v3 wire format
func (val *BacklogInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BacklogInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BacklogInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BacklogInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BacklogInfo
	switch t := that.(type) {
	case *BacklogInfo:
		that1 = t
	case BacklogInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskQueueStatus to the protobuf v3 wire format
func (val *TaskQueueStatus) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueueStatus from the protobuf v3 wire format
func (val *TaskQueueStatus) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueueStatus) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueueStatus values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueueStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueueStatus
	switch t := that.(type) {
	case *TaskQueueStatus:
		that1 = t
	case TaskQueueStatus:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskIdBlock to the protobuf v3 wire format
func (val *TaskIdBlock) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskIdBlock from the protobuf v3 wire format
func (val *TaskIdBlock) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskIdBlock) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskIdBlock values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskIdBlock) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskIdBlock
	switch t := that.(type) {
	case *TaskIdBlock:
		that1 = t
	case TaskIdBlock:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskQueuePartitionMetadata to the protobuf v3 wire format
func (val *TaskQueuePartitionMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueuePartitionMetadata from the protobuf v3 wire format
func (val *TaskQueuePartitionMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueuePartitionMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueuePartitionMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueuePartitionMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueuePartitionMetadata
	switch t := that.(type) {
	case *TaskQueuePartitionMetadata:
		that1 = t
	case TaskQueuePartitionMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type PollerInfo to the protobuf v3 wire format
func (val *PollerInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type PollerInfo from the protobuf v3 wire format
func (val *PollerInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *PollerInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two PollerInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *PollerInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *PollerInfo
	switch t := that.(type) {
	case *PollerInfo:
		that1 = t
	case PollerInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StickyExecutionAttributes to the protobuf v3 wire format
func (val *StickyExecutionAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StickyExecutionAttributes from the protobuf v3 wire format
func (val *StickyExecutionAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StickyExecutionAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StickyExecutionAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StickyExecutionAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StickyExecutionAttributes
	switch t := that.(type) {
	case *StickyExecutionAttributes:
		that1 = t
	case StickyExecutionAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CompatibleVersionSet to the protobuf v3 wire format
func (val *CompatibleVersionSet) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CompatibleVersionSet from the protobuf v3 wire format
func (val *CompatibleVersionSet) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CompatibleVersionSet) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CompatibleVersionSet values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CompatibleVersionSet) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CompatibleVersionSet
	switch t := that.(type) {
	case *CompatibleVersionSet:
		that1 = t
	case CompatibleVersionSet:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TaskQueueReachability to the protobuf v3 wire format
func (val *TaskQueueReachability) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TaskQueueReachability from the protobuf v3 wire format
func (val *TaskQueueReachability) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TaskQueueReachability) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TaskQueueReachability values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TaskQueueReachability) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TaskQueueReachability
	switch t := that.(type) {
	case *TaskQueueReachability:
		that1 = t
	case TaskQueueReachability:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BuildIdReachability to the protobuf v3 wire format
func (val *BuildIdReachability) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BuildIdReachability from the protobuf v3 wire format
func (val *BuildIdReachability) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BuildIdReachability) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BuildIdReachability values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BuildIdReachability) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BuildIdReachability
	switch t := that.(type) {
	case *BuildIdReachability:
		that1 = t
	case BuildIdReachability:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RampByPercentage to the protobuf v3 wire format
func (val *RampByPercentage) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RampByPercentage from the protobuf v3 wire format
func (val *RampByPercentage) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RampByPercentage) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RampByPercentage values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RampByPercentage) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RampByPercentage
	switch t := that.(type) {
	case *RampByPercentage:
		that1 = t
	case RampByPercentage:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RampByWorkerRatio to the protobuf v3 wire format
func (val *RampByWorkerRatio) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RampByWorkerRatio from the protobuf v3 wire format
func (val *RampByWorkerRatio) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RampByWorkerRatio) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RampByWorkerRatio values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RampByWorkerRatio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RampByWorkerRatio
	switch t := that.(type) {
	case *RampByWorkerRatio:
		that1 = t
	case RampByWorkerRatio:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BuildIdAssignmentRule to the protobuf v3 wire format
func (val *BuildIdAssignmentRule) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BuildIdAssignmentRule from the protobuf v3 wire format
func (val *BuildIdAssignmentRule) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BuildIdAssignmentRule) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BuildIdAssignmentRule values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BuildIdAssignmentRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BuildIdAssignmentRule
	switch t := that.(type) {
	case *BuildIdAssignmentRule:
		that1 = t
	case BuildIdAssignmentRule:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type CompatibleBuildIdRedirectRule to the protobuf v3 wire format
func (val *CompatibleBuildIdRedirectRule) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type CompatibleBuildIdRedirectRule from the protobuf v3 wire format
func (val *CompatibleBuildIdRedirectRule) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *CompatibleBuildIdRedirectRule) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CompatibleBuildIdRedirectRule values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CompatibleBuildIdRedirectRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CompatibleBuildIdRedirectRule
	switch t := that.(type) {
	case *CompatibleBuildIdRedirectRule:
		that1 = t
	case CompatibleBuildIdRedirectRule:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TimestampedBuildIdAssignmentRule to the protobuf v3 wire format
func (val *TimestampedBuildIdAssignmentRule) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TimestampedBuildIdAssignmentRule from the protobuf v3 wire format
func (val *TimestampedBuildIdAssignmentRule) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TimestampedBuildIdAssignmentRule) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimestampedBuildIdAssignmentRule values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimestampedBuildIdAssignmentRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimestampedBuildIdAssignmentRule
	switch t := that.(type) {
	case *TimestampedBuildIdAssignmentRule:
		that1 = t
	case TimestampedBuildIdAssignmentRule:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type TimestampedCompatibleBuildIdRedirectRule to the protobuf v3 wire format
func (val *TimestampedCompatibleBuildIdRedirectRule) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type TimestampedCompatibleBuildIdRedirectRule from the protobuf v3 wire format
func (val *TimestampedCompatibleBuildIdRedirectRule) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *TimestampedCompatibleBuildIdRedirectRule) Size() int {
	return proto.Size(val)
}

// Equal returns whether two TimestampedCompatibleBuildIdRedirectRule values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TimestampedCompatibleBuildIdRedirectRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TimestampedCompatibleBuildIdRedirectRule
	switch t := that.(type) {
	case *TimestampedCompatibleBuildIdRedirectRule:
		that1 = t
	case TimestampedCompatibleBuildIdRedirectRule:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
