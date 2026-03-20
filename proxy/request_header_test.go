package proxy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/deployment/v1"
	"go.temporal.io/api/taskqueue/v1"
	"go.temporal.io/api/workflowservice/v1"
)

// findHeader searches for a header key in the headers slice and returns its value
func findHeader(headers []string, key string) (string, bool) {
	for i := 0; i < len(headers); i += 2 {
		if i+1 < len(headers) && headers[i] == key {
			return headers[i+1], true
		}
	}
	return "", false
}

func TestExtractTemporalRequestHeaders_NamespaceAlwaysIncluded(t *testing.T) {
	req := &workflowservice.StartWorkflowExecutionRequest{
		Namespace:  "test-namespace",
		WorkflowId: "test-workflow",
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	// Namespace should always be included in the headers
	nsVal, found := findHeader(headers, "temporal-namespace")
	require.True(t, found, "Expected temporal-namespace header, but not found")
	require.Equal(t, "test-namespace", nsVal)
}

func TestExtractTemporalRequestHeaders_EmptyWorkflowId(t *testing.T) {
	req := &workflowservice.StartWorkflowExecutionRequest{
		Namespace:  "test-namespace",
		WorkflowId: "",
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	// Should still set namespace
	nsVal, found := findHeader(headers, "temporal-namespace")
	require.True(t, found, "Expected temporal-namespace header even with empty workflow_id")
	require.Equal(t, "test-namespace", nsVal)
}

func TestExtractTemporalRequestHeaders_SkipExistingHeaders(t *testing.T) {
	req := &workflowservice.StartWorkflowExecutionRequest{
		Namespace:  "test-namespace",
		WorkflowId: "test-workflow",
	}

	existingMD := metadata.MD{}
	existingMD.Set("temporal-namespace", "existing-namespace")
	existingMD.Set("temporal-resource-id", "existing-resource-id")

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request:          req,
		ExistingMetadata: existingMD,
	})
	require.NoError(t, err)

	// Should not add any headers since they already exist
	require.Empty(t, headers, "Expected no headers to be added when they already exist")
}

func TestExtractTemporalRequestHeaders_NilRequest(t *testing.T) {
	_, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: nil,
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "request cannot be nil")
}

func TestExtractTemporalRequestHeaders_ResourceIdWithPrefix(t *testing.T) {
	req := &workflowservice.StartWorkflowExecutionRequest{
		Namespace:  "test-namespace",
		WorkflowId: "my-workflow-123",
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	resVal, found := findHeader(headers, "temporal-resource-id")
	require.True(t, found, "Expected temporal-resource-id header")
	require.Equal(t, "workflow:my-workflow-123", resVal)
}

func TestExtractTemporalRequestHeaders_NestedFieldAccess(t *testing.T) {
	req := &workflowservice.DeleteWorkflowExecutionRequest{
		Namespace: "test-namespace",
		WorkflowExecution: &common.WorkflowExecution{
			WorkflowId: "nested-workflow",
		},
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	resVal, found := findHeader(headers, "temporal-resource-id")
	require.True(t, found, "Expected temporal-resource-id header for nested field")
	require.Equal(t, "workflow:nested-workflow", resVal)
}

func TestExtractTemporalRequestHeaders_NestedFieldNilParent(t *testing.T) {
	// WorkflowExecution is nil, so GetWorkflowId() should return ""
	req := &workflowservice.DeleteWorkflowExecutionRequest{
		Namespace:         "test-namespace",
		WorkflowExecution: nil,
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	_, found := findHeader(headers, "temporal-resource-id")
	require.False(t, found, "Expected no temporal-resource-id header when nested field is nil")

	// Namespace should still be set
	nsVal, found := findHeader(headers, "temporal-namespace")
	require.True(t, found)
	require.Equal(t, "test-namespace", nsVal)
}

func TestExtractTemporalRequestHeaders_DeploymentNestedField(t *testing.T) {
	req := &workflowservice.DeleteWorkerDeploymentVersionRequest{
		Namespace: "test-namespace",
		DeploymentVersion: &deployment.WorkerDeploymentVersion{
			DeploymentName: "my-deployment",
		},
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	resVal, found := findHeader(headers, "temporal-resource-id")
	require.True(t, found, "Expected temporal-resource-id header for deployment")
	require.Equal(t, "deployment:my-deployment", resVal)
}

func TestExtractTemporalRequestHeaders_TaskQueueNestedField(t *testing.T) {
	req := &workflowservice.DescribeTaskQueueRequest{
		Namespace: "test-namespace",
		TaskQueue: &taskqueue.TaskQueue{
			Name: "my-task-queue",
		},
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	resVal, found := findHeader(headers, "temporal-resource-id")
	require.True(t, found, "Expected temporal-resource-id header for task queue")
	require.Equal(t, "taskqueue:my-task-queue", resVal)
}

func TestExtractTemporalRequestHeaders_ScheduleResourceId(t *testing.T) {
	req := &workflowservice.CreateScheduleRequest{
		Namespace:  "test-namespace",
		ScheduleId: "my-schedule",
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	resVal, found := findHeader(headers, "temporal-resource-id")
	require.True(t, found, "Expected temporal-resource-id header for schedule")
	require.Equal(t, "schedule:my-schedule", resVal)
}

func TestExtractTemporalRequestHeaders_EmptyNamespace(t *testing.T) {
	req := &workflowservice.RecordActivityTaskHeartbeatRequest{
		TaskToken: []byte("token"),
		Namespace: "",
	}

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request: req,
	})
	require.NoError(t, err)

	_, found := findHeader(headers, "temporal-namespace")
	require.False(t, found, "Expected no temporal-namespace header when namespace is empty")
}
