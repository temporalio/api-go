package proxy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	command "go.temporal.io/api/command/v1"
	common "go.temporal.io/api/common/v1"
	history "go.temporal.io/api/history/v1"
	sdk "go.temporal.io/api/sdk/v1"
	workflowservice "go.temporal.io/api/workflowservice/v1"
	"google.golang.org/protobuf/proto"
)

const signalWithStartType = "temporal.api.workflowservice.v1.SignalWithStartWorkflowExecutionRequest"

var signalWithStartRequest = &workflowservice.SignalWithStartWorkflowExecutionRequest{
	Namespace:  "default",
	WorkflowId: "wf-id",
	SignalName: "my-signal",
	Input: &common.Payloads{Payloads: []*common.Payload{
		{Data: []byte("workflow-input")},
	}},
	SignalInput: &common.Payloads{Payloads: []*common.Payload{
		{Data: []byte("signal-input")},
	}},
	Memo: &common.Memo{Fields: map[string]*common.Payload{
		"memo-key": {Data: []byte("memo-value")},
	}},
	Header: &common.Header{Fields: map[string]*common.Payload{
		"header-key": {Data: []byte("header-value")},
	}},
	SearchAttributes: &common.SearchAttributes{IndexedFields: map[string]*common.Payload{
		"sa-key": {Data: []byte("sa-value")},
	}},
	UserMetadata: &sdk.UserMetadata{
		Summary: &common.Payload{Data: []byte("summary-value")},
		Details: &common.Payload{Data: []byte("details-value")},
	},
}

var visitedSignalWithStartRequest = &workflowservice.SignalWithStartWorkflowExecutionRequest{
	Namespace:  signalWithStartRequest.Namespace,
	WorkflowId: signalWithStartRequest.WorkflowId,
	SignalName: signalWithStartRequest.SignalName,
	Input: &common.Payloads{Payloads: []*common.Payload{
		{Data: []byte("visited-workflow-input")},
	}},
	SignalInput: &common.Payloads{Payloads: []*common.Payload{
		{Data: []byte("visited-signal-input")},
	}},
	Memo: &common.Memo{Fields: map[string]*common.Payload{
		"memo-key": {Data: []byte("visited-memo-value")},
	}},
	Header: &common.Header{Fields: map[string]*common.Payload{
		"header-key": {Data: []byte("visited-header-value")},
	}},
	SearchAttributes: &common.SearchAttributes{IndexedFields: map[string]*common.Payload{
		"sa-key": {Data: []byte("visited-sa-value")},
	}},
	UserMetadata: &sdk.UserMetadata{
		Summary: &common.Payload{Data: []byte("visited-summary-value")},
		Details: &common.Payload{Data: []byte("visited-details-value")},
	},
}

// buildSystemPayloadCommand builds a command whose Input is
// a SignalWithStartWorkflowExecutionRequest.
// The encoding and messageType parameters represent the "encoding" and
// "messageType" metadata of the Payload.
func buildSystemPayloadCommand(t *testing.T, encoding, messageType string) *command.Command {
	t.Helper()
	data, err := proto.Marshal(signalWithStartRequest)
	require.NoError(t, err)
	input := &common.Payload{
		Data: data,
		Metadata: map[string][]byte{
			SystemPayloadMetadataKey: []byte(systemPayloadMarkerValue),
			"encoding":               []byte(encoding),
			"messageType":            []byte(messageType),
		},
	}
	attrs := &command.ScheduleNexusOperationCommandAttributes{
		Endpoint:  "example-endpoint",
		Service:   "temporal.api.workflowservice.v1.WorkflowService",
		Operation: "SignalWithStartWorkflowExecution",
		Input:     input,
	}
	return &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: attrs,
		},
	}
}

// collectVisitor rewrites each payload to "visited-<data>" so callers can
// confirm write-back.
func collectVisitor(_ *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
	out := make([]*common.Payload, len(p))
	for i, pl := range p {
		out[i] = &common.Payload{Data: append([]byte("visited-"), pl.Data...)}
	}
	return out, nil
}

func trivialVisitor(_ *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
	return p, nil
}

