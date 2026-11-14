package update_test

import (
	"context"
	"testing"

	common "go.temporal.io/api/common/v1"
	failure "go.temporal.io/api/failure/v1"
	history "go.temporal.io/api/history/v1"
	"go.temporal.io/api/proxy"
	update "go.temporal.io/api/update/v1"
)

func TestDraftPayloadVisitors(t *testing.T) {
	for _, features := range [][]string{nil, {"field-feature"}, {"message-feature"}, {"field-feature", "message-feature"}} {
		p := &common.Payload{ExternalPayloads: []*common.Payload_ExternalPayloadDetails{{SizeBytes: 11}}}
		request := &update.Request{
			Input: &update.Input{Args: &common.Payloads{Payloads: []*common.Payload{p}}},
			DraftPayload: &update.DraftPayloadSubtype{
				Payload: p, Payloads: &common.Payloads{Payloads: []*common.Payload{p}},
				Failure: &failure.Failure{EncodedAttributes: p},
				Cycle:   &update.DraftPayloadCycle{Back: &update.DraftPayloadSubtype{Payload: p}},
				Nested:  &update.DraftPayloadSubtype_Nested{Payload: p},
			},
		}
		request.OmitDraftFields(func(feature string) bool {
			for _, enabled := range features {
				if enabled == feature {
					return true
				}
			}
			return false
		})
		event := &history.HistoryEvent{Attributes: &history.HistoryEvent_WorkflowExecutionUpdateAcceptedEventAttributes{
			WorkflowExecutionUpdateAcceptedEventAttributes: &history.WorkflowExecutionUpdateAcceptedEventAttributes{AcceptedRequest: request},
		}}
		count, size := 0, int64(0)
		if err := proxy.VisitPayloads(context.Background(), event, proxy.VisitPayloadsOptions{
			Visitor: func(_ *proxy.VisitPayloadsContext, payloads []*common.Payload) ([]*common.Payload, error) {
				count += len(payloads)
				for _, payload := range payloads {
					for _, external := range payload.ExternalPayloads {
						size += external.SizeBytes
					}
				}
				return payloads, nil
			},
		}); err != nil {
			t.Fatal(err)
		}
		want := 1
		if len(features) == 2 {
			want = 6
		}
		if count != want || size != int64(want)*11 {
			t.Fatalf("features %v: history payload count/size = %d/%d, want %d/%d", features, count, size, want, want*11)
		}
		failures := 0
		if err := proxy.VisitFailures(context.Background(), event, proxy.VisitFailuresOptions{
			Visitor: func(_ *proxy.VisitFailuresContext, _ *failure.Failure) error {
				failures++
				return nil
			},
		}); err != nil {
			t.Fatal(err)
		}
		wantFailures := 0
		if len(features) == 2 {
			wantFailures = 1
		}
		if failures != wantFailures {
			t.Fatalf("features %v: failures = %d, want %d", features, failures, wantFailures)
		}
	}
}
