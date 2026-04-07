package proxy

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"go.temporal.io/api/errordetails/v1"

	"github.com/stretchr/testify/require"
	"go.temporal.io/api/command/v1"
	"go.temporal.io/api/common/v1"
	_ "go.temporal.io/api/enums/v1"
	"go.temporal.io/api/export/v1"
	"go.temporal.io/api/failure/v1"
	_ "go.temporal.io/api/filter/v1"
	"go.temporal.io/api/history/v1"
	_ "go.temporal.io/api/namespace/v1"
	"go.temporal.io/api/protocol/v1"
	_ "go.temporal.io/api/query/v1"
	_ "go.temporal.io/api/schedule/v1"
	_ "go.temporal.io/api/taskqueue/v1"
	"go.temporal.io/api/update/v1"
	_ "go.temporal.io/api/version/v1"
	_ "go.temporal.io/api/workflow/v1"
	"go.temporal.io/api/workflowservice/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
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

func TestVisitPayloads_RepeatedPayload(t *testing.T) {
	root := &workflowservice.CountWorkflowExecutionsResponse_AggregationGroup{GroupValues: []*common.Payload{{Data: []byte("orig-val")}}}

	var count int
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			count += 1
			// Only mutate if the payloads has orig-val
			if len(p) == 1 && string(p[0].Data) == "orig-val" {
				return []*common.Payload{{Data: []byte("new-val")}}, nil
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.Equal(t, []*common.Payload{{Data: []byte("new-val")}}, root.GroupValues)
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

func TestVisitPayloads_RepeatedAny(t *testing.T) {
	msg, err := anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
		Payloads: []*common.Payload{{Data: []byte("orig-val")}},
	}}})
	require.NoError(t, err)
	root := &errordetails.MultiOperationExecutionFailure_OperationStatus{Details: []*anypb.Any{msg}}
	var anyCount int
	err = VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			anyCount++
			// Only mutate if the payloads has "test"
			if len(p) == 1 && string(p[0].Data) == "orig-val" {
				return []*common.Payload{{Data: []byte("new-val")}}, nil
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Equal(t, 1, anyCount)
	update1, err := root.Details[0].UnmarshalNew()

	require.NoError(t, err)
	require.Equal(t, "new-val", string(update1.(*update.Request).Input.Args.Payloads[0].Data))
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

func TestClientInterceptorGrpcFailures(t *testing.T) {
	require := require.New(t)

	server, err := startTestGRPCServer()
	require.NoError(err)

	inputs := inputPayloads()
	var inboundPayload *common.Payload
	var inboundFailure string
	failureMessage := "new failure message"

	interceptor, err := NewPayloadVisitorInterceptor(
		PayloadVisitorInterceptorOptions{
			Inbound: &VisitPayloadsOptions{
				Visitor: func(vpc *VisitPayloadsContext, payloads []*common.Payload) ([]*common.Payload, error) {
					inboundPayload = payloads[0]
					payloads[0] = &common.Payload{Data: []byte("new-val")}
					return payloads, nil
				},
			},
		},
	)
	require.NoError(err)

	failureInterceptor, err := NewFailureVisitorInterceptor(
		FailureVisitorInterceptorOptions{
			Inbound: &VisitFailuresOptions{Visitor: func(vpc *VisitFailuresContext, f *failure.Failure) error {
				inboundFailure = f.Message
				f.Message = failureMessage
				return nil
			}},
		})

	c, err := grpc.Dial(
		server.addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(interceptor, failureInterceptor),
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
	stat, ok := status.FromError(err)
	require.True(ok)
	for _, detail := range stat.Details() {
		multiOpFailure, ok := detail.(*errordetails.MultiOperationExecutionFailure)
		require.True(ok)
		payloads := &common.Payloads{}
		err = multiOpFailure.Statuses[0].Details[0].UnmarshalTo(payloads)
		require.NoError(err)

		newPayload := &common.Payload{Data: []byte("new-val")}
		require.True(proto.Equal(payloads.Payloads[0], newPayload))
	}

	_, err = client.QueryWorkflow(context.Background(), &workflowservice.QueryWorkflowRequest{})
	require.Error(err)
	require.Equal("test failure", inboundFailure)
	stat, ok = status.FromError(err)
	require.True(ok)
	for _, detail := range stat.Details() {
		queryFailure, ok := detail.(*errordetails.QueryFailedFailure)
		require.True(ok)
		require.Equal(failureMessage, queryFailure.Failure.Message)
	}

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
		return nil, err
	}
	operationStatus := &errordetails.MultiOperationExecutionFailure_OperationStatus{Details: []*anypb.Any{anyDetail}}
	multiOpFailure := errordetails.MultiOperationExecutionFailure{
		Statuses: []*errordetails.MultiOperationExecutionFailure_OperationStatus{operationStatus},
	}
	st := status.New(codes.Internal, "Operation failed due to a user error")

	stWithDetails, err := st.WithDetails(&multiOpFailure)
	if err != nil {
		return nil, st.Err()
	}

	return nil, stWithDetails.Err()
}

func (t *testGRPCServer) QueryWorkflow(
	ctx context.Context,
	req *workflowservice.QueryWorkflowRequest) (*workflowservice.QueryWorkflowResponse, error) {
	failureMessage := failure.Failure{
		Message: "test failure",
	}
	queryFailure := errordetails.QueryFailedFailure{Failure: &failureMessage}

	st := status.New(codes.Internal, "Operation failed due to a user error")

	stWithDetails, err := st.WithDetails(&queryFailure)
	if err != nil {
		return nil, st.Err()
	}

	return nil, stWithDetails.Err()
}

// Recursively crawl and test Payload(s) with Visitor
func populatePayload(root *proto.Message, msg proto.Message, require *require.Assertions, totalCount *int, count *int) {
	m := msg.ProtoReflect()
	fields := m.Descriptor().Fields()
	// Don't need to parse non-temporal types
	if !strings.HasPrefix(string(m.Descriptor().FullName()), "temporal.api.") && string(m.Descriptor().FullName()) != "google.protobuf.Any" {
		return
	}

	if m.Descriptor() == nil {
		panic("fail")
	}

	// Base case, ensure Visitor can reach Payload from root Message
	switch i := msg.(type) {
	case *common.Payload, *common.Payloads:
		*count++
		*totalCount++
		err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
			Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.Equal(1, *count)
				*count--
				return p, nil
			},
		})
		require.NoError(err)
		return
	case *anypb.Any:
		if i.TypeUrl == "" {
			// Set to a random proto struct we know contains a payload, to test if we
			// are able to recurse through Any to reach a payload
			newAny, err := anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
				Payloads: []*common.Payload{{Data: []byte("orig-val")}},
			}}})
			require.NoError(err)
			proto.Merge(msg, newAny)
		}
		*count++
		*totalCount++
		err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
			Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.Equal(1, *count)
				*count--
				return p, nil
			},
		})
		require.NoError(err)
		return
	}

	// Go through all fields, populating each then recursing into them to discover Payloads to test
	// with Visitor
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		value := m.Get(fd)

		if oneof := fd.ContainingOneof(); oneof != nil && fd.Kind() == protoreflect.MessageKind {
			newMsg := value.Message().New()
			m.Set(fd, protoreflect.ValueOf(newMsg))
			populatePayload(root, newMsg.Interface(), require, totalCount, count)
			// This ensures only 1 payload is set and discoverable from root at a time.
			m.Clear(fd)
		} else if fd.IsMap() {
			mapVal := m.Mutable(fd).Map()
			require.Equal(0, mapVal.Len())
			if fd.MapKey().Kind() == protoreflect.StringKind &&
				fd.MapValue().Kind() == protoreflect.MessageKind &&
				string(fd.MapValue().Message().FullName()) == "temporal.api.common.v1.Payload" {
				sampleKey := protoreflect.ValueOf("sample_key").MapKey()
				mapVal.Set(sampleKey, protoreflect.ValueOf(inputPayload().ProtoReflect()))
				mapVal.Range(func(key protoreflect.MapKey, val protoreflect.Value) bool {
					if fd.MapValue().Kind() == protoreflect.MessageKind {
						populatePayload(root, val.Message().Interface(), require, totalCount, count)
					}
					return true
				})
				mapVal.Clear(sampleKey)
			} else if fd.MapValue().Kind() == protoreflect.MessageKind {
				var sampleKey protoreflect.MapKey
				switch fd.MapKey().Kind() {
				case protoreflect.StringKind:
					sampleKey = protoreflect.ValueOf("sample_key").MapKey()
				case protoreflect.Int32Kind, protoreflect.Sint32Kind, protoreflect.Sfixed32Kind:
					sampleKey = protoreflect.ValueOf(int32(1)).MapKey()
				case protoreflect.Int64Kind, protoreflect.Sint64Kind, protoreflect.Uint64Kind:
					sampleKey = protoreflect.ValueOf(int64(1)).MapKey()
				case protoreflect.BoolKind:
					sampleKey = protoreflect.ValueOf(true).MapKey()
				default:
					fmt.Println("Skipping unsupported map key type:", fd.MapKey().Kind())
					return
				}
				mapVal.Set(sampleKey, mapVal.NewValue())
				mapVal.Range(func(key protoreflect.MapKey, val protoreflect.Value) bool {
					if fd.MapValue().Kind() == protoreflect.MessageKind {
						newMsg := val.Message()
						populatePayload(root, newMsg.Interface(), require, totalCount, count)
					}
					return true
				})
				// This ensures only 1 payload is set and discoverable from root at a time.
				mapVal.Clear(sampleKey)
			}
		} else if fd.IsList() {
			if fd.Kind() == protoreflect.MessageKind {
				listVal := m.Mutable(fd).List()
				require.Equal(0, listVal.Len())

				sampleVal := listVal.NewElement()
				listVal.Append(sampleVal)

				val := listVal.Get(0)
				require.True(val.Message().IsValid())
				require.Equal(1, listVal.Len())
				populatePayload(root, sampleVal.Message().Interface(), require, totalCount, count)
				// This ensures only 1 payload is set and discoverable from root at a time.
				listVal.Truncate(0)
			}
		} else {
			if fd.Kind() == protoreflect.MessageKind {
				// Avoid cycles
				if value.Message().Descriptor().FullName() == m.Descriptor().FullName() {
					continue
				}

				var newMsg protoreflect.Message
				newMsg = value.Message().New()
				m.Set(fd, protoreflect.ValueOf(newMsg))
				populatePayload(root, newMsg.Interface(), require, totalCount, count)
				// This ensures only 1 payload is set and discoverable from root at a time.
				m.Clear(fd)
			}
		}
		// Validate that all Payloads were found
		require.Equal(0, *count)
	}
}

