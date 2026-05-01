package workflowservice

import (
	"context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// ExampleWorkflowServiceClient is the client for experimental methods added by the Example feature.
// It calls the stable WorkflowService gRPC endpoint, so it works against any server
// that has the Example variant enabled via frontend.apiVariant.
type ExampleWorkflowServiceClient interface {
	Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error)
}

const WorkflowService_Echo_FullMethodName = "/temporal.api.workflowservice.v1.WorkflowService/Echo"

type exampleWorkflowServiceClient struct{ cc grpc.ClientConnInterface }

func NewExampleWorkflowServiceClient(cc grpc.ClientConnInterface) ExampleWorkflowServiceClient {
	return &exampleWorkflowServiceClient{cc}
}

func (c *exampleWorkflowServiceClient) Echo(ctx context.Context, in *EchoRequest, opts ...grpc.CallOption) (*EchoResponse, error) {
	out := new(EchoResponse)
	if err := c.cc.Invoke(ctx, WorkflowService_Echo_FullMethodName, in, out, opts...); err != nil {
		return nil, err
	}
	return out, nil
}

// ExampleWorkflowServiceServer must be implemented by the experimental handler registered via the variant registry.
type ExampleWorkflowServiceServer interface {
	Echo(context.Context, *EchoRequest) (*EchoResponse, error)
}

// UnimplementedExampleWorkflowServiceServer returns codes.Unimplemented for every method.
type UnimplementedExampleWorkflowServiceServer struct{}

func (UnimplementedExampleWorkflowServiceServer) Echo(context.Context, *EchoRequest) (*EchoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Echo not implemented")
}

func _ExampleWorkflowService_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EchoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExampleWorkflowServiceServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{Server: srv, FullMethod: WorkflowService_Echo_FullMethodName}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExampleWorkflowServiceServer).Echo(ctx, req.(*EchoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExampleWorkflowService_ServiceDesc is the overlay descriptor: its ServiceName matches the
// stable WorkflowService so registerServiceOverlay routes calls correctly.
var ExampleWorkflowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.api.workflowservice.v1.WorkflowService",
	HandlerType: (*ExampleWorkflowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{MethodName: "Echo", Handler: _ExampleWorkflowService_Echo_Handler},
	},
}
