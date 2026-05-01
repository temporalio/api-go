package main

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateExample(t *testing.T) {
	data := buildTestDescriptorSet(t)
	outDir := t.TempDir()
	gen := generator{
		resolveStableVersion: func(string) (moduleVersion, error) {
			return moduleVersion{Tag: "v1.2.3", GoVersion: "1.21"}, nil
		},
		skipModTidy: true,
	}
	err := gen.generate(data, "example", outDir)
	require.NoError(t, err)

	// Check service stubs (template-generated, not protoc-gen-go-grpc)
	svcFile := readFile(t, outDir, "experimental/workflowservice/v1/example_service.go")
	require.NotContains(t, svcFile, "//go:build experimental")
	require.Contains(t, svcFile, "ExampleWorkflowServiceClient interface")
	require.Contains(t, svcFile, "Echo(")
	require.Contains(t, svcFile, `"/temporal.api.workflowservice.v1.WorkflowService/Echo"`)
	require.Contains(t, svcFile, "ExampleWorkflowService_ServiceDesc")

	// Check message file
	msgFile := readFile(t, outDir, "experimental/workflowservice/v1/example_messages.pb.go")
	require.NotContains(t, msgFile, "//go:build experimental")
	require.Contains(t, msgFile, "type EchoRequest struct")
	require.Contains(t, msgFile, "type EchoResponse struct")

	// Check overlay file (testdata has experimental_field annotations)
	overlayFile := readFile(t, outDir, "experimental/workflowservice/v1/example_overlay.go")
	require.NotContains(t, overlayFile, "//go:build experimental")
	require.Contains(t, overlayFile, "GetStartWorkflowExecutionRequestOverlay(")

	commonOverlayFile := readFile(t, outDir, "experimental/common/v1/example_overlay.go")
	require.NotContains(t, commonOverlayFile, "//go:build experimental")
	require.Contains(t, commonOverlayFile, "GetLinkOverlay(")
	require.Contains(t, commonOverlayFile, "msg *stable.Link")

	commonMsgFile := readFile(t, outDir, "experimental/common/v1/example_messages.pb.go")
	require.Contains(t, commonMsgFile, "type LinkOverlay struct")
	require.Contains(t, commonMsgFile, "type NexusOperation struct")
	require.NotContains(t, commonMsgFile, "OneofWrappers")

	// Check enum file
	enumFile := readFile(t, outDir, "experimental/enums/v1/example_enum.go")
	require.NotContains(t, enumFile, "//go:build experimental")
	require.Contains(t, enumFile, "WORKFLOW_ID_CONFLICT_POLICY_FOO")

	// Check go.mod
	goMod := readFile(t, outDir, "experimental/go.mod")
	require.Contains(t, goMod, "module github.com/temporalio/api-go/experimental")
	require.Contains(t, goMod, "go.temporal.io/api v1.2.3")
}

func TestRunRequiresVariant(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := runMain(nil, bytes.NewReader(nil), &stdout, &stderr)
	require.Equal(t, 2, code)
	require.Contains(t, stderr.String(), "-variant")
	require.Empty(t, stdout.String())
}
