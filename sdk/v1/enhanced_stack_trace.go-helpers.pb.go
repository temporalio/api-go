// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package sdk

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type EnhancedStackTrace to the protobuf v3 wire format
func (val *EnhancedStackTrace) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type EnhancedStackTrace from the protobuf v3 wire format
func (val *EnhancedStackTrace) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *EnhancedStackTrace) Size() int {
	return proto.Size(val)
}

// Equal returns whether two EnhancedStackTrace values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *EnhancedStackTrace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *EnhancedStackTrace
	switch t := that.(type) {
	case *EnhancedStackTrace:
		that1 = t
	case EnhancedStackTrace:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StackTraceSDKInfo to the protobuf v3 wire format
func (val *StackTraceSDKInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StackTraceSDKInfo from the protobuf v3 wire format
func (val *StackTraceSDKInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StackTraceSDKInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StackTraceSDKInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StackTraceSDKInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StackTraceSDKInfo
	switch t := that.(type) {
	case *StackTraceSDKInfo:
		that1 = t
	case StackTraceSDKInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StackTraceFileSlice to the protobuf v3 wire format
func (val *StackTraceFileSlice) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StackTraceFileSlice from the protobuf v3 wire format
func (val *StackTraceFileSlice) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StackTraceFileSlice) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StackTraceFileSlice values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StackTraceFileSlice) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StackTraceFileSlice
	switch t := that.(type) {
	case *StackTraceFileSlice:
		that1 = t
	case StackTraceFileSlice:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StackTraceFileLocation to the protobuf v3 wire format
func (val *StackTraceFileLocation) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StackTraceFileLocation from the protobuf v3 wire format
func (val *StackTraceFileLocation) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StackTraceFileLocation) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StackTraceFileLocation values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StackTraceFileLocation) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StackTraceFileLocation
	switch t := that.(type) {
	case *StackTraceFileLocation:
		that1 = t
	case StackTraceFileLocation:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type StackTrace to the protobuf v3 wire format
func (val *StackTrace) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type StackTrace from the protobuf v3 wire format
func (val *StackTrace) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *StackTrace) Size() int {
	return proto.Size(val)
}

// Equal returns whether two StackTrace values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *StackTrace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *StackTrace
	switch t := that.(type) {
	case *StackTrace:
		that1 = t
	case StackTrace:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