func TestVisitPayloads_FailureCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.failure.v1.Failure") {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	var totalCount, count int
	populatePayload(&msg, msg, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(5, totalCount)
}

func TestVisitPayloads_UpdateRejectionCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.update.v1.Rejection") {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	var totalCount, count int
	populatePayload(&msg, msg, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(7, totalCount)
}

func TestVisitPayloads_PayloadsCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.query.v1.WorkflowQueryResult") {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	var totalCount, count int
	populatePayload(&msg, msg, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(6, totalCount)
}

func TestVisitPayloads_AnyCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if string(mt.Descriptor().FullName()) == "temporal.api.protocol.v1.Message" {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	var totalCount, count int
	populatePayload(&msg, msg, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(1, totalCount)
}

func TestVisitPayloads_CommandCount(t *testing.T) {
	require := require.New(t)
	// UserMetadata - 2
	// ScheduleActivityTaskCommandAttributes - 2
	// 		header - 1
	//  	payloads - 1
	//	CompleteWorkflowExecutionCommandAttributes - 1
	//	FailWorkflowExecutionCommandAttributes - 5
	//		failure - 5
	//	CancelWorkflowExecutionCommandAttributes - 1
	//	RecordMarkerCommandAttributes - 7
	//		details - 1
	//		header - 1
	//		failure - 5
	// 	ContinueAsNewWorkflowExecutionCommandAttributes - 10
	//		input - 1
	//		failure - 5
	//		last_completion_result - 1
	//		header - 1
	//		memo - 1
	//		SearchAttributes - 1
	//
	//	StartChildWorkflowExecutionCommandAttributes - 4
	//		input - 1
	//		header
	//		memo
	//		searchAttributes
	//	SignalExternalWorkflowExecutionCommandAttributes - 2
	//		header - 1
	//		input - 1
	//	UpsertWorkflowSearchAttributesCommandAttributes - 1
	//		searchAttributes - 1
	//	ModifyWorkflowPropertiesCommandAttributes - 1
	//		memo - 1
	//	ScheduleNexusOperationCommandAttributes - 1
	//		input - 1
	// TOTAL : 37
	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if string(mt.Descriptor().FullName()) == "temporal.api.command.v1.Command" {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	var totalCount, count int
	populatePayload(&msg, msg, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(37, totalCount)
}

func TestVisitPayloads_MapCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	var totalCount, count int

	// 	repeated temporal.api.command.v1.Command commands - 37
	// 	map<string, temporal.api.query.v1.WorkflowQueryResult> query_results - 6
	// 	repeated temporal.api.protocol.v1.Message messages - 1
	// TOTAL - 44
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if string(mt.Descriptor().FullName()) == "temporal.api.workflowservice.v1.RespondWorkflowTaskCompletedRequest" {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg1 := messageType.New().Interface().(proto.Message)
	totalCount = 0
	count = 0
	populatePayload(&msg1, msg1, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(44, totalCount)
}

func TestVisitPayloads_CountWorkflowExecutionsResponse(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	var totalCount, count int

	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if string(mt.Descriptor().FullName()) == "temporal.api.workflowservice.v1.CountWorkflowExecutionsResponse" {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg1 := messageType.New().Interface().(proto.Message)
	totalCount = 0
	count = 0
	populatePayload(&msg1, msg1, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(1, totalCount)
}

func TestVisitPayloads_ResponseCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if string(mt.Descriptor().FullName()) == "temporal.api.update.v1.Response" {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	var totalCount, count int
	populatePayload(&msg, msg, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(6, totalCount)
}

func TestVisitPayloads_OperationStatus(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if string(mt.Descriptor().FullName()) == "temporal.api.errordetails.v1.MultiOperationExecutionFailure.OperationStatus" {
			messageType = mt
		}
		return true
	})

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	var totalCount, count int
	populatePayload(&msg, msg, require, &totalCount, &count)

	require.Equal(0, count)
	require.Equal(1, totalCount)
}

func TestVisitPayloads_Everything(t *testing.T) {
	require := require.New(t)

	var messageType []protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		// The base case of passing Payload into the visitor is not supported.
		// See godoc for VisitPayloads
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.") && string(mt.Descriptor().FullName()) != "temporal.api.common.v1.Payload" {
			messageType = append(messageType, mt)
		}
		return true
	})
	for _, mt := range messageType {
		// Create empty instance and populate with test values
		msg := mt.New().Interface().(proto.Message)

		var totalCount, count int
		populatePayload(&msg, msg, require, &totalCount, &count)

		require.Equal(0, count)

	}
}

// contextHookKey is an unexported key for values injected by ContextHook tests.
type contextHookKey struct{}

func TestContextHook_NilHookIsNoop(t *testing.T) {
	root := &workflowservice.StartWorkflowExecutionRequest{
		Input: inputPayloads(),
	}
	var called bool
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		// ContextHook deliberately not set
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			called = true
			return p, nil
		},
	})
	require.NoError(t, err)
	require.True(t, called)
}

