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
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.temporal.io/api/command/v1"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/export/v1"
	"go.temporal.io/api/failure/v1"
	"go.temporal.io/api/history/v1"
	"go.temporal.io/api/protocol/v1"
	"go.temporal.io/api/update/v1"
	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

func inputPayloads() *common.Payloads {
	return &common.Payloads{
		Payloads: []*common.Payload{
			inputPayload(),
		},
	}
}

func inputPayload() *common.Payload {
	return &common.Payload{
		Metadata: map[string][]byte{
			"encoding": []byte("plain/json"),
		},
		Data: []byte("test"),
	}
}

func TestVisitPayloads(t *testing.T) {
	require := require.New(t)

	err := VisitPayloads(
		context.Background(),
		&workflowservice.StartWorkflowExecutionRequest{
			Input: inputPayloads(),
		},
		VisitPayloadsOptions{
			Visitor: func(vpc *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.False(vpc.SinglePayloadRequired)
				return p, nil
			},
		},
	)
	require.NoError(err)

	err = VisitPayloads(
		context.Background(),
		&workflowservice.StartWorkflowExecutionRequest{
			Header: &common.Header{
				Fields: map[string]*common.Payload{"test": inputPayload()},
			},
		},
		VisitPayloadsOptions{
			Visitor: func(vpc *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.True(vpc.SinglePayloadRequired)
				return p, nil
			},
		},
	)
	require.NoError(err)

	err = VisitPayloads(
		context.Background(),
		&export.WorkflowExecutions{Items: []*export.WorkflowExecution{{History: &history.History{
			Events: []*history.HistoryEvent{
				{
					Attributes: &history.HistoryEvent_WorkflowExecutionStartedEventAttributes{
						WorkflowExecutionStartedEventAttributes: &history.WorkflowExecutionStartedEventAttributes{
							Input: inputPayloads(),
						},
					},
				},
			},
		}}}},
		VisitPayloadsOptions{
			Visitor: func(vpc *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.False(vpc.SinglePayloadRequired)
				require.Equal([]byte("test"), p[0].Data)
				return p, nil
			},
		},
	)
	require.NoError(err)

	msg := &history.HistoryEvent{
		Attributes: &history.HistoryEvent_NexusOperationScheduledEventAttributes{
			NexusOperationScheduledEventAttributes: &history.NexusOperationScheduledEventAttributes{
				Input: inputPayload(),
			},
		},
	}
	err = VisitPayloads(
		context.Background(),
		msg,
		VisitPayloadsOptions{
			Visitor: func(vpc *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.True(vpc.SinglePayloadRequired)
				require.Equal([]byte("test"), p[0].Data)
				return []*common.Payload{{Data: []byte("visited")}}, nil
			},
		},
	)
	require.Equal([]byte("visited"), msg.GetNexusOperationScheduledEventAttributes().Input.Data)
	require.NoError(err)
}

func TestVisitPayloads_NestedParent(t *testing.T) {
	// Due to an invalid approach in the previous visitor, this test used to fail
	root := &command.StartChildWorkflowExecutionCommandAttributes{
		Header: &common.Header{
			Fields: map[string]*common.Payload{
				"header-key": {Data: []byte("header-value")},
			},
		},
		Input: &common.Payloads{
			Payloads: []*common.Payload{{Data: []byte("input-value")}},
		},
	}
	var headerParent, inputParent proto.Message
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			if len(p) == 1 {
				if string(p[0].Data) == "header-value" {
					headerParent = proto.Clone(ctx.Parent)
				} else if string(p[0].Data) == "input-value" {
					inputParent = proto.Clone(ctx.Parent)
				}
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.IsType(t, &common.Header{}, headerParent)
	require.IsType(t, &command.StartChildWorkflowExecutionCommandAttributes{}, inputParent)
}

func TestVisitPayloads_Any(t *testing.T) {
	// Due to us not visiting protos inside Any, this test used to fail
	msg1, err := anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
		Payloads: []*common.Payload{{Data: []byte("orig-val")}},
	}}})
	require.NoError(t, err)
	msg2, err := anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
		Payloads: []*common.Payload{{Data: []byte("orig-val-don't-touch")}},
	}}})
	require.NoError(t, err)
	msg3, err := anypb.New(&update.Response{Outcome: &update.Outcome{Value: &update.Outcome_Success{
		Success: &common.Payloads{
			Payloads: []*common.Payload{{Data: []byte("orig-val")}},
		},
	}}})
	require.NoError(t, err)
	root := &workflowservice.PollWorkflowTaskQueueResponse{
		Messages: []*protocol.Message{{Body: msg1}, {Body: msg2}, {Body: msg3}},
	}

	// Visit with any recursion enabled and only change orig-val
	err = VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			// Only mutate if the payloads has orig-val
			if len(p) == 1 && string(p[0].Data) == "orig-val" {
				return []*common.Payload{{Data: []byte("new-val")}}, nil
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	update1, err := root.Messages[0].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "new-val", string(update1.(*update.Request).Input.Args.Payloads[0].Data))
	update2, err := root.Messages[1].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val-don't-touch", string(update2.(*update.Request).Input.Args.Payloads[0].Data))
	update3, err := root.Messages[2].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "new-val", string(update3.(*update.Response).GetOutcome().GetSuccess().Payloads[0].Data))

	// Do the same test but with a do-nothing visitor and confirm unchanged
	msg1, err = anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
		Payloads: []*common.Payload{{Data: []byte("orig-val")}},
	}}})
	require.NoError(t, err)
	msg2, err = anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
		Payloads: []*common.Payload{{Data: []byte("orig-val-don't-touch")}},
	}}})
	require.NoError(t, err)
	msg3, err = anypb.New(&update.Response{Outcome: &update.Outcome{Value: &update.Outcome_Success{
		Success: &common.Payloads{
			Payloads: []*common.Payload{{Data: []byte("orig-val")}},
		},
	}}})
	require.NoError(t, err)
	root = &workflowservice.PollWorkflowTaskQueueResponse{
		Messages: []*protocol.Message{{Body: msg1}, {Body: msg2}, {Body: msg3}},
	}
	err = VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			// Only mutate if the payloads has orig-val
			if len(p) == 1 && string(p[0].Data) == "orig-val" {
				return []*common.Payload{{Data: []byte("new-val")}}, nil
			}
			return p, nil
		},
		WellKnownAnyVisitor: func(*VisitPayloadsContext, *anypb.Any) error { return nil },
	})
	require.NoError(t, err)
	update1, err = root.Messages[0].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val", string(update1.(*update.Request).Input.Args.Payloads[0].Data))
	update2, err = root.Messages[1].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val-don't-touch", string(update2.(*update.Request).Input.Args.Payloads[0].Data))
	update3, err = root.Messages[2].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val", string(update3.(*update.Response).GetOutcome().GetSuccess().Payloads[0].Data))
}

