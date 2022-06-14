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
	if !conf.GB.Global.Verify {
		return &v1.GetCaptchaResponse{}, errors.ErrNoNeedCaptcha()
	}

	id, bs64, err := tool.GetCaptcha()
	if err != nil {
		return &v1.GetCaptchaResponse{}, errors.ErrSystemBusy()
	}

	return &v1.GetCaptchaResponse{
		Key:    id,
		Verify: bs64,
	}, nil
}

// Login 登陆
func (r *RolesService) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	if !tool.VerifyNameFormat(req.Name) {
		return &v1.LoginResponse{}, errors.ErrInvalidUsername()
	}

	if !tool.VerifyPassFormat(req.Pass) {
		return &v1.LoginResponse{}, errors.ErrInvalidPass()
	}

	return r.uc.Login(ctx, req)
}

func (r *RolesService) CreateAdminUser(ctx context.Context, req *v1.CreateAdminUserRequest) (*v1.CreateAdminUserResponse, error) {
	if !tool.VerifyNameFormat(req.Name) {
		return &v1.CreateAdminUserResponse{}, errors.ErrInvalidUsername()
	}

	if !tool.VerifyPassFormat(req.Pass) {
		return &v1.CreateAdminUserResponse{}, errors.ErrInvalidPass()
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
		return &v1.AdminUserEditResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.AdminUserEdit(ctx, req)
}

func (r *RolesService) SetRoles(ctx context.Context, req *v1.SetRolesRequest) (*v1.SetRolesResponse, error) {
	if req.Uid <= 0 {
		return &v1.SetRolesResponse{}, errors.ErrInvalidUID()
	}

	if len(req.Rid) <= 0 {
		return &v1.SetRolesResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.SetRoles(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) SetRolesDelete(ctx context.Context, req *v1.SetRolesDeleteRequest) (*v1.SetRolesDeleteResponse, error) {
	if req.Id <= 0 {
		return &v1.SetRolesDeleteResponse{}, errors.ErrInvalidUID()
	}

	if len(req.Role) <= 0 {
		return &v1.SetRolesDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.SetRolesDelete(ctx, req)
}

func (r *RolesService) SetPermission(ctx context.Context, req *v1.SetPermissionRequest) (*v1.SetPermissionResponse, error) {
	if req.Uid <= 0 {
		return &v1.SetPermissionResponse{}, errors.ErrInvalidUID()
	}

	if len(req.Pid) <= 0 {
		return &v1.SetPermissionResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.SetPermission(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) SetPermissionDelete(ctx context.Context, req *v1.SetPermissionDeleteRequest) (*v1.SetPermissionDeleteResponse, error) {
	if req.Id <= 0 {
		return &v1.SetPermissionDeleteResponse{}, errors.ErrInvalidUID()
	}

	if len(req.Permission) <= 0 {
		return &v1.SetPermissionDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.SetPermissionDelete(ctx, req)
}

func (r *RolesService) AdminUserDelete(ctx context.Context, req *v1.AdminUserDeleteRequest) (*v1.AdminUserDeleteResponse, error) {
	if req.Id <= 0 {
		return &v1.AdminUserDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.AdminUserDelete(ctx, req)
}

func (r *RolesService) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (*v1.RoleCreateResponse, error) {
	if req.Name == "" || req.Code == "" {
		return &v1.RoleCreateResponse{}, errors.ErrInvalidParams()
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

	return r.uc.RoleList(ctx, req)
}

func (r *RolesService) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (*v1.RoleDeleteResponse, error) {
	if len(req.Id) <= 0 {
		return &v1.RoleDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoleDelete(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoleEdit(ctx context.Context, req *v1.RoleEditRequest) (*v1.RoleEditResponse, error) {
	if req.Id <= 0 || (req.Name == "" && req.Code == "") {
		return &v1.RoleEditResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoleEdit(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) PermissionCreate(ctx context.Context, req *v1.PermissionCreateRequest) (*v1.PermissionCreateResponse, error) {
	if req.Name == "" || req.Code == "" {
		return &v1.PermissionCreateResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.PermissionCreate(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) PermissionList(ctx context.Context, req *v1.PermissionListRequest) (*v1.PermissionListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	return r.uc.PermissionList(ctx, req)
}

func (r *RolesService) PermissionDelete(ctx context.Context, req *v1.PermissionDeleteRequest) (*v1.PermissionDeleteResponse, error) {
	if len(req.Id) <= 0 {
		return &v1.PermissionDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.PermissionDelete(ctx, req)
}

func (r *RolesService) PermissionEdit(ctx context.Context, req *v1.PermissionEditRequest) (*v1.PermissionEditResponse, error) {
	if req.Id <= 0 || (req.Name == "" && req.Code == "") {
		return &v1.PermissionEditResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.PermissionEdit(ctx, req)
}

func (r *RolesService) RouteCreate(ctx context.Context, req *v1.RouteCreateRequest) (*v1.RouteCreateResponse, error) {
	if req.Method <= 0 || req.Url == "" {
		return &v1.RouteCreateResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteCreate(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RouteList(ctx context.Context, req *v1.RouteListRequest) (*v1.RouteListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	return r.uc.RouteList(ctx, req)
}

func (r *RolesService) RouteEdit(ctx context.Context, req *v1.RouteEditRequest) (*v1.RouteEditResponse, error) {
	if req.Id <= 0 || (req.Method <= 0 && req.Url == "") {
		return &v1.RouteEditResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteEdit(ctx, req)
}

func (r *RolesService) RouteDelete(ctx context.Context, req *v1.RouteDeleteRequest) (*v1.RouteDeleteResponse, error) {
	if len(req.Id) <= 0 {
		return &v1.RouteDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteDelete(ctx, req)
}

func (r *RolesService) RouteRole(ctx context.Context, req *v1.RouteRoleRequest) (*v1.RouteRoleResponse, error) {
	if len(req.Role) <= 0 || req.Route <= 0 {
		return &v1.RouteRoleResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteRole(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RouteRoleDelete(ctx context.Context, req *v1.RouteRoleDeleteRequest) (*v1.RouteRoleDeleteResponse, error) {
	if req.Id <= 0 || len(req.Route) <= 0 {
		return &v1.RouteRoleDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteRoleDelete(ctx, req)
}

func (r *RolesService) RoutePermission(ctx context.Context, req *v1.RoutePermissionRequest) (*v1.RoutePermissionResponse, error) {
	if len(req.Permisson) <= 0 || req.Route <= 0 {
		return &v1.RoutePermissionResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoutePermission(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoutePermissionDelete(ctx context.Context, req *v1.RoutePermissionDeleteRequest) (*v1.RoutePermissionDeleteResponse, error) {
	if req.Id <= 0 || len(req.Permission) <= 0 {
		return &v1.RoutePermissionDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoutePermissionDelete(ctx, req)
}

func (r *RolesService) RouteDetails(ctx context.Context, req *v1.RouteDetailsRequest) (*v1.RouteDetailsResponse, error) {
	if req.Id <= 0 {
		return &v1.RouteDetailsResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteDetails(ctx, req)
}
