package temporalnexus_test

import (
	"net/url"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/nexus-rpc/sdk-go/nexus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	commonpb "go.temporal.io/api/common/v1"
	enumspb "go.temporal.io/api/enums/v1"
	nexuspb "go.temporal.io/api/nexus/v1"
	"go.temporal.io/api/temporalnexus"
	"google.golang.org/protobuf/testing/protocmp"
)

func newCallbackLink(executionType enumspb.ExecutionType, businessID string, componentPath []string) *commonpb.Link_Callback {
	return &commonpb.Link_Callback{
		Namespace: "ns",
		Execution: &commonpb.Execution{
			Type:       executionType,
			BusinessId: businessID,
			RunId:      "run-id",
		},
		ComponentPath: componentPath,
		RequestId:     "request-id",
	}
}

func TestConvertLinkWorkflowEventToNexusLink(t *testing.T) {
	type testcase struct {
		name      string
		input     *commonpb.Link_WorkflowEvent
		output    nexus.Link
		outputURL string
	}

	cases := []testcase{
		{
			name: "valid",
			input: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventId:   1,
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "eventID=1&eventType=WorkflowExecutionStarted&referenceType=EventReference",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/history?eventID=1&eventType=WorkflowExecutionStarted&referenceType=EventReference",
		},
		{
			name: "valid with angle bracket",
			input: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id>",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventId:   1,
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id>/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id%3E/run-id/history",
					RawQuery: "eventID=1&eventType=WorkflowExecutionStarted&referenceType=EventReference",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id%3E/run-id/history?eventID=1&eventType=WorkflowExecutionStarted&referenceType=EventReference",
		},
		{
			name: "valid with slash",
			input: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id/",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventId:   1,
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id//run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id%2F/run-id/history",
					RawQuery: "eventID=1&eventType=WorkflowExecutionStarted&referenceType=EventReference",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id%2F/run-id/history?eventID=1&eventType=WorkflowExecutionStarted&referenceType=EventReference",
		},
		{
			name: "valid event id missing",
			input: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "eventType=WorkflowExecutionStarted&referenceType=EventReference",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/history?eventType=WorkflowExecutionStarted&referenceType=EventReference",
		},
		{
			name: "valid request id",
			input: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_RequestIdRef{
					RequestIdRef: &commonpb.Link_WorkflowEvent_RequestIdReference{
						RequestId: "request-id",
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_OPTIONS_UPDATED,
					},
				},
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "eventType=WorkflowExecutionOptionsUpdated&referenceType=RequestIdReference&requestID=request-id",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/history?eventType=WorkflowExecutionOptionsUpdated&referenceType=RequestIdReference&requestID=request-id",
		},
		{
			name: "valid request id empty",
			input: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_RequestIdRef{
					RequestIdRef: &commonpb.Link_WorkflowEvent_RequestIdReference{
						RequestId: "",
					},
				},
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "eventType=Unspecified&referenceType=RequestIdReference&requestID=",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/history?eventType=Unspecified&referenceType=RequestIdReference&requestID=",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output := temporalnexus.ConvertLinkWorkflowEventToNexusLink(tc.input)
			require.Equal(t, tc.output, output)
			require.Equal(t, tc.outputURL, output.URL.String())
		})
	}
}

