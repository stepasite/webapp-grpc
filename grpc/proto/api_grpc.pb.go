// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_api

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

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	// unary call
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/api.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	// unary call
	Login(context.Context, *LoginRequest) (*LoginReply, error)
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) Login(context.Context, *LoginRequest) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/api.proto",
}

// StreamingNewsServiceClient is the client API for StreamingNewsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingNewsServiceClient interface {
	// server streaming call
	StreamNews(ctx context.Context, in *RepeatNewsRequest, opts ...grpc.CallOption) (StreamingNewsService_StreamNewsClient, error)
}

type streamingNewsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingNewsServiceClient(cc grpc.ClientConnInterface) StreamingNewsServiceClient {
	return &streamingNewsServiceClient{cc}
}

func (c *streamingNewsServiceClient) StreamNews(ctx context.Context, in *RepeatNewsRequest, opts ...grpc.CallOption) (StreamingNewsService_StreamNewsClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingNewsService_ServiceDesc.Streams[0], "/api.StreamingNewsService/StreamNews", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingNewsServiceStreamNewsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type StreamingNewsService_StreamNewsClient interface {
	Recv() (*NewsReply, error)
	grpc.ClientStream
}

type streamingNewsServiceStreamNewsClient struct {
	grpc.ClientStream
}

func (x *streamingNewsServiceStreamNewsClient) Recv() (*NewsReply, error) {
	m := new(NewsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingNewsServiceServer is the server API for StreamingNewsService service.
// All implementations must embed UnimplementedStreamingNewsServiceServer
// for forward compatibility
type StreamingNewsServiceServer interface {
	// server streaming call
	StreamNews(*RepeatNewsRequest, StreamingNewsService_StreamNewsServer) error
	mustEmbedUnimplementedStreamingNewsServiceServer()
}

// UnimplementedStreamingNewsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingNewsServiceServer struct {
}

func (UnimplementedStreamingNewsServiceServer) StreamNews(*RepeatNewsRequest, StreamingNewsService_StreamNewsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamNews not implemented")
}
func (UnimplementedStreamingNewsServiceServer) mustEmbedUnimplementedStreamingNewsServiceServer() {}

// UnsafeStreamingNewsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingNewsServiceServer will
// result in compilation errors.
type UnsafeStreamingNewsServiceServer interface {
	mustEmbedUnimplementedStreamingNewsServiceServer()
}

func RegisterStreamingNewsServiceServer(s grpc.ServiceRegistrar, srv StreamingNewsServiceServer) {
	s.RegisterService(&StreamingNewsService_ServiceDesc, srv)
}

func _StreamingNewsService_StreamNews_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RepeatNewsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamingNewsServiceServer).StreamNews(m, &streamingNewsServiceStreamNewsServer{stream})
}

type StreamingNewsService_StreamNewsServer interface {
	Send(*NewsReply) error
	grpc.ServerStream
}

type streamingNewsServiceStreamNewsServer struct {
	grpc.ServerStream
}

func (x *streamingNewsServiceStreamNewsServer) Send(m *NewsReply) error {
	return x.ServerStream.SendMsg(m)
}

// StreamingNewsService_ServiceDesc is the grpc.ServiceDesc for StreamingNewsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingNewsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.StreamingNewsService",
	HandlerType: (*StreamingNewsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamNews",
			Handler:       _StreamingNewsService_StreamNews_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/api.proto",
}