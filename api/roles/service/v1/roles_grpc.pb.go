// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: api/roles/service/v1/roles.proto

package v1

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

// RolesClient is the client API for Roles service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RolesClient interface {
	// 生成超级用户
	CreateSuperUser(ctx context.Context, in *CreateSuperUserRequest, opts ...grpc.CallOption) (*CreateSuperUserResponse, error)
	// 获取图片验证码
	GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaResponse, error)
	// 后台登陆
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// **********  超级管理员功能  ************
	// 创建后台用户
	CreateAdminUser(ctx context.Context, in *CreateAdminUserRequest, opts ...grpc.CallOption) (*CreateAdminUserResponse, error)
	// 后台用户列表
	AdminUserList(ctx context.Context, in *AdminUserListRequest, opts ...grpc.CallOption) (*AdminUserListResponse, error)
	// 后台用户详细
	AdminUserDetails(ctx context.Context, in *AdminUserDetailsRequest, opts ...grpc.CallOption) (*AdminUserDetailsResponse, error)
	// 编辑后台用户基本信息(用户名/密码/头像)
	AdminUserEdit(ctx context.Context, in *AdminUserEditRequest, opts ...grpc.CallOption) (*AdminUserEditResponse, error)
	// 用户分配角色
	SetRoles(ctx context.Context, in *SetRolesRequest, opts ...grpc.CallOption) (*SetRolesResponse, error)
	// 用户分配权限
	SetPermission(ctx context.Context, in *SetPermissionRequest, opts ...grpc.CallOption) (*SetPermissionResponse, error)
	// 删除后台用户
	AdminUserDelete(ctx context.Context, in *AdminUserDeleteRequest, opts ...grpc.CallOption) (*AdminUserDeleteResponse, error)
	// 创建角色
	RoleCreate(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error)
	// 角色列表
	RoleList(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error)
	// 删除角色
	RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error)
	// 编辑角色
	RoleEdit(ctx context.Context, in *RoleEditRequest, opts ...grpc.CallOption) (*RoleEditResponse, error)
	// 创建权限
	PermissionCreate(ctx context.Context, in *PermissionCreateRequest, opts ...grpc.CallOption) (*PermissionCreateResponse, error)
	// 权限列表
	PermissionList(ctx context.Context, in *PermissionListRequest, opts ...grpc.CallOption) (*PermissionListResponse, error)
	// 权限删除
	PermissionDelete(ctx context.Context, in *PermissionDeleteRequest, opts ...grpc.CallOption) (*PermissionDeleteResponse, error)
	// 更新权限
	PermissionEdit(ctx context.Context, in *PermissionEditRequest, opts ...grpc.CallOption) (*PermissionEditResponse, error)
	// *******  你的系统鉴权RPC接口  **********
	// 验证用户角色
	CheckRole(ctx context.Context, in *CheckRoleRequest, opts ...grpc.CallOption) (*CheckRoleResponse, error)
	// 验证用户权限
	CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error)
}

type rolesClient struct {
	cc grpc.ClientConnInterface
}

func NewRolesClient(cc grpc.ClientConnInterface) RolesClient {
	return &rolesClient{cc}
}

