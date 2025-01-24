// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
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

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/enums/v1/deployment.proto

package enums

import (
	reflect "reflect"
	"strconv"
	sync "sync"
	unsafe "unsafe"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Specify the reachability level for a deployment so users can decide if it is time to
// decommission the deployment.
type DeploymentReachability int32

const (
	// Reachability level is not specified.
	DEPLOYMENT_REACHABILITY_UNSPECIFIED DeploymentReachability = 0
	// The deployment is reachable by new and/or open workflows. The deployment cannot be
	// decommissioned safely.
	DEPLOYMENT_REACHABILITY_REACHABLE DeploymentReachability = 1
	// The deployment is not reachable by new or open workflows, but might be still needed by
	// Queries sent to closed workflows. The deployment can be decommissioned safely if user does
	// not query closed workflows.
	DEPLOYMENT_REACHABILITY_CLOSED_WORKFLOWS_ONLY DeploymentReachability = 2
	// The deployment is not reachable by any workflow because all the workflows who needed this
	// deployment went out of retention period. The deployment can be decommissioned safely.
	DEPLOYMENT_REACHABILITY_UNREACHABLE DeploymentReachability = 3
)

// Enum value maps for DeploymentReachability.
var (
	DeploymentReachability_name = map[int32]string{
		0: "DEPLOYMENT_REACHABILITY_UNSPECIFIED",
		1: "DEPLOYMENT_REACHABILITY_REACHABLE",
		2: "DEPLOYMENT_REACHABILITY_CLOSED_WORKFLOWS_ONLY",
		3: "DEPLOYMENT_REACHABILITY_UNREACHABLE",
	}
	DeploymentReachability_value = map[string]int32{
		"DEPLOYMENT_REACHABILITY_UNSPECIFIED":           0,
		"DEPLOYMENT_REACHABILITY_REACHABLE":             1,
		"DEPLOYMENT_REACHABILITY_CLOSED_WORKFLOWS_ONLY": 2,
		"DEPLOYMENT_REACHABILITY_UNREACHABLE":           3,
	}
)

func (x DeploymentReachability) Enum() *DeploymentReachability {
	p := new(DeploymentReachability)
	*p = x
	return p
}

func (x DeploymentReachability) String() string {
	switch x {
	case DEPLOYMENT_REACHABILITY_UNSPECIFIED:
		return "Unspecified"
	case DEPLOYMENT_REACHABILITY_REACHABLE:
		return "Reachable"
	case DEPLOYMENT_REACHABILITY_CLOSED_WORKFLOWS_ONLY:
		return "ClosedWorkflowsOnly"
	case DEPLOYMENT_REACHABILITY_UNREACHABLE:
		return "Unreachable"
	default:
		return strconv.Itoa(int(x))
	}

}

func (DeploymentReachability) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_deployment_proto_enumTypes[0].Descriptor()
}

func (DeploymentReachability) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_deployment_proto_enumTypes[0]
}

func (x DeploymentReachability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentReachability.Descriptor instead.
func (DeploymentReachability) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_deployment_proto_rawDescGZIP(), []int{0}
}

// (-- api-linter: core::0216::synonyms=disabled
//
//	aip.dev/not-precedent: Call this status because it is . --)
//
// Specify the drainage status for a Worker Deployment Version so users can decide whether they
// can safely decommission the version.
type VersionDrainageStatus int32

const (
	// Drainage Status is not specified.
	VERSION_DRAINAGE_STATUS_UNSPECIFIED VersionDrainageStatus = 0
	// The Worker Deployment Version is not used by new workflows but is still used by
	// open pinned workflows. The version cannot be decommissioned safely.
	VERSION_DRAINAGE_STATUS_DRAINING VersionDrainageStatus = 1
	// The Worker Deployment Version is not used by new or open workflows, but might be still needed by
	// Queries sent to closed workflows. The version can be decommissioned safely if user does
	// not query closed workflows. If the user does query closed workflows for some time x after
	// workflows are closed, they should decommission the version after it has been drained for that duration.
	VERSION_DRAINAGE_STATUS_DRAINED VersionDrainageStatus = 2
)

