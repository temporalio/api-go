package proxy

import (
	"fmt"

	"go.temporal.io/api/command/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
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

	if encoding := string(input.GetMetadata()["encoding"]); encoding != "binary/protobuf" {
		return fmt.Errorf(
			"system nexus envelope for operation %q must be encoded as binary/protobuf but got %q",
			attrs.GetOperation(), encoding,
		)
	}

	messageType := string(input.GetMetadata()["messageType"])
	if messageType == "" {
		return fmt.Errorf(
			"system nexus envelope for operation %q is missing the messageType metadata",
			attrs.GetOperation(),
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

	if err := visitPayloadsAndWait(ctx, options, msg, concState, msg); err != nil {
		return err
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
