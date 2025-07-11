// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.6
// source: user_permission.proto

package proto

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

// UserPermissionServiceClient is the client API for UserPermissionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserPermissionServiceClient interface {
	SetUserPermission(ctx context.Context, in *SetUserPermissionRequest, opts ...grpc.CallOption) (*SetUserPermissionResponse, error)
	GetUserPermissions(ctx context.Context, in *GetUserPermissionsRequest, opts ...grpc.CallOption) (*GetUserPermissionsResponse, error)
}

type userPermissionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserPermissionServiceClient(cc grpc.ClientConnInterface) UserPermissionServiceClient {
	return &userPermissionServiceClient{cc}
}

func (c *userPermissionServiceClient) SetUserPermission(ctx context.Context, in *SetUserPermissionRequest, opts ...grpc.CallOption) (*SetUserPermissionResponse, error) {
	out := new(SetUserPermissionResponse)
	err := c.cc.Invoke(ctx, "/user_permission.UserPermissionService/SetUserPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userPermissionServiceClient) GetUserPermissions(ctx context.Context, in *GetUserPermissionsRequest, opts ...grpc.CallOption) (*GetUserPermissionsResponse, error) {
	out := new(GetUserPermissionsResponse)
	err := c.cc.Invoke(ctx, "/user_permission.UserPermissionService/GetUserPermissions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserPermissionServiceServer is the server API for UserPermissionService service.
// All implementations should embed UnimplementedUserPermissionServiceServer
// for forward compatibility
type UserPermissionServiceServer interface {
	SetUserPermission(context.Context, *SetUserPermissionRequest) (*SetUserPermissionResponse, error)
	GetUserPermissions(context.Context, *GetUserPermissionsRequest) (*GetUserPermissionsResponse, error)
}

// UnimplementedUserPermissionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedUserPermissionServiceServer struct {
}

func (UnimplementedUserPermissionServiceServer) SetUserPermission(context.Context, *SetUserPermissionRequest) (*SetUserPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserPermission not implemented")
}
func (UnimplementedUserPermissionServiceServer) GetUserPermissions(context.Context, *GetUserPermissionsRequest) (*GetUserPermissionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPermissions not implemented")
}

// UnsafeUserPermissionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserPermissionServiceServer will
// result in compilation errors.
type UnsafeUserPermissionServiceServer interface {
	mustEmbedUnimplementedUserPermissionServiceServer()
}

func RegisterUserPermissionServiceServer(s grpc.ServiceRegistrar, srv UserPermissionServiceServer) {
	s.RegisterService(&UserPermissionService_ServiceDesc, srv)
}

func _UserPermissionService_SetUserPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPermissionServiceServer).SetUserPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_permission.UserPermissionService/SetUserPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPermissionServiceServer).SetUserPermission(ctx, req.(*SetUserPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserPermissionService_GetUserPermissions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPermissionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserPermissionServiceServer).GetUserPermissions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user_permission.UserPermissionService/GetUserPermissions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserPermissionServiceServer).GetUserPermissions(ctx, req.(*GetUserPermissionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserPermissionService_ServiceDesc is the grpc.ServiceDesc for UserPermissionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserPermissionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user_permission.UserPermissionService",
	HandlerType: (*UserPermissionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetUserPermission",
			Handler:    _UserPermissionService_SetUserPermission_Handler,
		},
		{
			MethodName: "GetUserPermissions",
			Handler:    _UserPermissionService_GetUserPermissions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_permission.proto",
}