func TestConvertNexusLinkToLinkWorkflowEvent(t *testing.T) {
	type testcase struct {
		name   string
		input  nexus.Link
		output *commonpb.Link_WorkflowEvent
		errMsg string
	}

	cases := []testcase{
		{
			name: "valid long event type",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventId:   1,
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
		},
		{
			name: "valid short event type",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=WorkflowExecutionStarted",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventId:   1,
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
		},
		{
			name: "valid with angle bracket",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id>/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id%2E/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id>",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventId:   1,
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
		},
		{
			name: "valid with slash",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id//run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id%2F/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id/",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventId:   1,
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
		},
		{
			name: "valid event id missing",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{
					EventRef: &commonpb.Link_WorkflowEvent_EventReference{
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
					},
				},
			},
		},
		{
			name: "invalid nil URL",
			input: nexus.Link{
				URL:  nil,
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "failed to parse link to Link_WorkflowEvent",
		},
		{
			name: "invalid scheme",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "random",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "failed to parse link to Link_WorkflowEvent",
		},
		{
			name: "invalid path missing history",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "failed to parse link to Link_WorkflowEvent",
		},
		{
			name: "invalid path missing namespace",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces//workflows/wf-id/run-id/history",
					RawPath:  "/namespaces//workflows/wf-id/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "failed to parse link to Link_WorkflowEvent",
		},
		{
			name: "invalid event type",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=EventReference&eventID=1&eventType=EVENT_TYPE_INVALID",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "failed to parse link to Link_WorkflowEvent",
		},
		{
			name: "valid request id long event type",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=RequestIdReference&requestID=request-id&eventType=EVENT_TYPE_WORKFLOW_EXECUTION_OPTIONS_UPDATED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_RequestIdRef{
					RequestIdRef: &commonpb.Link_WorkflowEvent_RequestIdReference{
						RequestId: "request-id",
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_OPTIONS_UPDATED,
					},
				},
			},
		},
		{
			name: "valid request id short event type",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=RequestIdReference&requestID=request-id&eventType=WorkflowExecutionOptionsUpdated",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_RequestIdRef{
					RequestIdRef: &commonpb.Link_WorkflowEvent_RequestIdReference{
						RequestId: "request-id",
						EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_OPTIONS_UPDATED,
					},
				},
			},
		},
		{
			name: "valid request id empty",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=RequestIdReference&requestID=&eventType=EVENT_TYPE_UNSPECIFIED",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			output: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_RequestIdRef{
					RequestIdRef: &commonpb.Link_WorkflowEvent_RequestIdReference{
						RequestId: "",
						EventType: enumspb.EVENT_TYPE_UNSPECIFIED,
					},
				},
			},
		},
		{
			name: "invalid request id reference missing event type",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/history",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/history",
					RawQuery: "referenceType=RequestIdReference&requestID=",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "failed to parse link to Link_WorkflowEvent",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := temporalnexus.ConvertNexusLinkToLinkWorkflowEvent(tc.input)
			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tc.output, output, protocmp.Transform()); diff != "" {
					assert.Fail(t, "Proto mismatch (-want +got):\n", diff)
				}
			}
		})
	}
}

func TestConvertLinkNexusOperationToNexusLink(t *testing.T) {
	type testcase struct {
		name      string
		input     *commonpb.Link_NexusOperation
		output    nexus.Link
		outputURL string
	}

	cases := []testcase{
		{
			name: "valid",
			input: &commonpb.Link_NexusOperation{
				Namespace:   "ns",
				OperationId: "op-id",
				RunId:       "run-id",
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/nexus-operations/op-id/run-id/details",
					RawPath: "/namespaces/ns/nexus-operations/op-id/run-id/details",
				},
				Type: "temporal.api.common.v1.Link.NexusOperation",
			},
			outputURL: "temporal:///namespaces/ns/nexus-operations/op-id/run-id/details",
		},
		{
			name: "valid with slash in operation id",
			input: &commonpb.Link_NexusOperation{
				Namespace:   "ns",
				OperationId: "op/id",
				RunId:       "run-id",
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/nexus-operations/op/id/run-id/details",
					RawPath: "/namespaces/ns/nexus-operations/op%2Fid/run-id/details",
				},
				Type: "temporal.api.common.v1.Link.NexusOperation",
			},
			outputURL: "temporal:///namespaces/ns/nexus-operations/op%2Fid/run-id/details",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output := temporalnexus.ConvertLinkNexusOperationToNexusLink(tc.input)
			require.Equal(t, tc.output, output)
			require.Equal(t, tc.outputURL, output.URL.String())
		})
	}
}

