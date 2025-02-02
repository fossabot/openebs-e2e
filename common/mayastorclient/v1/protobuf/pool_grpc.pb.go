// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: pool.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PoolRpcClient is the client API for PoolRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PoolRpcClient interface {
	// Storage pool related methods.
	//
	// Storage pool is made up of block devices disks and provides a storage
	// space for provisioning of replicas.
	CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*Pool, error)
	DestroyPool(ctx context.Context, in *DestroyPoolRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ExportPool(ctx context.Context, in *ExportPoolRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ImportPool(ctx context.Context, in *ImportPoolRequest, opts ...grpc.CallOption) (*Pool, error)
	ListPools(ctx context.Context, in *ListPoolOptions, opts ...grpc.CallOption) (*ListPoolsResponse, error)
}

type poolRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewPoolRpcClient(cc grpc.ClientConnInterface) PoolRpcClient {
	return &poolRpcClient{cc}
}

func (c *poolRpcClient) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*Pool, error) {
	out := new(Pool)
	err := c.cc.Invoke(ctx, "/mayastor.v1.PoolRpc/CreatePool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolRpcClient) DestroyPool(ctx context.Context, in *DestroyPoolRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mayastor.v1.PoolRpc/DestroyPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolRpcClient) ExportPool(ctx context.Context, in *ExportPoolRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/mayastor.v1.PoolRpc/ExportPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolRpcClient) ImportPool(ctx context.Context, in *ImportPoolRequest, opts ...grpc.CallOption) (*Pool, error) {
	out := new(Pool)
	err := c.cc.Invoke(ctx, "/mayastor.v1.PoolRpc/ImportPool", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *poolRpcClient) ListPools(ctx context.Context, in *ListPoolOptions, opts ...grpc.CallOption) (*ListPoolsResponse, error) {
	out := new(ListPoolsResponse)
	err := c.cc.Invoke(ctx, "/mayastor.v1.PoolRpc/ListPools", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PoolRpcServer is the server API for PoolRpc service.
// All implementations must embed UnimplementedPoolRpcServer
// for forward compatibility
type PoolRpcServer interface {
	// Storage pool related methods.
	//
	// Storage pool is made up of block devices disks and provides a storage
	// space for provisioning of replicas.
	CreatePool(context.Context, *CreatePoolRequest) (*Pool, error)
	DestroyPool(context.Context, *DestroyPoolRequest) (*emptypb.Empty, error)
	ExportPool(context.Context, *ExportPoolRequest) (*emptypb.Empty, error)
	ImportPool(context.Context, *ImportPoolRequest) (*Pool, error)
	ListPools(context.Context, *ListPoolOptions) (*ListPoolsResponse, error)
	mustEmbedUnimplementedPoolRpcServer()
}

// UnimplementedPoolRpcServer must be embedded to have forward compatible implementations.
type UnimplementedPoolRpcServer struct {
}

func (UnimplementedPoolRpcServer) CreatePool(context.Context, *CreatePoolRequest) (*Pool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedPoolRpcServer) DestroyPool(context.Context, *DestroyPoolRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DestroyPool not implemented")
}
func (UnimplementedPoolRpcServer) ExportPool(context.Context, *ExportPoolRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExportPool not implemented")
}
func (UnimplementedPoolRpcServer) ImportPool(context.Context, *ImportPoolRequest) (*Pool, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportPool not implemented")
}
func (UnimplementedPoolRpcServer) ListPools(context.Context, *ListPoolOptions) (*ListPoolsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPools not implemented")
}
func (UnimplementedPoolRpcServer) mustEmbedUnimplementedPoolRpcServer() {}

// UnsafePoolRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PoolRpcServer will
// result in compilation errors.
type UnsafePoolRpcServer interface {
	mustEmbedUnimplementedPoolRpcServer()
}

func RegisterPoolRpcServer(s grpc.ServiceRegistrar, srv PoolRpcServer) {
	s.RegisterService(&PoolRpc_ServiceDesc, srv)
}

func _PoolRpc_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolRpcServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mayastor.v1.PoolRpc/CreatePool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolRpcServer).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolRpc_DestroyPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DestroyPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolRpcServer).DestroyPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mayastor.v1.PoolRpc/DestroyPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolRpcServer).DestroyPool(ctx, req.(*DestroyPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolRpc_ExportPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExportPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolRpcServer).ExportPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mayastor.v1.PoolRpc/ExportPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolRpcServer).ExportPool(ctx, req.(*ExportPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolRpc_ImportPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImportPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolRpcServer).ImportPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mayastor.v1.PoolRpc/ImportPool",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolRpcServer).ImportPool(ctx, req.(*ImportPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PoolRpc_ListPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoolOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PoolRpcServer).ListPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/mayastor.v1.PoolRpc/ListPools",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PoolRpcServer).ListPools(ctx, req.(*ListPoolOptions))
	}
	return interceptor(ctx, in, info, handler)
}

// PoolRpc_ServiceDesc is the grpc.ServiceDesc for PoolRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PoolRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mayastor.v1.PoolRpc",
	HandlerType: (*PoolRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePool",
			Handler:    _PoolRpc_CreatePool_Handler,
		},
		{
			MethodName: "DestroyPool",
			Handler:    _PoolRpc_DestroyPool_Handler,
		},
		{
			MethodName: "ExportPool",
			Handler:    _PoolRpc_ExportPool_Handler,
		},
		{
			MethodName: "ImportPool",
			Handler:    _PoolRpc_ImportPool_Handler,
		},
		{
			MethodName: "ListPools",
			Handler:    _PoolRpc_ListPools_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pool.proto",
}
