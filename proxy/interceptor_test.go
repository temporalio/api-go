package proxy

import (
	"context"
	"fmt"
	"log"
	"net"
	"strings"
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
	return common.Payloads_builder{
		Payloads: []*common.Payload{
			inputPayload(),
		},
	}.Build()
}

func inputPayload() *common.Payload {
	return common.Payload_builder{
		Metadata: map[string][]byte{
			"encoding": []byte("plain/json"),
		},
		Data: []byte("test"),
	}.Build()
}

func TestVisitPayloads(t *testing.T) {
	require := require.New(t)

	err := VisitPayloads(
		context.Background(),
		workflowservice.StartWorkflowExecutionRequest_builder{
			Input: inputPayloads(),
		}.Build(),
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
		workflowservice.StartWorkflowExecutionRequest_builder{
			Header: common.Header_builder{
				Fields: map[string]*common.Payload{"test": inputPayload()},
			}.Build(),
		}.Build(),
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
		export.WorkflowExecutions_builder{Items: []*export.WorkflowExecution{export.WorkflowExecution_builder{History: history.History_builder{
			Events: []*history.HistoryEvent{
				history.HistoryEvent_builder{
					WorkflowExecutionStartedEventAttributes: history.WorkflowExecutionStartedEventAttributes_builder{
						Input: inputPayloads(),
					}.Build(),
				}.Build(),
			},
		}.Build()}.Build()}}.Build(),
		VisitPayloadsOptions{
			Visitor: func(vpc *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.False(vpc.SinglePayloadRequired)
				require.Equal([]byte("test"), p[0].GetData())
				return p, nil
			},
		},
	)
	require.NoError(err)

	msg := history.HistoryEvent_builder{
		NexusOperationScheduledEventAttributes: history.NexusOperationScheduledEventAttributes_builder{
			Input: inputPayload(),
		}.Build(),
	}.Build()
	err = VisitPayloads(
		context.Background(),
		msg,
		VisitPayloadsOptions{
			Visitor: func(vpc *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				require.True(vpc.SinglePayloadRequired)
				require.Equal([]byte("test"), p[0].GetData())
				return []*common.Payload{common.Payload_builder{Data: []byte("visited")}.Build()}, nil
			},
		},
	)
	require.Equal([]byte("visited"), msg.GetNexusOperationScheduledEventAttributes().GetInput().GetData())
	require.NoError(err)
}

