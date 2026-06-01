package proxy

import (
	workflowservice "go.temporal.io/api/workflowservice/v1"
	"google.golang.org/protobuf/proto"
)

const (
	// TemporalSystemNexusEndpoint is the well-known endpoint name for system Nexus operations.
	TemporalSystemNexusEndpoint = "__temporal_system"
)

// SystemNexusOpKey identifies a system Nexus operation by its (endpoint, operation) pair.
type SystemNexusOpKey struct {
	Endpoint  string
	Operation string
}

// SystemNexusOpTypes maps a system Nexus operation to the proto request and response
// types whose bytes are serialized in NexusOperationScheduled.Input and
// NexusOperationCompleted.Result.
type SystemNexusOpTypes struct {
	// NewRequest returns a fresh, zero-valued instance of the request proto.
	NewRequest func() proto.Message
	// NewResponse returns a fresh, zero-valued instance of the response proto.
	NewResponse func() proto.Message
}

// SystemNexusOperations is the canonical registry of known system Nexus operations.
// Consumers should use this rather than maintaining their own mapping.
var SystemNexusOperations = map[SystemNexusOpKey]SystemNexusOpTypes{
	{
		Endpoint:  TemporalSystemNexusEndpoint,
		Operation: "SignalWithStartWorkflowExecution",
	}: {
		NewRequest:  func() proto.Message { return &workflowservice.SignalWithStartWorkflowExecutionRequest{} },
		NewResponse: func() proto.Message { return &workflowservice.SignalWithStartWorkflowExecutionResponse{} },
	},
}