func TestConvertNexusLinkToLinkNexusOperation(t *testing.T) {
	type testcase struct {
		name   string
		input  nexus.Link
		output *commonpb.Link_NexusOperation
		errMsg string
	}

	cases := []testcase{
		{
			name: "valid",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/nexus-operations/op-id/run-id/details",
					RawPath: "/namespaces/ns/nexus-operations/op-id/run-id/details",
				},
				Type: "temporal.api.common.v1.Link.NexusOperation",
			},
			output: &commonpb.Link_NexusOperation{
				Namespace:   "ns",
				OperationId: "op-id",
				RunId:       "run-id",
			},
		},
		{
			name: "valid with encoded slash in operation id",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/nexus-operations/op/id/run-id/details",
					RawPath: "/namespaces/ns/nexus-operations/op%2Fid/run-id/details",
				},
				Type: "temporal.api.common.v1.Link.NexusOperation",
			},
			output: &commonpb.Link_NexusOperation{
				Namespace:   "ns",
				OperationId: "op/id",
				RunId:       "run-id",
			},
		},
		{
			name: "wrong link type",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "temporal",
					Path:   "/namespaces/ns/nexus-operations/op-id/run-id/details",
				},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "cannot parse link type",
		},
		{
			name: "invalid scheme",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "random",
					Path:   "/namespaces/ns/nexus-operations/op-id/run-id/details",
				},
				Type: "temporal.api.common.v1.Link.NexusOperation",
			},
			errMsg: "failed to parse link to Link_NexusOperation: invalid scheme",
		},
		{
			name: "malformed path",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "temporal",
					Path:   "/namespaces/ns/nexus-operations/op-id/run-id/",
				},
				Type: "temporal.api.common.v1.Link.NexusOperation",
			},
			errMsg: "failed to parse link to Link_NexusOperation: malformed URL path",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := temporalnexus.ConvertNexusLinkToLinkNexusOperation(tc.input)
			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
			} else {
				require.NoError(t, err)
				if diff := cmp.Diff(tc.output, output, protocmp.Transform()); diff != "" {
					assert.Fail(t, "Proto mismatch (-want +got):\n", diff)
				}
			}
		})
	}
}

func TestConvertLinkActivityToNexusLink(t *testing.T) {
	input := &commonpb.Link_Activity{
		Namespace:  "ns",
		ActivityId: "act-id",
		RunId:      "run-id",
	}

	output := temporalnexus.ConvertLinkActivityToNexusLink(input)
	require.Equal(t, nexus.Link{
		URL: &url.URL{
			Scheme:  "temporal",
			Path:    "/namespaces/ns/activities/act-id/run-id/details",
			RawPath: "/namespaces/ns/activities/act-id/run-id/details",
		},
		Type: "temporal.api.common.v1.Link.Activity",
	}, output)
	require.Equal(t, "temporal:///namespaces/ns/activities/act-id/run-id/details", output.URL.String())
}

func TestConvertNexusLinkToLinkActivity(t *testing.T) {
	type testcase struct {
		name     string
		input    nexus.Link
		expected *commonpb.Link_Activity
		errMsg   string
	}

	cases := []testcase{
		{
			name: "valid",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "temporal",
					Path:   "/namespaces/ns/activities/act-id/run-id/details",
				},
				Type: "temporal.api.common.v1.Link.Activity",
			},
			expected: &commonpb.Link_Activity{
				Namespace:  "ns",
				ActivityId: "act-id",
				RunId:      "run-id",
			},
		},
		{
			name: "round-trip with escaped path",
			input: temporalnexus.ConvertLinkActivityToNexusLink(&commonpb.Link_Activity{
				Namespace:  "ns/with/slash",
				ActivityId: "act id with space",
				RunId:      "run-id",
			}),
			expected: &commonpb.Link_Activity{
				Namespace:  "ns/with/slash",
				ActivityId: "act id with space",
				RunId:      "run-id",
			},
		},
		{
			name: "wrong type",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/activities/act-id"},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "cannot parse link type",
		},
		{
			name: "invalid scheme",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "http", Path: "/namespaces/ns/activities/act-id"},
				Type: "temporal.api.common.v1.Link.Activity",
			},
			errMsg: "invalid scheme",
		},
		{
			name: "malformed path",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/foo/act-id"},
				Type: "temporal.api.common.v1.Link.Activity",
			},
			errMsg: "malformed URL path",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := temporalnexus.ConvertNexusLinkToLinkActivity(tc.input)
			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
				return
			}
			require.NoError(t, err)
			if diff := cmp.Diff(tc.expected, out, protocmp.Transform()); diff != "" {
				assert.Fail(t, "Proto mismatch (-want +got):\n", diff)
			}
		})
	}
}