func TestContextHook_ContextPropagatedToVisitor(t *testing.T) {
	root := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: []*command.Command{
			{
				Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
					ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
						Input: inputPayloads(),
					},
				},
			},
		},
	}

	var capturedValue interface{}
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		ContextHook: func(ctx context.Context, msg proto.Message) (context.Context, error) {
			if _, ok := msg.(*command.ScheduleActivityTaskCommandAttributes); ok {
				return context.WithValue(ctx, contextHookKey{}, "schedule-activity"), nil
			}
			return ctx, nil
		},
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			capturedValue = ctx.Value(contextHookKey{})
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Equal(t, "schedule-activity", capturedValue)
}

func TestContextHook_ContextRestoredAfterSiblingMessage(t *testing.T) {
	// Two commands side-by-side: context set for first must not leak into second.
	root := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: []*command.Command{
			{
				Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
					ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
						Input: inputPayloads(),
					},
				},
			},
			{
				Attributes: &command.Command_CompleteWorkflowExecutionCommandAttributes{
					CompleteWorkflowExecutionCommandAttributes: &command.CompleteWorkflowExecutionCommandAttributes{
						Result: inputPayloads(),
					},
				},
			},
		},
	}

	var seenValues []interface{}
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		ContextHook: func(ctx context.Context, msg proto.Message) (context.Context, error) {
			if _, ok := msg.(*command.ScheduleActivityTaskCommandAttributes); ok {
				return context.WithValue(ctx, contextHookKey{}, "schedule-activity"), nil
			}
			return ctx, nil
		},
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			seenValues = append(seenValues, ctx.Value(contextHookKey{}))
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Len(t, seenValues, 2)
	require.Equal(t, "schedule-activity", seenValues[0], "first payload should be inside ScheduleActivity context")
	require.Nil(t, seenValues[1], "second payload should not see stale context from first command")
}

