// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.1.3

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
	CreateSuperUser(context.Context, *CreateSuperUserRequest) (*CreateSuperUserResponse, error)
	GetCaptcha(context.Context, *GetCaptchaRequest) (*GetCaptchaResponse, error)
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterRolesHTTPServer(s *http.Server, srv RolesHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/role/superuser/create", _Roles_CreateSuperUser0_HTTP_Handler(srv))
	r.GET("/v1/role/captcha", _Roles_GetCaptcha0_HTTP_Handler(srv))
	r.POST("/v1/role/superuser/create", _Roles_Login0_HTTP_Handler(srv))
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

type RolesHTTPClient interface {
	CreateSuperUser(ctx context.Context, req *CreateSuperUserRequest, opts ...http.CallOption) (rsp *CreateSuperUserResponse, err error)
	GetCaptcha(ctx context.Context, req *GetCaptchaRequest, opts ...http.CallOption) (rsp *GetCaptchaResponse, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *LoginResponse, err error)
}

type RolesHTTPClientImpl struct {
	cc *http.Client
}

func NewRolesHTTPClient(client *http.Client) RolesHTTPClient {
	return &RolesHTTPClientImpl{client}
}

func (c *RolesHTTPClientImpl) CreateSuperUser(ctx context.Context, in *CreateSuperUserRequest, opts ...http.CallOption) (*CreateSuperUserResponse, error) {
	var out CreateSuperUserResponse
	pattern := "/v1/role/superuser/create"
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
	pattern := "/v1/role/superuser/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation("/api.roles.service.v1.Roles/Login"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}