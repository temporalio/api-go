package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateExampleFromAPISHA(t *testing.T) {
	apiRepo := newExampleAPIRepo(t)
	sha := strings.TrimSpace(runCmd(t, apiRepo, "git", "rev-parse", "HEAD"))

	outDir := t.TempDir()
	gen := generator{
		resolveStableVersion: func(string) (moduleVersion, error) {
			return moduleVersion{Version: "v1.2.3", GoVersion: "1.21"}, nil
		},
		skipGoModTidy: true,
	}
	err := gen.generate(apiRepo, sha, "example", outDir)
	require.NoError(t, err)

	goMod, err := os.ReadFile(filepath.Join(outDir, "go.mod"))
	require.NoError(t, err)
	require.Contains(t, string(goMod), "module github.com/temporalio/api-go/experimental/example")
	require.Contains(t, string(goMod), "go.temporal.io/api v1.2.3")

	_, err = os.Stat(filepath.Join(outDir, "workflowservice", "v1", "echo_extensions.pb.go"))
	require.ErrorIs(t, err, os.ErrNotExist)

	grpcFile, err := os.ReadFile(filepath.Join(outDir, "workflowservice", "v1", "experimental_service_grpc.pb.go"))
	require.NoError(t, err)
	require.Contains(t, string(grpcFile), "type WorkflowServiceClient interface")
	require.Contains(t, string(grpcFile), "func NewWorkflowServiceClient(cc grpc.ClientConnInterface) WorkflowServiceClient")
	require.Contains(t, string(grpcFile), "type WorkflowServiceServer interface")
	require.Contains(t, string(grpcFile), `WorkflowService_Echo_FullMethodName = "/temporal.api.workflowservice.v1.WorkflowService/Echo"`)
	require.Contains(t, string(grpcFile), `ServiceName: "temporal.api.workflowservice.v1.WorkflowService"`)

	overlayFile, err := os.ReadFile(filepath.Join(outDir, "workflowservice", "v1", "overlay.go"))
	require.NoError(t, err)
	require.Contains(t, string(overlayFile), "func GetStartWorkflowExecutionRequestOverlay(")
	require.Contains(t, string(overlayFile), "func SetStartWorkflowExecutionRequestOverlay(")
	require.Contains(t, string(overlayFile), "func ClearStartWorkflowExecutionRequestOverlay(")
	require.Contains(t, string(overlayFile), "func overlayFieldNumbers(")
	require.Contains(t, string(overlayFile), "protowire.ConsumeFieldValue(")

	enumFile, err := os.ReadFile(filepath.Join(outDir, "enums", "v1", "enum.go"))
	require.NoError(t, err)
	require.Contains(t, string(enumFile), "WORKFLOW_ID_CONFLICT_POLICY_FOO stable.WorkflowIdConflictPolicy = 1000")

	pbFile, err := os.ReadFile(filepath.Join(outDir, "workflowservice", "v1", "request_response.pb.go"))
	require.NoError(t, err)
	require.Contains(t, string(pbFile), "type EchoRequest struct")
	require.Contains(t, string(pbFile), "type EchoResponse struct")
	require.Contains(t, string(pbFile), "type Foo struct")
	require.Contains(t, string(pbFile), "type StartWorkflowExecutionRequestOverlay struct")
	require.Contains(t, string(pbFile), `protobuf:"bytes,1000,opt,name=foo,proto3"`)
	require.Contains(t, string(pbFile), `protobuf:"varint,1009,opt,name=foo_policy`)
	require.Contains(t, string(pbFile), "FooById")
	require.Contains(t, string(pbFile), "FooPolicy")
	require.NotContains(t, string(pbFile), "type StartWorkflowExecutionRequest struct")
}

func TestGenerateClearsOutputDir(t *testing.T) {
	apiRepo := newExampleAPIRepo(t)
	sha := strings.TrimSpace(runCmd(t, apiRepo, "git", "rev-parse", "HEAD"))

	outDir := t.TempDir()
	staleFile := filepath.Join(outDir, "stale.txt")
	require.NoError(t, os.WriteFile(staleFile, []byte("stale"), 0o644))

	gen := generator{
		resolveStableVersion: func(string) (moduleVersion, error) {
			return moduleVersion{Version: "v1.2.3", GoVersion: "1.21"}, nil
		},
		skipGoModTidy: true,
	}
	require.NoError(t, gen.generate(apiRepo, sha, "example", outDir))
	_, err := os.Stat(staleFile)
	require.ErrorIs(t, err, os.ErrNotExist)
}

func TestRunRequiresVariantAndAPISHA(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	code := run(nil, &stdout, &stderr)
	require.Equal(t, 2, code)
	require.Contains(t, stderr.String(), "-api-sha")
	require.Empty(t, stdout.String())
}

func TestRunWritesVariantModuleToExperimentalDir(t *testing.T) {
	apiRepo := newExampleAPIRepo(t)
	sha := strings.TrimSpace(runCmd(t, apiRepo, "git", "rev-parse", "HEAD"))

	workDir := t.TempDir()
	prevWD, err := os.Getwd()
	require.NoError(t, err)
	require.NoError(t, os.Chdir(workDir))
	t.Cleanup(func() {
		require.NoError(t, os.Chdir(prevWD))
	})

	prevResolver := defaultStableVersionResolver
	defaultStableVersionResolver = func(string) (moduleVersion, error) {
		return moduleVersion{Version: "v1.2.3", GoVersion: "1.21"}, nil
	}
	t.Cleanup(func() {
		defaultStableVersionResolver = prevResolver
	})
	prevNewGenerator := newGenerator
	newGenerator = func() generator {
		return generator{
			resolveStableVersion: defaultStableVersionResolver,
			skipGoModTidy:        true,
		}
	}
	t.Cleanup(func() {
		newGenerator = prevNewGenerator
	})

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	code := run([]string{"-variant", "example", "-api-sha", sha, "-api-repo", apiRepo}, &stdout, &stderr)
	require.Equal(t, 0, code, stderr.String())
	require.Contains(t, stdout.String(), "generated experimental/example")
	require.Empty(t, stderr.String())

	goMod, err := os.ReadFile(filepath.Join(workDir, "experimental", "example", "go.mod"))
	require.NoError(t, err)
	require.Contains(t, string(goMod), "module github.com/temporalio/api-go/experimental/example")
}