func TestContextHook_ContextRestoredOnVisitorError(t *testing.T) {
	root := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: []*command.Command{
			{
				Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
					ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
						Input: inputPayloads(),
					},
				},
			},
		},
	}

	var ctxSeenByVisitor context.Context
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		ContextHook: func(ctx context.Context, msg proto.Message) (context.Context, error) {
			if _, ok := msg.(*command.ScheduleActivityTaskCommandAttributes); ok {
				return context.WithValue(ctx, contextHookKey{}, "hook-set"), nil
			}
			return ctx, nil
		},
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			ctxSeenByVisitor = ctx.Context
			return nil, fmt.Errorf("visitor error")
		},
	})
	require.Error(t, err)
	require.Equal(t, "hook-set", ctxSeenByVisitor.Value(contextHookKey{}), "visitor should have seen context set by hook")
}

func TestContextHook_ErrorPropagates(t *testing.T) {
	root := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: []*command.Command{
			{
				Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
					ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
						Input: inputPayloads(),
					},
				},
			},
		},
	}

	hookErr := fmt.Errorf("hook error")
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		ContextHook: func(ctx context.Context, msg proto.Message) (context.Context, error) {
			if _, ok := msg.(*command.ScheduleActivityTaskCommandAttributes); ok {
				return ctx, hookErr
			}
			return ctx, nil
		},
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			return p, nil
		},
	})
	require.ErrorIs(t, err, hookErr)
}

