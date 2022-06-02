package service

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/pkg/errors"
	"github.com/it-moom/moom-roles/pkg/tool"
)

// CreateSuperUser 生成超级用户
func (r *RolesService) CreateSuperUser(ctx context.Context, req *v1.CreateSuperUserRequest) (*v1.CreateSuperUserResponse, error) {
	return r.uc.CreateSuperUser(ctx)
}

// GetCaptcha 获取图片验证码
func (r *RolesService) GetCaptcha(ctx context.Context, req *v1.GetCaptchaRequest) (*v1.GetCaptchaResponse, error) {
	if !conf.GB.Verify {
		return &v1.GetCaptchaResponse{}, errors.ErrNoNeedCaptcha
	}

	id, bs64, err := tool.GetCaptcha()
	if err != nil {
		return &v1.GetCaptchaResponse{}, errors.ErrSystemBusy
	}

	return &v1.GetCaptchaResponse{
		Key:    id,
		Verify: bs64,
	}, nil
}

// Login 登陆
func (r *RolesService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if !tool.VerifyNameFormat(req.Name) {
		return &v1.LoginResponse{}, errors.ErrInvalidUsername
	}

	if !tool.VerifyPassFormat(req.Pass) {
		return &v1.LoginResponse{}, errors.ErrInvalidPass
	}

	return r.uc.Login(ctx, req)
}

func (r *RolesService) CreateAdminUser(ctx context.Context, req *v1.CreateAdminUserRequest) (*v1.CreateAdminUserResponse, error) {
	if !tool.VerifyNameFormat(req.Name) {
		return &v1.CreateAdminUserResponse{}, errors.ErrInvalidUsername
	}

	if !tool.VerifyPassFormat(req.Pass) {
		return &v1.CreateAdminUserResponse{}, errors.ErrInvalidPass
	}

	return r.uc.CreateAdminUser(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	return r.uc.AdminUserList(ctx, req)
}

func (r *RolesService) AdminUserInfos(ctx context.Context, req *v1.AdminUserInfosRequest) (*v1.AdminUserInfosResponse, error) {
	return r.uc.AdminUserInfos(ctx, r.GetUserID(ctx))
}

func (r *RolesService) AdminUserEdit(ctx context.Context, req *v1.AdminUserEditRequest) (*v1.AdminUserEditResponse, error) {
	if req.Name == "" && req.Pass == "" && req.Icon == "" {
		return &v1.AdminUserEditResponse{}, errors.ErrInvalidParams
	}

	return r.uc.AdminUserEdit(ctx, req)
}

func (r *RolesService) SetRoles(ctx context.Context, req *v1.SetRolesRequest) (*v1.SetRolesResponse, error) {
	if req.Uid <= 0 {
		return &v1.SetRolesResponse{}, errors.ErrInvalidUID
	}

	if len(req.Rid) <= 0 {
		return &v1.SetRolesResponse{}, errors.ErrInvalidParams
	}

	return r.uc.SetRoles(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) SetPermission(ctx context.Context, req *v1.SetPermissionRequest) (*v1.SetPermissionResponse, error) {
	if req.Uid <= 0 {
		return &v1.SetPermissionResponse{}, errors.ErrInvalidUID
	}

	if len(req.Pid) <= 0 {
		return &v1.SetPermissionResponse{}, errors.ErrInvalidParams
	}

	return r.uc.SetPermission(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) AdminUserDelete(ctx context.Context, req *v1.AdminUserDeleteRequest) (*v1.AdminUserDeleteResponse, error) {
	if req.Id <= 0 {
		return &v1.AdminUserDeleteResponse{}, errors.ErrInvalidParams
	}

	return &v1.AdminUserDeleteResponse{}, nil
}

func (r *RolesService) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (*v1.RoleCreateResponse, error) {
	if req.Name == "" || req.Code == "" {
		return &v1.RoleCreateResponse{}, errors.ErrInvalidParams
	}

	return r.uc.RoleCreate(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoleList(ctx context.Context, req *v1.RoleListRequest) (*v1.RoleListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	return &v1.RoleListResponse{}, nil
}

func (r *RolesService) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (*v1.RoleDeleteResponse, error) {
	if req.Id <= 0 {
		return &v1.RoleDeleteResponse{}, errors.ErrInvalidParams
	}

	return r.uc.RoleDelete(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoleEdit(ctx context.Context, req *v1.RoleEditRequest) (*v1.RoleEditResponse, error) {
	if req.Id <= 0 || (req.Name == "" && req.Code == "") {
		return &v1.RoleEditResponse{}, errors.ErrInvalidParams
	}

	return r.uc.RoleEdit(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) PermissionCreate(ctx context.Context, req *v1.PermissionCreateRequest) (*v1.PermissionCreateResponse, error) {
	return &v1.PermissionCreateResponse{}, nil
}

func (r *RolesService) PermissionList(ctx context.Context, req *v1.PermissionListRequest) (*v1.PermissionListResponse, error) {
	return &v1.PermissionListResponse{}, nil
}

func (r *RolesService) PermissionDelete(ctx context.Context, req *v1.PermissionDeleteRequest) (*v1.PermissionDeleteResponse, error) {
	return &v1.PermissionDeleteResponse{}, nil
}

func (r *RolesService) PermissionEdit(ctx context.Context, req *v1.PermissionEditRequest) (*v1.PermissionEditResponse, error) {
	return &v1.PermissionEditResponse{}, nil
}
