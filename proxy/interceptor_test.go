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

// getOneofOptions returns all possible field names for a oneof field
//func getOneofOptions(oneof protoreflect.OneofDescriptor) []protoreflect.FieldDescriptor {
//	var options []protoreflect.FieldDescriptor
//	for i := 0; i < oneof.Fields().Len(); i++ {
//		options = append(options, oneof.Fields().Get(i))
//	}
//	return options
//}

// TODO:
//📌 Expected Output
//
//If the Event message has a oneof field with ActivityScheduledEvent and TimerFiredEvent, this program will:
//
//Identify the oneof options.
//Set each oneof field separately.
//Ensure only one is active per test case.
//
//Example console output:
//
//Testing activity_scheduled
//Set event_type to activity_scheduled
//event_type:{activity_scheduled:{}}
//
//Testing timer_fired
//Set event_type to timer_fired
//event_type:{timer_fired:{}}

// testOneofValue populates a oneof field in a message and prints the result
//func testOneofValue(msg proto.Message, oneofField protoreflect.FieldDescriptor) {
//	v := msg.ProtoReflect()
//	oneofName := oneofField.ContainingOneof().Name()
//
//	// Create a new instance of the oneof field's type
//	fieldType := oneofField.Message()
//	if fieldType == nil {
//		fmt.Printf("Skipping %s: not a message type\n", oneofField.FullName())
//		return
//	}
//
//	fieldValue := fieldType.New()
//	v.Set(oneofField, protoreflect.ValueOfMessage(fieldValue))
//
//	fmt.Printf("Set %s to %s\n", oneofName, oneofField.Name())
//	fmt.Println(msg)
//}

// Track visited payloads
// We then use a reflection-based visitor to track which payloads were actually visited.
func trackVisitedPayloads(msg proto.Message, visited map[string]bool) {
	v := reflect.ValueOf(msg).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := v.Type().Field(i).Name

		if field.Kind() == reflect.Ptr && field.Type().Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
			visited[fieldName] = true
			trackVisitedPayloads(field.Interface().(proto.Message), visited)
		}
	}
}

// TODO: anything above this is curerntly not being used

// Track visited payloads
type visitedTracker struct {
	visited map[string]bool
}

func newVisitedTracker() *visitedTracker {
	return &visitedTracker{visited: make(map[string]bool)}
}

func (vt *visitedTracker) markVisited(fieldPath string) {
	vt.visited[fieldPath] = true
}

// Get all top-level messages from registered Protobuf descriptors
// This will return all registered message types, from which we can
// identify top-level messages like workflow activation and history event messages.
// getTopLevelMessages gathers all registered message types, filtering for temporal specific types.
func getTopLevelMessages() []proto.Message {
	var topMessages []proto.Message
	fmt.Println("getTopLevelMessages(): ")
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.failure.v1.ResetWorkflowFailureInfo") {
			//fmt.Println("\t", string(mt.Descriptor().FullName()), " ")
			fmt.Print(string(mt.Descriptor().FullName()))
			topMessages = append(topMessages, mt.New().Interface().(proto.Message))
		}
		return true
	})
	fmt.Println()
	return topMessages
}

//func createStructFromPath(path string) (proto.Message, error) {
//	parts := strings.Split(path, ".")
//
//	// Get the root struct type
//	var instance reflect.Value
//}

