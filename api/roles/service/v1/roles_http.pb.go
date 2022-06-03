// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.4

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type RolesHTTPServer interface {
	AdminUserDelete(context.Context, *AdminUserDeleteRequest) (*AdminUserDeleteResponse, error)
	AdminUserEdit(context.Context, *AdminUserEditRequest) (*AdminUserEditResponse, error)
	AdminUserInfos(context.Context, *AdminUserInfosRequest) (*AdminUserInfosResponse, error)
	AdminUserList(context.Context, *AdminUserListRequest) (*AdminUserListResponse, error)
	CreateAdminUser(context.Context, *CreateAdminUserRequest) (*CreateAdminUserResponse, error)
	CreateSuperUser(context.Context, *CreateSuperUserRequest) (*CreateSuperUserResponse, error)
	GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	PermissionCreate(context.Context, *PermissionCreateRequest) (*PermissionCreateResponse, error)
	PermissionDelete(context.Context, *PermissionDeleteRequest) (*PermissionDeleteResponse, error)
	PermissionEdit(context.Context, *PermissionEditRequest) (*PermissionEditResponse, error)
	PermissionList(context.Context, *PermissionListRequest) (*PermissionListResponse, error)
	RoleCreate(context.Context, *RoleCreateRequest) (*RoleCreateResponse, error)
	RoleDelete(context.Context, *RoleDeleteRequest) (*RoleDeleteResponse, error)
	RoleEdit(context.Context, *RoleEditRequest) (*RoleEditResponse, error)
	RoleList(context.Context, *RoleListRequest) (*RoleListResponse, error)
	RouteCreate(context.Context, *RouteCreateRequest) (*RouteCreateResponse, error)
	RoutePermission(context.Context, *RoutePermissionRequest) (*RoutePermissionResponse, error)
	RouteRole(context.Context, *RouteRoleRequest) (*RouteRoleResponse, error)
	SetPermission(context.Context, *SetPermissionRequest) (*SetPermissionResponse, error)
	SetRoles(context.Context, *SetRolesRequest) (*SetRolesResponse, error)
}

func RegisterRolesHTTPServer(s *http.Server, srv RolesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/auth/role/superuser/create", _Roles_CreateSuperUser0_HTTP_Handler(srv))
	r.GET("/v1/role/captcha", _Roles_GetCaptcha0_HTTP_Handler(srv))
	r.POST("/v1/role/login", _Roles_Login0_HTTP_Handler(srv))
	r.POST("/v1/role/user/create", _Roles_CreateAdminUser0_HTTP_Handler(srv))
	r.POST("/v1/role/user/list", _Roles_AdminUserList0_HTTP_Handler(srv))
	r.POST("/v1/role/user/details", _Roles_AdminUserInfos0_HTTP_Handler(srv))
	r.POST("/v1/role/user/update", _Roles_AdminUserEdit0_HTTP_Handler(srv))
	r.POST("/v1/role/user/setroles", _Roles_SetRoles0_HTTP_Handler(srv))
	r.POST("/v1/role/user/setpermissions", _Roles_SetPermission0_HTTP_Handler(srv))
	r.DELETE("/v1/role/user/{id}", _Roles_AdminUserDelete0_HTTP_Handler(srv))
	r.POST("/v1/role/create", _Roles_RoleCreate0_HTTP_Handler(srv))
	r.POST("/v1/role/list", _Roles_RoleList0_HTTP_Handler(srv))
	r.POST("/v1/role/delete", _Roles_RoleDelete0_HTTP_Handler(srv))
	r.POST("/v1/role/update", _Roles_RoleEdit0_HTTP_Handler(srv))
	r.POST("/v1/role/permission/create", _Roles_PermissionCreate0_HTTP_Handler(srv))
	r.POST("/v1/role/permission/list", _Roles_PermissionList0_HTTP_Handler(srv))
	r.POST("/v1/role/permission/delete", _Roles_PermissionDelete0_HTTP_Handler(srv))
	r.POST("/v1/role/permission/update", _Roles_PermissionEdit0_HTTP_Handler(srv))
	r.POST("/v1/role/route/create", _Roles_RouteCreate0_HTTP_Handler(srv))
	r.POST("/v1/role/route/setrole", _Roles_RouteRole0_HTTP_Handler(srv))
	r.POST("/v1/role/route/setpermission", _Roles_RoutePermission0_HTTP_Handler(srv))
}

