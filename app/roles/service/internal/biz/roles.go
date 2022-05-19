package biz

import (
	"context"
	"time"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
	"github.com/it-moom/moom-roles/pkg/errors"
	"github.com/it-moom/moom-roles/pkg/tool"
)

func (r *RolesUseCase) CreateSuperUser(ctx context.Context) (*v1.CreateSuperUserResponse, error) {
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

	if err := r.repo.CreateUser(ctx, data); err != nil {
		return &v1.CreateSuperUserResponse{}, errors.ErrSystemBusy
	}

	return &v1.CreateSuperUserResponse{}, nil
}