// findPayloadPaths recursively traverse fields and finds all payload paths
// TODO: memoize
func findPayloadPaths(root *proto.Message, msg proto.Message, parentPath string, payloadPaths *[]string, levels int) error { // memo map[reflect.Value]string
	if levels == 3 {
		return nil
	}
	value := reflect.ValueOf(msg).Elem()
	typ := value.Type()
	// TODO: use memoization to break the cycle early

	fmt.Println(levels, "\n\t[value]", value, "\n\t[type]", typ)
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)                  // value, AKA dummy value
		fieldType := typ.Field(i)                // {WorkflowTaskTimeout  *durationpb.Duration protobuf:"bytes,6,opt,name=workflow_task_timeout,json=workflowTaskTimeout,proto3" json:"workflow_task_timeout,omitempty" 56 [6] false}
		fieldName := fieldType.Name              // WorkflowTaskTimeout
		fullPath := parentPath + "." + fieldName // WorkflowExecutionContinuedAsNewEventAttributes.WorkflowTaskTimeout
		fmt.Println("\n\tfieldType:", fieldType)
		fmt.Println("\tfieldName:", fieldName)
		fmt.Println("\tfullPath:", fullPath)

		// If this is a Payload field, track it
		if field.Type().AssignableTo(reflect.TypeOf(&common.Payload{})) {
			fmt.Println("\n[Field]")
			if field.CanInterface() && field.Addr().Type().Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
				protoMsg := field.Addr().Interface().(proto.Message)
				fmt.Printf("Field %s is a proto.Message: %T\n", fieldType.Name, protoMsg)
			}
			err := VisitPayloads(context.Background(), field.Addr().Interface().(proto.Message), VisitPayloadsOptions{
				Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("ASDFASDFASDFASDFASD")
					return p, nil
				},
			})
			if err != nil {
				return err
			}
			*payloadPaths = append(*payloadPaths, fullPath)
			//memo[Type] = fullPath
		} else if field.Type().AssignableTo(reflect.TypeOf(&common.Payloads{})) {
			fmt.Println("[Payloads]")
			*payloadPaths = append(*payloadPaths, fullPath)
			//asdf := protoreflect.ValueOfString("WorkflowExecutionContinuedAsNewEventAttributes")
			err := VisitPayloads(context.Background(), field.Addr().Interface().(proto.Message), VisitPayloadsOptions{
				Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("ASDFASDFASDFASDFASD")
					return p, nil
				},
			})
			if err != nil {
				return err
			}
			//asdf := failure.ResetWorkflowFailureInfo{LastHeartbeatDetails: inputPayloads()}
		} else if field.Kind() == reflect.Ptr && field.Type().Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
			// Recursively check nested messages
			if field.IsNil() {
				field.Set(reflect.New(field.Type().Elem())) // Populate if nil
			}
			fmt.Println("\nRECURSING!!")
			//findPayloadPaths(field.Interface().(proto.Message), fullPath, payloadPaths, levels+1)
		}
	}
	return nil
}

