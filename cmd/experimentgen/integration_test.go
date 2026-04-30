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

	// Check service stubs (template-generated, not protoc-gen-go-grpc)
	svcFile := readFile(t, outDir, "workflowservice/v1/example_service_experimental.go")
	require.Contains(t, svcFile, "//go:build experimental")
	require.Contains(t, svcFile, "ExampleWorkflowServiceClient interface")
	require.Contains(t, svcFile, "Echo(")
	require.Contains(t, svcFile, `"/temporal.api.workflowservice.v1.WorkflowService/Echo"`)
	require.Contains(t, svcFile, "ExampleWorkflowService_ServiceDesc")

	// Check message file
	msgFile := readFile(t, outDir, "workflowservice/v1/example_messages_experimental.pb.go")
	require.Contains(t, msgFile, "//go:build experimental")
	require.Contains(t, msgFile, "type EchoRequest struct")
	require.Contains(t, msgFile, "type EchoResponse struct")

	// Check overlay file (testdata has experimental_field annotations)
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
