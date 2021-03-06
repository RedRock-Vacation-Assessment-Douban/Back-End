// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: changePW/changePW.proto

package changePW

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

// ChangePWClient is the client API for ChangePW service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChangePWClient interface {
	ChangePW(ctx context.Context, in *ChangeReq, opts ...grpc.CallOption) (*ChangeRes, error)
}

type changePWClient struct {
	cc grpc.ClientConnInterface
}

func NewChangePWClient(cc grpc.ClientConnInterface) ChangePWClient {
	return &changePWClient{cc}
}

func (c *changePWClient) ChangePW(ctx context.Context, in *ChangeReq, opts ...grpc.CallOption) (*ChangeRes, error) {
	out := new(ChangeRes)
	err := c.cc.Invoke(ctx, "/changePW.changePW/changePW", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChangePWServer is the server API for ChangePW service.
// All implementations must embed UnimplementedChangePWServer
// for forward compatibility
type ChangePWServer interface {
	ChangePW(context.Context, *ChangeReq) (*ChangeRes, error)
	mustEmbedUnimplementedChangePWServer()
}

// UnimplementedChangePWServer must be embedded to have forward compatible implementations.
type UnimplementedChangePWServer struct {
}

func (UnimplementedChangePWServer) ChangePW(context.Context, *ChangeReq) (*ChangeRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangePW not implemented")
}
func (UnimplementedChangePWServer) mustEmbedUnimplementedChangePWServer() {}

// UnsafeChangePWServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChangePWServer will
// result in compilation errors.
type UnsafeChangePWServer interface {
	mustEmbedUnimplementedChangePWServer()
}

func RegisterChangePWServer(s grpc.ServiceRegistrar, srv ChangePWServer) {
	s.RegisterService(&ChangePW_ServiceDesc, srv)
}

func _ChangePW_ChangePW_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChangePWServer).ChangePW(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/changePW.changePW/changePW",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChangePWServer).ChangePW(ctx, req.(*ChangeReq))
	}
	return interceptor(ctx, in, info, handler)
}

// ChangePW_ServiceDesc is the grpc.ServiceDesc for ChangePW service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChangePW_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "changePW.changePW",
	HandlerType: (*ChangePWServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "changePW",
			Handler:    _ChangePW_ChangePW_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "changePW/changePW.proto",
}