func TestVisitFailures(t *testing.T) {
	require := require.New(t)

	fail := &failure.Failure{}

	err := VisitFailures(
		context.Background(),
		&workflowservice.RespondActivityTaskFailedRequest{
			Failure: fail,
		},
		VisitFailuresOptions{
			Visitor: func(vfc *VisitFailuresContext, f *failure.Failure) error {
				require.Equal(fail, f)
				return nil
			},
		},
	)
	require.NoError(err)

	nestedFailure := &failure.Failure{Cause: fail}
	failureCount := 0

	err = VisitFailures(
		context.Background(),
		&workflowservice.RespondActivityTaskFailedRequest{
			Failure: nestedFailure,
		},
		VisitFailuresOptions{
			Visitor: func(vfc *VisitFailuresContext, f *failure.Failure) error {
				failureCount += 1
				return nil
			},
		},
	)
	require.NoError(err)
	require.Equal(2, failureCount)
}

func TestVisitFailuresAny(t *testing.T) {
	require := require.New(t)

	fail := &failure.Failure{
		Message: "test failure",
	}

	msg, err := anypb.New(&update.Response{Outcome: &update.Outcome{Value: &update.Outcome_Failure{
		Failure: fail,
	}}})
	require.NoError(err)

	req := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Messages: []*protocol.Message{{Body: msg}},
	}
	failureCount := 0
	err = VisitFailures(
		context.Background(),
		req,
		VisitFailuresOptions{
			Visitor: func(vfc *VisitFailuresContext, f *failure.Failure) error {
				failureCount += 1
				require.Equal("test failure", f.Message)
				f.EncodedAttributes = &common.Payload{Data: []byte("test failure")}
				f.Message = "encoded failure"
				return nil
			},
		},
	)
	require.NoError(err)
	require.Equal(1, failureCount)
	updateMsg, err := req.GetMessages()[0].GetBody().UnmarshalNew()
	require.NoError(err)
	require.Equal("encoded failure", updateMsg.(*update.Response).GetOutcome().GetFailure().GetMessage())
	require.Equal("test failure", string(updateMsg.(*update.Response).GetOutcome().GetFailure().EncodedAttributes.Data))

}

func TestClientInterceptor(t *testing.T) {
	require := require.New(t)

	server, err := startTestGRPCServer()
	require.NoError(err)

	inputs := inputPayloads()
	var inboundPayload *common.Payload
	var outboundPayload *common.Payload

	interceptor, err := NewPayloadVisitorInterceptor(
		PayloadVisitorInterceptorOptions{
			Outbound: &VisitPayloadsOptions{
				Visitor: func(vpc *VisitPayloadsContext, payloads []*common.Payload) ([]*common.Payload, error) {
					outboundPayload = payloads[0]
					return payloads, nil
				},
			},
			Inbound: &VisitPayloadsOptions{
				Visitor: func(vpc *VisitPayloadsContext, payloads []*common.Payload) ([]*common.Payload, error) {
					inboundPayload = payloads[0]
					return payloads, nil
				},
			},
		},
	)
	require.NoError(err)

	c, err := grpc.Dial(
		server.addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor),
	)
	require.NoError(err)

	client := workflowservice.NewWorkflowServiceClient(c)

	_, err = client.StartWorkflowExecution(
		context.Background(),
		&workflowservice.StartWorkflowExecutionRequest{
			Input: inputPayloads(),
		},
	)
	require.NoError(err)

	require.True(proto.Equal(inputs.Payloads[0], outboundPayload))

	_, err = client.PollActivityTaskQueue(
		context.Background(),
		&workflowservice.PollActivityTaskQueueRequest{},
	)
	require.NoError(err)

	require.True(proto.Equal(inputs.Payloads[0], inboundPayload))
}

