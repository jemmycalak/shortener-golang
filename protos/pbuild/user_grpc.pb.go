// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.2
// source: protos/user.proto

package pbuild

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

// UserServiceProtoClient is the client API for UserServiceProto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceProtoClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserProto, error)
	GetUserBy(ctx context.Context, in *GetUserByRequest, opts ...grpc.CallOption) (*UserProto, error)
	GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*UserProto, error)
	AutorizeToken(ctx context.Context, in *AuthorizeTokenRequest, opts ...grpc.CallOption) (*AuthorizeTokenResponse, error)
}

type userServiceProtoClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceProtoClient(cc grpc.ClientConnInterface) UserServiceProtoClient {
	return &userServiceProtoClient{cc}
}

func (c *userServiceProtoClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*UserProto, error) {
	out := new(UserProto)
	err := c.cc.Invoke(ctx, "/protos.UserServiceProto/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceProtoClient) GetUserBy(ctx context.Context, in *GetUserByRequest, opts ...grpc.CallOption) (*UserProto, error) {
	out := new(UserProto)
	err := c.cc.Invoke(ctx, "/protos.UserServiceProto/GetUserBy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceProtoClient) GetUserById(ctx context.Context, in *GetUserByIdRequest, opts ...grpc.CallOption) (*UserProto, error) {
	out := new(UserProto)
	err := c.cc.Invoke(ctx, "/protos.UserServiceProto/GetUserById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceProtoClient) AutorizeToken(ctx context.Context, in *AuthorizeTokenRequest, opts ...grpc.CallOption) (*AuthorizeTokenResponse, error) {
	out := new(AuthorizeTokenResponse)
	err := c.cc.Invoke(ctx, "/protos.UserServiceProto/AutorizeToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceProtoServer is the server API for UserServiceProto service.
// All implementations must embed UnimplementedUserServiceProtoServer
// for forward compatibility
type UserServiceProtoServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*UserProto, error)
	GetUserBy(context.Context, *GetUserByRequest) (*UserProto, error)
	GetUserById(context.Context, *GetUserByIdRequest) (*UserProto, error)
	AutorizeToken(context.Context, *AuthorizeTokenRequest) (*AuthorizeTokenResponse, error)
	mustEmbedUnimplementedUserServiceProtoServer()
}

// UnimplementedUserServiceProtoServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceProtoServer struct {
}

func (UnimplementedUserServiceProtoServer) CreateUser(context.Context, *CreateUserRequest) (*UserProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceProtoServer) GetUserBy(context.Context, *GetUserByRequest) (*UserProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserBy not implemented")
}
func (UnimplementedUserServiceProtoServer) GetUserById(context.Context, *GetUserByIdRequest) (*UserProto, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserById not implemented")
}
func (UnimplementedUserServiceProtoServer) AutorizeToken(context.Context, *AuthorizeTokenRequest) (*AuthorizeTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AutorizeToken not implemented")
}
func (UnimplementedUserServiceProtoServer) mustEmbedUnimplementedUserServiceProtoServer() {}

// UnsafeUserServiceProtoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceProtoServer will
// result in compilation errors.
type UnsafeUserServiceProtoServer interface {
	mustEmbedUnimplementedUserServiceProtoServer()
}

func RegisterUserServiceProtoServer(s grpc.ServiceRegistrar, srv UserServiceProtoServer) {
	s.RegisterService(&UserServiceProto_ServiceDesc, srv)
}

func _UserServiceProto_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceProtoServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserServiceProto/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceProtoServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceProto_GetUserBy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceProtoServer).GetUserBy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserServiceProto/GetUserBy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceProtoServer).GetUserBy(ctx, req.(*GetUserByRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceProto_GetUserById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceProtoServer).GetUserById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserServiceProto/GetUserById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceProtoServer).GetUserById(ctx, req.(*GetUserByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServiceProto_AutorizeToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthorizeTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceProtoServer).AutorizeToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protos.UserServiceProto/AutorizeToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceProtoServer).AutorizeToken(ctx, req.(*AuthorizeTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServiceProto_ServiceDesc is the grpc.ServiceDesc for UserServiceProto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServiceProto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protos.UserServiceProto",
	HandlerType: (*UserServiceProtoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserServiceProto_CreateUser_Handler,
		},
		{
			MethodName: "GetUserBy",
			Handler:    _UserServiceProto_GetUserBy_Handler,
		},
		{
			MethodName: "GetUserById",
			Handler:    _UserServiceProto_GetUserById_Handler,
		},
		{
			MethodName: "AutorizeToken",
			Handler:    _UserServiceProto_AutorizeToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/user.proto",
}
