// Package v1 re-exports the upstream Nexus annotations so both import paths
// use the same protobuf registration.
package v1

import upstream "github.com/nexus-rpc/nexus-proto-annotations/go/nexusannotations/v1"

type OperationOptions = upstream.OperationOptions

type ServiceOptions = upstream.ServiceOptions

// Extension fields to descriptorpb.ServiceOptions.
var (
	// optional nexusannotations.v1.ServiceOptions service = 8233;
	E_Service = upstream.E_Service
)

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional nexusannotations.v1.OperationOptions operation = 8234;
	E_Operation = upstream.E_Operation
)

var File_nexusannotations_v1_options_proto = upstream.File_nexusannotations_v1_options_proto
