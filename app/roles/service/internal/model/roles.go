package model

var (
	UserTableName            = "users"
	RoleTableName            = "roles"
	PermissionTableName      = "permissions"
	UserRoleTablename        = "user_roles"
	RolePermissionTablename  = "role_permissions"
	UserPermissionTablename  = "user_permissions"
	RouteTablename           = "routes"
	RouteRoleTablename       = "route_roles"
	RoutePermissionTablename = "route_permissions"
)

type User struct {
	Commom
	Name   string `gorm:"type:varchar(50);NOT NULL;UNIQUE;COMMENT:用户名"`
	Pass   string `gorm:"type:varchar(200);NOT NULL;COMMENT:登陆密码"`
	Type   int64  `gorm:"type:int(11);NOT NULL;DEFAULT 1;COMMENT:1-普通管理员 2-超级管理员"`
	Status int64  `gorm:"type:int(11);NOT NULL;DEFAULT 1;COMMENT:1-正常 2-冻结"`
	Icon   string `gorm:"type:varchar(500);COMMENT:头像"`
}

func (u *User) TableName() string {
	return UserTableName
}

type Role struct {
	Name string `gorm:"type:varchar(100);NOT NULL;UNIQUE;COMMENT:角色名称"`
	Code string `gorm:"type:varchar(50);NOT NULL;UNIQUE;COMMENT:角色代码"`

	Commom
}

func (r *Role) TableName() string {
	return RoleTableName
}

type Permission struct {
	Name string `gorm:"type:varchar(100);NOT NULL;UNIQUE;COMMENT:权限名称"`
	Code string `gorm:"type:varchar(50);NOT NULL;UNIQUE;COMMENT:权限代码"`

	Commom
}

func (p *Permission) TableName() string {
	return PermissionTableName
}

type UserRole struct {
	UID uint `gorm:"COMMENT:用户ID"`
	RID uint `gorm:"COMMENT:角色ID"`

	Commom
}

func (u *UserRole) TableName() string {
	return UserRoleTablename
}

type RolePermission struct {
	PID uint `gorm:"COMMENT:权限ID"`
	RID uint `gorm:"COMMENT:角色ID"`

	Commom
}

func (r *RolePermission) TableName() string {
	return RolePermissionTablename
}

// 兼容 用户直接获取某个单独的权限
type UserPermission struct {
	Commom
	UID uint `gorm:"COMMENT:用户ID"`
	PID uint `gorm:"COMMENT:权限ID"`
}

func (u *UserPermission) TableName() string {
	return UserPermissionTablename
}

type Route struct {
	Commom

	URL    string `gorm:"COMMENT:路由"`
	Method int64  `gorm:"COMMENT:对应路由方法：1-get 2-post 3-put 4-delete ..."`
}

func (u *Route) TableName() string {
	return RouteTablename
}

type RouteRole struct {
	Commom

	RID    int64 `gorm:"COMMENT:路由ID"`
	RoleID uint  `gorm:"COMMENT:角色ID"`
}

func (u *RouteRole) TableName() string {
	return RouteRoleTablename
}

type RoutePermission struct {
	Commom

	RID int64 `gorm:"COMMENT:路由ID"`
	PID uint  `gorm:"COMMENT:权限ID"`
}

func (u *RoutePermission) TableName() string {
	return RoutePermissionTablename
}
