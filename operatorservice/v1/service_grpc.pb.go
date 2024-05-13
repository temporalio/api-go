// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// plugins:
// - protoc-gen-go-grpc
// - protoc
// source: temporal/api/operatorservice/v1/service.proto

package operatorservice

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	OperatorService_AddSearchAttributes_FullMethodName      = "/temporal.api.operatorservice.v1.OperatorService/AddSearchAttributes"
	OperatorService_RemoveSearchAttributes_FullMethodName   = "/temporal.api.operatorservice.v1.OperatorService/RemoveSearchAttributes"
	OperatorService_ListSearchAttributes_FullMethodName     = "/temporal.api.operatorservice.v1.OperatorService/ListSearchAttributes"
	OperatorService_DeleteNamespace_FullMethodName          = "/temporal.api.operatorservice.v1.OperatorService/DeleteNamespace"
	OperatorService_AddOrUpdateRemoteCluster_FullMethodName = "/temporal.api.operatorservice.v1.OperatorService/AddOrUpdateRemoteCluster"
	OperatorService_RemoveRemoteCluster_FullMethodName      = "/temporal.api.operatorservice.v1.OperatorService/RemoveRemoteCluster"
	OperatorService_ListClusters_FullMethodName             = "/temporal.api.operatorservice.v1.OperatorService/ListClusters"
	OperatorService_GetNexusEndpoint_FullMethodName         = "/temporal.api.operatorservice.v1.OperatorService/GetNexusEndpoint"
	OperatorService_CreateNexusEndpoint_FullMethodName      = "/temporal.api.operatorservice.v1.OperatorService/CreateNexusEndpoint"
	OperatorService_UpdateNexusEndpoint_FullMethodName      = "/temporal.api.operatorservice.v1.OperatorService/UpdateNexusEndpoint"
	OperatorService_DeleteNexusEndpoint_FullMethodName      = "/temporal.api.operatorservice.v1.OperatorService/DeleteNexusEndpoint"
	OperatorService_ListNexusEndpoints_FullMethodName       = "/temporal.api.operatorservice.v1.OperatorService/ListNexusEndpoints"
)