// Finally, we assert that all expected payload paths were visited.
// Full test function
func validatePayloadTraversal() error {
	topMessages := getTopLevelMessages()
	//vt := newVisitedTracker()
	counter := 0

	for _, msg := range topMessages {
		if counter > 3 {
			break
		}
		fmt.Println("[", msg.ProtoReflect().Descriptor().FullName(), "]")

		// Find payload paths
		var payloadPaths []string
		findPayloadPaths(&msg, msg, reflect.TypeOf(msg).Elem().Name(), &payloadPaths, 0)
		fmt.Println("[payloadPaths]", payloadPaths)
		counter += 1

		//for _, path := range paths {
		//	// populate each path
		//}

		// run it over visitor to confirm visited TODO: not sure if this is a separate check or inside of the paths for loop

		//msg := messageType.New().Interface().(proto.Message) // Create empty instance
		//populatePayload(msg)                                 // Populate it with test values
		//
		//visited := make(map[string]bool)
		//trackVisitedPayloads(msg, visited) // Run visitor
		//
		//fmt.Printf("Testing %s - Visited payloads: %v\n", messageType.Descriptor().FullName(), visited)
		//
		//// Assert that at least some payloads were visited
		//if len(visited) == 0 {
		//	log.Fatalf("No payloads visited for %s", messageType.Descriptor().FullName())
		//}
		//fmt.Println()
	}

	var count int
	// TODO: change root
	root := &workflowservice.PollWorkflowTaskQueueResponse{
		Messages: []*protocol.Message{},
	}
	// Visit payloads, and add each count
	err := VisitPayloads(context.Background(), root, VisitPayloadsOptions{
		Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
			count += 1
			return p, nil
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func TestPlayground(t *testing.T) {
	require := require.New(t)
	require.True(true)
	err := validatePayloadTraversal()
	require.NoError(err)
	require.True(false)
}

// **Retrieve all possible options for a oneof field**
func getOneofOptions(md reflect.Type, oneofGroup string) []reflect.Type {
	var options []reflect.Type

	// Iterate over struct fields and find all belonging to this oneof group
	for i := 0; i < md.NumField(); i++ {
		field := md.Field(i)
		if field.Tag.Get("protobuf_oneof") == oneofGroup {
			options = append(options, field.Type)
		}
	}
	return options
}

// **Reset all fields in a oneof group (ensure only one is set)**
func resetOneofFields(v reflect.Value, oneofGroup string) {
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Tag.Get("protobuf_oneof") == oneofGroup {
			v.Field(i).Set(reflect.Zero(field.Type)) // Clear the field
		}
	}
}

// **Set a test value for a oneof field and print the result**
func testOneofValue(v reflect.Value, oneofGroup string, fieldType reflect.Type) {
	fmt.Printf("Testing oneof option: %s\n", fieldType.Name())

	// Create a new instance of the oneof type
	oneofValue := reflect.New(fieldType).Elem()

	// Set a sample value inside the oneof (e.g., set a payload field if applicable)
	if fieldType == reflect.TypeOf(&failure.ResetWorkflowFailureInfo{}) {
		oneofValue.FieldByName("LastHeartbeatDetails").Set(reflect.ValueOf(inputPayloads()))
	}

	// Set the value in the struct
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		if field.Tag.Get("protobuf_oneof") == oneofGroup {
			v.Field(i).Set(oneofValue)
			break
		}
	}

	fmt.Printf("Updated message: %+v\n", v.Interface())
}

// Populate message fields dynamically, using Go reflection.
// We now populate test values in message fields that could contain payloads.
func populatePayload(root *proto.Message, msg proto.Message, require *require.Assertions) {
	v := reflect.ValueOf(msg).Elem() // Get underlying struct
	fmt.Println("[v]", msg.ProtoReflect().Descriptor().FullName())
	fmt.Println("[v.type]", v.Type())
	var count int
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fmt.Println("\n[", v.Type().Field(i).Name, "]", field.Type())

		switch field.Kind() {
		case reflect.Ptr:
			if field.Type().Implements(reflect.TypeOf((*proto.Message)(nil)).Elem()) {
				newMsg := reflect.New(field.Type().Elem()) // Create new message instance
				fmt.Println("[newMsg.Elem()]", newMsg.Elem().Type())
				// can't print nil pointer
				if v.Type().AssignableTo(newMsg.Type()) {
					fmt.Println("Avoiding cycles for", newMsg.Type(), v.Type(), newMsg.Type().AssignableTo(v.Type()))
					continue
				}

				if newMsg.Type().AssignableTo(v.Type()) {
					fmt.Println("OTHER Avoiding cycles for", newMsg.Type(), v.Type(), newMsg.Type().AssignableTo(v.Type()))
					continue
				}
				// Avoid cycles
				if newMsg.Elem().Type() == v.Type() {
					fmt.Println("newMsg.Elem().Type() == v.Type()")
					continue
				}
				field.Set(newMsg)
				fmt.Println("RECURSING")
				populatePayload(root, newMsg.Interface().(proto.Message), require) // Recursively populate

				//if v.Type().Field(i).Tag.Get("protobuf_oneof") != "" {
				//	// Handle oneof by setting each possible value separately
				//	options := getOneofOptions(field.Type())
				//	for _, opt := range options {
				//		testOneofValue(msg, fieldName, opt)
				//	}
				//} else {
				//	populatePayload(field.Interface().(proto.Message))
				//}
			}
		}

		// Handle oneof by setting each possible value separately TODO: move to different part of function?
		if v.Type().Field(i).Tag.Get("protobuf_oneof") != "" {
			fmt.Printf("Found oneof field: %s (Group: %s)\n", v.Type().Field(i).Name, v.Type().Field(i).Tag.Get("protobuf_oneof"))
			oneofOptions := getOneofOptions(v.Type(), v.Type().Field(i).Tag.Get("protobuf_oneof"))
			fmt.Println("[oneofOptions]", oneofOptions)
		}
		// NOTE: must set Payloads value before calling this
		if field.Type().AssignableTo(reflect.TypeOf(&common.Payload{})) {
			count += 1
			fmt.Println("[Payload]")
			err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
				Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("\t\tFOUND PAYLOAD")
					count -= 1
					require.True(ctx.SinglePayloadRequired)
					return p, nil
				},
			})
			if err != nil {
				panic(err)
			}
		}
		if field.Type().AssignableTo(reflect.TypeOf(&common.Payloads{})) {
			count += 1
			fmt.Println("[Payloads]")
			err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
				Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("\t\tFOUND PAYLOADS")
					count -= 1
					require.False(ctx.SinglePayloadRequired)
					return p, nil
				},
			})
			if err != nil {
				panic(err)
			}
		}
	}
	// TODO: Figure out a way to verify that any Payloads have been found?
	if count != 0 {
		panic("fail, count should be 0")
	}
	fmt.Println("EXITING RECURSION")
}

