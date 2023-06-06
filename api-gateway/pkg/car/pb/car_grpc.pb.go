// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.1
// source: car.proto

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

// CarServiceClient is the client API for CarService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CarServiceClient interface {
	CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarResponse, error)
	ReadCar(ctx context.Context, in *ReadCarRequest, opts ...grpc.CallOption) (*ReadCarResponse, error)
	UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*UpdateCarResponse, error)
	DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarResponse, error)
	ListCar(ctx context.Context, in *ListCarRequest, opts ...grpc.CallOption) (CarService_ListCarClient, error)
}

type carServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCarServiceClient(cc grpc.ClientConnInterface) CarServiceClient {
	return &carServiceClient{cc}
}

func (c *carServiceClient) CreateCar(ctx context.Context, in *CreateCarRequest, opts ...grpc.CallOption) (*CreateCarResponse, error) {
	out := new(CreateCarResponse)
	err := c.cc.Invoke(ctx, "/car.CarService/CreateCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) ReadCar(ctx context.Context, in *ReadCarRequest, opts ...grpc.CallOption) (*ReadCarResponse, error) {
	out := new(ReadCarResponse)
	err := c.cc.Invoke(ctx, "/car.CarService/ReadCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) UpdateCar(ctx context.Context, in *UpdateCarRequest, opts ...grpc.CallOption) (*UpdateCarResponse, error) {
	out := new(UpdateCarResponse)
	err := c.cc.Invoke(ctx, "/car.CarService/UpdateCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) DeleteCar(ctx context.Context, in *DeleteCarRequest, opts ...grpc.CallOption) (*DeleteCarResponse, error) {
	out := new(DeleteCarResponse)
	err := c.cc.Invoke(ctx, "/car.CarService/DeleteCar", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *carServiceClient) ListCar(ctx context.Context, in *ListCarRequest, opts ...grpc.CallOption) (CarService_ListCarClient, error) {
	stream, err := c.cc.NewStream(ctx, &CarService_ServiceDesc.Streams[0], "/car.CarService/ListCar", opts...)
	if err != nil {
		return nil, err
	}
	x := &carServiceListCarClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CarService_ListCarClient interface {
	Recv() (*ListCarResponse, error)
	grpc.ClientStream
}

type carServiceListCarClient struct {
	grpc.ClientStream
}

func (x *carServiceListCarClient) Recv() (*ListCarResponse, error) {
	m := new(ListCarResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CarServiceServer is the server API for CarService service.
// All implementations must embed UnimplementedCarServiceServer
// for forward compatibility
type CarServiceServer interface {
	CreateCar(context.Context, *CreateCarRequest) (*CreateCarResponse, error)
	ReadCar(context.Context, *ReadCarRequest) (*ReadCarResponse, error)
	UpdateCar(context.Context, *UpdateCarRequest) (*UpdateCarResponse, error)
	DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarResponse, error)
	ListCar(*ListCarRequest, CarService_ListCarServer) error
	mustEmbedUnimplementedCarServiceServer()
}

// UnimplementedCarServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCarServiceServer struct {
}

func (UnimplementedCarServiceServer) CreateCar(context.Context, *CreateCarRequest) (*CreateCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCar not implemented")
}
func (UnimplementedCarServiceServer) ReadCar(context.Context, *ReadCarRequest) (*ReadCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCar not implemented")
}
func (UnimplementedCarServiceServer) UpdateCar(context.Context, *UpdateCarRequest) (*UpdateCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCar not implemented")
}
func (UnimplementedCarServiceServer) DeleteCar(context.Context, *DeleteCarRequest) (*DeleteCarResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCar not implemented")
}
func (UnimplementedCarServiceServer) ListCar(*ListCarRequest, CarService_ListCarServer) error {
	return status.Errorf(codes.Unimplemented, "method ListCar not implemented")
}
func (UnimplementedCarServiceServer) mustEmbedUnimplementedCarServiceServer() {}

// UnsafeCarServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CarServiceServer will
// result in compilation errors.
type UnsafeCarServiceServer interface {
	mustEmbedUnimplementedCarServiceServer()
}

func RegisterCarServiceServer(s grpc.ServiceRegistrar, srv CarServiceServer) {
	s.RegisterService(&CarService_ServiceDesc, srv)
}

func _CarService_CreateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).CreateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/CreateCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).CreateCar(ctx, req.(*CreateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_ReadCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).ReadCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/ReadCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).ReadCar(ctx, req.(*ReadCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_UpdateCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).UpdateCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/UpdateCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).UpdateCar(ctx, req.(*UpdateCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_DeleteCar_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCarRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CarServiceServer).DeleteCar(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/car.CarService/DeleteCar",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CarServiceServer).DeleteCar(ctx, req.(*DeleteCarRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CarService_ListCar_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ListCarRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CarServiceServer).ListCar(m, &carServiceListCarServer{stream})
}

type CarService_ListCarServer interface {
	Send(*ListCarResponse) error
	grpc.ServerStream
}

type carServiceListCarServer struct {
	grpc.ServerStream
}

func (x *carServiceListCarServer) Send(m *ListCarResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CarService_ServiceDesc is the grpc.ServiceDesc for CarService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CarService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "car.CarService",
	HandlerType: (*CarServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCar",
			Handler:    _CarService_CreateCar_Handler,
		},
		{
			MethodName: "ReadCar",
			Handler:    _CarService_ReadCar_Handler,
		},
		{
			MethodName: "UpdateCar",
			Handler:    _CarService_UpdateCar_Handler,
		},
		{
			MethodName: "DeleteCar",
			Handler:    _CarService_DeleteCar_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ListCar",
			Handler:       _CarService_ListCar_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "car.proto",
}