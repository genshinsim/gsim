// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: protos/backend/preview.proto

package preview

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
	Embed_Get_FullMethodName = "/preview.Embed/Get"
)

// EmbedClient is the client API for Embed service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmbedClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type embedClient struct {
	cc grpc.ClientConnInterface
}

func NewEmbedClient(cc grpc.ClientConnInterface) EmbedClient {
	return &embedClient{cc}
}

func (c *embedClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, Embed_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmbedServer is the server API for Embed service.
// All implementations must embed UnimplementedEmbedServer
// for forward compatibility
type EmbedServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedEmbedServer()
}

// UnimplementedEmbedServer must be embedded to have forward compatible implementations.
type UnimplementedEmbedServer struct {
}

func (UnimplementedEmbedServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedEmbedServer) mustEmbedUnimplementedEmbedServer() {}

// UnsafeEmbedServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmbedServer will
// result in compilation errors.
type UnsafeEmbedServer interface {
	mustEmbedUnimplementedEmbedServer()
}

func RegisterEmbedServer(s grpc.ServiceRegistrar, srv EmbedServer) {
	s.RegisterService(&Embed_ServiceDesc, srv)
}

func _Embed_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmbedServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Embed_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmbedServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Embed_ServiceDesc is the grpc.ServiceDesc for Embed service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Embed_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "preview.Embed",
	HandlerType: (*EmbedServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _Embed_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/backend/preview.proto",
}