// **Recursively populate Protobuf message fields, including oneof handling**
func populatePayloadProtoreflect(root *proto.Message, msg proto.Message, require *require.Assertions) {
	m := msg.ProtoReflect() // Get protoreflect message
	fmt.Println("\n[msg]", m.Descriptor().FullName())
	fields := m.Descriptor().Fields()
	var count int
	for i := 0; i < fields.Len(); i++ {
		fd := fields.Get(i)
		value := m.Get(fd)
		fmt.Printf("[%s] %s, %s\n", fd.Name(), value, value.Interface())

		// **Handle oneof fields**
		if oneof := fd.ContainingOneof(); oneof != nil {
			fmt.Printf("Found oneof field: %s (Group: %s)\n", fd.Name(), oneof.Name())

			//// Iterate over each possible oneof option
			//for i := 0; i < oneof.Fields().Len(); i++ {
			//	field := oneof.Fields().Get(i)
			//	fmt.Printf("  Trying oneof option: %s\n", field.Name())
			//
			//	// Create a new instance for the oneof field
			//	if field.Kind() == protoreflect.MessageKind {
			//		fmt.Println("TODO:")
			//		//newMsg := field.Message().New()
			//		//m.Set(fd, protoreflect.ValueOf(newMsg))
			//		//
			//		//// Print updated message
			//		//fmt.Println("  Updated message:", m)
			//		//populatePayload(root, newMsg.Interface(), require)
			//	}
			//}
			newMsg := value.Message().New()
			fmt.Println("RECURSING into (oneof)", fd.Name(), newMsg, newMsg.Interface())
			populatePayloadProtoreflect(root, newMsg.Interface(), require)
			// TODO: i don't think continue here is right? or maybe next "if" should be "else if"
			continue
		}

		// **Handle nested messages**
		if fd.Kind() == protoreflect.MessageKind && !fd.IsList() && !fd.IsMap() {
			// Avoid cycles
			if value.Message().Descriptor().FullName() == m.Descriptor().FullName() {
				fmt.Println("Avoiding cycles for", fd.Name())
				continue
			}

			// If field is not set, create a new message
			newMsg := value.Message().New()
			m.Set(fd, protoreflect.ValueOf(newMsg))
			fmt.Println("RECURSING into", fd.Name(), newMsg, newMsg.Interface())
			if !(fd.Message() != nil && fd.Message().FullName() == "temporal.api.common.v1.Payload" || fd.Message().FullName() == "temporal.api.common.v1.Payloads") {
				populatePayloadProtoreflect(root, newMsg.Interface(), require)
			}
		}

		// **Handle Payload fields**
		// TODO: There is an extra payload counted in ApplicationFailureInfo.details
		//  looks like when recursing into Payloads, it's counting an additional Payload??
		if fd.Message() != nil && fd.Message().FullName() == "temporal.api.common.v1.Payload" {
			fmt.Println("[aaa]", fd.Message(), fd.Message().FullName())
			count++
			fmt.Print("[Payload] - ")

			err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
				Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("FOUND")
					count--
					//require.True(ctx.SinglePayloadRequired)
					return p, nil
				},
			})
			if err != nil {
				panic(err)
			}
		}

		if fd.Message() != nil && fd.Message().FullName() == "temporal.api.common.v1.Payloads" {
			count++
			fmt.Print("[Payloads] - ")

			// TODO: How does this handle multiple payload/Payloads pairs
			err := VisitPayloads(context.Background(), *root, VisitPayloadsOptions{
				Visitor: func(ctx *VisitPayloadsContext, p []*common.Payload) ([]*common.Payload, error) {
					fmt.Println("FOUND")
					count--
					//require.False(ctx.SinglePayloadRequired)
					return p, nil
				},
			})
			if err != nil {
				panic(err)
			}
		}
	}

	// Validate that all Payloads were found
	if count != 0 {
		panic("Fail: count should be 0")
	}

	fmt.Println("EXITING RECURSION")
}
func TestSandbox(t *testing.T) {
	// TODO:
	//  test oneof - should be good
	//  test recursion
	//  test everything together
	require := require.New(t)

	var messageType protoreflect.MessageType
	protoregistry.GlobalTypes.RangeMessages(func(mt protoreflect.MessageType) bool {
		//if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.failure.v1.ResetWorkflowFailureInfo") {
		if strings.HasPrefix(string(mt.Descriptor().FullName()), "temporal.api.failure.v1.Failure") { // should have 5

			//fmt.Println("\t", string(mt.Descriptor().FullName()), " ")
			//fmt.Print(string(mt.Descriptor().FullName()))
			messageType = mt
		}
		return true
	})

	//fmt.Println("MessageType:", messageType)
	//Fields of  temporal.api.failure.v1.Failure
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.message, Number: 1, Cardinality: optional, Kind: string, HasJSONName: true, JSONName: "message"}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.source, Number: 2, Cardinality: optional, Kind: string, HasJSONName: true, JSONName: "source"}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.stack_trace, Number: 3, Cardinality: optional, Kind: string, HasJSONName: true, JSONName: "stackTrace"}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.encoded_attributes, Number: 20, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "encodedAttributes", HasPresence: true, Message: temporal.api.common.v1.Payload}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.cause, Number: 4, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "cause", HasPresence: true, Message: temporal.api.failure.v1.Failure}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.application_failure_info, Number: 5, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "applicationFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.ApplicationFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.timeout_failure_info, Number: 6, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "timeoutFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.TimeoutFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.canceled_failure_info, Number: 7, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "canceledFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.CanceledFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.terminated_failure_info, Number: 8, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "terminatedFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.TerminatedFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.server_failure_info, Number: 9, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "serverFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.ServerFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.reset_workflow_failure_info, Number: 10, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "resetWorkflowFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.ResetWorkflowFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.activity_failure_info, Number: 11, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "activityFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.ActivityFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.child_workflow_execution_failure_info, Number: 12, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "childWorkflowExecutionFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.ChildWorkflowExecutionFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.nexus_operation_execution_failure_info, Number: 13, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "nexusOperationExecutionFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.NexusOperationFailureInfo}
	//field FieldDescriptor{Syntax: proto3, FullName: temporal.api.failure.v1.Failure.nexus_handler_failure_info, Number: 14, Cardinality: optional, Kind: message, HasJSONName: true, JSONName: "nexusHandlerFailureInfo", HasPresence: true, Oneof: failure_info, Message: temporal.api.failure.v1.NexusHandlerFailureInfo}
	//md := messageType.Descriptor()
	//fmt.Println("Fields of ", md.FullName())
	//fields := md.Fields()
	//for i := 0; i < fields.Len(); i++ {
	//	field := fields.Get(i)
	//	fmt.Println("field", field)
	//}

	// Create empty instance and populate with test values
	msg := messageType.New().Interface().(proto.Message)
	fmt.Println("\nCalling populatePayload", msg.ProtoReflect().Descriptor().FullName())
	populatePayloadProtoreflect(&msg, msg, require)
	fmt.Println("[after populatePayload]", msg.ProtoReflect().Descriptor().FullName())

	// TODO: need to fail to print

	require.True(false)
}
