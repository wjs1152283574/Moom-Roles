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
	// 获取用户基本信息
	UserBaseInfos(ctx context.Context, uid int64) (model.User, error)
	// 编辑用户基本信息
	UserBaseEdit(ctx context.Context, user model.User) error
	// 获取用户角色列表
	UserRoleList(ctx context.Context, uid int64) ([]model.Role, error)
	// 获取用户权限列表
	UserPermissionList(ctx context.Context, uid int64) ([]model.Permission, error)
	// 设置用户角色
	SetRoles(ctx context.Context, uid, creator int64, rid []int64) error
	// 设置用户权限
	SetPermissions(ctx context.Context, uid, creator int64, pid []int64) error
	// 删除用户
	UserDelete(ctx context.Context, uid int64) error
}
