// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: yubikey/v1/yubikey.proto

package yubikey

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

// YubikeyServiceClient is the client API for YubikeyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type YubikeyServiceClient interface {
	GetSerials(ctx context.Context, in *GetSerialsRequest, opts ...grpc.CallOption) (*GetSerialsResponse, error)
	GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error)
}

type yubikeyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewYubikeyServiceClient(cc grpc.ClientConnInterface) YubikeyServiceClient {
	return &yubikeyServiceClient{cc}
}

func (c *yubikeyServiceClient) GetSerials(ctx context.Context, in *GetSerialsRequest, opts ...grpc.CallOption) (*GetSerialsResponse, error) {
	out := new(GetSerialsResponse)
	err := c.cc.Invoke(ctx, "/yubikey.v1.YubikeyService/GetSerials", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *yubikeyServiceClient) GetVersion(ctx context.Context, in *GetVersionRequest, opts ...grpc.CallOption) (*GetVersionResponse, error) {
	out := new(GetVersionResponse)
	err := c.cc.Invoke(ctx, "/yubikey.v1.YubikeyService/GetVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// YubikeyServiceServer is the server API for YubikeyService service.
// All implementations must embed UnimplementedYubikeyServiceServer
// for forward compatibility
type YubikeyServiceServer interface {
	GetSerials(context.Context, *GetSerialsRequest) (*GetSerialsResponse, error)
	GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error)
	mustEmbedUnimplementedYubikeyServiceServer()
}

// UnimplementedYubikeyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedYubikeyServiceServer struct {
}

func (UnimplementedYubikeyServiceServer) GetSerials(context.Context, *GetSerialsRequest) (*GetSerialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSerials not implemented")
}
func (UnimplementedYubikeyServiceServer) GetVersion(context.Context, *GetVersionRequest) (*GetVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVersion not implemented")
}
func (UnimplementedYubikeyServiceServer) mustEmbedUnimplementedYubikeyServiceServer() {}

// UnsafeYubikeyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to YubikeyServiceServer will
// result in compilation errors.
type UnsafeYubikeyServiceServer interface {
	mustEmbedUnimplementedYubikeyServiceServer()
}

func RegisterYubikeyServiceServer(s grpc.ServiceRegistrar, srv YubikeyServiceServer) {
	s.RegisterService(&YubikeyService_ServiceDesc, srv)
}

func _YubikeyService_GetSerials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSerialsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YubikeyServiceServer).GetSerials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yubikey.v1.YubikeyService/GetSerials",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YubikeyServiceServer).GetSerials(ctx, req.(*GetSerialsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _YubikeyService_GetVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(YubikeyServiceServer).GetVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yubikey.v1.YubikeyService/GetVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(YubikeyServiceServer).GetVersion(ctx, req.(*GetVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// YubikeyService_ServiceDesc is the grpc.ServiceDesc for YubikeyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var YubikeyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "yubikey.v1.YubikeyService",
	HandlerType: (*YubikeyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSerials",
			Handler:    _YubikeyService_GetSerials_Handler,
		},
		{
			MethodName: "GetVersion",
			Handler:    _YubikeyService_GetVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "yubikey/v1/yubikey.proto",
}
