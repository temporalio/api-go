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

package command

import "google.golang.org/protobuf/proto"

func (val *ScheduleActivityTaskCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleActivityTaskCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ScheduleActivityTaskCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ScheduleActivityTaskCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleActivityTaskCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleActivityTaskCommandAttributes
	switch t := that.(type) {
	case *ScheduleActivityTaskCommandAttributes:
		that1 = t
	case ScheduleActivityTaskCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RequestCancelActivityTaskCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RequestCancelActivityTaskCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *RequestCancelActivityTaskCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RequestCancelActivityTaskCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelActivityTaskCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelActivityTaskCommandAttributes
	switch t := that.(type) {
	case *RequestCancelActivityTaskCommandAttributes:
		that1 = t
	case RequestCancelActivityTaskCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StartTimerCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StartTimerCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *StartTimerCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartTimerCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartTimerCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartTimerCommandAttributes
	switch t := that.(type) {
	case *StartTimerCommandAttributes:
		that1 = t
	case StartTimerCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CompleteWorkflowExecutionCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CompleteWorkflowExecutionCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *CompleteWorkflowExecutionCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CompleteWorkflowExecutionCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CompleteWorkflowExecutionCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CompleteWorkflowExecutionCommandAttributes
	switch t := that.(type) {
	case *CompleteWorkflowExecutionCommandAttributes:
		that1 = t
	case CompleteWorkflowExecutionCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *FailWorkflowExecutionCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *FailWorkflowExecutionCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *FailWorkflowExecutionCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two FailWorkflowExecutionCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *FailWorkflowExecutionCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *FailWorkflowExecutionCommandAttributes
	switch t := that.(type) {
	case *FailWorkflowExecutionCommandAttributes:
		that1 = t
	case FailWorkflowExecutionCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CancelTimerCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CancelTimerCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *CancelTimerCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelTimerCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelTimerCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelTimerCommandAttributes
	switch t := that.(type) {
	case *CancelTimerCommandAttributes:
		that1 = t
	case CancelTimerCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *CancelWorkflowExecutionCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CancelWorkflowExecutionCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *CancelWorkflowExecutionCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two CancelWorkflowExecutionCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CancelWorkflowExecutionCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CancelWorkflowExecutionCommandAttributes
	switch t := that.(type) {
	case *CancelWorkflowExecutionCommandAttributes:
		that1 = t
	case CancelWorkflowExecutionCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RequestCancelExternalWorkflowExecutionCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RequestCancelExternalWorkflowExecutionCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *RequestCancelExternalWorkflowExecutionCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RequestCancelExternalWorkflowExecutionCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RequestCancelExternalWorkflowExecutionCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RequestCancelExternalWorkflowExecutionCommandAttributes
	switch t := that.(type) {
	case *RequestCancelExternalWorkflowExecutionCommandAttributes:
		that1 = t
	case RequestCancelExternalWorkflowExecutionCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SignalExternalWorkflowExecutionCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SignalExternalWorkflowExecutionCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *SignalExternalWorkflowExecutionCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two SignalExternalWorkflowExecutionCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SignalExternalWorkflowExecutionCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SignalExternalWorkflowExecutionCommandAttributes
	switch t := that.(type) {
	case *SignalExternalWorkflowExecutionCommandAttributes:
		that1 = t
	case SignalExternalWorkflowExecutionCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *UpsertWorkflowSearchAttributesCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *UpsertWorkflowSearchAttributesCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *UpsertWorkflowSearchAttributesCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpsertWorkflowSearchAttributesCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpsertWorkflowSearchAttributesCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpsertWorkflowSearchAttributesCommandAttributes
	switch t := that.(type) {
	case *UpsertWorkflowSearchAttributesCommandAttributes:
		that1 = t
	case UpsertWorkflowSearchAttributesCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ModifyWorkflowPropertiesCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ModifyWorkflowPropertiesCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ModifyWorkflowPropertiesCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ModifyWorkflowPropertiesCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ModifyWorkflowPropertiesCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ModifyWorkflowPropertiesCommandAttributes
	switch t := that.(type) {
	case *ModifyWorkflowPropertiesCommandAttributes:
		that1 = t
	case ModifyWorkflowPropertiesCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *RecordMarkerCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *RecordMarkerCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *RecordMarkerCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RecordMarkerCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RecordMarkerCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RecordMarkerCommandAttributes
	switch t := that.(type) {
	case *RecordMarkerCommandAttributes:
		that1 = t
	case RecordMarkerCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ContinueAsNewWorkflowExecutionCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ContinueAsNewWorkflowExecutionCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ContinueAsNewWorkflowExecutionCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ContinueAsNewWorkflowExecutionCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ContinueAsNewWorkflowExecutionCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ContinueAsNewWorkflowExecutionCommandAttributes
	switch t := that.(type) {
	case *ContinueAsNewWorkflowExecutionCommandAttributes:
		that1 = t
	case ContinueAsNewWorkflowExecutionCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StartChildWorkflowExecutionCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StartChildWorkflowExecutionCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *StartChildWorkflowExecutionCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StartChildWorkflowExecutionCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StartChildWorkflowExecutionCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StartChildWorkflowExecutionCommandAttributes
	switch t := that.(type) {
	case *StartChildWorkflowExecutionCommandAttributes:
		that1 = t
	case StartChildWorkflowExecutionCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ProtocolMessageCommandAttributes) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ProtocolMessageCommandAttributes) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *ProtocolMessageCommandAttributes) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ProtocolMessageCommandAttributes values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ProtocolMessageCommandAttributes) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ProtocolMessageCommandAttributes
	switch t := that.(type) {
	case *ProtocolMessageCommandAttributes:
		that1 = t
	case ProtocolMessageCommandAttributes:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *Command) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *Command) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

func (val *Command) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Command values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Command) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Command
	switch t := that.(type) {
	case *Command:
		that1 = t
	case Command:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
