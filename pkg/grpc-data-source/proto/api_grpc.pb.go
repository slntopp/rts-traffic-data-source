// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

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

// GrafanaQueryAPIClient is the client API for GrafanaQueryAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GrafanaQueryAPIClient interface {
	// Returns a list of all available dimensions
	ListDimensionKeys(ctx context.Context, in *ListDimensionKeysRequest, opts ...grpc.CallOption) (*ListDimensionKeysResponse, error)
	// Returns a list of all dimension values for a certain dimension
	ListDimensionValues(ctx context.Context, in *ListDimensionValuesRequest, opts ...grpc.CallOption) (*ListDimensionValuesResponse, error)
	// Returns all metrics from the system
	ListMetrics(ctx context.Context, in *ListMetricsRequest, opts ...grpc.CallOption) (*ListMetricsResponse, error)
	// Gets a metric's current value
	GetMetricValue(ctx context.Context, in *GetMetricValueRequest, opts ...grpc.CallOption) (*GetMetricValueResponse, error)
	// Gets the history of a metric's values
	GetMetricHistory(ctx context.Context, in *GetMetricHistoryRequest, opts ...grpc.CallOption) (*GetMetricHistoryResponse, error)
	// Gets the history of a metric's aggregated value
	GetMetricAggregate(ctx context.Context, in *GetMetricAggregateRequest, opts ...grpc.CallOption) (*GetMetricAggregateResponse, error)
}

type grafanaQueryAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewGrafanaQueryAPIClient(cc grpc.ClientConnInterface) GrafanaQueryAPIClient {
	return &grafanaQueryAPIClient{cc}
}

