package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateExample(t *testing.T) {
	data := buildTestDescriptorSet(t)
	outDir := t.TempDir()
	gen := generator{}
	err := gen.generate(data, "example", outDir)
	require.NoError(t, err)

	// Check grpc file
	grpcFile := readFile(t, outDir, "workflowservice/v1/example_service_experimental_grpc.pb.go")
	require.Contains(t, grpcFile, "//go:build experimental")
	require.Contains(t, grpcFile, "WorkflowServiceClient interface")
	require.Contains(t, grpcFile, "Echo(")

	// Check message file
	msgFile := readFile(t, outDir, "workflowservice/v1/example_messages_experimental.pb.go")
	require.Contains(t, msgFile, "//go:build experimental")
	require.Contains(t, msgFile, "type EchoRequest struct")
	require.Contains(t, msgFile, "type EchoResponse struct")

	// Check overlay file
	overlayFile := readFile(t, outDir, "workflowservice/v1/example_overlay_experimental.go")
	require.Contains(t, overlayFile, "//go:build experimental")
	require.Contains(t, overlayFile, "GetStartWorkflowExecutionRequestOverlay(")

	// Check enum file
	enumFile := readFile(t, outDir, "enums/v1/example_enum_experimental.go")
	require.Contains(t, enumFile, "//go:build experimental")
	require.Contains(t, enumFile, "WORKFLOW_ID_CONFLICT_POLICY_FOO")
}

func TestRunRequiresVariant(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := runMain(nil, bytes.NewReader(nil), &stdout, &stderr)
	require.Equal(t, 2, code)
	require.Contains(t, stderr.String(), "-variant")
	require.Empty(t, stdout.String())
}
