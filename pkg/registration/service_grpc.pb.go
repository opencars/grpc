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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	FindByNumber(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*Response, error)
	FindByVIN(ctx context.Context, in *VINRequest, opts ...grpc.CallOption) (*Response, error)
	FindByCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*Response, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) FindByNumber(ctx context.Context, in *NumberRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/registration.Service/FindByNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindByVIN(ctx context.Context, in *VINRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/registration.Service/FindByVIN", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) FindByCode(ctx context.Context, in *CodeRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/registration.Service/FindByCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	FindByNumber(context.Context, *NumberRequest) (*Response, error)
	FindByVIN(context.Context, *VINRequest) (*Response, error)
	FindByCode(context.Context, *CodeRequest) (*Response, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (*UnimplementedServiceServer) FindByNumber(context.Context, *NumberRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByNumber not implemented")
}
func (*UnimplementedServiceServer) FindByVIN(context.Context, *VINRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByVIN not implemented")
}
func (*UnimplementedServiceServer) FindByCode(context.Context, *CodeRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindByCode not implemented")
}
func (*UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

func RegisterServiceServer(s *grpc.Server, srv ServiceServer) {
	s.RegisterService(&_Service_serviceDesc, srv)
}

func _Service_FindByNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NumberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindByNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registration.Service/FindByNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindByNumber(ctx, req.(*NumberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindByVIN_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VINRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindByVIN(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registration.Service/FindByVIN",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindByVIN(ctx, req.(*VINRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_FindByCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).FindByCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registration.Service/FindByCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).FindByCode(ctx, req.(*CodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Service_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registration.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindByNumber",
			Handler:    _Service_FindByNumber_Handler,
		},
		{
			MethodName: "FindByVIN",
			Handler:    _Service_FindByVIN_Handler,
		},
		{
			MethodName: "FindByCode",
			Handler:    _Service_FindByCode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/registration/service.proto",
}
