// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package registration

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// VehicleServiceClient is the client API for VehicleService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VehicleServiceClient interface {
	FindByNumber(ctx context.Context, in *FindByNumberRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	FindByVIN(ctx context.Context, in *FindByVINRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
	FindByCode(ctx context.Context, in *FindByCodeRequest, opts ...grpc.CallOption) (*RegistrationResponse, error)
}

type vehicleServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVehicleServiceClient(cc grpc.ClientConnInterface) VehicleServiceClient {
	return &vehicleServiceClient{cc}
}

func (c *vehicleServiceClient) FindByNumber(ctx context.Context, in *FindByNumberRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/registration.VehicleService/FindByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) FindByVIN(ctx context.Context, in *FindByVINRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/registration.VehicleService/FindByVIN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vehicleServiceClient) FindByCode(ctx context.Context, in *FindByCodeRequest, opts ...grpc.CallOption) (*RegistrationResponse, error) {
	out := new(RegistrationResponse)
	err := c.cc.Invoke(ctx, "/registration.VehicleService/FindByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VehicleServiceServer is the server API for VehicleService service.
// All implementations must embed UnimplementedVehicleServiceServer
// for forward compatibility
type VehicleServiceServer interface {
	FindByNumber(context.Context, *FindByNumberRequest) (*RegistrationResponse, error)
	FindByVIN(context.Context, *FindByVINRequest) (*RegistrationResponse, error)
	FindByCode(context.Context, *FindByCodeRequest) (*RegistrationResponse, error)
	mustEmbedUnimplementedVehicleServiceServer()
}

// UnimplementedVehicleServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVehicleServiceServer struct {
}

func (*UnimplementedVehicleServiceServer) FindByNumber(context.Context, *FindByNumberRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByNumber not implemented")
}
func (*UnimplementedVehicleServiceServer) FindByVIN(context.Context, *FindByVINRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByVIN not implemented")
}
func (*UnimplementedVehicleServiceServer) FindByCode(context.Context, *FindByCodeRequest) (*RegistrationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCode not implemented")
}
func (*UnimplementedVehicleServiceServer) mustEmbedUnimplementedVehicleServiceServer() {}

func RegisterVehicleServiceServer(s *grpc.Server, srv VehicleServiceServer) {
	s.RegisterService(&_VehicleService_serviceDesc, srv)
}

func _VehicleService_FindByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByNumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).FindByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registration.VehicleService/FindByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).FindByNumber(ctx, req.(*FindByNumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_FindByVIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByVINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).FindByVIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registration.VehicleService/FindByVIN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).FindByVIN(ctx, req.(*FindByVINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VehicleService_FindByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindByCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VehicleServiceServer).FindByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registration.VehicleService/FindByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VehicleServiceServer).FindByCode(ctx, req.(*FindByCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VehicleService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registration.VehicleService",
	HandlerType: (*VehicleServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByNumber",
			Handler:    _VehicleService_FindByNumber_Handler,
		},
		{
			MethodName: "FindByVIN",
			Handler:    _VehicleService_FindByVIN_Handler,
		},
		{
			MethodName: "FindByCode",
			Handler:    _VehicleService_FindByCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/registration/service.proto",
}
