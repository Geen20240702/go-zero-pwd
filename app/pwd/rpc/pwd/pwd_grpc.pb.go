// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.19.4
// source: pwd.proto

package pwd

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Pwd_CreatePwd_FullMethodName  = "/pwd.Pwd/CreatePwd"
	Pwd_UpdatePwd_FullMethodName  = "/pwd.Pwd/UpdatePwd"
	Pwd_GetPwdList_FullMethodName = "/pwd.Pwd/GetPwdList"
	Pwd_DelPwd_FullMethodName     = "/pwd.Pwd/DelPwd"
)

// PwdClient is the client API for Pwd service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PwdClient interface {
	CreatePwd(ctx context.Context, in *CreatePwdReq, opts ...grpc.CallOption) (*CreatePwdResp, error)
	UpdatePwd(ctx context.Context, in *UpdatePwdReq, opts ...grpc.CallOption) (*UpdatePwdResp, error)
	GetPwdList(ctx context.Context, in *GetPwdListReq, opts ...grpc.CallOption) (*GetPwdListResp, error)
	DelPwd(ctx context.Context, in *DelPwdReq, opts ...grpc.CallOption) (*DelPwdResp, error)
}

type pwdClient struct {
	cc grpc.ClientConnInterface
}

func NewPwdClient(cc grpc.ClientConnInterface) PwdClient {
	return &pwdClient{cc}
}

func (c *pwdClient) CreatePwd(ctx context.Context, in *CreatePwdReq, opts ...grpc.CallOption) (*CreatePwdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreatePwdResp)
	err := c.cc.Invoke(ctx, Pwd_CreatePwd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pwdClient) UpdatePwd(ctx context.Context, in *UpdatePwdReq, opts ...grpc.CallOption) (*UpdatePwdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdatePwdResp)
	err := c.cc.Invoke(ctx, Pwd_UpdatePwd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pwdClient) GetPwdList(ctx context.Context, in *GetPwdListReq, opts ...grpc.CallOption) (*GetPwdListResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetPwdListResp)
	err := c.cc.Invoke(ctx, Pwd_GetPwdList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pwdClient) DelPwd(ctx context.Context, in *DelPwdReq, opts ...grpc.CallOption) (*DelPwdResp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DelPwdResp)
	err := c.cc.Invoke(ctx, Pwd_DelPwd_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PwdServer is the server API for Pwd service.
// All implementations must embed UnimplementedPwdServer
// for forward compatibility
type PwdServer interface {
	CreatePwd(context.Context, *CreatePwdReq) (*CreatePwdResp, error)
	UpdatePwd(context.Context, *UpdatePwdReq) (*UpdatePwdResp, error)
	GetPwdList(context.Context, *GetPwdListReq) (*GetPwdListResp, error)
	DelPwd(context.Context, *DelPwdReq) (*DelPwdResp, error)
	mustEmbedUnimplementedPwdServer()
}

// UnimplementedPwdServer must be embedded to have forward compatible implementations.
type UnimplementedPwdServer struct {
}

func (UnimplementedPwdServer) CreatePwd(context.Context, *CreatePwdReq) (*CreatePwdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePwd not implemented")
}
func (UnimplementedPwdServer) UpdatePwd(context.Context, *UpdatePwdReq) (*UpdatePwdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePwd not implemented")
}
func (UnimplementedPwdServer) GetPwdList(context.Context, *GetPwdListReq) (*GetPwdListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPwdList not implemented")
}
func (UnimplementedPwdServer) DelPwd(context.Context, *DelPwdReq) (*DelPwdResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DelPwd not implemented")
}
func (UnimplementedPwdServer) mustEmbedUnimplementedPwdServer() {}

// UnsafePwdServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PwdServer will
// result in compilation errors.
type UnsafePwdServer interface {
	mustEmbedUnimplementedPwdServer()
}

func RegisterPwdServer(s grpc.ServiceRegistrar, srv PwdServer) {
	s.RegisterService(&Pwd_ServiceDesc, srv)
}

func _Pwd_CreatePwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePwdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PwdServer).CreatePwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pwd_CreatePwd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PwdServer).CreatePwd(ctx, req.(*CreatePwdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pwd_UpdatePwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePwdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PwdServer).UpdatePwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pwd_UpdatePwd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PwdServer).UpdatePwd(ctx, req.(*UpdatePwdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pwd_GetPwdList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPwdListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PwdServer).GetPwdList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pwd_GetPwdList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PwdServer).GetPwdList(ctx, req.(*GetPwdListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Pwd_DelPwd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DelPwdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PwdServer).DelPwd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Pwd_DelPwd_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PwdServer).DelPwd(ctx, req.(*DelPwdReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Pwd_ServiceDesc is the grpc.ServiceDesc for Pwd service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Pwd_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pwd.Pwd",
	HandlerType: (*PwdServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePwd",
			Handler:    _Pwd_CreatePwd_Handler,
		},
		{
			MethodName: "UpdatePwd",
			Handler:    _Pwd_UpdatePwd_Handler,
		},
		{
			MethodName: "GetPwdList",
			Handler:    _Pwd_GetPwdList_Handler,
		},
		{
			MethodName: "DelPwd",
			Handler:    _Pwd_DelPwd_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pwd.proto",
}
