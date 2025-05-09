// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package rules

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type WorkflowRuleAction to the protobuf v3 wire format
func (val *WorkflowRuleAction) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowRuleAction from the protobuf v3 wire format
func (val *WorkflowRuleAction) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowRuleAction) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowRuleAction values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowRuleAction) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowRuleAction
	switch t := that.(type) {
	case *WorkflowRuleAction:
		that1 = t
	case WorkflowRuleAction:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowRuleSpec to the protobuf v3 wire format
func (val *WorkflowRuleSpec) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowRuleSpec from the protobuf v3 wire format
func (val *WorkflowRuleSpec) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowRuleSpec) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowRuleSpec values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowRuleSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowRuleSpec
	switch t := that.(type) {
	case *WorkflowRuleSpec:
		that1 = t
	case WorkflowRuleSpec:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkflowRule to the protobuf v3 wire format
func (val *WorkflowRule) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkflowRule from the protobuf v3 wire format
func (val *WorkflowRule) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkflowRule) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkflowRule values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkflowRule) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkflowRule
	switch t := that.(type) {
	case *WorkflowRule:
		that1 = t
	case WorkflowRule:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
