// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.6.1
// source: apps_svc.proto

package protos

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	AppsService_GetApps_FullMethodName       = "/protos.AppsService/GetApps"
	AppsService_GetApp_FullMethodName        = "/protos.AppsService/GetApp"
	AppsService_CreateApp_FullMethodName     = "/protos.AppsService/CreateApp"
	AppsService_RefreshAppKey_FullMethodName = "/protos.AppsService/RefreshAppKey"
	AppsService_DecodeAppKey_FullMethodName  = "/protos.AppsService/DecodeAppKey"
	AppsService_DeleteApp_FullMethodName     = "/protos.AppsService/DeleteApp"
)

// AppsServiceClient is the client API for AppsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppsServiceClient interface {
	GetApps(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppsResponse, error)
	GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*App, error)
	CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*App, error)
	RefreshAppKey(ctx context.Context, in *RefreshAppKeyRequest, opts ...grpc.CallOption) (*RefreshAppKeyResponse, error)
	DecodeAppKey(ctx context.Context, in *DecodeAppKeyRequest, opts ...grpc.CallOption) (*DecodeAppKeyResponse, error)
	DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type appsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAppsServiceClient(cc grpc.ClientConnInterface) AppsServiceClient {
	return &appsServiceClient{cc}
}

func (c *appsServiceClient) GetApps(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetAppsResponse, error) {
	out := new(GetAppsResponse)
	err := c.cc.Invoke(ctx, AppsService_GetApps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsServiceClient) GetApp(ctx context.Context, in *GetAppRequest, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, AppsService_GetApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsServiceClient) CreateApp(ctx context.Context, in *CreateAppRequest, opts ...grpc.CallOption) (*App, error) {
	out := new(App)
	err := c.cc.Invoke(ctx, AppsService_CreateApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsServiceClient) RefreshAppKey(ctx context.Context, in *RefreshAppKeyRequest, opts ...grpc.CallOption) (*RefreshAppKeyResponse, error) {
	out := new(RefreshAppKeyResponse)
	err := c.cc.Invoke(ctx, AppsService_RefreshAppKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsServiceClient) DecodeAppKey(ctx context.Context, in *DecodeAppKeyRequest, opts ...grpc.CallOption) (*DecodeAppKeyResponse, error) {
	out := new(DecodeAppKeyResponse)
	err := c.cc.Invoke(ctx, AppsService_DecodeAppKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *appsServiceClient) DeleteApp(ctx context.Context, in *DeleteAppRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, AppsService_DeleteApp_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppsServiceServer is the server API for AppsService service.
// All implementations must embed UnimplementedAppsServiceServer
// for forward compatibility
type AppsServiceServer interface {
	GetApps(context.Context, *empty.Empty) (*GetAppsResponse, error)
	GetApp(context.Context, *GetAppRequest) (*App, error)
	CreateApp(context.Context, *CreateAppRequest) (*App, error)
	RefreshAppKey(context.Context, *RefreshAppKeyRequest) (*RefreshAppKeyResponse, error)
	DecodeAppKey(context.Context, *DecodeAppKeyRequest) (*DecodeAppKeyResponse, error)
	DeleteApp(context.Context, *DeleteAppRequest) (*empty.Empty, error)
	mustEmbedUnimplementedAppsServiceServer()
}

// UnimplementedAppsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAppsServiceServer struct {
}

func (UnimplementedAppsServiceServer) GetApps(context.Context, *empty.Empty) (*GetAppsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApps not implemented")
}
func (UnimplementedAppsServiceServer) GetApp(context.Context, *GetAppRequest) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApp not implemented")
}
func (UnimplementedAppsServiceServer) CreateApp(context.Context, *CreateAppRequest) (*App, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateApp not implemented")
}
func (UnimplementedAppsServiceServer) RefreshAppKey(context.Context, *RefreshAppKeyRequest) (*RefreshAppKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshAppKey not implemented")
}
func (UnimplementedAppsServiceServer) DecodeAppKey(context.Context, *DecodeAppKeyRequest) (*DecodeAppKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecodeAppKey not implemented")
}
func (UnimplementedAppsServiceServer) DeleteApp(context.Context, *DeleteAppRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteApp not implemented")
}
func (UnimplementedAppsServiceServer) mustEmbedUnimplementedAppsServiceServer() {}

// UnsafeAppsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppsServiceServer will
// result in compilation errors.
type UnsafeAppsServiceServer interface {
	mustEmbedUnimplementedAppsServiceServer()
}

func RegisterAppsServiceServer(s grpc.ServiceRegistrar, srv AppsServiceServer) {
	s.RegisterService(&AppsService_ServiceDesc, srv)
}

func _AppsService_GetApps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).GetApps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_GetApps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).GetApps(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppsService_GetApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).GetApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_GetApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).GetApp(ctx, req.(*GetAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppsService_CreateApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).CreateApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_CreateApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).CreateApp(ctx, req.(*CreateAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppsService_RefreshAppKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshAppKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).RefreshAppKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_RefreshAppKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).RefreshAppKey(ctx, req.(*RefreshAppKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppsService_DecodeAppKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecodeAppKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).DecodeAppKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_DecodeAppKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).DecodeAppKey(ctx, req.(*DecodeAppKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AppsService_DeleteApp_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAppRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppsServiceServer).DeleteApp(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AppsService_DeleteApp_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppsServiceServer).DeleteApp(ctx, req.(*DeleteAppRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AppsService_ServiceDesc is the grpc.ServiceDesc for AppsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AppsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.AppsService",
	HandlerType: (*AppsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetApps",
			Handler:    _AppsService_GetApps_Handler,
		},
		{
			MethodName: "GetApp",
			Handler:    _AppsService_GetApp_Handler,
		},
		{
			MethodName: "CreateApp",
			Handler:    _AppsService_CreateApp_Handler,
		},
		{
			MethodName: "RefreshAppKey",
			Handler:    _AppsService_RefreshAppKey_Handler,
		},
		{
			MethodName: "DecodeAppKey",
			Handler:    _AppsService_DecodeAppKey_Handler,
		},
		{
			MethodName: "DeleteApp",
			Handler:    _AppsService_DeleteApp_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "apps_svc.proto",
}
