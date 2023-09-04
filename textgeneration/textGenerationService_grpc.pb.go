// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: textGenerationService.proto

package textgeneration

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

// TextGenerationServiceClient is the client API for TextGenerationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextGenerationServiceClient interface {
	/// Model Info
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
	/// Service discovery
	ServiceDiscovery(ctx context.Context, in *ServiceDiscoveryRequest, opts ...grpc.CallOption) (*ServiceDiscoveryResponse, error)
	/// Empties batch cache
	ClearCache(ctx context.Context, in *ClearCacheRequest, opts ...grpc.CallOption) (*ClearCacheResponse, error)
	/// Remove requests from a cached batch
	FilterBatch(ctx context.Context, in *FilterBatchRequest, opts ...grpc.CallOption) (*FilterBatchResponse, error)
	/// Warmup the model and compute max cache size
	Warmup(ctx context.Context, in *WarmupRequest, opts ...grpc.CallOption) (*WarmupResponse, error)
	/// Prefill batch and decode first token
	Prefill(ctx context.Context, in *PrefillRequest, opts ...grpc.CallOption) (*PrefillResponse, error)
	/// Decode token for a list of prefilled batches
	Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeResponse, error)
	/// Health check
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
}

type textGenerationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextGenerationServiceClient(cc grpc.ClientConnInterface) TextGenerationServiceClient {
	return &textGenerationServiceClient{cc}
}

func (c *textGenerationServiceClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerationServiceClient) ServiceDiscovery(ctx context.Context, in *ServiceDiscoveryRequest, opts ...grpc.CallOption) (*ServiceDiscoveryResponse, error) {
	out := new(ServiceDiscoveryResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/ServiceDiscovery", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerationServiceClient) ClearCache(ctx context.Context, in *ClearCacheRequest, opts ...grpc.CallOption) (*ClearCacheResponse, error) {
	out := new(ClearCacheResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/ClearCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerationServiceClient) FilterBatch(ctx context.Context, in *FilterBatchRequest, opts ...grpc.CallOption) (*FilterBatchResponse, error) {
	out := new(FilterBatchResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/FilterBatch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerationServiceClient) Warmup(ctx context.Context, in *WarmupRequest, opts ...grpc.CallOption) (*WarmupResponse, error) {
	out := new(WarmupResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/Warmup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerationServiceClient) Prefill(ctx context.Context, in *PrefillRequest, opts ...grpc.CallOption) (*PrefillResponse, error) {
	out := new(PrefillResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/Prefill", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerationServiceClient) Decode(ctx context.Context, in *DecodeRequest, opts ...grpc.CallOption) (*DecodeResponse, error) {
	out := new(DecodeResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/Decode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *textGenerationServiceClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, "/TextGenerationService/Health", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextGenerationServiceServer is the server API for TextGenerationService service.
// All implementations must embed UnimplementedTextGenerationServiceServer
// for forward compatibility
type TextGenerationServiceServer interface {
	/// Model Info
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	/// Service discovery
	ServiceDiscovery(context.Context, *ServiceDiscoveryRequest) (*ServiceDiscoveryResponse, error)
	/// Empties batch cache
	ClearCache(context.Context, *ClearCacheRequest) (*ClearCacheResponse, error)
	/// Remove requests from a cached batch
	FilterBatch(context.Context, *FilterBatchRequest) (*FilterBatchResponse, error)
	/// Warmup the model and compute max cache size
	Warmup(context.Context, *WarmupRequest) (*WarmupResponse, error)
	/// Prefill batch and decode first token
	Prefill(context.Context, *PrefillRequest) (*PrefillResponse, error)
	/// Decode token for a list of prefilled batches
	Decode(context.Context, *DecodeRequest) (*DecodeResponse, error)
	/// Health check
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	mustEmbedUnimplementedTextGenerationServiceServer()
}

// UnimplementedTextGenerationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTextGenerationServiceServer struct {
}

func (UnimplementedTextGenerationServiceServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedTextGenerationServiceServer) ServiceDiscovery(context.Context, *ServiceDiscoveryRequest) (*ServiceDiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceDiscovery not implemented")
}
func (UnimplementedTextGenerationServiceServer) ClearCache(context.Context, *ClearCacheRequest) (*ClearCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCache not implemented")
}
func (UnimplementedTextGenerationServiceServer) FilterBatch(context.Context, *FilterBatchRequest) (*FilterBatchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FilterBatch not implemented")
}
func (UnimplementedTextGenerationServiceServer) Warmup(context.Context, *WarmupRequest) (*WarmupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Warmup not implemented")
}
func (UnimplementedTextGenerationServiceServer) Prefill(context.Context, *PrefillRequest) (*PrefillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prefill not implemented")
}
func (UnimplementedTextGenerationServiceServer) Decode(context.Context, *DecodeRequest) (*DecodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decode not implemented")
}
func (UnimplementedTextGenerationServiceServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedTextGenerationServiceServer) mustEmbedUnimplementedTextGenerationServiceServer() {}

// UnsafeTextGenerationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextGenerationServiceServer will
// result in compilation errors.
type UnsafeTextGenerationServiceServer interface {
	mustEmbedUnimplementedTextGenerationServiceServer()
}

func RegisterTextGenerationServiceServer(s grpc.ServiceRegistrar, srv TextGenerationServiceServer) {
	s.RegisterService(&TextGenerationService_ServiceDesc, srv)
}

func _TextGenerationService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerationService_ServiceDiscovery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceDiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).ServiceDiscovery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/ServiceDiscovery",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).ServiceDiscovery(ctx, req.(*ServiceDiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerationService_ClearCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).ClearCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/ClearCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).ClearCache(ctx, req.(*ClearCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerationService_FilterBatch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FilterBatchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).FilterBatch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/FilterBatch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).FilterBatch(ctx, req.(*FilterBatchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerationService_Warmup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WarmupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).Warmup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/Warmup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).Warmup(ctx, req.(*WarmupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerationService_Prefill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).Prefill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/Prefill",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).Prefill(ctx, req.(*PrefillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerationService_Decode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).Decode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/Decode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).Decode(ctx, req.(*DecodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TextGenerationService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextGenerationServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextGenerationService/Health",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextGenerationServiceServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextGenerationService_ServiceDesc is the grpc.ServiceDesc for TextGenerationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextGenerationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TextGenerationService",
	HandlerType: (*TextGenerationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Info",
			Handler:    _TextGenerationService_Info_Handler,
		},
		{
			MethodName: "ServiceDiscovery",
			Handler:    _TextGenerationService_ServiceDiscovery_Handler,
		},
		{
			MethodName: "ClearCache",
			Handler:    _TextGenerationService_ClearCache_Handler,
		},
		{
			MethodName: "FilterBatch",
			Handler:    _TextGenerationService_FilterBatch_Handler,
		},
		{
			MethodName: "Warmup",
			Handler:    _TextGenerationService_Warmup_Handler,
		},
		{
			MethodName: "Prefill",
			Handler:    _TextGenerationService_Prefill_Handler,
		},
		{
			MethodName: "Decode",
			Handler:    _TextGenerationService_Decode_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _TextGenerationService_Health_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "textGenerationService.proto",
}