// Enum value maps for VersionDrainageStatus.
var (
	VersionDrainageStatus_name = map[int32]string{
		0: "VERSION_DRAINAGE_STATUS_UNSPECIFIED",
		1: "VERSION_DRAINAGE_STATUS_DRAINING",
		2: "VERSION_DRAINAGE_STATUS_DRAINED",
	}
	VersionDrainageStatus_value = map[string]int32{
		"VERSION_DRAINAGE_STATUS_UNSPECIFIED": 0,
		"VERSION_DRAINAGE_STATUS_DRAINING":    1,
		"VERSION_DRAINAGE_STATUS_DRAINED":     2,
	}
)

func (x VersionDrainageStatus) Enum() *VersionDrainageStatus {
	p := new(VersionDrainageStatus)
	*p = x
	return p
}

func (x VersionDrainageStatus) String() string {
	switch x {
	case VERSION_DRAINAGE_STATUS_UNSPECIFIED:
		return "Unspecified"
	case VERSION_DRAINAGE_STATUS_DRAINING:
		return "Draining"
	case VERSION_DRAINAGE_STATUS_DRAINED:
		return "Drained"
	default:
		return strconv.Itoa(int(x))
	}

}

func (VersionDrainageStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_deployment_proto_enumTypes[1].Descriptor()
}

func (VersionDrainageStatus) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_deployment_proto_enumTypes[1]
}

func (x VersionDrainageStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VersionDrainageStatus.Descriptor instead.
func (VersionDrainageStatus) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_deployment_proto_rawDescGZIP(), []int{1}
}

// Workflow Versioning Mode for a particular Worker Deployment Version. This value is
// configured by the app developer in the worker code.
type WorkflowVersioningMode int32

const (
	WORKFLOW_VERSIONING_MODE_UNSPECIFIED WorkflowVersioningMode = 0
	// Workflows processed by this Deployment Version will be unversioned and user
	// needs to use Patching to keep the new code compatible with prior versions.
	// This mode is recommended to be used along with Rolling Upgrade deployment
	// strategies.
	// Deployment Versions with this mode can not be set as the Current or Ramping
	// Version of their Deployment, and are not distinguished from each other for
	// task routing.
	WORKFLOW_VERSIONING_MODE_UNVERSIONED WorkflowVersioningMode = 1
	// Workflow Versioning Behaviors are enabled in this Deployment Version. Each
	// workflow type must choose between the Pinned and AutoUpgrade behaviors.
	// Depending on the chosen behavior user may or may not need to use Patching
	// to keep the new code compatible with prior versions. (see VersioningBehavior
	// enum.)
	// Deployment Versions with this mode can be set as the Current or Ramping
	// Version of their Deployment and hence are the best option for Blue/Green
	// and Rainbow strategies (but typically not suitable for Rolling upgrades.)
	WORKFLOW_VERSIONING_MODE_VERSIONING_BEHAVIORS WorkflowVersioningMode = 2
)

// Enum value maps for WorkflowVersioningMode.
var (
	WorkflowVersioningMode_name = map[int32]string{
		0: "WORKFLOW_VERSIONING_MODE_UNSPECIFIED",
		1: "WORKFLOW_VERSIONING_MODE_UNVERSIONED",
		2: "WORKFLOW_VERSIONING_MODE_VERSIONING_BEHAVIORS",
	}
	WorkflowVersioningMode_value = map[string]int32{
		"WORKFLOW_VERSIONING_MODE_UNSPECIFIED":          0,
		"WORKFLOW_VERSIONING_MODE_UNVERSIONED":          1,
		"WORKFLOW_VERSIONING_MODE_VERSIONING_BEHAVIORS": 2,
	}
)

func (x WorkflowVersioningMode) Enum() *WorkflowVersioningMode {
	p := new(WorkflowVersioningMode)
	*p = x
	return p
}

