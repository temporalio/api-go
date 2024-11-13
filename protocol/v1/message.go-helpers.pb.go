// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package protocol

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type Message to the protobuf v3 wire format
func (val *Message) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Message from the protobuf v3 wire format
func (val *Message) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Message) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Message values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Message) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Message
	switch t := that.(type) {
	case *Message:
		that1 = t
	case Message:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
