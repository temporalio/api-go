// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package history_test

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	history "go.temporal.io/api/history/v1"
	"google.golang.org/protobuf/testing/protocmp"
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
}
`

var longHistory = `
{
  "events": [
    {
      "eventId": 1,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowExecutionStarted",
      "workflowExecutionStartedEventAttributes": {
        "workflowType": {
          "name": "testReplayWorkflowFromFile"
        },
        "taskQueue": {
          "name": "taskQueue1"
        },
        "workflowRunTimeout": "60s",
        "workflowTaskTimeout": "60s",
        "identity": "temporal-cli@user-C02WC08UHTDG"
      }
    },
    {
      "eventId": 2,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowTaskScheduled",
      "workflowTaskScheduledEventAttributes": {
        "taskQueue": {
          "name": "taskQueue1"
        },
        "startToCloseTimeout": "60s",
        "attempt": 1
      }
    },
    {
      "eventId": 3,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowTaskStarted",
      "workflowTaskStartedEventAttributes": {
        "scheduledEventId": 2,
        "identity": "50114@user-C02WC08UHTDG@taskQueue1",
        "requestId": "b7403b35-b4b1-432f-84ff-01d66d060a87"
      }
    },
    {
      "eventId": 4,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowTaskCompleted",
      "workflowTaskCompletedEventAttributes": {
        "scheduledEventId": 2,
        "startedEventId": 3,
        "identity": "50114@user-C02WC08UHTDG@taskQueue1"
      }
    },
    {
      "eventId": 5,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "ActivityTaskScheduled",
      "activityTaskScheduledEventAttributes": {
        "activityId": "5",
        "activityType": {
          "name": "testActivityMultipleArgs"
        },
        "taskQueue": {
          "name": "taskQueue1"
        },
        "input": null,
        "scheduleToCloseTimeout": "120s",
        "scheduleToStartTimeout": "60s",
        "startToCloseTimeout": "60s",
        "heartbeatTimeout": "20s",
        "workflowTaskCompletedEventId": 4
      }
    },
    {
      "eventId": 6,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "ActivityTaskStarted",
      "version": -24,
      "taskId": 33554446,
      "activityTaskStartedEventAttributes": {
        "scheduledEventId": 5,
        "identity": "50114@user-C02WC08UHTDG@taskQueue1",
        "requestId": "45c4006a-ae7c-4392-baa6-c090857f884b",
        "attempt": 1
      }
    },
    {
      "eventId": 7,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "ActivityTaskCompleted",
      "version": -24,
      "taskId": 33554447,
      "activityTaskCompletedEventAttributes": {
        "result": null,
        "scheduledEventId": 5,
        "startedEventId": 6,
        "identity": "50114@user-C02WC08UHTDG@taskQueue1"
      }
    },
    {
      "eventId": 8,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowTaskScheduled",
      "version": -24,
      "taskId": 33554450,
      "workflowTaskScheduledEventAttributes": {
        "taskQueue": {
          "name": "longer-C02V60N3HTDG:33ab3ada-4636-4386-8575-81dd8dc02e9a"
        },
        "startToCloseTimeout": "10s",
        "attempt": 1
      }
    },
    {
      "eventId": 9,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowTaskStarted",
      "version": -24,
      "taskId": 33554454,
      "workflowTaskStartedEventAttributes": {
        "scheduledEventId": 8,
        "identity": "50114@user-C02WC08UHTDG@taskQueue1",
        "requestId": "cb1fdadf-f46b-4840-9b97-863f4b3b6b11"
      }
    },
    {
      "eventId": 10,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowTaskCompleted",
      "version": -24,
      "taskId": 33554457,
      "workflowTaskCompletedEventAttributes": {
        "scheduledEventId": 8,
        "startedEventId": 9,
        "identity": "50114@user-C02WC08UHTDG@taskQueue1",
        "binaryChecksum": "b2e32759177ccbb3e67ad7694aec233c"
      }
    },
    {
      "eventId": 11,
      "eventTime": "2020-07-30T00:30:02.971655189Z",
      "eventType": "WorkflowExecutionCompleted",
      "version": -24,
      "taskId": 33554458,
      "workflowExecutionCompletedEventAttributes": {
        "workflowTaskCompletedEventId": 10
      }
    }
  ]
}
`

type testCase struct {
	Name  string
	Input []byte
	Want  *history.History
	Error error
}

func TestLoadHistoryFromJSON_Compatible(t *testing.T) {
	t.Parallel()
	// Ensure both new and old enums deserialize the same way
	oldHist, err := history.LoadFromJSON(strings.NewReader(oldEnums))
	if err != nil {
		t.Errorf("Unexpected error loading old history json: %s", err)
	}
	newHist, err := history.LoadFromJSON(strings.NewReader(oldEnums))
	if err != nil {
		t.Errorf("Unexpected error loading new history json: %s", err)
	}
	if diff := cmp.Diff(oldHist, newHist, protocmp.Transform()); diff != "" {
		t.Errorf("LoadFromJSON() mismatch between old and new enum formats (-old +new):\n%v", diff)
	}
}

func TestLoadHistoryFromJSON_LastEventID(t *testing.T) {
	t.Parallel()

	hist, err := history.LoadFromJSON(strings.NewReader(longHistory))
	if err != nil {
		t.Errorf("Unexpected error loading history json: %s", err)
	}

	if len(hist.Events) != 11 {
		t.Errorf("Expected 11 history events, found %d", len(hist.Events))
	}

	hist5, err := history.LoadFromJSON(strings.NewReader(longHistory), int64(5))
	if err != nil {
		t.Errorf("Unexpected error loading history json: %s", err)
	}

	if len(hist5.Events) != 5 {
		t.Errorf("Expected 5 history events, found %d", len(hist5.Events))
	}
}