func decodeEnvelope(t *testing.T, input *common.Payload) *workflowservice.SignalWithStartWorkflowExecutionRequest {
	t.Helper()
	require.Equal(t, []byte("binary/protobuf"), input.GetMetadata()["encoding"])
	var req workflowservice.SignalWithStartWorkflowExecutionRequest
	require.NoError(t, proto.Unmarshal(input.GetData(), &req))
	return &req
}

//////////////////////// TESTS ////////////////////////

func TestSystemPayloadVisitsOnlyInnerPayloads(t *testing.T) {
	cmd := buildSystemPayloadCommand(t, "binary/protobuf", signalWithStartType)

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: collectVisitor,
	})
	require.NoError(t, err)

	req := decodeEnvelope(t, cmd.GetScheduleNexusOperationCommandAttributes().Input)
	require.True(t, proto.Equal(visitedSignalWithStartRequest, req))
}

func TestSystemPayloadVisitsInnerPayloadsConcurrent(t *testing.T) {
	cmd := buildSystemPayloadCommand(t, "binary/protobuf", signalWithStartType)

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		ConcurrencyLimit: 4,
		Visitor:          collectVisitor,
	})
	require.NoError(t, err)

	req := decodeEnvelope(t, cmd.GetScheduleNexusOperationCommandAttributes().Input)
	require.True(t, proto.Equal(visitedSignalWithStartRequest, req))
}

func TestSystemPayloadRejectsNonProtoBinaryEncoding(t *testing.T) {
	cmd := buildSystemPayloadCommand(t, "json/protobuf", signalWithStartType)

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: trivialVisitor,
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "binary/protobuf")
}

func TestSystemPayloadRejectsUnknownMessageType(t *testing.T) {
	cmd := buildSystemPayloadCommand(t, "binary/protobuf", "this isn't a valid message type")

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: trivialVisitor,
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "unknown message type")
}

func TestSystemPayloadRejectsMissingMessageType(t *testing.T) {
	cmd := buildSystemPayloadCommand(t, "binary/protobuf", "")

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: trivialVisitor,
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "missing")
}

func TestUnmarkedPayloadRetainsExistingBehavior(t *testing.T) {
	cmd := &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: &command.ScheduleNexusOperationCommandAttributes{
				Endpoint:  "my-endpoint",
				Service:   "my-service",
				Operation: "DoThing",
				Input:     &common.Payload{Data: []byte("user-payload")},
			},
		},
	}

	var seen []string
	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			for _, pl := range p {
				seen = append(seen, string(pl.Data))
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Equal(t, []string{"user-payload"}, seen)
}

func TestSystemPayloadInNexusOperationCompletedEvent(t *testing.T) {
	data, err := proto.Marshal(signalWithStartRequest)
	require.NoError(t, err)
	attrs := &history.NexusOperationCompletedEventAttributes{Result: &common.Payload{
		Data: data,
		Metadata: map[string][]byte{
			SystemPayloadMetadataKey: []byte(systemPayloadMarkerValue),
			"encoding":               []byte(binaryProtobufEncoding),
			"messageType":            []byte(signalWithStartType),
		},
	}}

	err = VisitPayloads(context.Background(), attrs, VisitPayloadsOptions{Visitor: collectVisitor})
	require.NoError(t, err)
	require.True(t, proto.Equal(visitedSignalWithStartRequest, decodeEnvelope(t, attrs.Result)))
}

func TestUnmarkedProtobufEnvelopeIsVisitedDirectly(t *testing.T) {
	data, err := proto.Marshal(signalWithStartRequest)
	require.NoError(t, err)
	outer := &common.Payload{Data: data, Metadata: map[string][]byte{
		"encoding":    []byte(binaryProtobufEncoding),
		"messageType": []byte(signalWithStartType),
	}}
	attrs := &history.NexusOperationCompletedEventAttributes{Result: outer}

	var visited [][]byte
	err = VisitPayloads(context.Background(), attrs, VisitPayloadsOptions{Visitor: func(_ *VisitPayloadsContext, payloads []*common.Payload) ([]*common.Payload, error) {
		for _, payload := range payloads {
			visited = append(visited, payload.Data)
		}
		return payloads, nil
	}})
	require.NoError(t, err)
	require.Equal(t, [][]byte{data}, visited)
}