func TestVisitPayloads_NestedParent(t *testing.T) {
	// Due to an invalid approach in the previous visitor, this test used to fail
	root := command.StartChildWorkflowExecutionCommandAttributes_builder{
		Header: common.Header_builder{
			Fields: map[string]*common.Payload{
				"header-key": common.Payload_builder{Data: []byte("header-value")}.Build(),
			},
		}.Build(),
		Input: common.Payloads_builder{
			Payloads: []*common.Payload{common.Payload_builder{Data: []byte("input-value")}.Build()},
		}.Build(),
	}.Build()
	var headerParent, inputParent proto.Message
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			if len(p) == 1 {
				if string(p[0].GetData()) == "header-value" {
					headerParent = proto.Clone(ctx.Parent)
				} else if string(p[0].GetData()) == "input-value" {
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
	root := workflowservice.CountWorkflowExecutionsResponse_AggregationGroup_builder{GroupValues: []*common.Payload{common.Payload_builder{Data: []byte("orig-val")}.Build()}}.Build()

	var count int
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			count += 1
			// Only mutate if the payloads has orig-val
			if len(p) == 1 && string(p[0].GetData()) == "orig-val" {
				return []*common.Payload{common.Payload_builder{Data: []byte("new-val")}.Build()}, nil
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Equal(t, 1, count)
	require.Equal(t, []*common.Payload{common.Payload_builder{Data: []byte("new-val")}.Build()}, root.GetGroupValues())
}

func TestVisitPayloads_Any(t *testing.T) {
	// Due to us not visiting protos inside Any, this test used to fail
	msg1, err := anypb.New(update.Request_builder{Input: update.Input_builder{Args: common.Payloads_builder{
		Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val")}.Build()},
	}.Build()}.Build()}.Build())
	require.NoError(t, err)
	msg2, err := anypb.New(update.Request_builder{Input: update.Input_builder{Args: common.Payloads_builder{
		Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val-don't-touch")}.Build()},
	}.Build()}.Build()}.Build())
	require.NoError(t, err)
	msg3, err := anypb.New(update.Response_builder{Outcome: update.Outcome_builder{Success: common.Payloads_builder{
		Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val")}.Build()},
	}.Build()}.Build()}.Build())
	require.NoError(t, err)
	root := workflowservice.PollWorkflowTaskQueueResponse_builder{
		Messages: []*protocol.Message{protocol.Message_builder{Body: msg1}.Build(), protocol.Message_builder{Body: msg2}.Build(), protocol.Message_builder{Body: msg3}.Build()},
	}.Build()

	// Visit with any recursion enabled and only change orig-val
	err = VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			// Only mutate if the payloads has orig-val
			if len(p) == 1 && string(p[0].GetData()) == "orig-val" {
				return []*common.Payload{common.Payload_builder{Data: []byte("new-val")}.Build()}, nil
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	update1, err := root.GetMessages()[0].GetBody().UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "new-val", string(update1.(*update.Request).GetInput().GetArgs().GetPayloads()[0].GetData()))
	update2, err := root.GetMessages()[1].GetBody().UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val-don't-touch", string(update2.(*update.Request).GetInput().GetArgs().GetPayloads()[0].GetData()))
	update3, err := root.GetMessages()[2].GetBody().UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "new-val", string(update3.(*update.Response).GetOutcome().GetSuccess().GetPayloads()[0].GetData()))

	// Do the same test but with a do-nothing visitor and confirm unchanged
	msg1, err = anypb.New(update.Request_builder{Input: update.Input_builder{Args: common.Payloads_builder{
		Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val")}.Build()},
	}.Build()}.Build()}.Build())
	require.NoError(t, err)
	msg2, err = anypb.New(update.Request_builder{Input: update.Input_builder{Args: common.Payloads_builder{
		Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val-don't-touch")}.Build()},
	}.Build()}.Build()}.Build())
	require.NoError(t, err)
	msg3, err = anypb.New(update.Response_builder{Outcome: update.Outcome_builder{Success: common.Payloads_builder{
		Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val")}.Build()},
	}.Build()}.Build()}.Build())
	require.NoError(t, err)
	root = workflowservice.PollWorkflowTaskQueueResponse_builder{
		Messages: []*protocol.Message{protocol.Message_builder{Body: msg1}.Build(), protocol.Message_builder{Body: msg2}.Build(), protocol.Message_builder{Body: msg3}.Build()},
	}.Build()
	err = VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			// Only mutate if the payloads has orig-val
			if len(p) == 1 && string(p[0].GetData()) == "orig-val" {
				return []*common.Payload{common.Payload_builder{Data: []byte("new-val")}.Build()}, nil
			}
			return p, nil
		},
		WellKnownAnyVisitor: func(*VisitPayloadsContext, *anypb.Any) error { return nil },
	})
	require.NoError(t, err)
	update1, err = root.GetMessages()[0].GetBody().UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val", string(update1.(*update.Request).GetInput().GetArgs().GetPayloads()[0].GetData()))
	update2, err = root.GetMessages()[1].GetBody().UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val-don't-touch", string(update2.(*update.Request).GetInput().GetArgs().GetPayloads()[0].GetData()))
	update3, err = root.GetMessages()[2].GetBody().UnmarshalNew()
	require.NoError(t, err)
	require.Equal(t, "orig-val", string(update3.(*update.Response).GetOutcome().GetSuccess().GetPayloads()[0].GetData()))
}

func TestVisitPayloads_RepeatedAny(t *testing.T) {
	msg, err := anypb.New(update.Request_builder{Input: update.Input_builder{Args: common.Payloads_builder{
		Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val")}.Build()},
	}.Build()}.Build()}.Build())
	require.NoError(t, err)
	root := errordetails.MultiOperationExecutionFailure_OperationStatus_builder{Details: []*anypb.Any{msg}}.Build()
	var anyCount int
	err = VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			anyCount++
			// Only mutate if the payloads has "test"
			if len(p) == 1 && string(p[0].GetData()) == "orig-val" {
				return []*common.Payload{common.Payload_builder{Data: []byte("new-val")}.Build()}, nil
			}
			return p, nil
		},
	})
	require.NoError(t, err)
	require.Equal(t, 1, anyCount)
	update1, err := root.GetDetails()[0].UnmarshalNew()

	require.NoError(t, err)
	require.Equal(t, "new-val", string(update1.(*update.Request).GetInput().GetArgs().GetPayloads()[0].GetData()))
}