func TestClientInterceptorFailures(t *testing.T) {
	require := require.New(t)

	server, err := startTestGRPCServer()
	require.NoError(err)

	inputs := inputPayloads()
	var inboundPayload *common.Payload
	//var outboundPayload *common.Payload

	interceptor, err := NewPayloadVisitorInterceptor(
		PayloadVisitorInterceptorOptions{
			Outbound: &VisitPayloadsOptions{
				Visitor: func(vpc *VisitPayloadsContext, payloads []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("[outbound]", vpc.Parent.ProtoReflect().Descriptor().Name())
					//outboundPayload = payloads[0]
					return payloads, nil
				},
			},
			Inbound: &VisitPayloadsOptions{
				Visitor: func(vpc *VisitPayloadsContext, payloads []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("[inbound]", vpc.Parent.ProtoReflect().Descriptor().Name())
					inboundPayload = payloads[0]
					return payloads, nil
				},
			},
		},
	)
	require.NoError(err)

	c, err := grpc.Dial(
		server.addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor),
	)
	require.NoError(err)

	client := workflowservice.NewWorkflowServiceClient(c)

	_, err = client.StartWorkflowExecution(
		context.Background(),
		&workflowservice.StartWorkflowExecutionRequest{
			Input: inputPayloads(),
		},
	)
	require.NoError(err)

	_, err = client.ExecuteMultiOperation(context.Background(), &workflowservice.ExecuteMultiOperationRequest{})
	// We expect that even though an error is returned, the Payload visitor visited the payload
	// included in the GRPC error details
	require.Error(err)
	require.True(proto.Equal(inputs.Payloads[0], inboundPayload))
}

type testGRPCServer struct {
	workflowservice.UnimplementedWorkflowServiceServer
	*grpc.Server
	listener                       net.Listener
	addr                           string
	startWorkflowExecutionRequest  *workflowservice.StartWorkflowExecutionRequest
	startWorkflowExecutionMetadata metadata.MD
}

func startTestGRPCServer() (*testGRPCServer, error) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return nil, err
	}
	t := &testGRPCServer{Server: grpc.NewServer(), listener: l, addr: l.Addr().String()}
	workflowservice.RegisterWorkflowServiceServer(t.Server, t)
	go func() {
		if err := t.Serve(l); err != nil {
			log.Fatal(err)
		}
	}()

	// Wait until get-system-info reports serving
	return t, t.waitUntilServing()
}

func (t *testGRPCServer) waitUntilServing() error {
	// Try 20 times, waiting 100ms between
	var lastErr error
	for i := 0; i < 20; i++ {
		conn, err := grpc.Dial(t.addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			lastErr = err
		} else {
			_, err := workflowservice.NewWorkflowServiceClient(conn).GetClusterInfo(
				context.Background(),
				&workflowservice.GetClusterInfoRequest{},
			)
			_ = conn.Close()
			if err != nil {
				lastErr = err
			} else {
				return nil
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
	return fmt.Errorf("failed waiting, last error: %w", lastErr)
}

func (t *testGRPCServer) Stop() {
	t.Server.Stop()
}

func (t *testGRPCServer) GetClusterInfo(
	context.Context,
	*workflowservice.GetClusterInfoRequest,
) (*workflowservice.GetClusterInfoResponse, error) {
	return &workflowservice.GetClusterInfoResponse{}, nil
}

func (t *testGRPCServer) StartWorkflowExecution(
	ctx context.Context,
	req *workflowservice.StartWorkflowExecutionRequest,
) (*workflowservice.StartWorkflowExecutionResponse, error) {
	t.startWorkflowExecutionRequest = req
	t.startWorkflowExecutionMetadata, _ = metadata.FromIncomingContext(ctx)
	return &workflowservice.StartWorkflowExecutionResponse{}, nil
}

func (t *testGRPCServer) PollActivityTaskQueue(
	ctx context.Context,
	req *workflowservice.PollActivityTaskQueueRequest,
) (*workflowservice.PollActivityTaskQueueResponse, error) {
	return &workflowservice.PollActivityTaskQueueResponse{
		Input: inputPayloads(),
	}, nil
}

func (t *testGRPCServer) ExecuteMultiOperation(
	ctx context.Context,
	req *workflowservice.ExecuteMultiOperationRequest) (*workflowservice.ExecuteMultiOperationResponse, error) {

	anyDetail, err := anypb.New(inputPayloads())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload to Any: %w", err)
	}

	st := status.New(codes.Internal, "Operation failed due to a user error")

	stWithDetails, err := st.WithDetails(anyDetail)
	if err != nil {
		return nil, st.Err()
	}

	return nil, stWithDetails.Err()
}
