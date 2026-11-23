package proxy

import (
	"fmt"

	"go.temporal.io/api/common/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
)

// SystemPayloadMetadataKey marks a binary protobuf payload whose data is a
// system envelope containing user payloads that must be visited recursively.
const SystemPayloadMetadataKey = "__temporal_system_payload"

const (
	systemPayloadMarkerValue = "true"
	binaryProtobufEncoding   = "binary/protobuf"
)

// visitSystemPayload unwraps marked system payloads, visits payloads in the
// decoded message, and writes the re-marshaled message back to the envelope.
// It returns true when the payload was handled as a system envelope.
func visitSystemPayload(
	ctx *VisitPayloadsContext,
	options *VisitPayloadsOptions,
	concState *payloadConcurrencyState,
	payload *common.Payload,
) (bool, error) {
	if payload == nil || string(payload.GetMetadata()[SystemPayloadMetadataKey]) != systemPayloadMarkerValue {
		return false, nil
	}

	if encoding := string(payload.GetMetadata()["encoding"]); encoding != binaryProtobufEncoding {
		return true, fmt.Errorf("system payload must be encoded as %s but got %q", binaryProtobufEncoding, encoding)
	}

	messageType := string(payload.GetMetadata()["messageType"])
	if messageType == "" {
		return true, fmt.Errorf("system payload is missing the messageType metadata")
	}

	messageTypeDescriptor, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(messageType))
	if err != nil {
		return true, fmt.Errorf("system payload references unknown message type %q: %w", messageType, err)
	}
	message := messageTypeDescriptor.New().Interface()
	if err := proto.Unmarshal(payload.GetData(), message); err != nil {
		return true, fmt.Errorf("failed to unmarshal system payload: %w", err)
	}

	if err := visitPayloadsAndWait(ctx, options, message, concState, message); err != nil {
		return true, err
	}

	data, err := proto.Marshal(message)
	if err != nil {
		return true, fmt.Errorf("failed to marshal system payload: %w", err)
	}
	payload.Data = data
	return true, nil
}
