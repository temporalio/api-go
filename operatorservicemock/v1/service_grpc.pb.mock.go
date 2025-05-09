// Code generated by MockGen. DO NOT EDIT.
// Source: operatorservice/v1/service_grpc.pb.go

// Package operatorservicemock is a generated GoMock package.
package operatorservicemock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	operatorservice "go.temporal.io/api/operatorservice/v1"
	grpc "google.golang.org/grpc"
)

// MockOperatorServiceClient is a mock of OperatorServiceClient interface.
type MockOperatorServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockOperatorServiceClientMockRecorder
}

// MockOperatorServiceClientMockRecorder is the mock recorder for MockOperatorServiceClient.
type MockOperatorServiceClientMockRecorder struct {
	mock *MockOperatorServiceClient
}

// NewMockOperatorServiceClient creates a new mock instance.
func NewMockOperatorServiceClient(ctrl *gomock.Controller) *MockOperatorServiceClient {
	mock := &MockOperatorServiceClient{ctrl: ctrl}
	mock.recorder = &MockOperatorServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorServiceClient) EXPECT() *MockOperatorServiceClientMockRecorder {
	return m.recorder
}

// AddOrUpdateRemoteCluster mocks base method.
func (m *MockOperatorServiceClient) AddOrUpdateRemoteCluster(ctx context.Context, in *operatorservice.AddOrUpdateRemoteClusterRequest, opts ...grpc.CallOption) (*operatorservice.AddOrUpdateRemoteClusterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddOrUpdateRemoteCluster", varargs...)
	ret0, _ := ret[0].(*operatorservice.AddOrUpdateRemoteClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrUpdateRemoteCluster indicates an expected call of AddOrUpdateRemoteCluster.
func (mr *MockOperatorServiceClientMockRecorder) AddOrUpdateRemoteCluster(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateRemoteCluster", reflect.TypeOf((*MockOperatorServiceClient)(nil).AddOrUpdateRemoteCluster), varargs...)
}

// AddSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) AddSearchAttributes(ctx context.Context, in *operatorservice.AddSearchAttributesRequest, opts ...grpc.CallOption) (*operatorservice.AddSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.AddSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSearchAttributes indicates an expected call of AddSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) AddSearchAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).AddSearchAttributes), varargs...)
}

// CreateNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) CreateNexusEndpoint(ctx context.Context, in *operatorservice.CreateNexusEndpointRequest, opts ...grpc.CallOption) (*operatorservice.CreateNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.CreateNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNexusEndpoint indicates an expected call of CreateNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) CreateNexusEndpoint(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).CreateNexusEndpoint), varargs...)
}

// DeleteNamespace mocks base method.
func (m *MockOperatorServiceClient) DeleteNamespace(ctx context.Context, in *operatorservice.DeleteNamespaceRequest, opts ...grpc.CallOption) (*operatorservice.DeleteNamespaceResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNamespace", varargs...)
	ret0, _ := ret[0].(*operatorservice.DeleteNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockOperatorServiceClientMockRecorder) DeleteNamespace(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockOperatorServiceClient)(nil).DeleteNamespace), varargs...)
}

// DeleteNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) DeleteNexusEndpoint(ctx context.Context, in *operatorservice.DeleteNexusEndpointRequest, opts ...grpc.CallOption) (*operatorservice.DeleteNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DeleteNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.DeleteNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNexusEndpoint indicates an expected call of DeleteNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) DeleteNexusEndpoint(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).DeleteNexusEndpoint), varargs...)
}

// GetNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) GetNexusEndpoint(ctx context.Context, in *operatorservice.GetNexusEndpointRequest, opts ...grpc.CallOption) (*operatorservice.GetNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.GetNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNexusEndpoint indicates an expected call of GetNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) GetNexusEndpoint(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).GetNexusEndpoint), varargs...)
}

// ListClusters mocks base method.
func (m *MockOperatorServiceClient) ListClusters(ctx context.Context, in *operatorservice.ListClustersRequest, opts ...grpc.CallOption) (*operatorservice.ListClustersResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListClusters", varargs...)
	ret0, _ := ret[0].(*operatorservice.ListClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockOperatorServiceClientMockRecorder) ListClusters(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockOperatorServiceClient)(nil).ListClusters), varargs...)
}

