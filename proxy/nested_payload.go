package proxy

import (
	"reflect"

	"go.temporal.io/api/workflowservice/v1/workflowservicenexus"
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
// It is built dynamically from the generated workflowservicenexus package so that
// new operations added to the proto definitions are picked up automatically without
// requiring changes to this file or downstream consumers.
var SystemNexusOperations = buildNexusServiceRegistry(
	TemporalSystemNexusEndpoint,
	workflowservicenexus.WorkflowService,
)

// buildNexusServiceRegistry uses reflection to iterate the fields of a generated
// Nexus service struct (e.g. workflowservicenexus.WorkflowService) and builds a
// registry entry for each field that implements the nexus.OperationReference
// interface (i.e. has Name(), InputType(), and OutputType() methods).
func buildNexusServiceRegistry(endpoint string, service any) map[SystemNexusOpKey]SystemNexusOpTypes {
	registry := make(map[SystemNexusOpKey]SystemNexusOpTypes)
	v := reflect.ValueOf(service)
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		nameMethod := field.MethodByName("Name")
		if !nameMethod.IsValid() {
			continue
		}
		inputTypeMethod := field.MethodByName("InputType")
		outputTypeMethod := field.MethodByName("OutputType")
		if !inputTypeMethod.IsValid() || !outputTypeMethod.IsValid() {
			continue
		}

		name := nameMethod.Call(nil)[0].String()
		inputType := inputTypeMethod.Call(nil)[0].Interface().(reflect.Type)
		outputType := outputTypeMethod.Call(nil)[0].Interface().(reflect.Type)

		registry[SystemNexusOpKey{Endpoint: endpoint, Operation: name}] = SystemNexusOpTypes{
			NewRequest:  newProtoFactory(inputType),
			NewResponse: newProtoFactory(outputType),
		}
	}
	return registry
}

// newProtoFactory returns a function that creates a new zero-valued proto.Message
// of the given type. The type should be a struct type (not a pointer); the returned
// function allocates a new instance and returns a pointer to it as proto.Message.
func newProtoFactory(t reflect.Type) func() proto.Message {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return func() proto.Message {
		return reflect.New(t).Interface().(proto.Message)
	}
}
