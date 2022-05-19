package biz

import (
	"context"

	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
)

type RolesRepo interface {
	// 批量新建用户
	CreateUser(ctx context.Context, data []model.User) error
}
