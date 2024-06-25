// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: superadmins.proto

package superadmin_service

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SuperadminServiceClient is the client API for SuperadminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SuperadminServiceClient interface {
	Create(ctx context.Context, in *CreateSuperadmin, opts ...grpc.CallOption) (*GetSuperadmin, error)
	GetByID(ctx context.Context, in *SuperadminPrimaryKey, opts ...grpc.CallOption) (*GetSuperadmin, error)
	GetList(ctx context.Context, in *GetListSuperadminRequest, opts ...grpc.CallOption) (*GetListSuperadminResponse, error)
	Update(ctx context.Context, in *UpdateSuperadmin, opts ...grpc.CallOption) (*GetSuperadmin, error)
	Delete(ctx context.Context, in *SuperadminPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error)
}

type superadminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSuperadminServiceClient(cc grpc.ClientConnInterface) SuperadminServiceClient {
	return &superadminServiceClient{cc}
}

func (c *superadminServiceClient) Create(ctx context.Context, in *CreateSuperadmin, opts ...grpc.CallOption) (*GetSuperadmin, error) {
	out := new(GetSuperadmin)
	err := c.cc.Invoke(ctx, "/superadmin_service_go.SuperadminService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superadminServiceClient) GetByID(ctx context.Context, in *SuperadminPrimaryKey, opts ...grpc.CallOption) (*GetSuperadmin, error) {
	out := new(GetSuperadmin)
	err := c.cc.Invoke(ctx, "/superadmin_service_go.SuperadminService/GetByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superadminServiceClient) GetList(ctx context.Context, in *GetListSuperadminRequest, opts ...grpc.CallOption) (*GetListSuperadminResponse, error) {
	out := new(GetListSuperadminResponse)
	err := c.cc.Invoke(ctx, "/superadmin_service_go.SuperadminService/GetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superadminServiceClient) Update(ctx context.Context, in *UpdateSuperadmin, opts ...grpc.CallOption) (*GetSuperadmin, error) {
	out := new(GetSuperadmin)
	err := c.cc.Invoke(ctx, "/superadmin_service_go.SuperadminService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *superadminServiceClient) Delete(ctx context.Context, in *SuperadminPrimaryKey, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/superadmin_service_go.SuperadminService/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SuperadminServiceServer is the server API for SuperadminService service.
// All implementations should embed UnimplementedSuperadminServiceServer
// for forward compatibility
type SuperadminServiceServer interface {
	Create(context.Context, *CreateSuperadmin) (*GetSuperadmin, error)
	GetByID(context.Context, *SuperadminPrimaryKey) (*GetSuperadmin, error)
	GetList(context.Context, *GetListSuperadminRequest) (*GetListSuperadminResponse, error)
	Update(context.Context, *UpdateSuperadmin) (*GetSuperadmin, error)
	Delete(context.Context, *SuperadminPrimaryKey) (*empty.Empty, error)
}

// UnimplementedSuperadminServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSuperadminServiceServer struct {
}

func (UnimplementedSuperadminServiceServer) Create(context.Context, *CreateSuperadmin) (*GetSuperadmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSuperadminServiceServer) GetByID(context.Context, *SuperadminPrimaryKey) (*GetSuperadmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByID not implemented")
}
func (UnimplementedSuperadminServiceServer) GetList(context.Context, *GetListSuperadminRequest) (*GetListSuperadminResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}
func (UnimplementedSuperadminServiceServer) Update(context.Context, *UpdateSuperadmin) (*GetSuperadmin, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSuperadminServiceServer) Delete(context.Context, *SuperadminPrimaryKey) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}

// UnsafeSuperadminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SuperadminServiceServer will
// result in compilation errors.
type UnsafeSuperadminServiceServer interface {
	mustEmbedUnimplementedSuperadminServiceServer()
}

func RegisterSuperadminServiceServer(s grpc.ServiceRegistrar, srv SuperadminServiceServer) {
	s.RegisterService(&SuperadminService_ServiceDesc, srv)
}

func _SuperadminService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSuperadmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperadminServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/superadmin_service_go.SuperadminService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperadminServiceServer).Create(ctx, req.(*CreateSuperadmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperadminService_GetByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuperadminPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperadminServiceServer).GetByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/superadmin_service_go.SuperadminService/GetByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperadminServiceServer).GetByID(ctx, req.(*SuperadminPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperadminService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetListSuperadminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperadminServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/superadmin_service_go.SuperadminService/GetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperadminServiceServer).GetList(ctx, req.(*GetListSuperadminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperadminService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateSuperadmin)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperadminServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/superadmin_service_go.SuperadminService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperadminServiceServer).Update(ctx, req.(*UpdateSuperadmin))
	}
	return interceptor(ctx, in, info, handler)
}

func _SuperadminService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SuperadminPrimaryKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SuperadminServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/superadmin_service_go.SuperadminService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SuperadminServiceServer).Delete(ctx, req.(*SuperadminPrimaryKey))
	}
	return interceptor(ctx, in, info, handler)
}

// SuperadminService_ServiceDesc is the grpc.ServiceDesc for SuperadminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SuperadminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "superadmin_service_go.SuperadminService",
	HandlerType: (*SuperadminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SuperadminService_Create_Handler,
		},
		{
			MethodName: "GetByID",
			Handler:    _SuperadminService_GetByID_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SuperadminService_GetList_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SuperadminService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SuperadminService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "superadmins.proto",
}
