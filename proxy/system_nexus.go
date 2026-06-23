package proxy

import (
	"fmt"

	"go.temporal.io/api/command/v1"
	"go.temporal.io/api/nexussystem"
	"google.golang.org/protobuf/proto"
)

// systemNexusServiceName is the Nexus service used for system Nexus envelopes
// routed to the workflow service. Together with nexussystem.Endpoint it
// identifies a ScheduleNexusOperationCommandAttributes whose Input is a system
// Nexus envelope.
const systemNexusServiceName = "temporal.api.workflowservice.v1.WorkflowService"

// Payload encoding metadata used to identify the proto-binary envelope.
const (
	payloadMetadataEncodingKey = "encoding"
	payloadEncodingProtoBinary = "binary/protobuf"
)

// isSystemNexusEnvelope reports whether the given schedule-nexus-operation
// command targets the system Nexus endpoint and service, and therefore carries
// a proto-message envelope in its Input rather than an opaque user payload.
func isSystemNexusEnvelope(attrs *command.ScheduleNexusOperationCommandAttributes) bool {
	if attrs == nil {
		return false
	}
	return attrs.GetEndpoint() == nexussystem.Endpoint &&
		attrs.GetService() == systemNexusServiceName
}

// visitScheduleNexusOperationInput visits the payloads referenced by a
// ScheduleNexusOperationCommandAttributes.Input field.
//
// For ordinary Nexus operations the Input is an opaque single payload and is
// visited directly. For system Nexus envelopes (endpoint __temporal_system,
// service WorkflowService) the Input is instead a proto-binary-serialized
// request message whose own fields contain the user payloads. In that case the
// envelope is deserialized, its inner payloads are visited recursively (so
// external storage and codecs apply to the inner payloads, not the envelope),
// and the message is re-serialized back into Input. The envelope itself is
// never offloaded or codec-encoded.
func visitScheduleNexusOperationInput(
	ctx *VisitPayloadsContext,
	options *VisitPayloadsOptions,
	concState *payloadConcurrencyState,
	attrs *command.ScheduleNexusOperationCommandAttributes,
) error {
	if attrs.Input == nil {
		return nil
	}

	if !isSystemNexusEnvelope(attrs) {
		// Ordinary Nexus operation: visit the Input as a single opaque payload.
		return visitPayload(ctx, options, attrs, concState, &attrs.Input)
	}

	return visitSystemNexusEnvelope(ctx, options, concState, attrs)
}

// visitSystemNexusEnvelope decodes the system Nexus envelope in attrs.Input,
// visits the payloads inside the decoded request message, and re-encodes it.
func visitSystemNexusEnvelope(
	ctx *VisitPayloadsContext,
	options *VisitPayloadsOptions,
	concState *payloadConcurrencyState,
	attrs *command.ScheduleNexusOperationCommandAttributes,
) error {
	msg, ok := nexussystem.InputMessage(attrs.GetService(), attrs.GetOperation())
	if !ok {
		return fmt.Errorf(
			"unknown system nexus operation %q for service %q",
			attrs.GetOperation(), attrs.GetService(),
		)
	}

	input := attrs.Input
	if encoding := string(input.GetMetadata()[payloadMetadataEncodingKey]); encoding != payloadEncodingProtoBinary {
		return fmt.Errorf(
			"system nexus envelope for operation %q must be encoded as %q, got %q",
			attrs.GetOperation(), payloadEncodingProtoBinary, encoding,
		)
	}

	if err := proto.Unmarshal(input.GetData(), msg); err != nil {
		return fmt.Errorf(
			"failed to unmarshal system nexus envelope for operation %q: %w",
			attrs.GetOperation(), err,
		)
	}

	// Visit the payloads contained within the decoded request message. We pass
	// the decoded message both as the parent and as the object to traverse so
	// that the generated visitor descends into its payload-bearing fields.
	//
	// In concurrent mode the visitor may spawn goroutines that write into the
	// decoded message's payload fields. We give those goroutines a sub-state
	// that shares the parent semaphore (to respect the global concurrency
	// limit) but has its own WaitGroup, so we can wait for exactly the
	// envelope's inner-payload goroutines before re-serializing -- mirroring how
	// the well-known Any visitor isolates child traversals.
	envState := concState
	if concState != nil {
		envState = &payloadConcurrencyState{sem: concState.sem}
	}

	if err := visitPayloads(ctx, options, msg, envState, msg); err != nil {
		return err
	}

	if envState != nil {
		envState.wg.Wait()
		if errPtr := envState.firstErr.Load(); errPtr != nil {
			return *errPtr
		}
	}

	data, err := proto.Marshal(msg)
	if err != nil {
		return fmt.Errorf(
			"failed to marshal system nexus envelope for operation %q: %w",
			attrs.GetOperation(), err,
		)
	}
	input.Data = data
	return nil
}
