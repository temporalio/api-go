package proxy

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/metadata"

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

	headers, err := ExtractTemporalRequestHeaders(context.Background(), ExtractHeadersOptions{
		Request:          req,
		ExistingMetadata: existingMD,
	})
	require.NoError(t, err)

	// Should not add any headers since they already exist
	require.Empty(t, headers, "Expected no headers to be added when they already exist")
}
