// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.12.4
// source: internal/api/grpc/proto/trace_collector.proto

package trace_collector_gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	TraceCollector_Create_FullMethodName = "/trace.collector.TraceCollector/Create"
	TraceCollector_Update_FullMethodName = "/trace.collector.TraceCollector/Update"
)

// TraceCollectorClient is the client API for TraceCollector service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TraceCollectorClient interface {
	Create(ctx context.Context, in *TraceCreateRequest, opts ...grpc.CallOption) (*TraceCollectorResponse, error)
	Update(ctx context.Context, in *TraceUpdateRequest, opts ...grpc.CallOption) (*TraceCollectorResponse, error)
}

type traceCollectorClient struct {
	cc grpc.ClientConnInterface
}

func NewTraceCollectorClient(cc grpc.ClientConnInterface) TraceCollectorClient {
	return &traceCollectorClient{cc}
}

func (c *traceCollectorClient) Create(ctx context.Context, in *TraceCreateRequest, opts ...grpc.CallOption) (*TraceCollectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TraceCollectorResponse)
	err := c.cc.Invoke(ctx, TraceCollector_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceCollectorClient) Update(ctx context.Context, in *TraceUpdateRequest, opts ...grpc.CallOption) (*TraceCollectorResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TraceCollectorResponse)
	err := c.cc.Invoke(ctx, TraceCollector_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TraceCollectorServer is the server API for TraceCollector service.
// All implementations must embed UnimplementedTraceCollectorServer
// for forward compatibility.
type TraceCollectorServer interface {
	Create(context.Context, *TraceCreateRequest) (*TraceCollectorResponse, error)
	Update(context.Context, *TraceUpdateRequest) (*TraceCollectorResponse, error)
	mustEmbedUnimplementedTraceCollectorServer()
}

// UnimplementedTraceCollectorServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedTraceCollectorServer struct{}

func (UnimplementedTraceCollectorServer) Create(context.Context, *TraceCreateRequest) (*TraceCollectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTraceCollectorServer) Update(context.Context, *TraceUpdateRequest) (*TraceCollectorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTraceCollectorServer) mustEmbedUnimplementedTraceCollectorServer() {}
func (UnimplementedTraceCollectorServer) testEmbeddedByValue()                        {}

// UnsafeTraceCollectorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TraceCollectorServer will
// result in compilation errors.
type UnsafeTraceCollectorServer interface {
	mustEmbedUnimplementedTraceCollectorServer()
}

func RegisterTraceCollectorServer(s grpc.ServiceRegistrar, srv TraceCollectorServer) {
	// If the following call pancis, it indicates UnimplementedTraceCollectorServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&TraceCollector_ServiceDesc, srv)
}

func _TraceCollector_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TraceCollector_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).Create(ctx, req.(*TraceCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TraceCollector_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TraceUpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TraceCollectorServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TraceCollector_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TraceCollectorServer).Update(ctx, req.(*TraceUpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TraceCollector_ServiceDesc is the grpc.ServiceDesc for TraceCollector service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TraceCollector_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "trace.collector.TraceCollector",
	HandlerType: (*TraceCollectorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _TraceCollector_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TraceCollector_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/api/grpc/proto/trace_collector.proto",
}
