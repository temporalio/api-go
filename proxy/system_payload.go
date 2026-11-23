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

// SystemPayloadErrorReason identifies why a marked system payload is not a valid system envelope.
type SystemPayloadErrorReason int

const (
	// SystemPayloadErrorReasonUnspecified is the zero value. VisitPayloads does not return it.
	SystemPayloadErrorReasonUnspecified SystemPayloadErrorReason = iota
	// SystemPayloadErrorUnsupportedEncoding means that the encoding metadata is not binary/protobuf.
	SystemPayloadErrorUnsupportedEncoding
	// SystemPayloadErrorMissingMessageType means that the messageType metadata is empty or not set.
	SystemPayloadErrorMissingMessageType
	// SystemPayloadErrorUnknownMessageType means that the global protobuf registry does not contain the
	// messageType. A process with a newer API version can know the type.
	SystemPayloadErrorUnknownMessageType
	// SystemPayloadErrorMalformedData means that the payload data does not unmarshal as the messageType.
	SystemPayloadErrorMalformedData
)

// SystemPayloadError is returned by VisitPayloads when a payload has the SystemPayloadMetadataKey marker
// but the payload is not a valid system envelope. VisitPayloads does not wrap errors from the visitor for the
// payloads in a valid envelope in this type.
type SystemPayloadError struct {
	Reason SystemPayloadErrorReason
	// MessageType is the messageType metadata of the payload. It is empty when Reason is
	// SystemPayloadErrorMissingMessageType.
	MessageType string
	msg         string
	cause       error
}

func (e *SystemPayloadError) Error() string {
	if e.cause != nil {
		return e.msg + ": " + e.cause.Error()
	}
	return e.msg
}

func (e *SystemPayloadError) Unwrap() error {
	return e.cause
}

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
		return true, &SystemPayloadError{
			Reason:      SystemPayloadErrorUnsupportedEncoding,
			MessageType: string(payload.GetMetadata()["messageType"]),
			msg:         fmt.Sprintf("system payload must be encoded as %s but got %q", binaryProtobufEncoding, encoding),
		}
	}

	messageType := string(payload.GetMetadata()["messageType"])
	if messageType == "" {
		return true, &SystemPayloadError{
			Reason: SystemPayloadErrorMissingMessageType,
			msg:    "system payload is missing the messageType metadata",
		}
	}

	messageTypeDescriptor, err := protoregistry.GlobalTypes.FindMessageByName(protoreflect.FullName(messageType))
	if err != nil {
		return true, &SystemPayloadError{
			Reason:      SystemPayloadErrorUnknownMessageType,
			MessageType: messageType,
			msg:         fmt.Sprintf("system payload references unknown message type %q", messageType),
			cause:       err,
		}
	}
	message := messageTypeDescriptor.New().Interface()
	if err := proto.Unmarshal(payload.GetData(), message); err != nil {
		return true, &SystemPayloadError{
			Reason:      SystemPayloadErrorMalformedData,
			MessageType: messageType,
			msg:         fmt.Sprintf("failed to unmarshal system payload %q", messageType),
			cause:       err,
		}
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