// OperatorServiceClient is the client API for OperatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperatorServiceClient interface {
	// AddSearchAttributes add custom search attributes.
	//
	// Returns ALREADY_EXISTS status code if a Search Attribute with any of the specified names already exists
	// Returns INTERNAL status code with temporal.api.errordetails.v1.SystemWorkflowFailure in Error Details if registration process fails,
	AddSearchAttributes(ctx context.Context, in *AddSearchAttributesRequest, opts ...grpc.CallOption) (*AddSearchAttributesResponse, error)
	// RemoveSearchAttributes removes custom search attributes.
	//
	// Returns NOT_FOUND status code if a Search Attribute with any of the specified names is not registered
	RemoveSearchAttributes(ctx context.Context, in *RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*RemoveSearchAttributesResponse, error)
	// ListSearchAttributes returns comprehensive information about search attributes.
	ListSearchAttributes(ctx context.Context, in *ListSearchAttributesRequest, opts ...grpc.CallOption) (*ListSearchAttributesResponse, error)
	// DeleteNamespace synchronously deletes a namespace and asynchronously reclaims all namespace resources.
	DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error)
	// AddOrUpdateRemoteCluster adds or updates remote cluster.
	AddOrUpdateRemoteCluster(ctx context.Context, in *AddOrUpdateRemoteClusterRequest, opts ...grpc.CallOption) (*AddOrUpdateRemoteClusterResponse, error)
	// RemoveRemoteCluster removes remote cluster.
	RemoveRemoteCluster(ctx context.Context, in *RemoveRemoteClusterRequest, opts ...grpc.CallOption) (*RemoveRemoteClusterResponse, error)
	// ListClusters returns information about Temporal clusters.
	ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error)
	// Get a registered Nexus endpoint by ID. The returned version can be used for optimistic updates.
	GetNexusEndpoint(ctx context.Context, in *GetNexusEndpointRequest, opts ...grpc.CallOption) (*GetNexusEndpointResponse, error)
	// Create a Nexus endpoint. This will fail if an endpoint with the same name is already registered with a status of
	// ALREADY_EXISTS.
	// Returns the created endpoint with its initial version. You may use this version for subsequent updates.
	CreateNexusEndpoint(ctx context.Context, in *CreateNexusEndpointRequest, opts ...grpc.CallOption) (*CreateNexusEndpointResponse, error)
	// Optimistically update a Nexus endpoint based on provided version as obtained via the `GetNexusEndpoint` or
	// `ListNexusEndpointResponse` APIs. This will fail with a status of FAILED_PRECONDITION if the version does not
	// match.
	// Returns the updated endpoint with its updated version. You may use this version for subsequent updates. You don't
	// need to increment the version yourself. The server will increment the version for you after each update.
	UpdateNexusEndpoint(ctx context.Context, in *UpdateNexusEndpointRequest, opts ...grpc.CallOption) (*UpdateNexusEndpointResponse, error)
	// Delete an incoming Nexus service by ID.
	DeleteNexusEndpoint(ctx context.Context, in *DeleteNexusEndpointRequest, opts ...grpc.CallOption) (*DeleteNexusEndpointResponse, error)
	// List all Nexus endpoints for the cluster, sorted by ID in ascending order. Set page_token in the request to the
	// next_page_token field of the previous response to get the next page of results. An empty next_page_token
	// indicates that there are no more results. During pagination, a newly added service with an ID lexicographically
	// earlier than the previous page's last endpoint's ID may be missed.
	ListNexusEndpoints(ctx context.Context, in *ListNexusEndpointsRequest, opts ...grpc.CallOption) (*ListNexusEndpointsResponse, error)
}

type operatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOperatorServiceClient(cc grpc.ClientConnInterface) OperatorServiceClient {
	return &operatorServiceClient{cc}
}