func (c *grafanaQueryAPIClient) ListDimensionKeys(ctx context.Context, in *ListDimensionKeysRequest, opts ...grpc.CallOption) (*ListDimensionKeysResponse, error) {
	out := new(ListDimensionKeysResponse)
	err := c.cc.Invoke(ctx, "/grafana.GrafanaQueryAPI/ListDimensionKeys", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grafanaQueryAPIClient) ListDimensionValues(ctx context.Context, in *ListDimensionValuesRequest, opts ...grpc.CallOption) (*ListDimensionValuesResponse, error) {
	out := new(ListDimensionValuesResponse)
	err := c.cc.Invoke(ctx, "/grafana.GrafanaQueryAPI/ListDimensionValues", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grafanaQueryAPIClient) ListMetrics(ctx context.Context, in *ListMetricsRequest, opts ...grpc.CallOption) (*ListMetricsResponse, error) {
	out := new(ListMetricsResponse)
	err := c.cc.Invoke(ctx, "/grafana.GrafanaQueryAPI/ListMetrics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grafanaQueryAPIClient) GetMetricValue(ctx context.Context, in *GetMetricValueRequest, opts ...grpc.CallOption) (*GetMetricValueResponse, error) {
	out := new(GetMetricValueResponse)
	err := c.cc.Invoke(ctx, "/grafana.GrafanaQueryAPI/GetMetricValue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grafanaQueryAPIClient) GetMetricHistory(ctx context.Context, in *GetMetricHistoryRequest, opts ...grpc.CallOption) (*GetMetricHistoryResponse, error) {
	out := new(GetMetricHistoryResponse)
	err := c.cc.Invoke(ctx, "/grafana.GrafanaQueryAPI/GetMetricHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grafanaQueryAPIClient) GetMetricAggregate(ctx context.Context, in *GetMetricAggregateRequest, opts ...grpc.CallOption) (*GetMetricAggregateResponse, error) {
	out := new(GetMetricAggregateResponse)
	err := c.cc.Invoke(ctx, "/grafana.GrafanaQueryAPI/GetMetricAggregate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GrafanaQueryAPIServer is the server API for GrafanaQueryAPI service.
// All implementations must embed UnimplementedGrafanaQueryAPIServer
// for forward compatibility
type GrafanaQueryAPIServer interface {
	// Returns a list of all available dimensions
	ListDimensionKeys(context.Context, *ListDimensionKeysRequest) (*ListDimensionKeysResponse, error)
	// Returns a list of all dimension values for a certain dimension
	ListDimensionValues(context.Context, *ListDimensionValuesRequest) (*ListDimensionValuesResponse, error)
	// Returns all metrics from the system
	ListMetrics(context.Context, *ListMetricsRequest) (*ListMetricsResponse, error)
	// Gets a metric's current value
	GetMetricValue(context.Context, *GetMetricValueRequest) (*GetMetricValueResponse, error)
	// Gets the history of a metric's values
	GetMetricHistory(context.Context, *GetMetricHistoryRequest) (*GetMetricHistoryResponse, error)
	// Gets the history of a metric's aggregated value
	GetMetricAggregate(context.Context, *GetMetricAggregateRequest) (*GetMetricAggregateResponse, error)
	mustEmbedUnimplementedGrafanaQueryAPIServer()
}

// UnimplementedGrafanaQueryAPIServer must be embedded to have forward compatible implementations.
type UnimplementedGrafanaQueryAPIServer struct {
}

func (UnimplementedGrafanaQueryAPIServer) ListDimensionKeys(context.Context, *ListDimensionKeysRequest) (*ListDimensionKeysResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDimensionKeys not implemented")
}
func (UnimplementedGrafanaQueryAPIServer) ListDimensionValues(context.Context, *ListDimensionValuesRequest) (*ListDimensionValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDimensionValues not implemented")
}
func (UnimplementedGrafanaQueryAPIServer) ListMetrics(context.Context, *ListMetricsRequest) (*ListMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetrics not implemented")
}
func (UnimplementedGrafanaQueryAPIServer) GetMetricValue(context.Context, *GetMetricValueRequest) (*GetMetricValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricValue not implemented")
}
func (UnimplementedGrafanaQueryAPIServer) GetMetricHistory(context.Context, *GetMetricHistoryRequest) (*GetMetricHistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricHistory not implemented")
}
func (UnimplementedGrafanaQueryAPIServer) GetMetricAggregate(context.Context, *GetMetricAggregateRequest) (*GetMetricAggregateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetricAggregate not implemented")
}
func (UnimplementedGrafanaQueryAPIServer) mustEmbedUnimplementedGrafanaQueryAPIServer() {}

// UnsafeGrafanaQueryAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GrafanaQueryAPIServer will
// result in compilation errors.
type UnsafeGrafanaQueryAPIServer interface {
	mustEmbedUnimplementedGrafanaQueryAPIServer()
}

func RegisterGrafanaQueryAPIServer(s grpc.ServiceRegistrar, srv GrafanaQueryAPIServer) {
	s.RegisterService(&GrafanaQueryAPI_ServiceDesc, srv)
}

func _GrafanaQueryAPI_ListDimensionKeys_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDimensionKeysRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrafanaQueryAPIServer).ListDimensionKeys(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafana.GrafanaQueryAPI/ListDimensionKeys",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrafanaQueryAPIServer).ListDimensionKeys(ctx, req.(*ListDimensionKeysRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrafanaQueryAPI_ListDimensionValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDimensionValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrafanaQueryAPIServer).ListDimensionValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafana.GrafanaQueryAPI/ListDimensionValues",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrafanaQueryAPIServer).ListDimensionValues(ctx, req.(*ListDimensionValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrafanaQueryAPI_ListMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrafanaQueryAPIServer).ListMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafana.GrafanaQueryAPI/ListMetrics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrafanaQueryAPIServer).ListMetrics(ctx, req.(*ListMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrafanaQueryAPI_GetMetricValue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricValueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrafanaQueryAPIServer).GetMetricValue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafana.GrafanaQueryAPI/GetMetricValue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrafanaQueryAPIServer).GetMetricValue(ctx, req.(*GetMetricValueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrafanaQueryAPI_GetMetricHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrafanaQueryAPIServer).GetMetricHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafana.GrafanaQueryAPI/GetMetricHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrafanaQueryAPIServer).GetMetricHistory(ctx, req.(*GetMetricHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GrafanaQueryAPI_GetMetricAggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricAggregateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrafanaQueryAPIServer).GetMetricAggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grafana.GrafanaQueryAPI/GetMetricAggregate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrafanaQueryAPIServer).GetMetricAggregate(ctx, req.(*GetMetricAggregateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GrafanaQueryAPI_ServiceDesc is the grpc.ServiceDesc for GrafanaQueryAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GrafanaQueryAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grafana.GrafanaQueryAPI",
	HandlerType: (*GrafanaQueryAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDimensionKeys",
			Handler:    _GrafanaQueryAPI_ListDimensionKeys_Handler,
		},
		{
			MethodName: "ListDimensionValues",
			Handler:    _GrafanaQueryAPI_ListDimensionValues_Handler,
		},
		{
			MethodName: "ListMetrics",
			Handler:    _GrafanaQueryAPI_ListMetrics_Handler,
		},
		{
			MethodName: "GetMetricValue",
			Handler:    _GrafanaQueryAPI_GetMetricValue_Handler,
		},
		{
			MethodName: "GetMetricHistory",
			Handler:    _GrafanaQueryAPI_GetMetricHistory_Handler,
		},
		{
			MethodName: "GetMetricAggregate",
			Handler:    _GrafanaQueryAPI_GetMetricAggregate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/grpc-data-source/proto/api.proto",
}
