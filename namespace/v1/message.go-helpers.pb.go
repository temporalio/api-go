// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package namespace

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type NamespaceInfo to the protobuf v3 wire format
func (val *NamespaceInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NamespaceInfo from the protobuf v3 wire format
func (val *NamespaceInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NamespaceInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NamespaceInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceInfo
	switch t := that.(type) {
	case *NamespaceInfo:
		that1 = t
	case NamespaceInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NamespaceConfig to the protobuf v3 wire format
func (val *NamespaceConfig) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NamespaceConfig from the protobuf v3 wire format
func (val *NamespaceConfig) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NamespaceConfig) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NamespaceConfig values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceConfig
	switch t := that.(type) {
	case *NamespaceConfig:
		that1 = t
	case NamespaceConfig:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BadBinaries to the protobuf v3 wire format
func (val *BadBinaries) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BadBinaries from the protobuf v3 wire format
func (val *BadBinaries) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BadBinaries) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BadBinaries values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BadBinaries) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BadBinaries
	switch t := that.(type) {
	case *BadBinaries:
		that1 = t
	case BadBinaries:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type BadBinaryInfo to the protobuf v3 wire format
func (val *BadBinaryInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type BadBinaryInfo from the protobuf v3 wire format
func (val *BadBinaryInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *BadBinaryInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two BadBinaryInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *BadBinaryInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *BadBinaryInfo
	switch t := that.(type) {
	case *BadBinaryInfo:
		that1 = t
	case BadBinaryInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateNamespaceInfo to the protobuf v3 wire format
func (val *UpdateNamespaceInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateNamespaceInfo from the protobuf v3 wire format
func (val *UpdateNamespaceInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateNamespaceInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateNamespaceInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateNamespaceInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateNamespaceInfo
	switch t := that.(type) {
	case *UpdateNamespaceInfo:
		that1 = t
	case UpdateNamespaceInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type NamespaceFilter to the protobuf v3 wire format
func (val *NamespaceFilter) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type NamespaceFilter from the protobuf v3 wire format
func (val *NamespaceFilter) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *NamespaceFilter) Size() int {
	return proto.Size(val)
}

// Equal returns whether two NamespaceFilter values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *NamespaceFilter) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *NamespaceFilter
	switch t := that.(type) {
	case *NamespaceFilter:
		that1 = t
	case NamespaceFilter:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
