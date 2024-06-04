// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
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
	defer endSrv.Close()
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
