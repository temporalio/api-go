// Code generated by protoc-gen-go-helpers. DO NOT EDIT.
package deployment

import (
	"google.golang.org/protobuf/proto"
)

// Marshal an object of type WorkerDeploymentOptions to the protobuf v3 wire format
func (val *WorkerDeploymentOptions) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkerDeploymentOptions from the protobuf v3 wire format
func (val *WorkerDeploymentOptions) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkerDeploymentOptions) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkerDeploymentOptions values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkerDeploymentOptions) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkerDeploymentOptions
	switch t := that.(type) {
	case *WorkerDeploymentOptions:
		that1 = t
	case WorkerDeploymentOptions:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type Deployment to the protobuf v3 wire format
func (val *Deployment) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type Deployment from the protobuf v3 wire format
func (val *Deployment) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *Deployment) Size() int {
	return proto.Size(val)
}

// Equal returns whether two Deployment values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *Deployment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *Deployment
	switch t := that.(type) {
	case *Deployment:
		that1 = t
	case Deployment:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeploymentInfo to the protobuf v3 wire format
func (val *DeploymentInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeploymentInfo from the protobuf v3 wire format
func (val *DeploymentInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeploymentInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeploymentInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeploymentInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeploymentInfo
	switch t := that.(type) {
	case *DeploymentInfo:
		that1 = t
	case DeploymentInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type UpdateDeploymentMetadata to the protobuf v3 wire format
func (val *UpdateDeploymentMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type UpdateDeploymentMetadata from the protobuf v3 wire format
func (val *UpdateDeploymentMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *UpdateDeploymentMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two UpdateDeploymentMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *UpdateDeploymentMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *UpdateDeploymentMetadata
	switch t := that.(type) {
	case *UpdateDeploymentMetadata:
		that1 = t
	case UpdateDeploymentMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type DeploymentListInfo to the protobuf v3 wire format
func (val *DeploymentListInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type DeploymentListInfo from the protobuf v3 wire format
func (val *DeploymentListInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *DeploymentListInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two DeploymentListInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *DeploymentListInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *DeploymentListInfo
	switch t := that.(type) {
	case *DeploymentListInfo:
		that1 = t
	case DeploymentListInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkerDeploymentVersionInfo to the protobuf v3 wire format
func (val *WorkerDeploymentVersionInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkerDeploymentVersionInfo from the protobuf v3 wire format
func (val *WorkerDeploymentVersionInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkerDeploymentVersionInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkerDeploymentVersionInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkerDeploymentVersionInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkerDeploymentVersionInfo
	switch t := that.(type) {
	case *WorkerDeploymentVersionInfo:
		that1 = t
	case WorkerDeploymentVersionInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VersionDrainageInfo to the protobuf v3 wire format
func (val *VersionDrainageInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VersionDrainageInfo from the protobuf v3 wire format
func (val *VersionDrainageInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VersionDrainageInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersionDrainageInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersionDrainageInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersionDrainageInfo
	switch t := that.(type) {
	case *VersionDrainageInfo:
		that1 = t
	case VersionDrainageInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkerDeploymentInfo to the protobuf v3 wire format
func (val *WorkerDeploymentInfo) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkerDeploymentInfo from the protobuf v3 wire format
func (val *WorkerDeploymentInfo) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkerDeploymentInfo) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkerDeploymentInfo values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkerDeploymentInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkerDeploymentInfo
	switch t := that.(type) {
	case *WorkerDeploymentInfo:
		that1 = t
	case WorkerDeploymentInfo:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type WorkerDeploymentVersion to the protobuf v3 wire format
func (val *WorkerDeploymentVersion) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type WorkerDeploymentVersion from the protobuf v3 wire format
func (val *WorkerDeploymentVersion) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *WorkerDeploymentVersion) Size() int {
	return proto.Size(val)
}

// Equal returns whether two WorkerDeploymentVersion values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *WorkerDeploymentVersion) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *WorkerDeploymentVersion
	switch t := that.(type) {
	case *WorkerDeploymentVersion:
		that1 = t
	case WorkerDeploymentVersion:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type VersionMetadata to the protobuf v3 wire format
func (val *VersionMetadata) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type VersionMetadata from the protobuf v3 wire format
func (val *VersionMetadata) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *VersionMetadata) Size() int {
	return proto.Size(val)
}

// Equal returns whether two VersionMetadata values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *VersionMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *VersionMetadata
	switch t := that.(type) {
	case *VersionMetadata:
		that1 = t
	case VersionMetadata:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}

// Marshal an object of type RoutingConfig to the protobuf v3 wire format
func (val *RoutingConfig) Marshal() ([]byte, error) {
	return proto.Marshal(val)
}

// Unmarshal an object of type RoutingConfig from the protobuf v3 wire format
func (val *RoutingConfig) Unmarshal(buf []byte) error {
	return proto.Unmarshal(buf, val)
}

// Size returns the size of the object, in bytes, once serialized
func (val *RoutingConfig) Size() int {
	return proto.Size(val)
}

// Equal returns whether two RoutingConfig values are equivalent by recursively
// comparing the message's fields.
// For more information see the documentation for
// https://pkg.go.dev/google.golang.org/protobuf/proto#Equal
func (this *RoutingConfig) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	var that1 *RoutingConfig
	switch t := that.(type) {
	case *RoutingConfig:
		that1 = t
	case RoutingConfig:
		that1 = &t
	default:
		return false
	}

	return proto.Equal(this, that1)
}