func TestConvertLinkWorkflowToNexusLink(t *testing.T) {
	type testcase struct {
		name      string
		input     *commonpb.Link_Workflow
		output    nexus.Link
		outputURL string
	}

	cases := []testcase{
		{
			name: "valid",
			input: &commonpb.Link_Workflow{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/workflows/wf-id/run-id",
					RawPath: "/namespaces/ns/workflows/wf-id/run-id",
				},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id",
		},
		{
			name: "valid with reason",
			input: &commonpb.Link_Workflow{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reason:     "rejected update",
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id",
					RawQuery: "reason=rejected+update",
				},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id?reason=rejected+update",
		},
		{
			name: "valid with slash in workflow id",
			input: &commonpb.Link_Workflow{
				Namespace:  "ns",
				WorkflowId: "wf/id",
				RunId:      "run-id",
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/workflows/wf/id/run-id",
					RawPath: "/namespaces/ns/workflows/wf%2Fid/run-id",
				},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf%2Fid/run-id",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output := temporalnexus.ConvertLinkWorkflowToNexusLink(tc.input)
			require.Equal(t, tc.output, output)
			require.Equal(t, tc.outputURL, output.URL.String())
		})
	}
}

func TestConvertNexusLinkToLinkWorkflow(t *testing.T) {
	type testcase struct {
		name     string
		input    nexus.Link
		expected *commonpb.Link_Workflow
		errMsg   string
	}

	cases := []testcase{
		{
			name: "valid",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "temporal",
					Path:   "/namespaces/ns/workflows/wf-id/run-id",
				},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			expected: &commonpb.Link_Workflow{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
			},
		},
		{
			name: "valid with reason",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id",
					RawQuery: "reason=rejected+update",
				},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			expected: &commonpb.Link_Workflow{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reason:     "rejected update",
			},
		},
		{
			name: "round-trip with escaped path and reason",
			input: temporalnexus.ConvertLinkWorkflowToNexusLink(&commonpb.Link_Workflow{
				Namespace:  "ns/with/slash",
				WorkflowId: "wf id with space",
				RunId:      "run-id",
				Reason:     "reason with = and &",
			}),
			expected: &commonpb.Link_Workflow{
				Namespace:  "ns/with/slash",
				WorkflowId: "wf id with space",
				RunId:      "run-id",
				Reason:     "reason with = and &",
			},
		},
		{
			name: "wrong type",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/workflows/wf-id/run-id"},
				Type: "temporal.api.common.v1.Link.WorkflowEvent",
			},
			errMsg: "cannot parse link type",
		},
		{
			name: "nil URL",
			input: nexus.Link{
				URL:  nil,
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			errMsg: "empty URL",
		},
		{
			name: "invalid scheme",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "http", Path: "/namespaces/ns/workflows/wf-id/run-id"},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			errMsg: "invalid scheme",
		},
		{
			name: "malformed path with history suffix",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/workflows/wf-id/run-id/history"},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			errMsg: "malformed URL path",
		},
		{
			name: "malformed path missing run id",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/workflows/wf-id"},
				Type: "temporal.api.common.v1.Link.Workflow",
			},
			errMsg: "malformed URL path",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := temporalnexus.ConvertNexusLinkToLinkWorkflow(tc.input)
			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
				return
			}
			require.NoError(t, err)
			if diff := cmp.Diff(tc.expected, out, protocmp.Transform()); diff != "" {
				assert.Fail(t, "Proto mismatch (-want +got):\n", diff)
			}
		})
	}
}

