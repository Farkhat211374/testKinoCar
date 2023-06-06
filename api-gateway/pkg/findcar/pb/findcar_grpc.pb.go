// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: findcar.proto

package pb

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

// FindcarServiceClient is the client API for FindcarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FindcarServiceClient interface {
	CreateFindcar(ctx context.Context, in *CreateFindcarRequest, opts ...grpc.CallOption) (*CreateFindcarResponse, error)
	ReadFindcar(ctx context.Context, in *ReadFindcarRequest, opts ...grpc.CallOption) (*ReadFindcarResponse, error)
	UpdateFindcar(ctx context.Context, in *UpdateFindcarRequest, opts ...grpc.CallOption) (*UpdateFindcarResponse, error)
	DeleteFindcar(ctx context.Context, in *DeleteFindcarRequest, opts ...grpc.CallOption) (*DeleteFindcarResponse, error)
}

type findcarServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFindcarServiceClient(cc grpc.ClientConnInterface) FindcarServiceClient {
	return &findcarServiceClient{cc}
}

func (c *findcarServiceClient) CreateFindcar(ctx context.Context, in *CreateFindcarRequest, opts ...grpc.CallOption) (*CreateFindcarResponse, error) {
	out := new(CreateFindcarResponse)
	err := c.cc.Invoke(ctx, "/findcar.FindcarService/CreateFindcar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *findcarServiceClient) ReadFindcar(ctx context.Context, in *ReadFindcarRequest, opts ...grpc.CallOption) (*ReadFindcarResponse, error) {
	out := new(ReadFindcarResponse)
	err := c.cc.Invoke(ctx, "/findcar.FindcarService/ReadFindcar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *findcarServiceClient) UpdateFindcar(ctx context.Context, in *UpdateFindcarRequest, opts ...grpc.CallOption) (*UpdateFindcarResponse, error) {
	out := new(UpdateFindcarResponse)
	err := c.cc.Invoke(ctx, "/findcar.FindcarService/UpdateFindcar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *findcarServiceClient) DeleteFindcar(ctx context.Context, in *DeleteFindcarRequest, opts ...grpc.CallOption) (*DeleteFindcarResponse, error) {
	out := new(DeleteFindcarResponse)
	err := c.cc.Invoke(ctx, "/findcar.FindcarService/DeleteFindcar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FindcarServiceServer is the server API for FindcarService service.
// All implementations must embed UnimplementedFindcarServiceServer
// for forward compatibility
type FindcarServiceServer interface {
	CreateFindcar(context.Context, *CreateFindcarRequest) (*CreateFindcarResponse, error)
	ReadFindcar(context.Context, *ReadFindcarRequest) (*ReadFindcarResponse, error)
	UpdateFindcar(context.Context, *UpdateFindcarRequest) (*UpdateFindcarResponse, error)
	DeleteFindcar(context.Context, *DeleteFindcarRequest) (*DeleteFindcarResponse, error)
	mustEmbedUnimplementedFindcarServiceServer()
}

// UnimplementedFindcarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFindcarServiceServer struct {
}

func (UnimplementedFindcarServiceServer) CreateFindcar(context.Context, *CreateFindcarRequest) (*CreateFindcarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFindcar not implemented")
}
func (UnimplementedFindcarServiceServer) ReadFindcar(context.Context, *ReadFindcarRequest) (*ReadFindcarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadFindcar not implemented")
}
func (UnimplementedFindcarServiceServer) UpdateFindcar(context.Context, *UpdateFindcarRequest) (*UpdateFindcarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFindcar not implemented")
}
func (UnimplementedFindcarServiceServer) DeleteFindcar(context.Context, *DeleteFindcarRequest) (*DeleteFindcarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFindcar not implemented")
}
func (UnimplementedFindcarServiceServer) mustEmbedUnimplementedFindcarServiceServer() {}

// UnsafeFindcarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FindcarServiceServer will
// result in compilation errors.
type UnsafeFindcarServiceServer interface {
	mustEmbedUnimplementedFindcarServiceServer()
}

func RegisterFindcarServiceServer(s grpc.ServiceRegistrar, srv FindcarServiceServer) {
	s.RegisterService(&FindcarService_ServiceDesc, srv)
}

func _FindcarService_CreateFindcar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFindcarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FindcarServiceServer).CreateFindcar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/findcar.FindcarService/CreateFindcar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FindcarServiceServer).CreateFindcar(ctx, req.(*CreateFindcarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FindcarService_ReadFindcar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadFindcarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FindcarServiceServer).ReadFindcar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/findcar.FindcarService/ReadFindcar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FindcarServiceServer).ReadFindcar(ctx, req.(*ReadFindcarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FindcarService_UpdateFindcar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateFindcarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FindcarServiceServer).UpdateFindcar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/findcar.FindcarService/UpdateFindcar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FindcarServiceServer).UpdateFindcar(ctx, req.(*UpdateFindcarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FindcarService_DeleteFindcar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFindcarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FindcarServiceServer).DeleteFindcar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/findcar.FindcarService/DeleteFindcar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FindcarServiceServer).DeleteFindcar(ctx, req.(*DeleteFindcarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FindcarService_ServiceDesc is the grpc.ServiceDesc for FindcarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FindcarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "findcar.FindcarService",
	HandlerType: (*FindcarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFindcar",
			Handler:    _FindcarService_CreateFindcar_Handler,
		},
		{
			MethodName: "ReadFindcar",
			Handler:    _FindcarService_ReadFindcar_Handler,
		},
		{
			MethodName: "UpdateFindcar",
			Handler:    _FindcarService_UpdateFindcar_Handler,
		},
		{
			MethodName: "DeleteFindcar",
			Handler:    _FindcarService_DeleteFindcar_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "findcar.proto",
}
