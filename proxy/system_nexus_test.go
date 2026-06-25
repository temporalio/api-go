package proxy

import (
	"context"
	"sort"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
	command "go.temporal.io/api/command/v1"
	common "go.temporal.io/api/common/v1"
	sdk "go.temporal.io/api/sdk/v1"
	workflowservice "go.temporal.io/api/workflowservice/v1"
	"google.golang.org/protobuf/proto"
)

const (
	systemNexusService   = "temporal.api.workflowservice.v1.WorkflowService"
	systemNexusSignalOp  = "SignalWithStartWorkflowExecution"
	protoBinaryEncoding  = "binary/protobuf"
	jsonProtoEncoding    = "json/protobuf"
	collectingVisitorTag = "visited-"
	signalWithStartType  = "temporal.api.workflowservice.v1.SignalWithStartWorkflowExecutionRequest"
)

// signalWithStartEnvelope builds a SignalWithStartWorkflowExecutionRequest with
// a payload in every payload-bearing field, marshals it proto-binary, and wraps
// it in a *common.Payload with the given encoding and messageType metadata.
func signalWithStartEnvelope(t *testing.T, encoding string) *common.Payload {
	t.Helper()
	return signalWithStartEnvelopeWithType(t, encoding, signalWithStartType)
}

// signalWithStartEnvelopeWithType is like signalWithStartEnvelope but allows
// overriding the messageType metadata (e.g. to an unknown type).
func signalWithStartEnvelopeWithType(t *testing.T, encoding, messageType string) *common.Payload {
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
	metadata := map[string][]byte{"encoding": []byte(encoding)}
	if messageType != "" {
		metadata["messageType"] = []byte(messageType)
	}
	return &common.Payload{
		Metadata: metadata,
		Data:     data,
	}
}

func scheduleSystemNexusCommand(input *common.Payload) *command.ScheduleNexusOperationCommandAttributes {
	return &command.ScheduleNexusOperationCommandAttributes{
		Endpoint:  "__temporal_system",
		Service:   systemNexusService,
		Operation: systemNexusSignalOp,
		Input:     input,
	}
}

// collectVisitor records the data of every payload it sees and rewrites each to
// "visited-<data>" so callers can confirm write-back.
func collectVisitor(seen *[]string, mu *sync.Mutex) func(*VisitPayloadsContext, []*common.Payload) ([]*common.Payload, error) {
	return func(_ *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
		out := make([]*common.Payload, len(p))
		for i, pl := range p {
			mu.Lock()
			*seen = append(*seen, string(pl.Data))
			mu.Unlock()
			out[i] = &common.Payload{Data: append([]byte(collectingVisitorTag), pl.Data...)}
		}
		return out, nil
	}
}

func decodeEnvelope(t *testing.T, input *common.Payload) *workflowservice.SignalWithStartWorkflowExecutionRequest {
	t.Helper()
	require.Equal(t, []byte(protoBinaryEncoding), input.GetMetadata()["encoding"])
	var req workflowservice.SignalWithStartWorkflowExecutionRequest
	require.NoError(t, proto.Unmarshal(input.GetData(), &req))
	return &req
}

func TestSystemNexusEnvelopeVisitsInnerPayloads(t *testing.T) {
	attrs := scheduleSystemNexusCommand(signalWithStartEnvelope(t, protoBinaryEncoding))

	var seen []string
	var mu sync.Mutex
	err := VisitPayloads(context.Background(), &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: attrs,
		},
	}, VisitPayloadsOptions{
		Visitor: collectVisitor(&seen, &mu),
	})
	require.NoError(t, err)

	sort.Strings(seen)
	require.Equal(t, []string{
		"details-value",
		"header-value",
		"memo-value",
		"sa-value",
		"signal-input",
		"summary-value",
		"workflow-input",
	}, seen)

	// The envelope itself must remain a proto-binary payload and round-trip,
	// with the inner payloads rewritten by the visitor.
	req := decodeEnvelope(t, attrs.Input)
	require.Equal(t, []byte("visited-workflow-input"), req.Input.Payloads[0].Data)
	require.Equal(t, []byte("visited-signal-input"), req.SignalInput.Payloads[0].Data)
	require.Equal(t, []byte("visited-memo-value"), req.Memo.Fields["memo-key"].Data)
	require.Equal(t, []byte("visited-header-value"), req.Header.Fields["header-key"].Data)
	require.Equal(t, []byte("visited-sa-value"), req.SearchAttributes.IndexedFields["sa-key"].Data)
	require.Equal(t, []byte("visited-summary-value"), req.UserMetadata.Summary.Data)
	require.Equal(t, []byte("visited-details-value"), req.UserMetadata.Details.Data)
}

func TestSystemNexusEnvelopeVisitsInnerPayloadsConcurrent(t *testing.T) {
	attrs := scheduleSystemNexusCommand(signalWithStartEnvelope(t, protoBinaryEncoding))

	var seen []string
	var mu sync.Mutex
	err := VisitPayloads(context.Background(), &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: attrs,
		},
	}, VisitPayloadsOptions{
		ConcurrencyLimit: 4,
		Visitor:          collectVisitor(&seen, &mu),
	})
	require.NoError(t, err)

	require.Len(t, seen, 7)
	req := decodeEnvelope(t, attrs.Input)
	require.Equal(t, []byte("visited-workflow-input"), req.Input.Payloads[0].Data)
	require.Equal(t, []byte("visited-details-value"), req.UserMetadata.Details.Data)
}

func TestSystemNexusEnvelopeRejectsNonProtoBinaryEncoding(t *testing.T) {
	attrs := scheduleSystemNexusCommand(signalWithStartEnvelope(t, jsonProtoEncoding))

	err := VisitPayloads(context.Background(), &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: attrs,
		},
	}, VisitPayloadsOptions{
		Visitor: func(_ *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			return p, nil
		},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "binary/protobuf")
}

func TestSystemNexusEnvelopeRejectsUnknownMessageType(t *testing.T) {
	attrs := scheduleSystemNexusCommand(
		signalWithStartEnvelopeWithType(t, protoBinaryEncoding, "temporal.api.workflowservice.v1.NoSuchMessage"),
	)

	err := VisitPayloads(context.Background(), &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: attrs,
		},
	}, VisitPayloadsOptions{
		Visitor: func(_ *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			return p, nil
		},
	})
	require.Error(t, err)
	require.ErrorContains(t, err, "unknown message type")
}

func TestSystemNexusEnvelopeRejectsMissingMessageType(t *testing.T) {
	attrs := scheduleSystemNexusCommand(
		signalWithStartEnvelopeWithType(t, protoBinaryEncoding, ""),
	)

	err := VisitPayloads(context.Background(), &command.Command{
		Attributes: &command.Command_ScheduleNexusOperationCommandAttributes{
			ScheduleNexusOperationCommandAttributes: attrs,
		},
	}, VisitPayloadsOptions{
		Visitor: func(_ *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			return p, nil
		},
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
