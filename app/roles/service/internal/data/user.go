package data

import (
	"context"
	"fmt"
	"time"

	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
	"github.com/it-moom/moom-roles/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (r *UserRepo) CreateUser(ctx context.Context, data []model.User) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range data {
			err := tx.Create(&v).Error
			if err != nil {
				return errors.ErrMuiltpleUserName(err)
			}
		}

		return nil
	})
}

func (r *UserRepo) CheckUser(ctx context.Context, name string) (model.User, error) {
	var user model.User
	if err := r.data.db.Table(model.UserTableName).Where("name = ?", name).First(&user).Error; err != nil {
		return user, errors.ErrUserNotExit(err)
	}

	return user, nil
}

func (r *UserRepo) CheckUserByID(ctx context.Context, uid int64) (model.User, error) {
	var user model.User
	if err := r.data.db.Table(model.UserTableName).Where("id = ?", uid).First(&user).Error; err != nil {
		return user, errors.ErrUserNotExit(err)
	}

	return user, nil
}

func (r *UserRepo) UserList(ctx context.Context, name, cname string, page, limit int32, typ, status []int32) ([]model.User, int64, error) {
	var users []model.User
	var total int64
	tx := r.data.db.Table(model.UserTableName)

	if name != "" {
		keys := fmt.Sprintf("%%%s%%", name)
		tx.Where("name like ?", keys)
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
		err = errors.ErrSystemBusy(err)
	}

	return users, total, err
}

// 获取用户基本信息
func (r *UserRepo) UserBaseInfos(ctx context.Context, uid int32) (model.User, error) {
	var user model.User
	user.ID = uint(uid)
	err := r.data.db.First(&user).Error
	if err != nil {
		return user, errors.ErrUserNotExit(err)
	}

	return user, nil
}

// 获取用户角色列表
func (r *UserRepo) UserRoleList(ctx context.Context, uid int32) ([]model.Role, error) {
	var role []model.Role
	userRoleSql := r.data.db.Model(&model.UserRole{}).Where("user = ?", uid).Select("`role`")
	err := r.data.db.Model(&model.Role{}).Where("id in ( ? )", userRoleSql).Scan(&role).Error
	if err != nil {
		return role, errors.ErrSystemBusy(err)
	}

	return role, nil
}

// 获取用户权限列表
func (r *UserRepo) UserPermissionList(ctx context.Context, uid int32) ([]model.Permission, error) {
	var permissions []model.Permission
	userPermissionSql := r.data.db.Model(&model.UserPermission{}).Where("user = ?", uid).Select("`permission`")
	err := r.data.db.Model(&model.Permission{}).Where("id in ( ? )", userPermissionSql).Scan(&permissions).Error
	if err != nil {
		return permissions, errors.ErrSystemBusy(err)
	}

	return permissions, nil
}

// 编辑用户基本信息
func (r *UserRepo) UserBaseEdit(ctx context.Context, user model.User) error {
	err := r.data.db.Updates(&user).Error
	if err != nil {
		return errors.ErrSystemBusy(err)
	}

	return nil
}

func (r *UserRepo) UserDelete(ctx context.Context, ids []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			var user model.User
			user.ID = uint(id)
			tx.First(&user)
			err := tx.Unscoped().Where("id = ?", id).Delete(&user).Error
			if err != nil {
				return errors.ErrSystemBusy(err)
			}

			// 同时清空关联表数据
			if err := tx.Unscoped().Where("user = ?", id).Delete(&model.UserRole{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
			if err := tx.Unscoped().Where("user = ?", id).Delete(&model.UserPermission{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}

		return nil
	})
}

func (r *UserRepo) SetRolesDelete(ctx context.Context, uid int32, role []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range role {
			var r model.Role
			if err := tx.Model(&model.Role{}).Where("id = ?", v).First(&r).Error; err != nil {
				return errors.ErrRoleNotExit(err)
			}

			if err := tx.Unscoped().Where("user = ? and `role` = ?", uid, v).Delete(&model.UserRole{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

func (r *UserRepo) SetPermissionDelete(ctx context.Context, uid int32, permission []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range permission {
			var r model.Permission
			if err := tx.Model(&model.Permission{}).Where("id = ?", v).First(&r).Error; err != nil {
				return errors.ErrPermissionNotExit(err)
			}

			if err := tx.Unscoped().Where("user = ? and permission = ?", uid, v).Delete(&model.UserPermission{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

// 设置用户角色
func (r *UserRepo) SetRoles(ctx context.Context, uid, creator int32, rid []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range rid {
			var r model.Role
			if err := tx.Model(&model.Role{}).Where("id = ?", v).First(&r).Error; err != nil {
				return errors.ErrRoleNotExit(err)
			}

			var userRole model.UserRole
			userRole.User = uint(uid)
			userRole.CreatedTime = time.Now().Unix()
			userRole.CreatorID = int64(creator)
			userRole.Role = uint(v)
			if err := tx.Create(&userRole).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}

		return nil
	})
}

func (r *UserRepo) SetPermissions(ctx context.Context, uid, creator int32, pid []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range pid {
			var r model.Permission
			if err := tx.Model(&model.Permission{}).Where("id = ?", v).First(&r).Error; err != nil {
				return errors.ErrPermissionNotExit(err)
			}

			var userRole model.UserPermission
			userRole.User = uint(uid)
			userRole.CreatedTime = time.Now().Unix()
			userRole.CreatorID = int64(creator)
			userRole.Permission = uint(v)
			if err := tx.Create(&userRole).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}

		return nil
	})
}

func (r *UserRepo) IsSuperUser(ctx context.Context, uid int64) (result bool) {
	var user model.User
	if err := r.data.db.Where("id = ?", uid).First(&user).Error; err != nil {
		return
	}

	if user.Type == 2 && user.Status == 1 {
		return true
	}

	return
}

// 更新用户状态:1-正常 2-冻结
func (r *UserRepo) UpdateUserStatus(ctx context.Context, uid, status int64) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", uid).First(&user).Error
		if err != nil {
			return errors.ErrSystemBusy(err)
		}

		user.Status = status

		err = tx.Updates(&user).Error
		if err != nil {
			return errors.ErrSystemBusy(err)
		}

		return nil
	})
}
