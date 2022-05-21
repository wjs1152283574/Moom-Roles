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