func _Roles_CreateSuperUser0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSuperUserRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/CreateSuperUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSuperUser(ctx, req.(*CreateSuperUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSuperUserResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_GetCaptcha0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetCaptchaRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/GetCaptcha")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetCaptcha(ctx, req.(*GetCaptchaRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetCaptchaResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_Login0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/Login")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*LoginResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_CreateAdminUser0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateAdminUserRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/CreateAdminUser")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateAdminUser(ctx, req.(*CreateAdminUserRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateAdminUserResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_AdminUserList0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminUserListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/AdminUserList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminUserList(ctx, req.(*AdminUserListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminUserListResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_AdminUserInfos0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminUserInfosRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/AdminUserInfos")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminUserInfos(ctx, req.(*AdminUserInfosRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminUserInfosResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_AdminUserEdit0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminUserEditRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/AdminUserEdit")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminUserEdit(ctx, req.(*AdminUserEditRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminUserEditResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_SetRoles0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetRolesRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/SetRoles")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetRoles(ctx, req.(*SetRolesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetRolesResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_SetPermission0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SetPermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/SetPermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SetPermission(ctx, req.(*SetPermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SetPermissionResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_AdminUserDelete0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AdminUserDeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/AdminUserDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AdminUserDelete(ctx, req.(*AdminUserDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AdminUserDeleteResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_RoleCreate0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/RoleCreate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleCreate(ctx, req.(*RoleCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_RoleList0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/RoleList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleList(ctx, req.(*RoleListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleListResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_RoleDelete0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleDeleteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/RoleDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleDelete(ctx, req.(*RoleDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleDeleteResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_RoleEdit0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoleEditRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/RoleEdit")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoleEdit(ctx, req.(*RoleEditRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoleEditResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_PermissionCreate0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/PermissionCreate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionCreate(ctx, req.(*PermissionCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_PermissionList0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionListRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/PermissionList")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionList(ctx, req.(*PermissionListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionListResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_PermissionDelete0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionDeleteRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/PermissionDelete")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionDelete(ctx, req.(*PermissionDeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionDeleteResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_PermissionEdit0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in PermissionEditRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/PermissionEdit")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.PermissionEdit(ctx, req.(*PermissionEditRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*PermissionEditResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_RouteCreate0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RouteCreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/RouteCreate")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RouteCreate(ctx, req.(*RouteCreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RouteCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_RouteRole0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RouteRoleRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/RouteRole")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RouteRole(ctx, req.(*RouteRoleRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RouteRoleResponse)
		return ctx.Result(200, reply)
	}
}

func _Roles_RoutePermission0_HTTP_Handler(srv RolesHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RoutePermissionRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/api.roles.service.v1.Roles/RoutePermission")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RoutePermission(ctx, req.(*RoutePermissionRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RoutePermissionResponse)
		return ctx.Result(200, reply)
	}
}

type RolesHTTPClient interface {
	AdminUserDelete(ctx context.Context, req *AdminUserDeleteRequest, opts ...http.CallOption) (rsp *AdminUserDeleteResponse, err error)
	AdminUserEdit(ctx context.Context, req *AdminUserEditRequest, opts ...http.CallOption) (rsp *AdminUserEditResponse, err error)
	AdminUserInfos(ctx context.Context, req *AdminUserInfosRequest, opts ...http.CallOption) (rsp *AdminUserInfosResponse, err error)
	AdminUserList(ctx context.Context, req *AdminUserListRequest, opts ...http.CallOption) (rsp *AdminUserListResponse, err error)
	CreateAdminUser(ctx context.Context, req *CreateAdminUserRequest, opts ...http.CallOption) (rsp *CreateAdminUserResponse, err error)
	CreateSuperUser(ctx context.Context, req *CreateSuperUserRequest, opts ...http.CallOption) (rsp *CreateSuperUserResponse, err error)
	GetCaptcha(ctx context.Context, req *GetCaptchaRequest, opts ...http.CallOption) (rsp *GetCaptchaResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
	PermissionCreate(ctx context.Context, req *PermissionCreateRequest, opts ...http.CallOption) (rsp *PermissionCreateResponse, err error)
	PermissionDelete(ctx context.Context, req *PermissionDeleteRequest, opts ...http.CallOption) (rsp *PermissionDeleteResponse, err error)
	PermissionEdit(ctx context.Context, req *PermissionEditRequest, opts ...http.CallOption) (rsp *PermissionEditResponse, err error)
	PermissionList(ctx context.Context, req *PermissionListRequest, opts ...http.CallOption) (rsp *PermissionListResponse, err error)
	RoleCreate(ctx context.Context, req *RoleCreateRequest, opts ...http.CallOption) (rsp *RoleCreateResponse, err error)
	RoleDelete(ctx context.Context, req *RoleDeleteRequest, opts ...http.CallOption) (rsp *RoleDeleteResponse, err error)
	RoleEdit(ctx context.Context, req *RoleEditRequest, opts ...http.CallOption) (rsp *RoleEditResponse, err error)
	RoleList(ctx context.Context, req *RoleListRequest, opts ...http.CallOption) (rsp *RoleListResponse, err error)
	RouteCreate(ctx context.Context, req *RouteCreateRequest, opts ...http.CallOption) (rsp *RouteCreateResponse, err error)
	RoutePermission(ctx context.Context, req *RoutePermissionRequest, opts ...http.CallOption) (rsp *RoutePermissionResponse, err error)
	RouteRole(ctx context.Context, req *RouteRoleRequest, opts ...http.CallOption) (rsp *RouteRoleResponse, err error)
	SetPermission(ctx context.Context, req *SetPermissionRequest, opts ...http.CallOption) (rsp *SetPermissionResponse, err error)
	SetRoles(ctx context.Context, req *SetRolesRequest, opts ...http.CallOption) (rsp *SetRolesResponse, err error)
}

type RolesHTTPClientImpl struct {
	cc *http.Client
}

func NewRolesHTTPClient(client *http.Client) RolesHTTPClient {
	return &RolesHTTPClientImpl{client}
}

func (c *RolesHTTPClientImpl) AdminUserDelete(ctx context.Context, in *AdminUserDeleteRequest, opts ...http.CallOption) (*AdminUserDeleteResponse, error) {
	var out AdminUserDeleteResponse
	pattern := "/v1/role/user/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/AdminUserDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) AdminUserEdit(ctx context.Context, in *AdminUserEditRequest, opts ...http.CallOption) (*AdminUserEditResponse, error) {
	var out AdminUserEditResponse
	pattern := "/v1/role/user/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/AdminUserEdit"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) AdminUserInfos(ctx context.Context, in *AdminUserInfosRequest, opts ...http.CallOption) (*AdminUserInfosResponse, error) {
	var out AdminUserInfosResponse
	pattern := "/v1/role/user/details"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/AdminUserInfos"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) AdminUserList(ctx context.Context, in *AdminUserListRequest, opts ...http.CallOption) (*AdminUserListResponse, error) {
	var out AdminUserListResponse
	pattern := "/v1/role/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/AdminUserList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) CreateAdminUser(ctx context.Context, in *CreateAdminUserRequest, opts ...http.CallOption) (*CreateAdminUserResponse, error) {
	var out CreateAdminUserResponse
	pattern := "/v1/role/user/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/CreateAdminUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) CreateSuperUser(ctx context.Context, in *CreateSuperUserRequest, opts ...http.CallOption) (*CreateSuperUserResponse, error) {
	var out CreateSuperUserResponse
	pattern := "/v1/auth/role/superuser/create"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/CreateSuperUser"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) GetCaptcha(ctx context.Context, in *GetCaptchaRequest, opts ...http.CallOption) (*GetCaptchaResponse, error) {
	var out GetCaptchaResponse
	pattern := "/v1/role/captcha"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/GetCaptcha"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*LoginResponse, error) {
	var out LoginResponse
	pattern := "/v1/role/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) PermissionCreate(ctx context.Context, in *PermissionCreateRequest, opts ...http.CallOption) (*PermissionCreateResponse, error) {
	var out PermissionCreateResponse
	pattern := "/v1/role/permission/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/PermissionCreate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) PermissionDelete(ctx context.Context, in *PermissionDeleteRequest, opts ...http.CallOption) (*PermissionDeleteResponse, error) {
	var out PermissionDeleteResponse
	pattern := "/v1/role/permission/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/PermissionDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) PermissionEdit(ctx context.Context, in *PermissionEditRequest, opts ...http.CallOption) (*PermissionEditResponse, error) {
	var out PermissionEditResponse
	pattern := "/v1/role/permission/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/PermissionEdit"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) PermissionList(ctx context.Context, in *PermissionListRequest, opts ...http.CallOption) (*PermissionListResponse, error) {
	var out PermissionListResponse
	pattern := "/v1/role/permission/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/PermissionList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) RoleCreate(ctx context.Context, in *RoleCreateRequest, opts ...http.CallOption) (*RoleCreateResponse, error) {
	var out RoleCreateResponse
	pattern := "/v1/role/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/RoleCreate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) RoleDelete(ctx context.Context, in *RoleDeleteRequest, opts ...http.CallOption) (*RoleDeleteResponse, error) {
	var out RoleDeleteResponse
	pattern := "/v1/role/delete"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/RoleDelete"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) RoleEdit(ctx context.Context, in *RoleEditRequest, opts ...http.CallOption) (*RoleEditResponse, error) {
	var out RoleEditResponse
	pattern := "/v1/role/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/RoleEdit"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) RoleList(ctx context.Context, in *RoleListRequest, opts ...http.CallOption) (*RoleListResponse, error) {
	var out RoleListResponse
	pattern := "/v1/role/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/RoleList"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) RouteCreate(ctx context.Context, in *RouteCreateRequest, opts ...http.CallOption) (*RouteCreateResponse, error) {
	var out RouteCreateResponse
	pattern := "/v1/role/route/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/RouteCreate"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) RoutePermission(ctx context.Context, in *RoutePermissionRequest, opts ...http.CallOption) (*RoutePermissionResponse, error) {
	var out RoutePermissionResponse
	pattern := "/v1/role/route/setpermission"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/RoutePermission"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) RouteRole(ctx context.Context, in *RouteRoleRequest, opts ...http.CallOption) (*RouteRoleResponse, error) {
	var out RouteRoleResponse
	pattern := "/v1/role/route/setrole"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/RouteRole"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) SetPermission(ctx context.Context, in *SetPermissionRequest, opts ...http.CallOption) (*SetPermissionResponse, error) {
	var out SetPermissionResponse
	pattern := "/v1/role/user/setpermissions"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/SetPermission"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *RolesHTTPClientImpl) SetRoles(ctx context.Context, in *SetRolesRequest, opts ...http.CallOption) (*SetRolesResponse, error) {
	var out SetRolesResponse
	pattern := "/v1/role/user/setroles"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/SetRoles"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