func TestVisitFailures(t *testing.T) {
	require := require.New(t)

	fail := &failure.Failure{}

	err := VisitFailures(
		context.Background(),
		workflowservice.RespondActivityTaskFailedRequest_builder{
			Failure: fail,
		}.Build(),
		VisitFailuresOptions{
			Visitor: func(vfc *VisitFailuresContext, f *failure.Failure) error {
				require.Equal(fail, f)
				return nil
			},
		},
	)
	require.NoError(err)

	nestedFailure := failure.Failure_builder{Cause: fail}.Build()
	failureCount := 0

	err = VisitFailures(
		context.Background(),
		workflowservice.RespondActivityTaskFailedRequest_builder{
			Failure: nestedFailure,
		}.Build(),
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

	fail := failure.Failure_builder{
		Message: "test failure",
	}.Build()

	msg, err := anypb.New(update.Response_builder{Outcome: update.Outcome_builder{Failure: proto.ValueOrDefault(fail)}.Build()}.Build())
	require.NoError(err)

	req := workflowservice.RespondWorkflowTaskCompletedRequest_builder{
		Messages: []*protocol.Message{protocol.Message_builder{Body: msg}.Build()},
	}.Build()
	failureCount := 0
	err = VisitFailures(
		context.Background(),
		req,
		VisitFailuresOptions{
			Visitor: func(vfc *VisitFailuresContext, f *failure.Failure) error {
				failureCount += 1
				require.Equal("test failure", f.GetMessage())
				f.SetEncodedAttributes(common.Payload_builder{Data: []byte("test failure")}.Build())
				f.SetMessage("encoded failure")
				return nil
			},
		},
	)
	require.NoError(err)
	require.Equal(1, failureCount)
	updateMsg, err := req.GetMessages()[0].GetBody().UnmarshalNew()
	require.NoError(err)
	require.Equal("encoded failure", updateMsg.(*update.Response).GetOutcome().GetFailure().GetMessage())
	require.Equal("test failure", string(updateMsg.(*update.Response).GetOutcome().GetFailure().GetEncodedAttributes().GetData()))

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
		workflowservice.StartWorkflowExecutionRequest_builder{
			Input: inputPayloads(),
		}.Build(),
	)
	require.NoError(err)

	require.True(proto.Equal(inputs.GetPayloads()[0], outboundPayload))

	_, err = client.PollActivityTaskQueue(
		context.Background(),
		&workflowservice.PollActivityTaskQueueRequest{},
	)
	require.NoError(err)

	require.True(proto.Equal(inputs.GetPayloads()[0], inboundPayload))
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
					payloads[0] = common.Payload_builder{Data: []byte("new-val")}.Build()
					return payloads, nil
				},
			},
		},
	)
	require.NoError(err)

	failureInterceptor, err := NewFailureVisitorInterceptor(
		FailureVisitorInterceptorOptions{
			Inbound: &VisitFailuresOptions{Visitor: func(vpc *VisitFailuresContext, f *failure.Failure) error {
				inboundFailure = f.GetMessage()
				f.SetMessage(failureMessage)
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
		workflowservice.StartWorkflowExecutionRequest_builder{
			Input: inputPayloads(),
		}.Build(),
	)
	require.NoError(err)

	_, err = client.ExecuteMultiOperation(context.Background(), &workflowservice.ExecuteMultiOperationRequest{})
	// We expect that even though an error is returned, the Payload visitor visited the payload
	// included in the GRPC error details
	require.Error(err)
	require.True(proto.Equal(inputs.GetPayloads()[0], inboundPayload))
	stat, ok := status.FromError(err)
	require.True(ok)
	for _, detail := range stat.Details() {
		multiOpFailure, ok := detail.(*errordetails.MultiOperationExecutionFailure)
		require.True(ok)
		payloads := &common.Payloads{}
		err = multiOpFailure.GetStatuses()[0].GetDetails()[0].UnmarshalTo(payloads)
		require.NoError(err)

		newPayload := common.Payload_builder{Data: []byte("new-val")}.Build()
		require.True(proto.Equal(payloads.GetPayloads()[0], newPayload))
	}

	_, err = client.QueryWorkflow(context.Background(), &workflowservice.QueryWorkflowRequest{})
	require.Error(err)
	require.Equal("test failure", inboundFailure)
	stat, ok = status.FromError(err)
	require.True(ok)
	for _, detail := range stat.Details() {
		queryFailure, ok := detail.(*errordetails.QueryFailedFailure)
		require.True(ok)
		require.Equal(failureMessage, queryFailure.GetFailure().GetMessage())
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
	return workflowservice.PollActivityTaskQueueResponse_builder{
		Input: inputPayloads(),
	}.Build(), nil
}

func (t *testGRPCServer) ExecuteMultiOperation(
	ctx context.Context,
	req *workflowservice.ExecuteMultiOperationRequest) (*workflowservice.ExecuteMultiOperationResponse, error) {
	anyDetail, err := anypb.New(inputPayloads())
	if err != nil {
		return nil, err
	}
	operationStatus := errordetails.MultiOperationExecutionFailure_OperationStatus_builder{Details: []*anypb.Any{anyDetail}}.Build()
	multiOpFailure := errordetails.MultiOperationExecutionFailure_builder{
		Statuses: []*errordetails.MultiOperationExecutionFailure_OperationStatus{operationStatus},
	}.Build()
	st := status.New(codes.Internal, "Operation failed due to a user error")

	stWithDetails, err := st.WithDetails(multiOpFailure)
	if err != nil {
		return nil, st.Err()
	}

	return nil, stWithDetails.Err()
}

func (t *testGRPCServer) QueryWorkflow(
	ctx context.Context,
	req *workflowservice.QueryWorkflowRequest) (*workflowservice.QueryWorkflowResponse, error) {
	failureMessage := failure.Failure_builder{
		Message: "test failure",
	}.Build()
	queryFailure := errordetails.QueryFailedFailure_builder{Failure: failureMessage}.Build()

	st := status.New(codes.Internal, "Operation failed due to a user error")

	stWithDetails, err := st.WithDetails(queryFailure)
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
			newAny, err := anypb.New(update.Request_builder{Input: update.Input_builder{Args: common.Payloads_builder{
				Payloads: []*common.Payload{common.Payload_builder{Data: []byte("orig-val")}.Build()},
			}.Build()}.Build()}.Build())
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