func TestContextHook_SearchAttributesVisitedWhenNotSkipped(t *testing.T) {
	root := &workflowservice.StartWorkflowExecutionRequest{
		SearchAttributes: &common.SearchAttributes{
			IndexedFields: map[string]*common.Payload{
				"key": {Data: []byte("sa-val")},
			},
		},
		Input: inputPayloads(),
	}

	hookFiredForSearchAttrs := false
	var visitedData []string
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		// SkipSearchAttributes not set — default false
		ContextHook: func(ctx context.Context, msg proto.Message) (context.Context, error) {
			if _, ok := msg.(*common.SearchAttributes); ok {
				hookFiredForSearchAttrs = true
			}
			return ctx, nil
		},
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			for _, pl := range p {
				visitedData = append(visitedData, string(pl.Data))
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.True(t, hookFiredForSearchAttrs, "ContextHook should fire for SearchAttributes when SkipSearchAttributes=false")
	require.Contains(t, visitedData, "sa-val")
	require.Contains(t, visitedData, "test")
}

func TestContextHook_SkipSearchAttributesRespected(t *testing.T) {
	root := &workflowservice.StartWorkflowExecutionRequest{
		SearchAttributes: &common.SearchAttributes{
			IndexedFields: map[string]*common.Payload{
				"key": {Data: []byte("sa-val")},
			},
		},
		Input: inputPayloads(),
	}

	hookFiredForSearchAttrs := false
	var visitedData []string
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		SkipSearchAttributes: true,
		ContextHook: func(ctx context.Context, msg proto.Message) (context.Context, error) {
			if _, ok := msg.(*common.SearchAttributes); ok {
				hookFiredForSearchAttrs = true
			}
			return ctx, nil
		},
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			for _, pl := range p {
				visitedData = append(visitedData, string(pl.Data))
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.False(t, hookFiredForSearchAttrs, "ContextHook must not fire for SearchAttributes when SkipSearchAttributes=true")
	require.NotContains(t, visitedData, "sa-val")
	require.Contains(t, visitedData, "test")
}

func TestVisitPayloadsConcurrent(t *testing.T) {
	// Build a message that exercises all three payload containers:
	//   - WorkflowExecutionStartedEventAttributes.Input (*common.Payloads) → visited as a slice
	//   - WorkflowExecutionStartedEventAttributes.Header.Fields (map[string]*common.Payload) → visited as map
	//   - NexusOperationScheduledEventAttributes.Input (*common.Payload) — single payload field
	msg := &history.History{
		Events: []*history.HistoryEvent{
			{
				Attributes: &history.HistoryEvent_WorkflowExecutionStartedEventAttributes{
					WorkflowExecutionStartedEventAttributes: &history.WorkflowExecutionStartedEventAttributes{
						Input: &common.Payloads{
							Payloads: []*common.Payload{
								{Data: []byte("payloads-0")},
								{Data: []byte("payloads-1")},
							},
						},
						Header: &common.Header{
							Fields: map[string]*common.Payload{
								"k1": {Data: []byte("map-k1")},
								"k2": {Data: []byte("map-k2")},
							},
						},
					},
				},
			},
			{
				Attributes: &history.HistoryEvent_NexusOperationScheduledEventAttributes{
					NexusOperationScheduledEventAttributes: &history.NexusOperationScheduledEventAttributes{
						Input: &common.Payload{Data: []byte("nexus-input")},
					},
				},
			},
		},
	}

	var visited sync.Map

	err := VisitPayloads(context.Background(), msg, VisitPayloadsOptions{
		ConcurrencyLimit: 4,
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			out := make([]*common.Payload, len(p))
			for i, pl := range p {
				visited.Store(string(pl.Data), true)
				out[i] = &common.Payload{Data: append([]byte("visited-"), pl.Data...)}
			}
			return out, nil
		},
	})
	require.NoError(t, err)

	// All original payloads must have been visited.
	for _, key := range []string{"payloads-0", "payloads-1", "map-k1", "map-k2", "nexus-input"} {
		_, ok := visited.Load(key)
		require.True(t, ok, "payload %q not visited", key)
	}

	// Results must be written back.
	startedAttrs := msg.Events[0].GetWorkflowExecutionStartedEventAttributes()
	nexusAttrs := msg.Events[1].GetNexusOperationScheduledEventAttributes()
	require.Equal(t, []byte("visited-payloads-0"), startedAttrs.Input.Payloads[0].Data)
	require.Equal(t, []byte("visited-payloads-1"), startedAttrs.Input.Payloads[1].Data)
	require.Equal(t, []byte("visited-map-k1"), startedAttrs.Header.Fields["k1"].Data)
	require.Equal(t, []byte("visited-map-k2"), startedAttrs.Header.Fields["k2"].Data)
	require.Equal(t, []byte("visited-nexus-input"), nexusAttrs.Input.Data)
}

func TestVisitPayloadsConcurrentMaxInflight(t *testing.T) {
	const limit = 3
	const total = 20

	payloads := make([]*common.Payload, total)
	for i := range payloads {
		payloads[i] = &common.Payload{Data: []byte(fmt.Sprintf("p%d", i))}
	}
	msg := &common.Payloads{Payloads: payloads}
	// Wrap in a message that contains a *common.Payloads field.
	req := &workflowservice.StartWorkflowExecutionRequest{
		Input: msg,
	}

	var inflight atomic.Int64
	var maxSeen atomic.Int64

	blocker := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		// Unblock all goroutines after a short delay.
		<-blocker
	}()

	err := VisitPayloads(context.Background(), req, VisitPayloadsOptions{
		ConcurrencyLimit: limit,
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			cur := inflight.Add(1)
			defer inflight.Add(-1)
			for {
				old := maxSeen.Load()
				if cur <= old || maxSeen.CompareAndSwap(old, cur) {
					break
				}
			}
			return p, nil
		},
	})
	close(blocker)
	wg.Wait()
	require.NoError(t, err)
	require.LessOrEqual(t, maxSeen.Load(), int64(limit), "concurrent inflight exceeded ConcurrencyLimit")
}