func TestConvertLinkCallbackToNexusLink(t *testing.T) {
	type testcase struct {
		name      string
		input     *commonpb.Link_Callback
		output    nexus.Link
		outputURL string
		errMsg    string
	}

	cases := []testcase{
		{
			name:  "valid workflow",
			input: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", nil),
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawPath: "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
		},
		{
			name:  "valid activity",
			input: newCallbackLink(enumspb.EXECUTION_TYPE_ACTIVITY, "act-id", nil),
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/activities/act-id/run-id/callbacks/request-id",
					RawPath: "/namespaces/ns/activities/act-id/run-id/callbacks/request-id",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns/activities/act-id/run-id/callbacks/request-id",
		},
		{
			name:  "valid nexus operation",
			input: newCallbackLink(enumspb.EXECUTION_TYPE_NEXUS_OPERATION, "op-id", nil),
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/nexus-operations/op-id/run-id/callbacks/request-id",
					RawPath: "/namespaces/ns/nexus-operations/op-id/run-id/callbacks/request-id",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns/nexus-operations/op-id/run-id/callbacks/request-id",
		},
		{
			// An empty component path adds no query parameters, the same as a nil one.
			name:  "valid with empty component path",
			input: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{}),
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawPath: "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
		},
		{
			name:  "valid with escaped component path",
			input: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{"A:B", "C/D", "E_F"}),
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawQuery: "componentPath=A%3AB&componentPath=C%2FD&componentPath=E_F",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/callbacks/request-id?componentPath=A%3AB&componentPath=C%2FD&componentPath=E_F",
		},
		{
			// Each segment is its own query parameter, so a delimiter inside a segment cannot be
			// mistaken for a separator. Segments end in a user-supplied ID, so a comma is reachable.
			name:  "valid with delimiters in component path segment",
			input: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{"Updates", "id,with,commas"}),
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawQuery: "componentPath=Updates&componentPath=id%2Cwith%2Ccommas",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/callbacks/request-id?componentPath=Updates&componentPath=id%2Cwith%2Ccommas",
		},
		{
			// Only a path whose order differs from its sorted order proves the order is preserved.
			name:  "valid with unsorted component path",
			input: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{"zulu", "alpha"}),
			output: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawPath:  "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawQuery: "componentPath=zulu&componentPath=alpha",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns/workflows/wf-id/run-id/callbacks/request-id?componentPath=zulu&componentPath=alpha",
		},
		{
			// IDs are user-supplied and may contain characters that would otherwise change the
			// shape of the path.
			name: "valid with escaped ids",
			input: &commonpb.Link_Callback{
				Namespace: "ns/with-slash",
				Execution: &commonpb.Execution{
					Type:       enumspb.EXECUTION_TYPE_WORKFLOW,
					BusinessId: "wf-id/callbacks/fake",
					RunId:      "run-id",
				},
				RequestId: "request id?x=1",
			},
			output: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/with-slash/workflows/wf-id/callbacks/fake/run-id/callbacks/request id?x=1",
					RawPath: "/namespaces/ns%2Fwith-slash/workflows/wf-id%2Fcallbacks%2Ffake/run-id/callbacks/request%20id%3Fx=1",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			outputURL: "temporal:///namespaces/ns%2Fwith-slash/workflows/wf-id%2Fcallbacks%2Ffake/run-id/callbacks/request%20id%3Fx=1",
		},
		{
			name:   "unsupported execution type",
			input:  newCallbackLink(enumspb.EXECUTION_TYPE_UNSPECIFIED, "wf-id", nil),
			errMsg: "unsupported execution type",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			output, err := temporalnexus.ConvertLinkCallbackToNexusLink(tc.input)
			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tc.output, output)
			require.Equal(t, tc.outputURL, output.URL.String())
		})
	}
}