func (x WorkflowVersioningMode) String() string {
	switch x {
	case WORKFLOW_VERSIONING_MODE_UNSPECIFIED:
		return "Unspecified"
	case WORKFLOW_VERSIONING_MODE_UNVERSIONED:
		return "Unversioned"
	case WORKFLOW_VERSIONING_MODE_VERSIONING_BEHAVIORS:
		return "VersioningBehaviors"
	default:
		return strconv.Itoa(int(x))
	}

}

func (WorkflowVersioningMode) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_deployment_proto_enumTypes[2].Descriptor()
}

func (WorkflowVersioningMode) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_deployment_proto_enumTypes[2]
}

func (x WorkflowVersioningMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowVersioningMode.Descriptor instead.
func (WorkflowVersioningMode) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_deployment_proto_rawDescGZIP(), []int{2}
}

var File_temporal_api_enums_v1_deployment_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_deployment_proto_rawDesc = string([]byte{
	0x0a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0xc4, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x61, 0x63, 0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x27, 0x0a, 0x23, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e,
	0x54, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x52,
	0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x31, 0x0a, 0x2d, 0x44, 0x45,
	0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42,
	0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44, 0x5f, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x27, 0x0a,
	0x23, 0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x4d, 0x45, 0x4e, 0x54, 0x5f, 0x52, 0x45, 0x41, 0x43,
	0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x55, 0x4e, 0x52, 0x45, 0x41, 0x43, 0x48,
	0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x2a, 0x8b, 0x01, 0x0a, 0x15, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x44, 0x72, 0x61, 0x69, 0x6e, 0x61, 0x67, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x27, 0x0a, 0x23, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x52, 0x41, 0x49,
	0x4e, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x56, 0x45, 0x52,
	0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x52, 0x41, 0x49, 0x4e, 0x41, 0x47, 0x45, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x41, 0x49, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12,
	0x23, 0x0a, 0x1f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x52, 0x41, 0x49, 0x4e,
	0x41, 0x47, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x52, 0x41, 0x49, 0x4e,
	0x45, 0x44, 0x10, 0x02, 0x2a, 0x9f, 0x01, 0x0a, 0x16, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x4d, 0x6f, 0x64, 0x65, 0x12,
	0x28, 0x0a, 0x24, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x45, 0x52, 0x53,
	0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47,
	0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x31, 0x0a, 0x2d, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f,
	0x56, 0x45, 0x52, 0x53, 0x49, 0x4f, 0x4e, 0x49, 0x4e, 0x47, 0x5f, 0x42, 0x45, 0x48, 0x41, 0x56,
	0x49, 0x4f, 0x52, 0x53, 0x10, 0x02, 0x42, 0x87, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65,
	0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x76, 0x31, 0x42, 0x0f, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73,
	0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f,
	0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_temporal_api_enums_v1_deployment_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_deployment_proto_rawDescData []byte
)

func file_temporal_api_enums_v1_deployment_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_deployment_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_deployment_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_deployment_proto_rawDesc), len(file_temporal_api_enums_v1_deployment_proto_rawDesc)))
	})
	return file_temporal_api_enums_v1_deployment_proto_rawDescData
}

var file_temporal_api_enums_v1_deployment_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_temporal_api_enums_v1_deployment_proto_goTypes = []any{
	(DeploymentReachability)(0), // 0: temporal.api.enums.v1.DeploymentReachability
	(VersionDrainageStatus)(0),  // 1: temporal.api.enums.v1.VersionDrainageStatus
	(WorkflowVersioningMode)(0), // 2: temporal.api.enums.v1.WorkflowVersioningMode
}
var file_temporal_api_enums_v1_deployment_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_deployment_proto_init() }
func file_temporal_api_enums_v1_deployment_proto_init() {
	if File_temporal_api_enums_v1_deployment_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_temporal_api_enums_v1_deployment_proto_rawDesc), len(file_temporal_api_enums_v1_deployment_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_deployment_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_deployment_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_deployment_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_deployment_proto = out.File
	file_temporal_api_enums_v1_deployment_proto_goTypes = nil
	file_temporal_api_enums_v1_deployment_proto_depIdxs = nil
}