func TestVisitPayloadsConcurrentBarrier(t *testing.T) {
	// Prove that at least ConcurrencyLimit visitors run truly concurrently by
	// blocking each visitor at a barrier until exactly that many have entered.
	const limit = 4
	commands := make([]*command.Command, limit)
	for i := 0; i < limit; i++ {
		commands[i] = &command.Command{
			Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
				ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
					Input: &common.Payloads{
						Payloads: []*common.Payload{{Data: []byte(fmt.Sprintf("p%d", i))}},
					},
				},
			},
		}
	}
	req := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: commands,
	}

	var entered atomic.Int64
	barrier := make(chan struct{})

	err := VisitPayloads(context.Background(), req, VisitPayloadsOptions{
		ConcurrencyLimit: limit,
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			if entered.Add(1) == limit {
				close(barrier)
			}
			<-barrier
			return p, nil
		},
	})
	require.NoError(t, err)
}

func TestVisitPayloadsSequentialCancellationIgnored(t *testing.T) {
	// In sequential mode, context cancellation is not checked between visits.
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	msg := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: []*command.Command{
			{
				Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
					ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
						Input: &common.Payloads{Payloads: []*common.Payload{{Data: []byte("a")}}},
					},
				},
			},
			{
				Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
					ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
						Input: &common.Payloads{Payloads: []*common.Payload{{Data: []byte("b")}}},
					},
				},
			},
		},
	}

	var visited []string
	err := VisitPayloads(ctx, msg, VisitPayloadsOptions{
		ConcurrencyLimit: 1,
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			visited = append(visited, string(p[0].Data))
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Equal(t, []string{"a", "b"}, visited, "sequential mode must visit all payloads regardless of cancellation")
}

func TestVisitPayloadsConcurrentCancellation(t *testing.T) {
	// In concurrent mode, context cancellation is detected at semaphore
	// acquisition, so traversal stops promptly without the Visitor needing to
	// check the context.
	ctx, cancel := context.WithCancel(context.Background())

	// Block visitors until we cancel, keeping the semaphore full.
	const limit = 2
	allEntered := make(chan struct{})
	unblock := make(chan struct{})

	var enteredCount atomic.Int64
	makeCommand := func(data string) *command.Command {
		return &command.Command{
			Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
				ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
					Input: &common.Payloads{Payloads: []*common.Payload{{Data: []byte(data)}}},
				},
			},
		}
	}
	msg := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: []*command.Command{
			makeCommand("a"),
			makeCommand("b"),
			makeCommand("c"),
		},
	}

	done := make(chan error, 1)
	go func() {
		done <- VisitPayloads(ctx, msg, VisitPayloadsOptions{
			ConcurrencyLimit: limit,
			Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				if enteredCount.Add(1) == limit {
					close(allEntered)
				}
				<-unblock
				return p, nil
			},
		})
	}()

	<-allEntered
	cancel()
	close(unblock)

	err := <-done
	require.ErrorIs(t, err, context.Canceled)
}

