package biz

import (
	"context"

	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
)

type RolesRepo interface {
	// 批量新建用户
	CreateUser(ctx context.Context, data []model.User) error
	// 检测用户是否存在
	CheckUser(ctx context.Context, name string) (model.User, error)
	// 存入redis ttl:second
	RedisSet(ctx context.Context, key, val string, ttl int64) error
	// 获取用户列表
	UserList(ctx context.Context, name, cname string, page, limit int64, typ, status []int64) ([]model.User, int64, error)
}
