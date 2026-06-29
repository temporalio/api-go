package proxy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	command "go.temporal.io/api/command/v1"
	common "go.temporal.io/api/common/v1"
	sdk "go.temporal.io/api/sdk/v1"
	workflowservice "go.temporal.io/api/workflowservice/v1"
	"google.golang.org/protobuf/proto"
)

const signalWithStartType = "temporal.api.workflowservice.v1.SignalWithStartWorkflowExecutionRequest"

// buildSystemNexusCommand builds a system Nexus command whose Input is
// a SignalWithStartWorkflowExecutionRequest.
// The encoding and messageType parameters represent the "encoding" and
// "messageType" metadata of the Payload.
func buildSystemNexusCommand(t *testing.T, encoding, messageType string) *command.Command {
	t.Helper()
	req := &workflowservice.SignalWithStartWorkflowExecutionRequest{
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
	data, err := proto.Marshal(req)
	require.NoError(t, err)
	input := &common.Payload{
		Data: data,
		Metadata: map[string][]byte{
			"encoding":    []byte(encoding),
			"messageType": []byte(messageType),
		},
	}
	attrs := &command.ScheduleNexusOperationCommandAttributes{
		Endpoint:  "__temporal_system",
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

func TestSystemNexusEnvelopeVisitsInnerPayloads(t *testing.T) {
	cmd := buildSystemNexusCommand(t, "binary/protobuf", signalWithStartType)

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: collectVisitor,
	})
	require.NoError(t, err)

	// The envelope itself must remain a proto-binary payload and round-trip,
	// with the inner payloads rewritten by the visitor.
	req := decodeEnvelope(t, cmd.GetScheduleNexusOperationCommandAttributes().Input)
	require.Equal(t, []byte("visited-workflow-input"), req.Input.Payloads[0].Data)
	require.Equal(t, []byte("visited-signal-input"), req.SignalInput.Payloads[0].Data)
	require.Equal(t, []byte("visited-memo-value"), req.Memo.Fields["memo-key"].Data)
	require.Equal(t, []byte("visited-header-value"), req.Header.Fields["header-key"].Data)
	require.Equal(t, []byte("visited-sa-value"), req.SearchAttributes.IndexedFields["sa-key"].Data)
	require.Equal(t, []byte("visited-summary-value"), req.UserMetadata.Summary.Data)
	require.Equal(t, []byte("visited-details-value"), req.UserMetadata.Details.Data)
}

func TestSystemNexusEnvelopeVisitsInnerPayloadsConcurrent(t *testing.T) {
	cmd := buildSystemNexusCommand(t, "binary/protobuf", signalWithStartType)

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		ConcurrencyLimit: 4,
		Visitor:          collectVisitor,
	})
	require.NoError(t, err)

	req := decodeEnvelope(t, cmd.GetScheduleNexusOperationCommandAttributes().Input)
	require.Equal(t, []byte("visited-workflow-input"), req.Input.Payloads[0].Data)
	require.Equal(t, []byte("visited-details-value"), req.UserMetadata.Details.Data)
}

func TestSystemNexusEnvelopeRejectsNonProtoBinaryEncoding(t *testing.T) {
	cmd := buildSystemNexusCommand(t, "json/protobuf", signalWithStartType)

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: trivialVisitor,
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "binary/protobuf")
}

func TestSystemNexusEnvelopeRejectsUnknownMessageType(t *testing.T) {
	cmd := buildSystemNexusCommand(t, "binary/protobuf", "temporal.api.workflowservice.v1.NoSuchMessage")

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: trivialVisitor,
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "unknown message type")
}

func TestSystemNexusEnvelopeRejectsMissingMessageType(t *testing.T) {
	cmd := buildSystemNexusCommand(t, "binary/protobuf", "")

	err := VisitPayloads(context.Background(), cmd, VisitPayloadsOptions{
		Visitor: trivialVisitor,
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "missing")
}

// TestNonSystemNexusInputVisitedAsSinglePayload confirms that an ordinary (non
// system) Nexus operation continues to have its Input visited as one opaque
// payload rather than being decoded as an envelope.
func TestNonSystemNexusInputVisitedAsSinglePayload(t *testing.T) {
	attrs := &command.ScheduleNexusOperationCommandAttributes{
		Endpoint:  "my-endpoint",
		Service:   "my.Service",
		Operation: "DoThing",
		Input:     &common.Payload{Data: []byte("user-payload")},
	}

	var seen []string
	var single bool
	err := VisitPayloads(context.Background(), &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: attrs,
		},
	}, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			single = ctx.SinglePayloadRequired
			for _, pl := range p {
				seen = append(seen, string(pl.Data))
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.True(t, single, "ordinary Nexus Input should be visited as a single payload")
	require.Equal(t, []string{"user-payload"}, seen)
	require.Equal(t, []byte("user-payload"), attrs.Input.Data)
}