// ListNexusEndpoints mocks base method.
func (m *MockOperatorServiceClient) ListNexusEndpoints(ctx context.Context, in *operatorservice.ListNexusEndpointsRequest, opts ...grpc.CallOption) (*operatorservice.ListNexusEndpointsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListNexusEndpoints", varargs...)
	ret0, _ := ret[0].(*operatorservice.ListNexusEndpointsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNexusEndpoints indicates an expected call of ListNexusEndpoints.
func (mr *MockOperatorServiceClientMockRecorder) ListNexusEndpoints(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNexusEndpoints", reflect.TypeOf((*MockOperatorServiceClient)(nil).ListNexusEndpoints), varargs...)
}

// ListSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) ListSearchAttributes(ctx context.Context, in *operatorservice.ListSearchAttributesRequest, opts ...grpc.CallOption) (*operatorservice.ListSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.ListSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSearchAttributes indicates an expected call of ListSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) ListSearchAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).ListSearchAttributes), varargs...)
}

// RemoveRemoteCluster mocks base method.
func (m *MockOperatorServiceClient) RemoveRemoteCluster(ctx context.Context, in *operatorservice.RemoveRemoteClusterRequest, opts ...grpc.CallOption) (*operatorservice.RemoveRemoteClusterResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveRemoteCluster", varargs...)
	ret0, _ := ret[0].(*operatorservice.RemoveRemoteClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveRemoteCluster indicates an expected call of RemoveRemoteCluster.
func (mr *MockOperatorServiceClientMockRecorder) RemoveRemoteCluster(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRemoteCluster", reflect.TypeOf((*MockOperatorServiceClient)(nil).RemoveRemoteCluster), varargs...)
}

// RemoveSearchAttributes mocks base method.
func (m *MockOperatorServiceClient) RemoveSearchAttributes(ctx context.Context, in *operatorservice.RemoveSearchAttributesRequest, opts ...grpc.CallOption) (*operatorservice.RemoveSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RemoveSearchAttributes", varargs...)
	ret0, _ := ret[0].(*operatorservice.RemoveSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSearchAttributes indicates an expected call of RemoveSearchAttributes.
func (mr *MockOperatorServiceClientMockRecorder) RemoveSearchAttributes(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSearchAttributes", reflect.TypeOf((*MockOperatorServiceClient)(nil).RemoveSearchAttributes), varargs...)
}

// UpdateNexusEndpoint mocks base method.
func (m *MockOperatorServiceClient) UpdateNexusEndpoint(ctx context.Context, in *operatorservice.UpdateNexusEndpointRequest, opts ...grpc.CallOption) (*operatorservice.UpdateNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateNexusEndpoint", varargs...)
	ret0, _ := ret[0].(*operatorservice.UpdateNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNexusEndpoint indicates an expected call of UpdateNexusEndpoint.
func (mr *MockOperatorServiceClientMockRecorder) UpdateNexusEndpoint(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNexusEndpoint", reflect.TypeOf((*MockOperatorServiceClient)(nil).UpdateNexusEndpoint), varargs...)
}

// MockOperatorServiceServer is a mock of OperatorServiceServer interface.
type MockOperatorServiceServer struct {
	operatorservice.UnimplementedOperatorServiceServer
	ctrl     *gomock.Controller
	recorder *MockOperatorServiceServerMockRecorder
}

var _ operatorservice.OperatorServiceServer = (*MockOperatorServiceServer)(nil)

// MockOperatorServiceServerMockRecorder is the mock recorder for MockOperatorServiceServer.
type MockOperatorServiceServerMockRecorder struct {
	mock *MockOperatorServiceServer
}

// NewMockOperatorServiceServer creates a new mock instance.
func NewMockOperatorServiceServer(ctrl *gomock.Controller) *MockOperatorServiceServer {
	mock := &MockOperatorServiceServer{ctrl: ctrl}
	mock.recorder = &MockOperatorServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOperatorServiceServer) EXPECT() *MockOperatorServiceServerMockRecorder {
	return m.recorder
}

// AddOrUpdateRemoteCluster mocks base method.
func (m *MockOperatorServiceServer) AddOrUpdateRemoteCluster(arg0 context.Context, arg1 *operatorservice.AddOrUpdateRemoteClusterRequest) (*operatorservice.AddOrUpdateRemoteClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrUpdateRemoteCluster", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.AddOrUpdateRemoteClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrUpdateRemoteCluster indicates an expected call of AddOrUpdateRemoteCluster.
func (mr *MockOperatorServiceServerMockRecorder) AddOrUpdateRemoteCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrUpdateRemoteCluster", reflect.TypeOf((*MockOperatorServiceServer)(nil).AddOrUpdateRemoteCluster), arg0, arg1)
}

// AddSearchAttributes mocks base method.
func (m *MockOperatorServiceServer) AddSearchAttributes(arg0 context.Context, arg1 *operatorservice.AddSearchAttributesRequest) (*operatorservice.AddSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddSearchAttributes", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.AddSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddSearchAttributes indicates an expected call of AddSearchAttributes.
func (mr *MockOperatorServiceServerMockRecorder) AddSearchAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSearchAttributes", reflect.TypeOf((*MockOperatorServiceServer)(nil).AddSearchAttributes), arg0, arg1)
}

// CreateNexusEndpoint mocks base method.
func (m *MockOperatorServiceServer) CreateNexusEndpoint(arg0 context.Context, arg1 *operatorservice.CreateNexusEndpointRequest) (*operatorservice.CreateNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNexusEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.CreateNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNexusEndpoint indicates an expected call of CreateNexusEndpoint.
func (mr *MockOperatorServiceServerMockRecorder) CreateNexusEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNexusEndpoint", reflect.TypeOf((*MockOperatorServiceServer)(nil).CreateNexusEndpoint), arg0, arg1)
}

// DeleteNamespace mocks base method.
func (m *MockOperatorServiceServer) DeleteNamespace(arg0 context.Context, arg1 *operatorservice.DeleteNamespaceRequest) (*operatorservice.DeleteNamespaceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespace", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.DeleteNamespaceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNamespace indicates an expected call of DeleteNamespace.
func (mr *MockOperatorServiceServerMockRecorder) DeleteNamespace(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespace", reflect.TypeOf((*MockOperatorServiceServer)(nil).DeleteNamespace), arg0, arg1)
}

// DeleteNexusEndpoint mocks base method.
func (m *MockOperatorServiceServer) DeleteNexusEndpoint(arg0 context.Context, arg1 *operatorservice.DeleteNexusEndpointRequest) (*operatorservice.DeleteNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNexusEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.DeleteNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNexusEndpoint indicates an expected call of DeleteNexusEndpoint.
func (mr *MockOperatorServiceServerMockRecorder) DeleteNexusEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNexusEndpoint", reflect.TypeOf((*MockOperatorServiceServer)(nil).DeleteNexusEndpoint), arg0, arg1)
}

// GetNexusEndpoint mocks base method.
func (m *MockOperatorServiceServer) GetNexusEndpoint(arg0 context.Context, arg1 *operatorservice.GetNexusEndpointRequest) (*operatorservice.GetNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNexusEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.GetNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNexusEndpoint indicates an expected call of GetNexusEndpoint.
func (mr *MockOperatorServiceServerMockRecorder) GetNexusEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNexusEndpoint", reflect.TypeOf((*MockOperatorServiceServer)(nil).GetNexusEndpoint), arg0, arg1)
}

// ListClusters mocks base method.
func (m *MockOperatorServiceServer) ListClusters(arg0 context.Context, arg1 *operatorservice.ListClustersRequest) (*operatorservice.ListClustersResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListClusters", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.ListClustersResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListClusters indicates an expected call of ListClusters.
func (mr *MockOperatorServiceServerMockRecorder) ListClusters(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListClusters", reflect.TypeOf((*MockOperatorServiceServer)(nil).ListClusters), arg0, arg1)
}

// ListNexusEndpoints mocks base method.
func (m *MockOperatorServiceServer) ListNexusEndpoints(arg0 context.Context, arg1 *operatorservice.ListNexusEndpointsRequest) (*operatorservice.ListNexusEndpointsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNexusEndpoints", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.ListNexusEndpointsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNexusEndpoints indicates an expected call of ListNexusEndpoints.
func (mr *MockOperatorServiceServerMockRecorder) ListNexusEndpoints(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNexusEndpoints", reflect.TypeOf((*MockOperatorServiceServer)(nil).ListNexusEndpoints), arg0, arg1)
}

// ListSearchAttributes mocks base method.
func (m *MockOperatorServiceServer) ListSearchAttributes(arg0 context.Context, arg1 *operatorservice.ListSearchAttributesRequest) (*operatorservice.ListSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSearchAttributes", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.ListSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSearchAttributes indicates an expected call of ListSearchAttributes.
func (mr *MockOperatorServiceServerMockRecorder) ListSearchAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSearchAttributes", reflect.TypeOf((*MockOperatorServiceServer)(nil).ListSearchAttributes), arg0, arg1)
}

// RemoveRemoteCluster mocks base method.
func (m *MockOperatorServiceServer) RemoveRemoteCluster(arg0 context.Context, arg1 *operatorservice.RemoveRemoteClusterRequest) (*operatorservice.RemoveRemoteClusterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveRemoteCluster", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.RemoveRemoteClusterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveRemoteCluster indicates an expected call of RemoveRemoteCluster.
func (mr *MockOperatorServiceServerMockRecorder) RemoveRemoteCluster(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveRemoteCluster", reflect.TypeOf((*MockOperatorServiceServer)(nil).RemoveRemoteCluster), arg0, arg1)
}

// RemoveSearchAttributes mocks base method.
func (m *MockOperatorServiceServer) RemoveSearchAttributes(arg0 context.Context, arg1 *operatorservice.RemoveSearchAttributesRequest) (*operatorservice.RemoveSearchAttributesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveSearchAttributes", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.RemoveSearchAttributesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveSearchAttributes indicates an expected call of RemoveSearchAttributes.
func (mr *MockOperatorServiceServerMockRecorder) RemoveSearchAttributes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveSearchAttributes", reflect.TypeOf((*MockOperatorServiceServer)(nil).RemoveSearchAttributes), arg0, arg1)
}

// UpdateNexusEndpoint mocks base method.
func (m *MockOperatorServiceServer) UpdateNexusEndpoint(arg0 context.Context, arg1 *operatorservice.UpdateNexusEndpointRequest) (*operatorservice.UpdateNexusEndpointResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNexusEndpoint", arg0, arg1)
	ret0, _ := ret[0].(*operatorservice.UpdateNexusEndpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateNexusEndpoint indicates an expected call of UpdateNexusEndpoint.
func (mr *MockOperatorServiceServerMockRecorder) UpdateNexusEndpoint(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNexusEndpoint", reflect.TypeOf((*MockOperatorServiceServer)(nil).UpdateNexusEndpoint), arg0, arg1)
}

// mustEmbedUnimplementedOperatorServiceServer mocks base method.
func (m *MockOperatorServiceServer) mustEmbedUnimplementedOperatorServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedOperatorServiceServer")
}

// mustEmbedUnimplementedOperatorServiceServer indicates an expected call of mustEmbedUnimplementedOperatorServiceServer.
func (mr *MockOperatorServiceServerMockRecorder) mustEmbedUnimplementedOperatorServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedOperatorServiceServer", reflect.TypeOf((*MockOperatorServiceServer)(nil).mustEmbedUnimplementedOperatorServiceServer))
}

// MockUnsafeOperatorServiceServer is a mock of UnsafeOperatorServiceServer interface.
type MockUnsafeOperatorServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeOperatorServiceServerMockRecorder
}

// MockUnsafeOperatorServiceServerMockRecorder is the mock recorder for MockUnsafeOperatorServiceServer.
type MockUnsafeOperatorServiceServerMockRecorder struct {
	mock *MockUnsafeOperatorServiceServer
}

// NewMockUnsafeOperatorServiceServer creates a new mock instance.
func NewMockUnsafeOperatorServiceServer(ctrl *gomock.Controller) *MockUnsafeOperatorServiceServer {
	mock := &MockUnsafeOperatorServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeOperatorServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeOperatorServiceServer) EXPECT() *MockUnsafeOperatorServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedOperatorServiceServer mocks base method.
func (m *MockUnsafeOperatorServiceServer) mustEmbedUnimplementedOperatorServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedOperatorServiceServer")
}

// mustEmbedUnimplementedOperatorServiceServer indicates an expected call of mustEmbedUnimplementedOperatorServiceServer.
func (mr *MockUnsafeOperatorServiceServerMockRecorder) mustEmbedUnimplementedOperatorServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedOperatorServiceServer", reflect.TypeOf((*MockUnsafeOperatorServiceServer)(nil).mustEmbedUnimplementedOperatorServiceServer))
}
