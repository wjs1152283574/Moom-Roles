package service

import (
	"context"

	pb "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/pkg/errors"
	"github.com/it-moom/moom-roles/pkg/tool"
)

// CreateSuperUser 生成超级用户
func (r *RolesService) CreateSuperUser(ctx context.Context, req *pb.CreateSuperUserRequest) (*pb.CreateSuperUserResponse, error) {
	return r.uc.CreateSuperUser(ctx)
}

// GetCaptcha 获取图片验证码
func (r *RolesService) GetCaptcha(ctx context.Context, req *pb.GetCaptchaRequest) (*pb.GetCaptchaResponse, error) {
	if !conf.GB.Verify {
		return &pb.GetCaptchaResponse{}, errors.ErrNoNeedCaptcha
	}

	id, bs64, err := tool.GetCaptcha()
	if err != nil {
		return &pb.GetCaptchaResponse{}, errors.ErrSystemBusy
	}

	return &pb.GetCaptchaResponse{
		Key:    id,
		Verify: bs64,
	}, nil
}

// Login 登陆
func (r *RolesService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if !tool.VerifyNameFormat(req.Name) {
		return &pb.LoginResponse{}, errors.ErrInvalidUsername
	}

	if !tool.VerifyPassFormat(req.Pass) {
		return &pb.LoginResponse{}, errors.ErrInvalidPass
	}

	return r.uc.Login(ctx, req)
}

func (r *RolesService) CreateAdminUser(ctx context.Context, req *pb.CreateAdminUserRequest) (*pb.CreateAdminUserResponse, error) {
	if !tool.VerifyNameFormat(req.Name) {
		return &pb.CreateAdminUserResponse{}, errors.ErrInvalidUsername
	}

	if !tool.VerifyPassFormat(req.Pass) {
		return &pb.CreateAdminUserResponse{}, errors.ErrInvalidPass
	}

	return r.uc.CreateAdminUser(ctx, req)
}

func (r *RolesService) AdminUserList(ctx context.Context, req *pb.AdminUserListRequest) (*pb.AdminUserListResponse, error) {
	return &pb.AdminUserListResponse{}, nil
}

func (r *RolesService) AdminUserDetails(ctx context.Context, req *pb.AdminUserDetailsRequest) (*pb.AdminUserDetailsResponse, error) {
	return &pb.AdminUserDetailsResponse{}, nil
}

func (r *RolesService) AdminUserEdit(ctx context.Context, req *pb.AdminUserEditRequest) (*pb.AdminUserEditResponse, error) {
	return &pb.AdminUserEditResponse{}, nil
}

func (r *RolesService) SetRoles(ctx context.Context, req *pb.SetRolesRequest) (*pb.SetRolesResponse, error) {
	return &pb.SetRolesResponse{}, nil
}

func (r *RolesService) SetPermission(ctx context.Context, req *pb.SetPermissionRequest) (*pb.SetPermissionResponse, error) {
	return &pb.SetPermissionResponse{}, nil
}

func (r *RolesService) AdminUserDelete(ctx context.Context, req *pb.AdminUserDeleteRequest) (*pb.AdminUserDeleteResponse, error) {
	return &pb.AdminUserDeleteResponse{}, nil
}

func (r *RolesService) RoleCreate(ctx context.Context, req *pb.RoleCreateRequest) (*pb.RoleCreateResponse, error) {
	return &pb.RoleCreateResponse{}, nil
}

func (r *RolesService) RoleList(ctx context.Context, req *pb.RoleListRequest) (*pb.RoleListResponse, error) {
	return &pb.RoleListResponse{}, nil
}

func (r *RolesService) RoleDelete(ctx context.Context, req *pb.RoleDeleteRequest) (*pb.RoleDeleteResponse, error) {
	return &pb.RoleDeleteResponse{}, nil
}

func (r *RolesService) RoleEdit(ctx context.Context, req *pb.RoleEditRequest) (*pb.RoleEditResponse, error) {
	return &pb.RoleEditResponse{}, nil
}

func (r *RolesService) PermissionCreate(ctx context.Context, req *pb.PermissionCreateRequest) (*pb.PermissionCreateResponse, error) {
	return &pb.PermissionCreateResponse{}, nil
}

func (r *RolesService) PermissionList(ctx context.Context, req *pb.PermissionListRequest) (*pb.PermissionListResponse, error) {
	return &pb.PermissionListResponse{}, nil
}

func (r *RolesService) PermissionDelete(ctx context.Context, req *pb.PermissionDeleteRequest) (*pb.PermissionDeleteResponse, error) {
	return &pb.PermissionDeleteResponse{}, nil
}

func (r *RolesService) PermissionEdit(ctx context.Context, req *pb.PermissionEditRequest) (*pb.PermissionEditResponse, error) {
	return &pb.PermissionEditResponse{}, nil
}
