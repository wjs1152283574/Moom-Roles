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
	// 设置用户角色--解除
	SetRolesDelete(ctx context.Context, uid int64, role []int64) error
	// 设置用户权限
	SetPermissions(ctx context.Context, uid, creator int64, pid []int64) error
	// 设置用户权限--解除
	SetPermissionDelete(ctx context.Context, uid int64, permission []int64) error
	// 删除用户
	UserDelete(ctx context.Context, uid int64) error
	// 创建角色
	RoleCreate(ctx context.Context, creator int64, name, code string) error
	// 角色列表
	RoleList(ctx context.Context, page, limit int64, name, code string) ([]model.Role, int64, error)
	// 删除角色
	RoleDelete(ctx context.Context, id, creator int64) error
	// 编辑角色
	RoleEdit(ctx context.Context, id, creator int64, name, code string) error
	// 创建权限
	PermissionCreate(ctx context.Context, creator int64, name, code string) error
	// 权限列表
	PermissionList(ctx context.Context, page, limit int64, name, code string) ([]model.Permission, int64, error)
	// 删除权限
	PermissionDelete(ctx context.Context, id int64) error
	// 编辑权限
	PermissionEdit(ctx context.Context, id int64, name, code string) error
	// 创建路由
	RouteCreate(ctx context.Context, uid, method int64, url string) error
	// 路由列表
	RouteList(ctx context.Context, page, limit, method int64, url string) ([]model.Route, int64, error)
	// 路由编辑
	RouteEdit(ctx context.Context, id, method int64, url string) error
	// 删除路由
	RouteDelete(ctx context.Context, id int64) error
	// 设置路由角色
	RouteRole(ctx context.Context, uid, route int64, role []int64) error
	// 解除路由角色
	RouteRoleDelete(ctx context.Context, id int64, role []int64) error
	// 设置路由权限
	RoutePermission(ctx context.Context, uid, route int64, permission []int64) error
	// 解除路由权限
	RoutePermissionDelete(ctx context.Context, id int64, permission []int64) error
	// 路由详细
	RouteDetails(ctx context.Context, routeID int64) (route model.Route, role []model.Role, permission []model.Permission, err error)

	Rpc
}

type Rpc interface {
	CheckRole(ctx context.Context, uid int64, code string) error
	CheckPermission(ctx context.Context, uid int64, code string) error
}
