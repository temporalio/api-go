package proxy

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/api/workflowservice/v1"
)

func TestSystemNexusOperations_SignalWithStartRegistered(t *testing.T) {
	key := SystemNexusOpKey{
		Endpoint:  TemporalSystemNexusEndpoint,
		Operation: "SignalWithStartWorkflowExecution",
	}
	types, ok := SystemNexusOperations[key]
	require.True(t, ok, "SignalWithStartWorkflowExecution must be registered")

	req := types.NewRequest()
	_, ok = req.(*workflowservice.SignalWithStartWorkflowExecutionRequest)
	require.True(t, ok, "request type must be SignalWithStartWorkflowExecutionRequest")

	resp := types.NewResponse()
	_, ok = resp.(*workflowservice.SignalWithStartWorkflowExecutionResponse)
	require.True(t, ok, "response type must be SignalWithStartWorkflowExecutionResponse")
}
