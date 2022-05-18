package model

var (
	UserTableName           = "users"
	RoleTableName           = "roles"
	PermissionTableName     = "permissions"
	UserRoleTablename       = "user_roles"
	RolePermissionTablename = "role_permissions"
	UserPermissionTablename = "user_permissions"
)

type Commom struct {
	CreatorID   int64  `gorm:"type:int(11);COMMENT:创建者ID"`
	CreatorName string `gorm:"type:varchar(50);COMMENT:创建者名称（列表冗余）"`

	UpdatedTime int64 `gorm:"type:bigint(20);COMMENT:最后修改时间"`
	CreatedTime int64 `gorm:"type:bigint(20);COMMENT:创建时间"`
	DeleteTime  int64 `gorm:"type:bigint(20);COMMENT:删除时间"`
}

type User struct {
	ID     uint   `gorm:"primarykey"`
	Name   string `gorm:"type:varchar(50);NOT NULL;UNIQUE;COMMENT:用户名"`
	Pass   string `gorm:"type:varchar(200);NOT NULL;COMMENT:登陆密码"`
	Type   int64  `gorm:"type:int(11);NOT NULL;DEFAULT 1;COMMENT:1-普通管理员 2-超级管理员"`
	Status int64  `gorm:"type:int(11);NOT NULL;DEFAULT 1;COMMENT:1-正常 2-冻结"`
	Icon   string `gorm:"type:varchar(500);COMMENT:头像"`

	Commom
}

func (u *User) TableName() string {
	return UserTableName
}

type Role struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(100);NOT NULL;UNIQUE;COMMENT:角色名称"`
	Code string `gorm:"type:varchar(50);NOT NULL;UNIQUE;COMMENT:角色代码"`

	Commom
}

func (r *Role) TableName() string {
	return RoleTableName
}

type Permission struct {
	ID   uint   `gorm:"primarykey"`
	Name string `gorm:"type:varchar(100);NOT NULL;UNIQUE;COMMENT:权限名称"`
	Code string `gorm:"type:varchar(50);NOT NULL;UNIQUE;COMMENT:权限代码"`

	Commom
}

func (p *Permission) TableName() string {
	return PermissionTableName
}

type UserRole struct {
	ID  uint `gorm:"primarykey"`
	UID uint `gorm:"COMMENT:用户ID"`
	RID uint `gorm:"COMMENT:角色ID"`

	Commom
}

func (u *UserRole) TableName() string {
	return UserRoleTablename
}

type RolePermission struct {
	ID  uint `gorm:"primarykey"`
	PID uint `gorm:"COMMENT:权限ID"`
	RID uint `gorm:"COMMENT:角色ID"`

	Commom
}

func (r *RolePermission) TableName() string {
	return RolePermissionTablename
}

// 兼容 用户直接获取某个单独的权限
type UserPermission struct {
	ID  uint `gorm:"primarykey"`
	UID uint `gorm:"COMMENT:用户ID"`
	PID uint `gorm:"COMMENT:权限ID"`

	Commom
}

func (u *UserPermission) TableName() string {
	return UserPermissionTablename
}
