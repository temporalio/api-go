package proxy

import (
	"fmt"

	"go.temporal.io/api/command/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// Payload encoding/type metadata used to decode the proto-binary envelope.
const (
	payloadMetadataEncodingKey    = "encoding"
	payloadMetadataMessageTypeKey = "messageType"
	payloadEncodingProtoBinary    = "binary/protobuf"
)

// visitSystemNexusEnvelope decodes the system Nexus envelope in attrs.Input,
// visits the payloads inside the decoded request message, and re-encodes it.
//
// The envelope's proto message type is taken from the payload's "messageType"
// metadata, so no operation registry is required. The envelope must be encoded
// as binary/protobuf. The inner payloads (and only those) are passed to the
// visitor, so external storage and codecs apply to them and not to the envelope
// itself, which is never offloaded or codec-encoded.
func visitSystemNexusEnvelope(
	ctx *VisitPayloadsContext,
	options *VisitPayloadsOptions,
	concState *payloadConcurrencyState,
	attrs *command.ScheduleNexusOperationCommandAttributes,
) error {
	input := attrs.Input

	if encoding := string(input.GetMetadata()[payloadMetadataEncodingKey]); encoding != payloadEncodingProtoBinary {
		return fmt.Errorf(
			"system nexus envelope for operation %q must be encoded as %q, got %q",
			attrs.GetOperation(), payloadEncodingProtoBinary, encoding,
		)
	}

	messageType := string(input.GetMetadata()[payloadMetadataMessageTypeKey])
	if messageType == "" {
		return fmt.Errorf(
			"system nexus envelope for operation %q is missing the %q metadata",
			attrs.GetOperation(), payloadMetadataMessageTypeKey,
		)
	}

	mt, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(messageType))
	if err != nil {
		return fmt.Errorf(
			"system nexus envelope for operation %q references unknown message type %q: %w",
			attrs.GetOperation(), messageType, err,
		)
	}
	msg := mt.New().Interface()

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
