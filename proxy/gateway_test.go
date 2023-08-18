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

package proxy_test

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"strings"
	"testing"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/stretchr/testify/require"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/proxy"
	"go.temporal.io/api/query/v1"
	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGRPCGatewayProxy(t *testing.T) {
	ctx := context.Background()
	srv, err := StartTestServer(ctx)
	require.NoError(t, err)
	defer srv.Close()

	// Send gRPC request and confirm response
	srv.nextQueryResp = &workflowservice.QueryWorkflowResponse{
		QueryResult: &common.Payloads{
			Payloads: []*common.Payload{{
				Metadata: map[string][]byte{"encoding": []byte("json/plain")},
				Data:     []byte(`"resp-data"`),
			}},
		},
	}
	grpcConn, err := grpc.DialContext(
		ctx,
		srv.grpcListener.Addr().String(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	require.NoError(t, err)
	defer grpcConn.Close()
	grpcClient := workflowservice.NewWorkflowServiceClient(grpcConn)
	grpcReq := &workflowservice.QueryWorkflowRequest{
		Namespace: "my-namespace",
		Execution: &common.WorkflowExecution{WorkflowId: "my-workflow"},
		Query: &query.WorkflowQuery{
			QueryType: "my-query",
			QueryArgs: &common.Payloads{
				Payloads: []*common.Payload{{
					Metadata: map[string][]byte{"encoding": []byte("json/plain")},
					Data:     []byte(`"req-data"`),
				}},
			},
		},
	}
	grpcResp, err := grpcClient.QueryWorkflow(ctx, grpcReq)
	require.NoError(t, err)
	require.True(t, proto.Equal(srv.lastQueryReq, grpcReq))
	require.True(t, proto.Equal(srv.nextQueryResp, grpcResp))

	// Send HTTP request with ugly payloads
	srv.lastQueryReq = nil
	httpResp, err := http.Post(
		fmt.Sprintf(
			"http://%v/api/v1/namespaces/my-namespace/workflows/my-workflow/query/my-query?noPayloadShorthand",
			srv.httpListener.Addr(),
		),
		"application/json",
		strings.NewReader(`{
			"query": {
				"queryArgs": {
					"payloads": [{
						"metadata": {
							"encoding": "anNvbi9wbGFpbg=="
						},
						"data": "InJlcS1kYXRhIg=="
					}]
				}
			}
		}`),
	)
	require.NoError(t, err)
	httpRespBody, err := io.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, httpResp.StatusCode, http.StatusOK, "Bad status: %s", httpRespBody)
	require.True(t, proto.Equal(srv.lastQueryReq, grpcReq))
	require.JSONEq(t, `{
		"queryResult": {
			"payloads": [{
				"metadata": {
					"encoding": "anNvbi9wbGFpbg=="
				},
				"data": "InJlc3AtZGF0YSI="
			}]
		}
	}`, string(httpRespBody))

	// Now try with shorthand payload notation
	srv.lastQueryReq = nil
	httpResp, err = http.Post(
		fmt.Sprintf(
			"http://%v/api/v1/namespaces/my-namespace/workflows/my-workflow/query/my-query",
			srv.httpListener.Addr(),
		),
		"application/json",
		strings.NewReader(`{ "query": { "queryArgs": ["req-data"] } }`),
	)
	require.NoError(t, err)
	httpRespBody, err = io.ReadAll(httpResp.Body)
	httpResp.Body.Close()
	require.NoError(t, err)
	require.Equal(t, httpResp.StatusCode, http.StatusOK, "Bad status: %s", httpRespBody)
	require.True(t, proto.Equal(srv.lastQueryReq, grpcReq))
	require.JSONEq(t, `{ "queryResult": ["resp-data"] }`, string(httpRespBody))
}

type testServer struct {
	workflowservice.UnimplementedWorkflowServiceServer
	grpcListener net.Listener
	grpcServer   *grpc.Server
	httpListener net.Listener
	httpServer   http.Server
	serveErrCh   chan error

	lastQueryReq  *workflowservice.QueryWorkflowRequest
	nextQueryResp *workflowservice.QueryWorkflowResponse
	nextQueryErr  error
}

var shorthandMarshaler runtime.Marshaler
var nonShorthandMarshaler runtime.Marshaler

func init() {
	if m, err := proxy.NewJSONPBMarshaler(proxy.JSONPBMarshalerOptions{}); err != nil {
		panic(err)
	} else if u, err := proxy.NewJSONPBUnmarshaler(proxy.JSONPBUnmarshalerOptions{}); err != nil {
		panic(err)
	} else {
		shorthandMarshaler = proxy.NewGRPCGatewayJSONPBMarshaler(m, u)
	}
	if m, err := proxy.NewJSONPBMarshaler(proxy.JSONPBMarshalerOptions{DisablePayloadShorthand: true}); err != nil {
		panic(err)
	} else if u, err := proxy.NewJSONPBUnmarshaler(proxy.JSONPBUnmarshalerOptions{DisablePayloadShorthand: true}); err != nil {
		panic(err)
	} else {
		nonShorthandMarshaler = proxy.NewGRPCGatewayJSONPBMarshaler(m, u)
	}
}

func StartTestServer(ctx context.Context) (*testServer, error) {
	serveMux := runtime.NewServeMux(
		runtime.WithMarshalerOption("application/json+no-payload-shorthand", nonShorthandMarshaler),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, shorthandMarshaler),
	)
	t := &testServer{
		serveErrCh: make(chan error, 2),
		grpcServer: grpc.NewServer(),
	}
	t.httpServer.Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.URL.Query()["noPayloadShorthand"]; ok {
			r.Header.Set("Accept", "application/json+no-payload-shorthand")
		}
		serveMux.ServeHTTP(w, r)
	})

	// Start gRPC server
	var err error
	if t.grpcListener, err = net.Listen("tcp", "127.0.0.1:0"); err != nil {
		return nil, err
	}
	workflowservice.RegisterWorkflowServiceServer(t.grpcServer, t)
	go func() { t.serveErrCh <- t.grpcServer.Serve(t.grpcListener) }()

	// Start HTTP server
	if t.httpListener, err = net.Listen("tcp", "127.0.0.1:0"); err != nil {
		t.grpcServer.Stop()
		return nil, err
	}
	if err = workflowservice.RegisterWorkflowServiceHandlerServer(ctx, serveMux, t); err != nil {
		t.grpcServer.Stop()
		t.httpListener.Close()
	}
	go func() { t.serveErrCh <- t.httpServer.Serve(t.httpListener) }()
	return t, nil
}

func (t *testServer) Close() error {
	t.grpcServer.Stop()
	return t.httpServer.Close()
}

func (t *testServer) QueryWorkflow(
	ctx context.Context,
	req *workflowservice.QueryWorkflowRequest,
) (*workflowservice.QueryWorkflowResponse, error) {
	t.lastQueryReq = req
	return t.nextQueryResp, t.nextQueryErr
}
