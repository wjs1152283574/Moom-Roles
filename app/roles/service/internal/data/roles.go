package data

import (
	"context"
	"time"

	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
	"github.com/it-moom/moom-roles/pkg/errors"
	"gorm.io/gorm"
)

func (r *UserRepo) CreateUser(ctx context.Context, data []model.User) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range data {
			err := tx.Create(&v).Error
			if err != nil {
				return errors.ErrMuiltpleUserName
			}
		}

		return nil
	})
}

func (r *UserRepo) CheckUser(ctx context.Context, name string) (model.User, error) {
	var user model.User
	if err := r.data.db.Table(model.UserTableName).Where("name = ?", name).First(&user).Error; err != nil {
		return user, errors.ErrUserNotExit
	}

	return user, nil
}

func (r *UserRepo) RedisSet(ctx context.Context, key, val string, ttl int64) error {
	if err := r.data.rd.Set(ctx, key, val, time.Duration(ttl)).Err(); err != nil {
		return errors.ErrSystemBusy
	}

	return nil
}

func (r *UserRepo) UserList(ctx context.Context, name, cname string, page, limit int64, typ, status []int64) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	tx := r.data.db.Table(model.UserTableName)

	if name != "" {
		tx.Where("name = ?", name)
	}

	if cname != "" {
		tx.Where("creator_name = ?", cname)
	}

	if len(typ) > 0 {
		tx.Where("type in ?", typ)
	}

	if len(status) > 0 {
		tx.Where("status in ?", status)
	}

	tx.Count(&total)
	tx.Offset((int(page) - 1) * int(limit)).Limit(int(limit))
	tx.Order("created_time DESC") // 默认创建时间倒序
	err := tx.Scan(&users).Error
	if err != nil {
		r.log.Errorf("[UserList] get list err : %v", err)
		err = errors.ErrSystemBusy
	}

	return users, total, err
}

// 获取用户基本信息
func (r *UserRepo) UserBaseInfos(ctx context.Context, uid int64) (model.User, error) {
	var user model.User
	user.ID = uint(uid)
	err := r.data.db.First(&user).Error
	if err != nil {
		return user, errors.ErrUserNotExit
	}

	return user, nil
}

// 获取用户角色列表
func (r *UserRepo) UserRoleList(ctx context.Context, uid int64) ([]model.Role, error) {
	var role []model.Role
	err := r.data.db.Exec("SELECT * FROM roles WHERE id in (SELECT r_id FROM user_roles WHERE uid=?)", uid).Scan(&role).Error
	if err != nil {
		return role, errors.ErrSystemBusy
	}

	return role, nil
}

// 获取用户权限列表
func (r *UserRepo) UserPermissionList(ctx context.Context, uid int64) ([]model.Permission, error) {
	var permissions []model.Permission
	err := r.data.db.Exec("SELECT * FROM permissions WHERE id in (SELECT p_id FROM user_permissions WHERE uid=?)", uid).Scan(&permissions).Error
	if err != nil {
		return permissions, errors.ErrSystemBusy
	}

	return permissions, nil
}

// 编辑用户基本信息
func (r *UserRepo) UserBaseEdit(ctx context.Context, user model.User) error {
	err := r.data.db.Updates(&user).Error
	if err != nil {
		return errors.ErrSystemBusy
	}

	return nil
}

// 设置用户角色
func (r *UserRepo) SetRoles(ctx context.Context, uid, creator int64, rid []int64) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range rid {
			var userRole model.UserRole
			userRole.UID = uint(uid)
			userRole.CreatedTime = time.Now().Unix()
			userRole.CreatorID = creator
			userRole.RID = uint(v)
			if err := tx.Create(&userRole).Error; err != nil {
				return errors.ErrSystemBusy
			}
		}

		return nil
	})
}

func (r *UserRepo) SetPermissions(ctx context.Context, uid, creator int64, pid []int64) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range pid {
			var userRole model.UserPermission
			userRole.UID = uint(uid)
			userRole.CreatedTime = time.Now().Unix()
			userRole.CreatorID = creator
			userRole.PID = uint(v)
			if err := tx.Create(&userRole).Error; err != nil {
				return errors.ErrSystemBusy
			}
		}

		return nil
	})
}

func (r *UserRepo) UserDelete(ctx context.Context, uid int64) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Exec("delete from users where id = ?", uid).Error
		if err != nil {
			return errors.ErrSystemBusy
		}

		return nil
	})
}