func TestConvertNexusLinkToLinkCallback(t *testing.T) {
	type testcase struct {
		name     string
		input    nexus.Link
		expected *commonpb.Link_Callback
		errMsg   string
	}

	cases := []testcase{
		{
			name: "valid workflow",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "temporal",
					Path:   "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			expected: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", nil),
		},
		{
			name: "valid activity",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "temporal",
					Path:   "/namespaces/ns/activities/act-id/run-id/callbacks/request-id",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			expected: newCallbackLink(enumspb.EXECUTION_TYPE_ACTIVITY, "act-id", nil),
		},
		{
			name: "valid nexus operation",
			input: nexus.Link{
				URL: &url.URL{
					Scheme: "temporal",
					Path:   "/namespaces/ns/nexus-operations/op-id/run-id/callbacks/request-id",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			expected: newCallbackLink(enumspb.EXECUTION_TYPE_NEXUS_OPERATION, "op-id", nil),
		},
		{
			name: "valid with escaped component path",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawQuery: "componentPath=A%3AB&componentPath=C%2FD&componentPath=E_F",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			expected: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{"A:B", "C/D", "E_F"}),
		},
		{
			name: "valid with delimiters in component path segment",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawQuery: "componentPath=Updates&componentPath=id%2Cwith%2Ccommas",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			expected: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{"Updates", "id,with,commas"}),
		},
		{
			// Only a path whose order differs from its sorted order proves the order is preserved.
			name: "valid with unsorted component path",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:   "temporal",
					Path:     "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id",
					RawQuery: "componentPath=zulu&componentPath=alpha",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			expected: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{"zulu", "alpha"}),
		},
		{
			name: "valid with escaped ids",
			input: nexus.Link{
				URL: &url.URL{
					Scheme:  "temporal",
					Path:    "/namespaces/ns/with-slash/workflows/wf-id/callbacks/fake/run-id/callbacks/request id?x=1",
					RawPath: "/namespaces/ns%2Fwith-slash/workflows/wf-id%2Fcallbacks%2Ffake/run-id/callbacks/request%20id%3Fx=1",
				},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			expected: &commonpb.Link_Callback{
				Namespace: "ns/with-slash",
				Execution: &commonpb.Execution{
					Type:       enumspb.EXECUTION_TYPE_WORKFLOW,
					BusinessId: "wf-id/callbacks/fake",
					RunId:      "run-id",
				},
				RequestId: "request id?x=1",
			},
		},
		{
			name: "wrong type",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id"},
				Type: "temporal.api.common.v1.Link.Activity",
			},
			errMsg: "cannot parse link type",
		},
		{
			name: "nil URL",
			input: nexus.Link{
				URL:  nil,
				Type: "temporal.api.common.v1.Link.Callback",
			},
			errMsg: "empty URL",
		},
		{
			name: "invalid scheme",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "https", Path: "/namespaces/ns/workflows/wf-id/run-id/callbacks/request-id"},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			errMsg: "invalid scheme",
		},
		{
			name: "malformed path (missing request id)",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/workflows/wf-id/run-id/callbacks"},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			errMsg: "malformed URL path",
		},
		{
			name: "unsupported execution type",
			input: nexus.Link{
				URL:  &url.URL{Scheme: "temporal", Path: "/namespaces/ns/schedules/sched-id/run-id/callbacks/request-id"},
				Type: "temporal.api.common.v1.Link.Callback",
			},
			errMsg: `unsupported execution type: "schedules"`,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			out, err := temporalnexus.ConvertNexusLinkToLinkCallback(tc.input)
			if tc.errMsg != "" {
				require.ErrorContains(t, err, tc.errMsg)
				return
			}
			require.NoError(t, err)
			if diff := cmp.Diff(tc.expected, out, protocmp.Transform()); diff != "" {
				assert.Fail(t, "Proto mismatch (-want +got):\n", diff)
			}
		})
	}
}

func TestCommonLinkToNexusLinkRoundTrip(t *testing.T) {
	cases := []struct {
		name string
		link *commonpb.Link
	}{
		{
			name: "workflow event",
			link: &commonpb.Link{Variant: &commonpb.Link_WorkflowEvent_{WorkflowEvent: &commonpb.Link_WorkflowEvent{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reference: &commonpb.Link_WorkflowEvent_EventRef{EventRef: &commonpb.Link_WorkflowEvent_EventReference{
					EventId:   1,
					EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
				}},
			}}},
		},
		{
			name: "nexus operation",
			link: &commonpb.Link{Variant: &commonpb.Link_NexusOperation_{NexusOperation: &commonpb.Link_NexusOperation{
				Namespace:   "ns",
				OperationId: "op-id",
				RunId:       "run-id",
			}}},
		},
		{
			name: "activity",
			link: &commonpb.Link{Variant: &commonpb.Link_Activity_{Activity: &commonpb.Link_Activity{
				Namespace:  "ns",
				ActivityId: "act-id",
				RunId:      "run-id",
			}}},
		},
		{
			name: "workflow",
			link: &commonpb.Link{Variant: &commonpb.Link_Workflow_{Workflow: &commonpb.Link_Workflow{
				Namespace:  "ns",
				WorkflowId: "wf-id",
				RunId:      "run-id",
				Reason:     "rejected update",
			}}},
		},
		{
			name: "callback",
			link: &commonpb.Link{Variant: &commonpb.Link_Callback_{Callback: newCallbackLink(enumspb.EXECUTION_TYPE_WORKFLOW, "wf-id", []string{"Updates", "update-id"})}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			nexusLink, ok := temporalnexus.CommonLinkToNexusLink(tc.link)
			require.True(t, ok)

			commonLink, ok := temporalnexus.NexusLinkToCommonLink(nexusLink)
			require.True(t, ok)

			if diff := cmp.Diff(tc.link, commonLink, protocmp.Transform()); diff != "" {
				assert.Fail(t, "Proto mismatch (-want +got):\n"+diff)
			}
		})
	}
}

