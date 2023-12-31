// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: user.proto

package user

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
	User_Register_FullMethodName     = "/user.user/Register"
	User_Login_FullMethodName        = "/user.user/Login"
	User_UserInfo_FullMethodName     = "/user.user/UserInfo"
	User_AddCount_FullMethodName     = "/user.user/AddCount"
	User_UserIsExists_FullMethodName = "/user.user/UserIsExists"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error)
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error)
	UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error)
	AddCount(ctx context.Context, in *AddCountReq, opts ...grpc.CallOption) (*AddCountResp, error)
	UserIsExists(ctx context.Context, in *UserIsExistsReq, opts ...grpc.CallOption) (*UserIsExistsResp, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterReq, opts ...grpc.CallOption) (*RegisterResp, error) {
	out := new(RegisterResp)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginResp, error) {
	out := new(LoginResp)
	err := c.cc.Invoke(ctx, User_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserInfo(ctx context.Context, in *UserInfoReq, opts ...grpc.CallOption) (*UserInfoResp, error) {
	out := new(UserInfoResp)
	err := c.cc.Invoke(ctx, User_UserInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) AddCount(ctx context.Context, in *AddCountReq, opts ...grpc.CallOption) (*AddCountResp, error) {
	out := new(AddCountResp)
	err := c.cc.Invoke(ctx, User_AddCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) UserIsExists(ctx context.Context, in *UserIsExistsReq, opts ...grpc.CallOption) (*UserIsExistsResp, error) {
	out := new(UserIsExistsResp)
	err := c.cc.Invoke(ctx, User_UserIsExists_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *RegisterReq) (*RegisterResp, error)
	Login(context.Context, *LoginReq) (*LoginResp, error)
	UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error)
	AddCount(context.Context, *AddCountReq) (*AddCountResp, error)
	UserIsExists(context.Context, *UserIsExistsReq) (*UserIsExistsResp, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *RegisterReq) (*RegisterResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginReq) (*LoginResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) UserInfo(context.Context, *UserInfoReq) (*UserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserInfo not implemented")
}
func (UnimplementedUserServer) AddCount(context.Context, *AddCountReq) (*AddCountResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCount not implemented")
}
func (UnimplementedUserServer) UserIsExists(context.Context, *UserIsExistsReq) (*UserIsExistsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserIsExists not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserInfo(ctx, req.(*UserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_AddCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).AddCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_AddCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).AddCount(ctx, req.(*AddCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_UserIsExists_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserIsExistsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).UserIsExists(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_UserIsExists_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).UserIsExists(ctx, req.(*UserIsExistsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.user",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "UserInfo",
			Handler:    _User_UserInfo_Handler,
		},
		{
			MethodName: "AddCount",
			Handler:    _User_AddCount_Handler,
		},
		{
			MethodName: "UserIsExists",
			Handler:    _User_UserIsExists_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