func (c *rolesClient) CreateSuperUser(ctx context.Context, in *CreateSuperUserRequest, opts ...grpc.CallOption) (*CreateSuperUserResponse, error) {
	out := new(CreateSuperUserResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/CreateSuperUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...grpc.CallOption) (*GetCaptchaResponse, error) {
	out := new(GetCaptchaResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/GetCaptcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) CreateAdminUser(ctx context.Context, in *CreateAdminUserRequest, opts ...grpc.CallOption) (*CreateAdminUserResponse, error) {
	out := new(CreateAdminUserResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/CreateAdminUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) AdminUserList(ctx context.Context, in *AdminUserListRequest, opts ...grpc.CallOption) (*AdminUserListResponse, error) {
	out := new(AdminUserListResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/AdminUserList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) AdminUserDetails(ctx context.Context, in *AdminUserDetailsRequest, opts ...grpc.CallOption) (*AdminUserDetailsResponse, error) {
	out := new(AdminUserDetailsResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/AdminUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) AdminUserEdit(ctx context.Context, in *AdminUserEditRequest, opts ...grpc.CallOption) (*AdminUserEditResponse, error) {
	out := new(AdminUserEditResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/AdminUserEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) SetRoles(ctx context.Context, in *SetRolesRequest, opts ...grpc.CallOption) (*SetRolesResponse, error) {
	out := new(SetRolesResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/SetRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) SetPermission(ctx context.Context, in *SetPermissionRequest, opts ...grpc.CallOption) (*SetPermissionResponse, error) {
	out := new(SetPermissionResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/SetPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) AdminUserDelete(ctx context.Context, in *AdminUserDeleteRequest, opts ...grpc.CallOption) (*AdminUserDeleteResponse, error) {
	out := new(AdminUserDeleteResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/AdminUserDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RoleCreate(ctx context.Context, in *RoleCreateRequest, opts ...grpc.CallOption) (*RoleCreateResponse, error) {
	out := new(RoleCreateResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/RoleCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RoleList(ctx context.Context, in *RoleListRequest, opts ...grpc.CallOption) (*RoleListResponse, error) {
	out := new(RoleListResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/RoleList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...grpc.CallOption) (*RoleDeleteResponse, error) {
	out := new(RoleDeleteResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/RoleDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) RoleEdit(ctx context.Context, in *RoleEditRequest, opts ...grpc.CallOption) (*RoleEditResponse, error) {
	out := new(RoleEditResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/RoleEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) PermissionCreate(ctx context.Context, in *PermissionCreateRequest, opts ...grpc.CallOption) (*PermissionCreateResponse, error) {
	out := new(PermissionCreateResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/PermissionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) PermissionList(ctx context.Context, in *PermissionListRequest, opts ...grpc.CallOption) (*PermissionListResponse, error) {
	out := new(PermissionListResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/PermissionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) PermissionDelete(ctx context.Context, in *PermissionDeleteRequest, opts ...grpc.CallOption) (*PermissionDeleteResponse, error) {
	out := new(PermissionDeleteResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/PermissionDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) PermissionEdit(ctx context.Context, in *PermissionEditRequest, opts ...grpc.CallOption) (*PermissionEditResponse, error) {
	out := new(PermissionEditResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/PermissionEdit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) CheckRole(ctx context.Context, in *CheckRoleRequest, opts ...grpc.CallOption) (*CheckRoleResponse, error) {
	out := new(CheckRoleResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/CheckRole", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rolesClient) CheckPermission(ctx context.Context, in *CheckPermissionRequest, opts ...grpc.CallOption) (*CheckPermissionResponse, error) {
	out := new(CheckPermissionResponse)
	err := c.cc.Invoke(ctx, "/api.roles.service.v1.Roles/CheckPermission", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RolesServer is the server API for Roles service.
// All implementations must embed UnimplementedRolesServer
// for forward compatibility
type RolesServer interface {
	// 生成超级用户
	CreateSuperUser(context.Context, *CreateSuperUserRequest) (*CreateSuperUserResponse, error)
	// 获取图片验证码
	GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaResponse, error)
	// 后台登陆
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// **********  超级管理员功能  ************
	// 创建后台用户
	CreateAdminUser(context.Context, *CreateAdminUserRequest) (*CreateAdminUserResponse, error)
	// 后台用户列表
	AdminUserList(context.Context, *AdminUserListRequest) (*AdminUserListResponse, error)
	// 后台用户详细
	AdminUserDetails(context.Context, *AdminUserDetailsRequest) (*AdminUserDetailsResponse, error)
	// 编辑后台用户基本信息(用户名/密码/头像)
	AdminUserEdit(context.Context, *AdminUserEditRequest) (*AdminUserEditResponse, error)
	// 用户分配角色
	SetRoles(context.Context, *SetRolesRequest) (*SetRolesResponse, error)
	// 用户分配权限
	SetPermission(context.Context, *SetPermissionRequest) (*SetPermissionResponse, error)
	// 删除后台用户
	AdminUserDelete(context.Context, *AdminUserDeleteRequest) (*AdminUserDeleteResponse, error)
	// 创建角色
	RoleCreate(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error)
	// 角色列表
	RoleList(context.Context, *RoleListRequest) (*RoleListResponse, error)
	// 删除角色
	RoleDelete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error)
	// 编辑角色
	RoleEdit(context.Context, *RoleEditRequest) (*RoleEditResponse, error)
	// 创建权限
	PermissionCreate(context.Context, *PermissionCreateRequest) (*PermissionCreateResponse, error)
	// 权限列表
	PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error)
	// 权限删除
	PermissionDelete(context.Context, *PermissionDeleteRequest) (*PermissionDeleteResponse, error)
	// 更新权限
	PermissionEdit(context.Context, *PermissionEditRequest) (*PermissionEditResponse, error)
	// *******  你的系统鉴权RPC接口  **********
	// 验证用户角色
	CheckRole(context.Context, *CheckRoleRequest) (*CheckRoleResponse, error)
	// 验证用户权限
	CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error)
	mustEmbedUnimplementedRolesServer()
}

// UnimplementedRolesServer must be embedded to have forward compatible implementations.
type UnimplementedRolesServer struct {
}

func (UnimplementedRolesServer) CreateSuperUser(context.Context, *CreateSuperUserRequest) (*CreateSuperUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSuperUser not implemented")
}
func (UnimplementedRolesServer) GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcha not implemented")
}
func (UnimplementedRolesServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedRolesServer) CreateAdminUser(context.Context, *CreateAdminUserRequest) (*CreateAdminUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAdminUser not implemented")
}
func (UnimplementedRolesServer) AdminUserList(context.Context, *AdminUserListRequest) (*AdminUserListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUserList not implemented")
}
func (UnimplementedRolesServer) AdminUserDetails(context.Context, *AdminUserDetailsRequest) (*AdminUserDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUserDetails not implemented")
}
func (UnimplementedRolesServer) AdminUserEdit(context.Context, *AdminUserEditRequest) (*AdminUserEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUserEdit not implemented")
}
func (UnimplementedRolesServer) SetRoles(context.Context, *SetRolesRequest) (*SetRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetRoles not implemented")
}
func (UnimplementedRolesServer) SetPermission(context.Context, *SetPermissionRequest) (*SetPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetPermission not implemented")
}
func (UnimplementedRolesServer) AdminUserDelete(context.Context, *AdminUserDeleteRequest) (*AdminUserDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminUserDelete not implemented")
}
func (UnimplementedRolesServer) RoleCreate(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleCreate not implemented")
}
func (UnimplementedRolesServer) RoleList(context.Context, *RoleListRequest) (*RoleListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleList not implemented")
}
func (UnimplementedRolesServer) RoleDelete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleDelete not implemented")
}
func (UnimplementedRolesServer) RoleEdit(context.Context, *RoleEditRequest) (*RoleEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RoleEdit not implemented")
}
func (UnimplementedRolesServer) PermissionCreate(context.Context, *PermissionCreateRequest) (*PermissionCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionCreate not implemented")
}
func (UnimplementedRolesServer) PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionList not implemented")
}
func (UnimplementedRolesServer) PermissionDelete(context.Context, *PermissionDeleteRequest) (*PermissionDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionDelete not implemented")
}
func (UnimplementedRolesServer) PermissionEdit(context.Context, *PermissionEditRequest) (*PermissionEditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PermissionEdit not implemented")
}
func (UnimplementedRolesServer) CheckRole(context.Context, *CheckRoleRequest) (*CheckRoleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckRole not implemented")
}
func (UnimplementedRolesServer) CheckPermission(context.Context, *CheckPermissionRequest) (*CheckPermissionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckPermission not implemented")
}
func (UnimplementedRolesServer) mustEmbedUnimplementedRolesServer() {}

// UnsafeRolesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RolesServer will
// result in compilation errors.
type UnsafeRolesServer interface {
	mustEmbedUnimplementedRolesServer()
}

func RegisterRolesServer(s grpc.ServiceRegistrar, srv RolesServer) {
	s.RegisterService(&Roles_ServiceDesc, srv)
}

func _Roles_CreateSuperUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSuperUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).CreateSuperUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/CreateSuperUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).CreateSuperUser(ctx, req.(*CreateSuperUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_GetCaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).GetCaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/GetCaptcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).GetCaptcha(ctx, req.(*GetCaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_CreateAdminUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAdminUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).CreateAdminUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/CreateAdminUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).CreateAdminUser(ctx, req.(*CreateAdminUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_AdminUserList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUserListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).AdminUserList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/AdminUserList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).AdminUserList(ctx, req.(*AdminUserListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_AdminUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUserDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).AdminUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/AdminUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).AdminUserDetails(ctx, req.(*AdminUserDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_AdminUserEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUserEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).AdminUserEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/AdminUserEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).AdminUserEdit(ctx, req.(*AdminUserEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_SetRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).SetRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/SetRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).SetRoles(ctx, req.(*SetRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_SetPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).SetPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/SetPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).SetPermission(ctx, req.(*SetPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_AdminUserDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AdminUserDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).AdminUserDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/AdminUserDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).AdminUserDelete(ctx, req.(*AdminUserDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_RoleCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).RoleCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/RoleCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).RoleCreate(ctx, req.(*RoleCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_RoleList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).RoleList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/RoleList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).RoleList(ctx, req.(*RoleListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_RoleDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).RoleDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/RoleDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).RoleDelete(ctx, req.(*RoleDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_RoleEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RoleEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).RoleEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/RoleEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).RoleEdit(ctx, req.(*RoleEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_PermissionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).PermissionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/PermissionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).PermissionCreate(ctx, req.(*PermissionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_PermissionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).PermissionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/PermissionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).PermissionList(ctx, req.(*PermissionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_PermissionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).PermissionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/PermissionDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).PermissionDelete(ctx, req.(*PermissionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_PermissionEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PermissionEditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).PermissionEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/PermissionEdit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).PermissionEdit(ctx, req.(*PermissionEditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_CheckRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRoleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).CheckRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/CheckRole",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).CheckRole(ctx, req.(*CheckRoleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Roles_CheckPermission_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckPermissionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RolesServer).CheckPermission(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.roles.service.v1.Roles/CheckPermission",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RolesServer).CheckPermission(ctx, req.(*CheckPermissionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Roles_ServiceDesc is the grpc.ServiceDesc for Roles service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Roles_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.roles.service.v1.Roles",
	HandlerType: (*RolesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSuperUser",
			Handler:    _Roles_CreateSuperUser_Handler,
		},
		{
			MethodName: "GetCaptcha",
			Handler:    _Roles_GetCaptcha_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _Roles_Login_Handler,
		},
		{
			MethodName: "CreateAdminUser",
			Handler:    _Roles_CreateAdminUser_Handler,
		},
		{
			MethodName: "AdminUserList",
			Handler:    _Roles_AdminUserList_Handler,
		},
		{
			MethodName: "AdminUserDetails",
			Handler:    _Roles_AdminUserDetails_Handler,
		},
		{
			MethodName: "AdminUserEdit",
			Handler:    _Roles_AdminUserEdit_Handler,
		},
		{
			MethodName: "SetRoles",
			Handler:    _Roles_SetRoles_Handler,
		},
		{
			MethodName: "SetPermission",
			Handler:    _Roles_SetPermission_Handler,
		},
		{
			MethodName: "AdminUserDelete",
			Handler:    _Roles_AdminUserDelete_Handler,
		},
		{
			MethodName: "RoleCreate",
			Handler:    _Roles_RoleCreate_Handler,
		},
		{
			MethodName: "RoleList",
			Handler:    _Roles_RoleList_Handler,
		},
		{
			MethodName: "RoleDelete",
			Handler:    _Roles_RoleDelete_Handler,
		},
		{
			MethodName: "RoleEdit",
			Handler:    _Roles_RoleEdit_Handler,
		},
		{
			MethodName: "PermissionCreate",
			Handler:    _Roles_PermissionCreate_Handler,
		},
		{
			MethodName: "PermissionList",
			Handler:    _Roles_PermissionList_Handler,
		},
		{
			MethodName: "PermissionDelete",
			Handler:    _Roles_PermissionDelete_Handler,
		},
		{
			MethodName: "PermissionEdit",
			Handler:    _Roles_PermissionEdit_Handler,
		},
		{
			MethodName: "CheckRole",
			Handler:    _Roles_CheckRole_Handler,
		},
		{
			MethodName: "CheckPermission",
			Handler:    _Roles_CheckPermission_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/roles/service/v1/roles.proto",
}
