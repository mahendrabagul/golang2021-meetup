// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package employee

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

// EmployeeServiceClient is the client API for EmployeeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmployeeServiceClient interface {
	GetDetails(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error)
}

type employeeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmployeeServiceClient(cc grpc.ClientConnInterface) EmployeeServiceClient {
	return &employeeServiceClient{cc}
}

func (c *employeeServiceClient) GetDetails(ctx context.Context, in *EmployeeRequest, opts ...grpc.CallOption) (*EmployeeResponse, error) {
	out := new(EmployeeResponse)
	err := c.cc.Invoke(ctx, "/employee.EmployeeService/getDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmployeeServiceServer is the server API for EmployeeService service.
// All implementations must embed UnimplementedEmployeeServiceServer
// for forward compatibility
type EmployeeServiceServer interface {
	GetDetails(context.Context, *EmployeeRequest) (*EmployeeResponse, error)
	mustEmbedUnimplementedEmployeeServiceServer()
}

// UnimplementedEmployeeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmployeeServiceServer struct {
}

func (UnimplementedEmployeeServiceServer) GetDetails(context.Context, *EmployeeRequest) (*EmployeeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetails not implemented")
}
func (UnimplementedEmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {}

// UnsafeEmployeeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmployeeServiceServer will
// result in compilation errors.
type UnsafeEmployeeServiceServer interface {
	mustEmbedUnimplementedEmployeeServiceServer()
}

func RegisterEmployeeServiceServer(s grpc.ServiceRegistrar, srv EmployeeServiceServer) {
	s.RegisterService(&EmployeeService_ServiceDesc, srv)
}

func _EmployeeService_GetDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmployeeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmployeeServiceServer).GetDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/employee.EmployeeService/getDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmployeeServiceServer).GetDetails(ctx, req.(*EmployeeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmployeeService_ServiceDesc is the grpc.ServiceDesc for EmployeeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmployeeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "employee.EmployeeService",
	HandlerType: (*EmployeeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "getDetails",
			Handler:    _EmployeeService_GetDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobufs/employee.proto",
}