func TestCommonLinkToNexusLinkUnsupported(t *testing.T) {
	// nil variant is not convertible.
	_, ok := temporalnexus.CommonLinkToNexusLink(&commonpb.Link{})
	require.False(t, ok)

	// A nil inner message for a known variant is not convertible.
	_, ok = temporalnexus.CommonLinkToNexusLink(&commonpb.Link{
		Variant: &commonpb.Link_Activity_{},
	})
	require.False(t, ok)
}

func TestNexusLinkToCommonLinkUnsupported(t *testing.T) {
	// An unknown link type is not convertible.
	_, ok := temporalnexus.NexusLinkToCommonLink(&nexuspb.Link{
		Url:  "temporal:///namespaces/ns/activities/act-id/run-id/details",
		Type: "temporal.api.common.v1.Link.SomethingElse",
	})
	require.False(t, ok)

	// A malformed URL for a known type is not convertible.
	_, ok = temporalnexus.NexusLinkToCommonLink(&nexuspb.Link{
		Url:  "temporal:///namespaces/ns/bogus/act-id/run-id/details",
		Type: "temporal.api.common.v1.Link.Activity",
	})
	require.False(t, ok)
}

// Every link type builds its path through the same encoder, so escaping is checked for all of
// them: a space in a path segment is %20, never the form-encoded +, and a literal + survives.
func TestLinkIDEscapingRoundTrip(t *testing.T) {
	const id = "id with space+plus/slash%percent"
	cases := []struct {
		name string
		link *commonpb.Link
	}{
		{
			name: "workflow event",
			link: &commonpb.Link{Variant: &commonpb.Link_WorkflowEvent_{WorkflowEvent: &commonpb.Link_WorkflowEvent{
				Namespace:  id,
				WorkflowId: id,
				RunId:      id,
				Reference: &commonpb.Link_WorkflowEvent_EventRef{EventRef: &commonpb.Link_WorkflowEvent_EventReference{
					EventType: enumspb.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED,
				}},
			}}},
		},
		{
			name: "workflow",
			link: &commonpb.Link{Variant: &commonpb.Link_Workflow_{Workflow: &commonpb.Link_Workflow{
				Namespace:  id,
				WorkflowId: id,
				RunId:      id,
			}}},
		},
		{
			name: "nexus operation",
			link: &commonpb.Link{Variant: &commonpb.Link_NexusOperation_{NexusOperation: &commonpb.Link_NexusOperation{
				Namespace:   id,
				OperationId: id,
				RunId:       id,
			}}},
		},
		{
			name: "activity",
			link: &commonpb.Link{Variant: &commonpb.Link_Activity_{Activity: &commonpb.Link_Activity{
				Namespace:  id,
				ActivityId: id,
				RunId:      id,
			}}},
		},
		{
			name: "callback",
			link: &commonpb.Link{Variant: &commonpb.Link_Callback_{Callback: &commonpb.Link_Callback{
				Namespace: id,
				Execution: &commonpb.Execution{
					Type:       enumspb.EXECUTION_TYPE_ACTIVITY,
					BusinessId: id,
					RunId:      id,
				},
				ComponentPath: []string{id},
				RequestId:     id,
			}}},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			nexusLink, ok := temporalnexus.CommonLinkToNexusLink(tc.link)
			require.True(t, ok)
			u, err := url.Parse(nexusLink.GetUrl())
			require.NoError(t, err)
			require.Contains(t, u.EscapedPath(), "/id%20with%20space+plus%2Fslash%25percent/")

			commonLink, ok := temporalnexus.NexusLinkToCommonLink(nexusLink)
			require.True(t, ok)
			if diff := cmp.Diff(tc.link, commonLink, protocmp.Transform()); diff != "" {
				assert.Fail(t, "Proto mismatch (-want +got):\n"+diff)
			}
		})
	}
}
