// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/backend/result.proto

package result

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

// ResultStoreClient is the client API for ResultStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResultStoreClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	SetTTL(ctx context.Context, in *SetTTLRequest, opts ...grpc.CallOption) (*SetTTLResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error)
}

type resultStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewResultStoreClient(cc grpc.ClientConnInterface) ResultStoreClient {
	return &resultStoreClient{cc}
}

func (c *resultStoreClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/result.ResultStore/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultStoreClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, "/result.ResultStore/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultStoreClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, "/result.ResultStore/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultStoreClient) SetTTL(ctx context.Context, in *SetTTLRequest, opts ...grpc.CallOption) (*SetTTLResponse, error) {
	out := new(SetTTLResponse)
	err := c.cc.Invoke(ctx, "/result.ResultStore/SetTTL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/result.ResultStore/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultStoreClient) Random(ctx context.Context, in *RandomRequest, opts ...grpc.CallOption) (*RandomResponse, error) {
	out := new(RandomResponse)
	err := c.cc.Invoke(ctx, "/result.ResultStore/Random", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultStoreServer is the server API for ResultStore service.
// All implementations must embed UnimplementedResultStoreServer
// for forward compatibility
type ResultStoreServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	SetTTL(context.Context, *SetTTLRequest) (*SetTTLResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Random(context.Context, *RandomRequest) (*RandomResponse, error)
	mustEmbedUnimplementedResultStoreServer()
}

// UnimplementedResultStoreServer must be embedded to have forward compatible implementations.
type UnimplementedResultStoreServer struct {
}

func (UnimplementedResultStoreServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResultStoreServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedResultStoreServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResultStoreServer) SetTTL(context.Context, *SetTTLRequest) (*SetTTLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetTTL not implemented")
}
func (UnimplementedResultStoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResultStoreServer) Random(context.Context, *RandomRequest) (*RandomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Random not implemented")
}
func (UnimplementedResultStoreServer) mustEmbedUnimplementedResultStoreServer() {}

// UnsafeResultStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResultStoreServer will
// result in compilation errors.
type UnsafeResultStoreServer interface {
	mustEmbedUnimplementedResultStoreServer()
}

func RegisterResultStoreServer(s grpc.ServiceRegistrar, srv ResultStoreServer) {
	s.RegisterService(&ResultStore_ServiceDesc, srv)
}

func _ResultStore_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultStoreServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultStore/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultStoreServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultStoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultStore/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultStoreServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultStore_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultStoreServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultStore/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultStoreServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultStore_SetTTL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetTTLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultStoreServer).SetTTL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultStore/SetTTL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultStoreServer).SetTTL(ctx, req.(*SetTTLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultStore/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResultStore_Random_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RandomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultStoreServer).Random(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/result.ResultStore/Random",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultStoreServer).Random(ctx, req.(*RandomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResultStore_ServiceDesc is the grpc.ServiceDesc for ResultStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResultStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "result.ResultStore",
	HandlerType: (*ResultStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _ResultStore_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _ResultStore_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResultStore_Update_Handler,
		},
		{
			MethodName: "SetTTL",
			Handler:    _ResultStore_SetTTL_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResultStore_Delete_Handler,
		},
		{
			MethodName: "Random",
			Handler:    _ResultStore_Random_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/backend/result.proto",
}
