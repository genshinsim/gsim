// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: protos/backend/compute.proto

package compute

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

// ComputeClient is the client API for Compute service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ComputeClient interface {
	Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error)
}

type computeClient struct {
	cc grpc.ClientConnInterface
}

func NewComputeClient(cc grpc.ClientConnInterface) ComputeClient {
	return &computeClient{cc}
}

func (c *computeClient) Run(ctx context.Context, in *RunRequest, opts ...grpc.CallOption) (*RunResponse, error) {
	out := new(RunResponse)
	err := c.cc.Invoke(ctx, "/compute.Compute/Run", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ComputeServer is the server API for Compute service.
// All implementations must embed UnimplementedComputeServer
// for forward compatibility
type ComputeServer interface {
	Run(context.Context, *RunRequest) (*RunResponse, error)
	mustEmbedUnimplementedComputeServer()
}

// UnimplementedComputeServer must be embedded to have forward compatible implementations.
type UnimplementedComputeServer struct {
}

func (UnimplementedComputeServer) Run(context.Context, *RunRequest) (*RunResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Run not implemented")
}
func (UnimplementedComputeServer) mustEmbedUnimplementedComputeServer() {}

// UnsafeComputeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ComputeServer will
// result in compilation errors.
type UnsafeComputeServer interface {
	mustEmbedUnimplementedComputeServer()
}

func RegisterComputeServer(s grpc.ServiceRegistrar, srv ComputeServer) {
	s.RegisterService(&Compute_ServiceDesc, srv)
}

func _Compute_Run_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ComputeServer).Run(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/compute.Compute/Run",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ComputeServer).Run(ctx, req.(*RunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Compute_ServiceDesc is the grpc.ServiceDesc for Compute service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Compute_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "compute.Compute",
	HandlerType: (*ComputeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Run",
			Handler:    _Compute_Run_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/backend/compute.proto",
}