func TestVisitPayloadsConcurrentCancellationDrainsGoroutines(t *testing.T) {
	// Verify that VisitPayloads waits for all already-spawned goroutines to
	// complete before returning, even when the context is cancelled mid-traversal.
	ctx, cancel := context.WithCancel(context.Background())

	const limit = 2
	allEntered := make(chan struct{})
	unblock := make(chan struct{})

	var enteredCount atomic.Int64
	var inflight atomic.Int64

	makeCommand := func(data string) *command.Command {
		return &command.Command{
			Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
				ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
					Input: &common.Payloads{Payloads: []*common.Payload{{Data: []byte(data)}}},
				},
			},
		}
	}
	msg := &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: []*command.Command{
			makeCommand("a"),
			makeCommand("b"),
			makeCommand("c"),
		},
	}

	err := func() error {
		done := make(chan error, 1)
		go func() {
			done <- VisitPayloads(ctx, msg, VisitPayloadsOptions{
				ConcurrencyLimit: limit,
				Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
					inflight.Add(1)
					if enteredCount.Add(1) == limit {
						close(allEntered)
					}
					<-unblock
					inflight.Add(-1)
					return p, nil
				},
			})
		}()
		<-allEntered
		cancel()
		close(unblock)
		return <-done
	}()

	require.ErrorIs(t, err, context.Canceled)
	require.Equal(t, int64(0), inflight.Load(), "all in-flight goroutines must complete before VisitPayloads returns")
}

