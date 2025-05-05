package temporalproto_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"

	enums "go.temporal.io/api/enums/v1"
	historypb "go.temporal.io/api/history/v1"
	"go.temporal.io/api/temporalproto"
)

var oldEnums = `
{
    "events": [
        {
            "eventId": "1",
            "eventTime": "2020-12-10T22:18:41.248154805Z",
            "eventType": "WorkflowExecutionStarted",
            "taskId": "4195749",
            "workflowExecutionStartedEventAttributes": {
                "workflowType": {
                    "name": "CancelTimerAfterActivity"
                },
                "taskQueue": {
                    "name": "tq-1-TestIntegrationSuite/TestCancelTimerAfterActivity",
                    "kind": "Normal"
                },
                "workflowExecutionTimeout": "15s",
                "workflowRunTimeout": "15s",
                "workflowTaskTimeout": "1s",
                "originalExecutionRunId": "b3481d08-84c0-4afa-bcea-0f29a5c62b2f",
                "identity": "535538@monolith-linux@",
                "firstExecutionRunId": "b3481d08-84c0-4afa-bcea-0f29a5c62b2f",
                "attempt": 1,
                "workflowExecutionExpirationTime": "2020-12-10T22:18:56.248Z",
                "firstWorkflowTaskBackoff": "0s",
                "header": {}
            }
        }
    ]
}
`

var newEnums = `
{
    "events": [
        {
            "eventId": "1",
            "eventTime": "2020-12-10T22:18:41.248154805Z",
            "eventType": "EVENT_TYPE_WORKFLOW_EXECUTION_STARTED",
            "taskId": "4195749",
            "workflowExecutionStartedEventAttributes": {
                "workflowType": {
                    "name": "CancelTimerAfterActivity"
                },
                "taskQueue": {
                    "name": "tq-1-TestIntegrationSuite/TestCancelTimerAfterActivity",
                    "kind": "TASK_QUEUE_KIND_NORMAL"
                },
                "workflowExecutionTimeout": "15s",
                "workflowRunTimeout": "15s",
                "workflowTaskTimeout": "1s",
                "originalExecutionRunId": "b3481d08-84c0-4afa-bcea-0f29a5c62b2f",
                "identity": "535538@monolith-linux@",
                "firstExecutionRunId": "b3481d08-84c0-4afa-bcea-0f29a5c62b2f",
                "attempt": 1,
                "workflowExecutionExpirationTime": "2020-12-10T22:18:56.248Z",
                "firstWorkflowTaskBackoff": "0s",
                "header": {}
            }
        }
    ]
}`

var newNestedFailure = `
{
    "events": [
        {
            "eventId": "1",
            "eventTime": "2020-12-10T22:18:41.248154805Z",
            "eventType": "EVENT_TYPE_WORKFLOW_EXECUTION_FAILED",
            "taskId": "4195749",
            "workflowExecutionFailedEventAttributes": {
                "failure": {
                    "message": "Outer failure",
                    "cause": {
                        "message": "Nested failure"
                    }
                }
            }
        }
    ]
}`

func TestUnmarshalJSON(t *testing.T) {
	t.Parallel()
	require := require.New(t)

	u := temporalproto.CustomJSONUnmarshalOptions{
		DiscardUnknown: true,
	}
	var hist historypb.History
	require.NoError(u.Unmarshal([]byte(newEnums), &hist))
	require.Len(hist.Events, 1)

	ev := hist.Events[0]
	require.Equal(ev.EventType, enums.EVENT_TYPE_WORKFLOW_EXECUTION_STARTED)
	require.Equal(ev.GetWorkflowExecutionStartedEventAttributes().TaskQueue.Kind, enums.TASK_QUEUE_KIND_NORMAL)
}

func TestUnmarshalJSON_Compatible(t *testing.T) {
	t.Parallel()
	u := temporalproto.CustomJSONUnmarshalOptions{
		DiscardUnknown: true,
	}
	// Ensure both new and old enums deserialize the same way
	var oldHist, newHist historypb.History
	require.NoError(t, u.Unmarshal([]byte(oldEnums), &oldHist))
	require.NoError(t, u.Unmarshal([]byte(newEnums), &newHist))
	if !proto.Equal(&oldHist, &newHist) {
		t.Errorf("LoadFromJSON() mismatch between old and new enum formats:\n%#v\n%#v", &oldHist, &newHist)
	}
}

func TestUnmarshalJSON_NestedType(t *testing.T) {
	t.Parallel()

	u := temporalproto.CustomJSONUnmarshalOptions{
		DiscardUnknown: true,
	}
	require := require.New(t)
	var newHist historypb.History
	require.NoError(u.Unmarshal([]byte(newNestedFailure), &newHist))
	require.Len(newHist.Events, 1)

	wfFail := newHist.Events[0].GetWorkflowExecutionFailedEventAttributes()
	require.NotNil(wfFail, "Expected workflow execution failure event, found %s", newHist.Events[0].EventType)
	require.NotNil(wfFail.Failure)
	require.Equal(wfFail.Failure.Message, "Outer failure")
	require.NotNil(wfFail.Failure.Cause)
	require.NotNil(wfFail.Failure.Cause.Message, "Inner failure")
}
