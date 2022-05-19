package biz

import (
	"context"
	"strconv"
	"time"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
	"github.com/it-moom/moom-roles/pkg/errors"
	"github.com/it-moom/moom-roles/pkg/key"
	"github.com/it-moom/moom-roles/pkg/tool"
)

func (r *RolesUseCase) CreateSuperUser(ctx context.Context) (*v1.CreateSuperUserResponse, error) {
	// 组装用户数据
	var data []model.User
	for _, v := range conf.SU.SuperUser {
		data = append(data, model.User{
			Name:   v.Name,
			Pass:   tool.Base64Md5(v.Pass),
			Type:   2,
			Status: 1,
			Commom: model.Commom{CreatedTime: time.Now().Unix()},
		})

	}
	// 新建用户
	if err := r.repo.CreateUser(ctx, data); err != nil {
		return &v1.CreateSuperUserResponse{}, errors.ErrSystemBusy
	}

	return &v1.CreateSuperUserResponse{}, nil
}

func (r *RolesUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// 需要检验验证码
	if conf.GB.Verify {
		if !tool.Verify(req.Key, req.Val) {
			return &v1.LoginResponse{}, errors.ErrInvalidVerifyCode
		}
	}
	// 检验用户合法性
	user, err := r.repo.CheckUser(ctx, req.Name)
	if err != nil {
		return &v1.LoginResponse{}, err
	}
	// 检验用户密码
	if tool.Base64Md5(req.Pass) != user.Pass {
		return &v1.LoginResponse{}, errors.ErrInvalidPass
	}
	// 生成token
	token, err := tool.NewJWT(conf.GB.TokenScreat).CreateToken(strconv.Itoa(int(user.ID)), conf.GB.Issuer, conf.GB.TokenTtl)
	if err != nil {
		return &v1.LoginResponse{}, err
	}
	// 存入redis
	err = r.repo.RedisSet(ctx, key.TokenKey(uint64(user.ID)), token, conf.GB.TokenTtl)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		Token: token,
	}, nil
}
