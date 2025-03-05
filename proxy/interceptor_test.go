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
	"log"
	"net"
	"reflect"
	"strings"
	"testing"
	"time"

	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"

	"github.com/stretchr/testify/require"
	"go.temporal.io/api/command/v1"
	"go.temporal.io/api/common/v1"
	"go.temporal.io/api/export/v1"
	"go.temporal.io/api/failure/v1"
	"go.temporal.io/api/history/v1"
	"go.temporal.io/api/protocol/v1"
	"go.temporal.io/api/update/v1"
	"go.temporal.io/api/workflowservice/v1"

	// TODO: Find way to ensure all go.temporal.io/api packages are imported
	//   for https://pkg.go.dev/google.golang.org/protobuf/reflect/protoregistry#GlobalTypes
	// Chatgpt generated
	//_ "go.temporal.io/api/common/v1"
	//_ "go.temporal.io/api/enums/v1"
	//_ "go.temporal.io/api/failure/v1"
	//_ "go.temporal.io/api/filter/v1"
	//_ "go.temporal.io/api/history/v1"
	//_ "go.temporal.io/api/namespace/v1"
	//_ "go.temporal.io/api/query/v1"
	//_ "go.temporal.io/api/schedule/v1"
	//_ "go.temporal.io/api/taskqueue/v1"
	//_ "go.temporal.io/api/version/v1"
	//_ "go.temporal.io/api/workflow/v1"
	//_ "go.temporal.io/api/workflowservice/v1"
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
				fmt.Println("hi")
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

// **Recursively populate Protobuf message fields, including oneof handling**
func populatePayload(root *proto.Message, msg proto.Message, require *require.Assertions, totalCount *int, count *int) {
	m := msg.ProtoReflect() // Get protoreflect message
	fields := m.Descriptor().Fields()
	fmt.Println("\n[msg]", m.Descriptor().FullName(), m.Descriptor(), fields.Len())
	fmt.Println("[root]", (*root).ProtoReflect().Descriptor())

	// Don't need to parse non-temporal types
	if !strings.HasPrefix(string(m.Descriptor().FullName()), "temporal.api.") && string(m.Descriptor().FullName()) != "google.protobuf.Any" {
		fmt.Println("EXITING RECURSION-")
		return
	}

	if m.Descriptor() == nil {
		panic("fail")
	}

	switch i := msg.(type) {
	case *common.Payload, *common.Payloads:
		*count++
		*totalCount++
		fmt.Print("[Payload(s)] - ")

		err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
			Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				fmt.Print("FOUND")
				require.Equal(1, *count)
				*count--
				return p, nil
			},
		})
		fmt.Println()
		require.NoError(err)
		return
	case *anypb.Any:
		fmt.Println("ANYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYYY")
		if i.TypeUrl == "" {
			fmt.Println("⚠️ Any field is empty. Assigning a default value...")

			// **Create a sample Payloads message**
			payloads := &common.Payloads{
				Payloads: []*common.Payload{
					{Data: []byte("dummy_data")},
				},
			}

			// **Wrap Payloads inside Any**
			newAny, err := anypb.New(payloads)
			if err != nil {
				fmt.Println("❌ Failed to wrap Payloads into Any:", err)
				return
			}

			// **Set the new Any value**
			*i = *newAny
			fmt.Println("✅ Successfully set Any with Payloads")
		}
		*count++
		*totalCount++
		fmt.Print("[Payload(s)] - ")

		err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
			Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
				fmt.Println("FOUND")
				require.Equal(1, *count)
				*count--
				return p, nil
			},
		})
		require.NoError(err)
		return
	}

	fmt.Println("\tfields", fields)
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		value := m.Get(fd)
		fmt.Printf("[%s]\n", fd.Name())
		fmt.Printf("\t%s, %s, %t, %t\n", fd.Name(), fd.Kind(), fd.IsList(), fd.IsMap())
		fmt.Printf("\t%v\n", value.Interface())

		if oneof := fd.ContainingOneof(); oneof != nil && fd.Kind() == protoreflect.MessageKind {
			fmt.Println("\toneof")
			//fmt.Printf("Found oneof field: %s (Group: %s)\n", fd.Name(), oneof.Name())
			newMsg := value.Message().New()
			m.Set(fd, protoreflect.ValueOf(newMsg))
			//fmt.Println("RECURSING into (oneof)", fd.Name(), newMsg, newMsg.Interface())
			fmt.Println("RECURSING oneof")
			populatePayload(root, newMsg.Interface(), require, totalCount, count)
			// This ensures only 1 payload is set and discoverable from root at a time.
			m.Clear(fd)
		} else if fd.Kind() == protoreflect.MessageKind && fd.IsMap() {
			fmt.Println("\tMap")
			mapVal := m.Mutable(fd).Map()
			require.Equal(0, mapVal.Len())
			// TODO: should only have to handle string and maybe int32 for keys,
			if fd.MapKey().Kind() == protoreflect.StringKind &&
				fd.MapValue().Kind() == protoreflect.MessageKind &&
				string(fd.MapValue().Message().FullName()) == "temporal.api.common.v1.Payload" {
				//fmt.Println("✅ Found a map<string, Payload>! len", mapVal.Len(), "IsValid()", mapVal.IsValid())
				sampleKey := protoreflect.ValueOf("sample_key").MapKey()
				mapVal.Set(sampleKey, protoreflect.ValueOf(inputPayload().ProtoReflect()))
				mapVal.Range(func(key protoreflect.MapKey, val protoreflect.Value) bool {
					//fmt.Printf("  Map Key: %v\n", key)

					// **Handle map of messages**
					if fd.MapValue().Kind() == protoreflect.MessageKind {
						// TODO: Why are we testing with newMsg that isn't being set to mapVal?
						newMsg := val.Message().New()
						fmt.Println("RECURSING map2")
						populatePayload(root, newMsg.Interface(), require, totalCount, count)
					}
					return true
				})
				mapVal.Clear(sampleKey)
				//}
			} else if fd.MapValue().Kind() == protoreflect.MessageKind {
				fmt.Println("mapkind123")
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
				fmt.Println("[sampleKey]", sampleKey, sampleKey.Interface())
				fmt.Println("[mapVal]", reflect.TypeOf(mapVal.NewValue()))
				// TODO: This is not setting the value properly?
				mapVal.Set(sampleKey, mapVal.NewValue())
				fmt.Println("map len", mapVal.Len())
				mapVal.Range(func(key protoreflect.MapKey, val protoreflect.Value) bool {
					if fd.MapValue().Kind() == protoreflect.MessageKind {
						newMsg := val.Message()
						//fmt.Println("RECURSING into map value", fd.Name(), key, newMsg, newMsg.Interface())
						fmt.Println("RECURSING map1")
						populatePayload(root, newMsg.Interface(), require, totalCount, count)
					}
					return true
				})
				//populatePayload(root, newMsg.Interface(), require, totalCount, count)
				mapVal.Clear(sampleKey)
			}
		} else if fd.Kind() == protoreflect.MessageKind && fd.IsList() {
			listVal := m.Mutable(fd).List()
			// We expect this list to be empty
			require.Equal(0, listVal.Len())

			// Create a sample val
			sampleVal := listVal.NewElement()
			listVal.Append(sampleVal)

			val := listVal.Get(0)

			require.True(val.Message().IsValid())

			newMsg := val.Message()
			fmt.Println("RECURSING list")
			populatePayload(root, newMsg.Interface(), require, totalCount, count)
			listVal.Truncate(0)
		} else if fd.Kind() == protoreflect.MessageKind && !fd.IsList() && !fd.IsMap() {
			fmt.Println("\tMessageKind")
			// Avoid cycles
			if value.Message().Descriptor().FullName() == m.Descriptor().FullName() {
				fmt.Println("Avoiding cycles for", fd.Name(), m.Descriptor().FullName())
				continue
			}

			var newMsg protoreflect.Message
			// If field is not set, create a new message
			newMsg = value.Message().New()
			m.Set(fd, protoreflect.ValueOf(newMsg))
			fmt.Println("RECURSING")
			populatePayload(root, newMsg.Interface(), require, totalCount, count)
			//}

			// This ensures only 1 payload is set and discoverable from root at a time.
			m.Clear(fd)
		}
	}

	// Validate that all Payloads were found
	//require.Equal(0, *count)

	fmt.Println("EXITING RECURSION")
}

func TestFailureCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.failure.v1.Failure") { // should have 5
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

func TestUpdateRejectionCount(t *testing.T) {
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
	require.False(true)
}

func TestPayloadsCount(t *testing.T) {
	require := require.New(t)

	// answer - Payloads
	// failure - 5
	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.query.v1.WorkflowQueryResult") { // should have 5
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

func TestAnyCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.protocol.v1.Message") {
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

func TestMapCount(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	var totalCount, count int

	// 	repeated temporal.api.command.v1.Command commands = 2; - 37
	// 	map<string, temporal.api.query.v1.WorkflowQueryResult> query_results - 6
	//		answer - 1
	//		failure - 5
	// 	repeated temporal.api.protocol.v1.Message messages - 1
	//  	google.protobuf.Any body - 1
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

func TestCommandCount(t *testing.T) {
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
	//require.True(false)
}

func TestEverything(t *testing.T) {
	require := require.New(t)

	var messageType []protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.") && string(mt.Descriptor().FullName()) != "temporal.api.common.v1.Payload" { // 2 from request, 7 total
			messageType = append(messageType, mt)
		}
		return true
	})
	fmt.Println("messageType", messageType)
	for _, mt := range messageType {
		// Create empty instance and populate with test values
		msg := mt.New().Interface().(proto.Message)

		var totalCount, count int
		populatePayload(&msg, msg, require, &totalCount, &count)

		require.Equal(0, count)

	}
	//require.True(false)
}

func TestResponseCount(t *testing.T) {
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

func TestSandbox(t *testing.T) {
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		// sdk.UserMetadata works, gets 2
		//
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
	require.Equal(36, totalCount)
	require.True(false)
}
