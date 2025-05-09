// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package version

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type ReleaseInfo to the protobuf v3 wire format
func (val *ReleaseInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type ReleaseInfo from the protobuf v3 wire format
func (val *ReleaseInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *ReleaseInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two ReleaseInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *ReleaseInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *ReleaseInfo
	switch t := that.(type) {
	case *ReleaseInfo:
		that1 = t
	case ReleaseInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Alert to the protobuf v3 wire format
func (val *Alert) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Alert from the protobuf v3 wire format
func (val *Alert) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Alert) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Alert values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Alert) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Alert
	switch t := that.(type) {
	case *Alert:
		that1 = t
	case Alert:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VersionInfo to the protobuf v3 wire format
func (val *VersionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VersionInfo from the protobuf v3 wire format
func (val *VersionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VersionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersionInfo
	switch t := that.(type) {
	case *VersionInfo:
		that1 = t
	case VersionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