func TestVisitPayloadsConcurrentError(t *testing.T) {
	visitorErr := fmt.Errorf("visitor error")

	// *common.Payloads path: one goroutine per command's Input field.
	const limit = 4
	commands := make([]*command.Command, limit)
	for i := 0; i < limit; i++ {
		commands[i] = &command.Command{
			Attributes: &command.Command_ScheduleActivityTaskCommandAttributes{
				ScheduleActivityTaskCommandAttributes: &command.ScheduleActivityTaskCommandAttributes{
					Input: &common.Payloads{
						Payloads: []*common.Payload{{Data: []byte(fmt.Sprintf("p%d", i))}},
					},
				},
			},
		}
	}
	err := VisitPayloads(context.Background(), &workflowservice.RespondWorkflowTaskCompletedRequest{
		Commands: commands,
	}, VisitPayloadsOptions{
		ConcurrencyLimit: limit,
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			if string(p[0].Data) == "p2" {
				return nil, visitorErr
			}
			return p, nil
		},
	})
	require.ErrorIs(t, err, visitorErr)

	// map[string]*common.Payload path: one goroutine per map entry.
	fields := make(map[string]*common.Payload, limit)
	for i := 0; i < limit; i++ {
		fields[fmt.Sprintf("k%d", i)] = &common.Payload{Data: []byte(fmt.Sprintf("v%d", i))}
	}
	err = VisitPayloads(context.Background(), &workflowservice.StartWorkflowExecutionRequest{
		Header: &common.Header{Fields: fields},
	}, VisitPayloadsOptions{
		ConcurrencyLimit: limit,
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			if string(p[0].Data) == "v2" {
				return nil, visitorErr
			}
			return p, nil
		},
	})
	require.ErrorIs(t, err, visitorErr)
}

func TestVisitPayloadsConcurrentAny(t *testing.T) {
	// Verify that defaultWellKnownAnyVisitor correctly visits and re-marshals.
	msg1, err := anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
		Payloads: []*common.Payload{{Data: []byte("any-a")}},
	}}})
	require.NoError(t, err)
	msg2, err := anypb.New(&update.Request{Input: &update.Input{Args: &common.Payloads{
		Payloads: []*common.Payload{{Data: []byte("any-b")}},
	}}})
	require.NoError(t, err)
	msg3, err := anypb.New(&update.Response{Outcome: &update.Outcome{Value: &update.Outcome_Success{
		Success: &common.Payloads{
			Payloads: []*common.Payload{{Data: []byte("any-c")}},
		},
	}}})
	require.NoError(t, err)

	root := &workflowservice.PollWorkflowTaskQueueResponse{
		Messages: []*protocol.Message{{Body: msg1}, {Body: msg2}, {Body: msg3}},
	}

	err = VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		ConcurrencyLimit: 4,
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			out := make([]*common.Payload, len(p))
			for i, pl := range p {
				out[i] = &common.Payload{Data: append([]byte("visited-"), pl.Data...)}
			}
			return out, nil
		},
	})
	require.NoError(t, err)

	// All three Any payloads must have been visited and re-marshaled correctly.
	got1, err := root.Messages[0].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "visited-any-a", string(got1.(*update.Request).Input.Args.Payloads[0].Data))

	got2, err := root.Messages[1].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "visited-any-b", string(got2.(*update.Request).Input.Args.Payloads[0].Data))

	got3, err := root.Messages[2].Body.UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "visited-any-c", string(got3.(*update.Response).GetOutcome().GetSuccess().Payloads[0].Data))
}

func TestVisitPayloadsLimit1IsSequential(t *testing.T) {
	// ConcurrencyLimit <= 1 must produce correct results identical to the default.
	msg := &workflowservice.StartWorkflowExecutionRequest{
		Input: &common.Payloads{
			Payloads: []*common.Payload{
				{Data: []byte("a")},
				{Data: []byte("b")},
			},
		},
		Header: &common.Header{
			Fields: map[string]*common.Payload{
				"h": {Data: []byte("c")},
			},
		},
	}

	for _, limit := range []int{0, 1} {
		msg2 := proto.Clone(msg).(*workflowservice.StartWorkflowExecutionRequest)
		err := VisitPayloads(context.Background(), msg2, VisitPayloadsOptions{
			ConcurrencyLimit: limit,
			Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				out := make([]*common.Payload, len(p))
				for i, pl := range p {
					out[i] = &common.Payload{Data: append([]byte("x"), pl.Data...)}
				}
				return out, nil
			},
		})
		require.NoError(t, err, "limit=%d", limit)
		require.Equal(t, []byte("xa"), msg2.Input.Payloads[0].Data, "limit=%d", limit)
		require.Equal(t, []byte("xb"), msg2.Input.Payloads[1].Data, "limit=%d", limit)
		require.Equal(t, []byte("xc"), msg2.Header.Fields["h"].Data, "limit=%d", limit)
	}
}
