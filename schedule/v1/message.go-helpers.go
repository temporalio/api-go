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

package schedule

import "google.golang.org/protobuf/proto"

func (val *CalendarSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *CalendarSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two CalendarSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *CalendarSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *CalendarSpec
	switch t := that.(type) {
	case *CalendarSpec:
		that1 = t
	case CalendarSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *Range) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *Range) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two Range values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Range) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Range
	switch t := that.(type) {
	case *Range:
		that1 = t
	case Range:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *StructuredCalendarSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *StructuredCalendarSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two StructuredCalendarSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StructuredCalendarSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StructuredCalendarSpec
	switch t := that.(type) {
	case *StructuredCalendarSpec:
		that1 = t
	case StructuredCalendarSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *IntervalSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *IntervalSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two IntervalSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *IntervalSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *IntervalSpec
	switch t := that.(type) {
	case *IntervalSpec:
		that1 = t
	case IntervalSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScheduleSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScheduleSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleSpec
	switch t := that.(type) {
	case *ScheduleSpec:
		that1 = t
	case ScheduleSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SchedulePolicies) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SchedulePolicies) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two SchedulePolicies values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SchedulePolicies) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SchedulePolicies
	switch t := that.(type) {
	case *SchedulePolicies:
		that1 = t
	case SchedulePolicies:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScheduleAction) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleAction) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScheduleAction values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleAction
	switch t := that.(type) {
	case *ScheduleAction:
		that1 = t
	case ScheduleAction:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScheduleActionResult) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleActionResult) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScheduleActionResult values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleActionResult) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleActionResult
	switch t := that.(type) {
	case *ScheduleActionResult:
		that1 = t
	case ScheduleActionResult:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScheduleState) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleState) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScheduleState values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleState) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleState
	switch t := that.(type) {
	case *ScheduleState:
		that1 = t
	case ScheduleState:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *TriggerImmediatelyRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *TriggerImmediatelyRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two TriggerImmediatelyRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *TriggerImmediatelyRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *TriggerImmediatelyRequest
	switch t := that.(type) {
	case *TriggerImmediatelyRequest:
		that1 = t
	case TriggerImmediatelyRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *BackfillRequest) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *BackfillRequest) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two BackfillRequest values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BackfillRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BackfillRequest
	switch t := that.(type) {
	case *BackfillRequest:
		that1 = t
	case BackfillRequest:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *SchedulePatch) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *SchedulePatch) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two SchedulePatch values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *SchedulePatch) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *SchedulePatch
	switch t := that.(type) {
	case *SchedulePatch:
		that1 = t
	case SchedulePatch:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScheduleInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScheduleInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleInfo
	switch t := that.(type) {
	case *ScheduleInfo:
		that1 = t
	case ScheduleInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *Schedule) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *Schedule) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two Schedule values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Schedule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Schedule
	switch t := that.(type) {
	case *Schedule:
		that1 = t
	case Schedule:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScheduleListInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleListInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScheduleListInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleListInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleListInfo
	switch t := that.(type) {
	case *ScheduleListInfo:
		that1 = t
	case ScheduleListInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
func (val *ScheduleListEntry) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

func (val *ScheduleListEntry) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Equal returns whether two ScheduleListEntry values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ScheduleListEntry) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ScheduleListEntry
	switch t := that.(type) {
	case *ScheduleListEntry:
		that1 = t
	case ScheduleListEntry:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