func (c *operatorServiceClient) AddSearchAttributes(ctx context.Context, in *AddSearchAttributesRequest, opts ...grpc.CallOption) (*AddSearchAttributesResponse, error) {
	out := new(AddSearchAttributesResponse)
	err := c.cc.Invoke(ctx, OperatorService_AddSearchAttributes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) RemoveSearchAttributes(ctx context.Context, in *RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*RemoveSearchAttributesResponse, error) {
	out := new(RemoveSearchAttributesResponse)
	err := c.cc.Invoke(ctx, OperatorService_RemoveSearchAttributes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListSearchAttributes(ctx context.Context, in *ListSearchAttributesRequest, opts ...grpc.CallOption) (*ListSearchAttributesResponse, error) {
	out := new(ListSearchAttributesResponse)
	err := c.cc.Invoke(ctx, OperatorService_ListSearchAttributes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DeleteNamespace(ctx context.Context, in *DeleteNamespaceRequest, opts ...grpc.CallOption) (*DeleteNamespaceResponse, error) {
	out := new(DeleteNamespaceResponse)
	err := c.cc.Invoke(ctx, OperatorService_DeleteNamespace_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) AddOrUpdateRemoteCluster(ctx context.Context, in *AddOrUpdateRemoteClusterRequest, opts ...grpc.CallOption) (*AddOrUpdateRemoteClusterResponse, error) {
	out := new(AddOrUpdateRemoteClusterResponse)
	err := c.cc.Invoke(ctx, OperatorService_AddOrUpdateRemoteCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) RemoveRemoteCluster(ctx context.Context, in *RemoveRemoteClusterRequest, opts ...grpc.CallOption) (*RemoveRemoteClusterResponse, error) {
	out := new(RemoveRemoteClusterResponse)
	err := c.cc.Invoke(ctx, OperatorService_RemoveRemoteCluster_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListClusters(ctx context.Context, in *ListClustersRequest, opts ...grpc.CallOption) (*ListClustersResponse, error) {
	out := new(ListClustersResponse)
	err := c.cc.Invoke(ctx, OperatorService_ListClusters_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) GetNexusEndpoint(ctx context.Context, in *GetNexusEndpointRequest, opts ...grpc.CallOption) (*GetNexusEndpointResponse, error) {
	out := new(GetNexusEndpointResponse)
	err := c.cc.Invoke(ctx, OperatorService_GetNexusEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) CreateNexusEndpoint(ctx context.Context, in *CreateNexusEndpointRequest, opts ...grpc.CallOption) (*CreateNexusEndpointResponse, error) {
	out := new(CreateNexusEndpointResponse)
	err := c.cc.Invoke(ctx, OperatorService_CreateNexusEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) UpdateNexusEndpoint(ctx context.Context, in *UpdateNexusEndpointRequest, opts ...grpc.CallOption) (*UpdateNexusEndpointResponse, error) {
	out := new(UpdateNexusEndpointResponse)
	err := c.cc.Invoke(ctx, OperatorService_UpdateNexusEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) DeleteNexusEndpoint(ctx context.Context, in *DeleteNexusEndpointRequest, opts ...grpc.CallOption) (*DeleteNexusEndpointResponse, error) {
	out := new(DeleteNexusEndpointResponse)
	err := c.cc.Invoke(ctx, OperatorService_DeleteNexusEndpoint_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operatorServiceClient) ListNexusEndpoints(ctx context.Context, in *ListNexusEndpointsRequest, opts ...grpc.CallOption) (*ListNexusEndpointsResponse, error) {
	out := new(ListNexusEndpointsResponse)
	err := c.cc.Invoke(ctx, OperatorService_ListNexusEndpoints_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperatorServiceServer is the server API for OperatorService service.
// All implementations must embed UnimplementedOperatorServiceServer
// for forward compatibility
type OperatorServiceServer interface {
	// AddSearchAttributes add custom search attributes.
	//
	// Returns ALREADY_EXISTS status code if a Search Attribute with any of the specified names already exists
	// Returns INTERNAL status code with temporal.api.errordetails.v1.SystemWorkflowFailure in Error Details if registration process fails,
	AddSearchAttributes(context.Context, *AddSearchAttributesRequest) (*AddSearchAttributesResponse, error)
	// RemoveSearchAttributes removes custom search attributes.
	//
	// Returns NOT_FOUND status code if a Search Attribute with any of the specified names is not registered
	RemoveSearchAttributes(context.Context, *RemoveSearchAttributesRequest) (*RemoveSearchAttributesResponse, error)
	// ListSearchAttributes returns comprehensive information about search attributes.
	ListSearchAttributes(context.Context, *ListSearchAttributesRequest) (*ListSearchAttributesResponse, error)
	// DeleteNamespace synchronously deletes a namespace and asynchronously reclaims all namespace resources.
	DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error)
	// AddOrUpdateRemoteCluster adds or updates remote cluster.
	AddOrUpdateRemoteCluster(context.Context, *AddOrUpdateRemoteClusterRequest) (*AddOrUpdateRemoteClusterResponse, error)
	// RemoveRemoteCluster removes remote cluster.
	RemoveRemoteCluster(context.Context, *RemoveRemoteClusterRequest) (*RemoveRemoteClusterResponse, error)
	// ListClusters returns information about Temporal clusters.
	ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error)
	// Get a registered Nexus endpoint by ID. The returned version can be used for optimistic updates.
	GetNexusEndpoint(context.Context, *GetNexusEndpointRequest) (*GetNexusEndpointResponse, error)
	// Create a Nexus endpoint. This will fail if an endpoint with the same name is already registered with a status of
	// ALREADY_EXISTS.
	// Returns the created endpoint with its initial version. You may use this version for subsequent updates.
	CreateNexusEndpoint(context.Context, *CreateNexusEndpointRequest) (*CreateNexusEndpointResponse, error)
	// Optimistically update a Nexus endpoint based on provided version as obtained via the `GetNexusEndpoint` or
	// `ListNexusEndpointResponse` APIs. This will fail with a status of FAILED_PRECONDITION if the version does not
	// match.
	// Returns the updated endpoint with its updated version. You may use this version for subsequent updates. You don't
	// need to increment the version yourself. The server will increment the version for you after each update.
	UpdateNexusEndpoint(context.Context, *UpdateNexusEndpointRequest) (*UpdateNexusEndpointResponse, error)
	// Delete an incoming Nexus service by ID.
	DeleteNexusEndpoint(context.Context, *DeleteNexusEndpointRequest) (*DeleteNexusEndpointResponse, error)
	// List all Nexus endpoints for the cluster, sorted by ID in ascending order. Set page_token in the request to the
	// next_page_token field of the previous response to get the next page of results. An empty next_page_token
	// indicates that there are no more results. During pagination, a newly added service with an ID lexicographically
	// earlier than the previous page's last endpoint's ID may be missed.
	ListNexusEndpoints(context.Context, *ListNexusEndpointsRequest) (*ListNexusEndpointsResponse, error)
	mustEmbedUnimplementedOperatorServiceServer()
}

// UnimplementedOperatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOperatorServiceServer struct {
}

func (UnimplementedOperatorServiceServer) AddSearchAttributes(context.Context, *AddSearchAttributesRequest) (*AddSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSearchAttributes not implemented")
}
func (UnimplementedOperatorServiceServer) RemoveSearchAttributes(context.Context, *RemoveSearchAttributesRequest) (*RemoveSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSearchAttributes not implemented")
}
func (UnimplementedOperatorServiceServer) ListSearchAttributes(context.Context, *ListSearchAttributesRequest) (*ListSearchAttributesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSearchAttributes not implemented")
}
func (UnimplementedOperatorServiceServer) DeleteNamespace(context.Context, *DeleteNamespaceRequest) (*DeleteNamespaceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNamespace not implemented")
}
func (UnimplementedOperatorServiceServer) AddOrUpdateRemoteCluster(context.Context, *AddOrUpdateRemoteClusterRequest) (*AddOrUpdateRemoteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddOrUpdateRemoteCluster not implemented")
}
func (UnimplementedOperatorServiceServer) RemoveRemoteCluster(context.Context, *RemoveRemoteClusterRequest) (*RemoveRemoteClusterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveRemoteCluster not implemented")
}
func (UnimplementedOperatorServiceServer) ListClusters(context.Context, *ListClustersRequest) (*ListClustersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListClusters not implemented")
}
func (UnimplementedOperatorServiceServer) GetNexusEndpoint(context.Context, *GetNexusEndpointRequest) (*GetNexusEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNexusEndpoint not implemented")
}
func (UnimplementedOperatorServiceServer) CreateNexusEndpoint(context.Context, *CreateNexusEndpointRequest) (*CreateNexusEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNexusEndpoint not implemented")
}
func (UnimplementedOperatorServiceServer) UpdateNexusEndpoint(context.Context, *UpdateNexusEndpointRequest) (*UpdateNexusEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateNexusEndpoint not implemented")
}
func (UnimplementedOperatorServiceServer) DeleteNexusEndpoint(context.Context, *DeleteNexusEndpointRequest) (*DeleteNexusEndpointResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteNexusEndpoint not implemented")
}
func (UnimplementedOperatorServiceServer) ListNexusEndpoints(context.Context, *ListNexusEndpointsRequest) (*ListNexusEndpointsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNexusEndpoints not implemented")
}
func (UnimplementedOperatorServiceServer) mustEmbedUnimplementedOperatorServiceServer() {}

// UnsafeOperatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperatorServiceServer will
// result in compilation errors.
type UnsafeOperatorServiceServer interface {
	mustEmbedUnimplementedOperatorServiceServer()
}

func RegisterOperatorServiceServer(s grpc.ServiceRegistrar, srv OperatorServiceServer) {
	s.RegisterService(&OperatorService_ServiceDesc, srv)
}

func _OperatorService_AddSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).AddSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_AddSearchAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).AddSearchAttributes(ctx, req.(*AddSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_RemoveSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).RemoveSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_RemoveSearchAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).RemoveSearchAttributes(ctx, req.(*RemoveSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListSearchAttributes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSearchAttributesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListSearchAttributes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_ListSearchAttributes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListSearchAttributes(ctx, req.(*ListSearchAttributesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DeleteNamespace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNamespaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DeleteNamespace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_DeleteNamespace_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DeleteNamespace(ctx, req.(*DeleteNamespaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_AddOrUpdateRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddOrUpdateRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).AddOrUpdateRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_AddOrUpdateRemoteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).AddOrUpdateRemoteCluster(ctx, req.(*AddOrUpdateRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_RemoveRemoteCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveRemoteClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).RemoveRemoteCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_RemoveRemoteCluster_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).RemoveRemoteCluster(ctx, req.(*RemoveRemoteClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_ListClusters_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListClusters(ctx, req.(*ListClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_GetNexusEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNexusEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).GetNexusEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_GetNexusEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).GetNexusEndpoint(ctx, req.(*GetNexusEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_CreateNexusEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNexusEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).CreateNexusEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_CreateNexusEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).CreateNexusEndpoint(ctx, req.(*CreateNexusEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_UpdateNexusEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNexusEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).UpdateNexusEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_UpdateNexusEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).UpdateNexusEndpoint(ctx, req.(*UpdateNexusEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_DeleteNexusEndpoint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNexusEndpointRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).DeleteNexusEndpoint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_DeleteNexusEndpoint_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).DeleteNexusEndpoint(ctx, req.(*DeleteNexusEndpointRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperatorService_ListNexusEndpoints_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNexusEndpointsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperatorServiceServer).ListNexusEndpoints(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OperatorService_ListNexusEndpoints_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperatorServiceServer).ListNexusEndpoints(ctx, req.(*ListNexusEndpointsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperatorService_ServiceDesc is the grpc.ServiceDesc for OperatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.api.operatorservice.v1.OperatorService",
	HandlerType: (*OperatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddSearchAttributes",
			Handler:    _OperatorService_AddSearchAttributes_Handler,
		},
		{
			MethodName: "RemoveSearchAttributes",
			Handler:    _OperatorService_RemoveSearchAttributes_Handler,
		},
		{
			MethodName: "ListSearchAttributes",
			Handler:    _OperatorService_ListSearchAttributes_Handler,
		},
		{
			MethodName: "DeleteNamespace",
			Handler:    _OperatorService_DeleteNamespace_Handler,
		},
		{
			MethodName: "AddOrUpdateRemoteCluster",
			Handler:    _OperatorService_AddOrUpdateRemoteCluster_Handler,
		},
		{
			MethodName: "RemoveRemoteCluster",
			Handler:    _OperatorService_RemoveRemoteCluster_Handler,
		},
		{
			MethodName: "ListClusters",
			Handler:    _OperatorService_ListClusters_Handler,
		},
		{
			MethodName: "GetNexusEndpoint",
			Handler:    _OperatorService_GetNexusEndpoint_Handler,
		},
		{
			MethodName: "CreateNexusEndpoint",
			Handler:    _OperatorService_CreateNexusEndpoint_Handler,
		},
		{
			MethodName: "UpdateNexusEndpoint",
			Handler:    _OperatorService_UpdateNexusEndpoint_Handler,
		},
		{
			MethodName: "DeleteNexusEndpoint",
			Handler:    _OperatorService_DeleteNexusEndpoint_Handler,
		},
		{
			MethodName: "ListNexusEndpoints",
			Handler:    _OperatorService_ListNexusEndpoints_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/api/operatorservice/v1/service.proto",
}
