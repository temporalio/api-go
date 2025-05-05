package proxy

import (
	"context"
	"net"
	"testing"

	"github.com/stretchr/testify/require"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func TestProxyMetadataForward(t *testing.T) {
	// Create an end server
	endSrv, err := startTestGRPCServer()
	require.NoError(t, err)
	defer endSrv.Stop()
	endConn, err := grpc.NewClient(endSrv.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	require.NoError(t, err)
	defer endConn.Close()

	// Create a proxy
	proxyImpl, err := NewWorkflowServiceProxyServer(WorkflowServiceProxyOptions{
		Client: workflowservice.NewWorkflowServiceClient(endConn),
	})
	require.NoError(t, err)
	proxyListener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	proxySrv := grpc.NewServer()
	workflowservice.RegisterWorkflowServiceServer(proxySrv, proxyImpl)
	go func() {
		if err := proxySrv.Serve(proxyListener); err != nil {
			t.Logf("Failed serving: %v", err)
		}
	}()
	defer proxySrv.Stop()

	// Create client to proxy
	clientConn, err := grpc.NewClient(
		proxyListener.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	defer clientConn.Close()
	client := workflowservice.NewWorkflowServiceClient(clientConn)

	// Make call with metadata and confirm properly set
	ctx := metadata.AppendToOutgoingContext(context.Background(), "my-header", "my-header-value")
	_, err = client.StartWorkflowExecution(ctx, &workflowservice.StartWorkflowExecutionRequest{
		WorkflowType: &common.WorkflowType{Name: "my-workflow-1"},
	})
	require.NoError(t, err)
	require.Equal(t, "my-workflow-1", endSrv.startWorkflowExecutionRequest.WorkflowType.Name)
	require.Equal(t, []string{"my-header-value"}, endSrv.startWorkflowExecutionMetadata.Get("my-header"))
	// Also make sure that authority is proper and didn't get overridden
	require.Equal(t, []string{endSrv.addr}, endSrv.startWorkflowExecutionMetadata.Get(":authority"))
}
