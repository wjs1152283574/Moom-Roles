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

func (r *UserRepo) RedisSet(ctx context.Context, key, val string, ttl int32) error {
	if err := r.data.rd.Set(ctx, key, val, time.Duration(ttl)).Err(); err != nil {
		return errors.ErrSystemBusy(err)
	}

	return nil
}

func (r *UserRepo) UserList(ctx context.Context, name, cname string, page, limit int32, typ, status []int32) ([]model.User, int64, error) {
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

// 设置用户角色
func (r *UserRepo) SetRoles(ctx context.Context, uid, creator int32, rid []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range rid {
			var r model.Role
			if err := tx.Model(&model.Role{}).Where("id = ?", v).First(&r); err != nil {
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
			if err := tx.Model(&model.Permission{}).Where("id = ?", v).First(&r); err != nil {
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

func (r *UserRepo) UserDelete(ctx context.Context, uid int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		var user model.User
		user.ID = uint(uid)
		tx.First(&user)
		err := tx.Unscoped().Where("id = ?", uid).Delete(&user).Error
		if err != nil {
			return errors.ErrSystemBusy(err)
		}

		return nil
	})
}

func (r *UserRepo) RoleCreate(ctx context.Context, creator int32, name, code string) error {
	var role model.Role
	role.Name = name
	role.Code = code
	role.Commom = model.Commom{CreatorID: int64(creator), CreatedTime: time.Now().Unix()}
	if err := r.data.db.Table(model.RoleTableName).Create(&role).Error; err != nil {
		return errors.ErrSystemBusy(err)
	}

	return nil
}

func (r *UserRepo) RoleList(ctx context.Context, page, limit int32, name, code string) ([]model.Role, int64, error) {
	var roles []model.Role
	var total int64
	tx := r.data.db.Table(model.RoleTableName)
	if name != "" {
		key := fmt.Sprintf("%%%s%%", name)
		tx.Where("name like ?", key)
	}

	if code != "" {
		key := fmt.Sprintf("%%%s%%", code)
		tx.Where("code like ?", key)
	}

	tx.Count(&total)
	tx.Offset((int(page) - 1) * int(limit)).Limit(int(limit))

	err := tx.Scan(&roles).Error
	if err != nil {
		r.log.Errorf("[RoleList] get list err : %v", err)
	}

	return roles, total, nil
}

func (r *UserRepo) RoleDelete(ctx context.Context, ids []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			err := tx.Unscoped().Where("id = ?", id).Delete(&model.Role{}).Error
			if err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

func (r *UserRepo) RoleEdit(ctx context.Context, id, creator int32, name, code string) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		var role model.Role
		err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", id).First(&role).Error
		if err != nil {
			return errors.ErrSystemBusy(err)
		}

		if name != "" {
			role.Name = name
		}
		if code != "" {
			role.Code = code
		}
		role.UpdatedTime = time.Now().Unix()
		if err := tx.Updates(&role).Error; err != nil {
			return errors.ErrSystemBusy(err)
		}

		return nil
	})
}

func (r *UserRepo) PermissionCreate(ctx context.Context, creator int32, name, code string) error {
	var per model.Permission
	per.Code = code
	per.Name = name
	per.Commom = model.Commom{CreatorID: int64(creator), CreatedTime: time.Now().Unix()}
	if err := r.data.db.Create(&per).Error; err != nil {
		return errors.ErrMuiltiRecord(err)
	}

	return nil
}

func (r *UserRepo) PermissionList(ctx context.Context, page, limit int32, name, code string) ([]model.Permission, int64, error) {
	var pers []model.Permission
	var total int64
	tx := r.data.db.Table(model.PermissionTableName)

	if name != "" {
		key := fmt.Sprintf("%%%s%%", name)
		tx.Where("name like ?", key)
	}

	if code != "" {
		key := fmt.Sprintf("%%%s%%", code)
		tx.Where("code like ?", key)
	}

	tx.Count(&total)
	tx.Offset((int(page) - 1) * int(limit)).Limit(int(limit))

	err := tx.Scan(&pers).Error
	if err != nil {
		r.log.Errorf("[RoleList] get list err : %v", err)
	}

	return pers, total, nil
}

func (r *UserRepo) PermissionDelete(ctx context.Context, ids []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			err := tx.Unscoped().Where("id = ?", id).Delete(&model.Permission{}).Error
			if err != nil {
				return errors.ErrSystemBusy(err)
			}

		}
		return nil
	})
}

func (r *UserRepo) PermissionEdit(ctx context.Context, id int32, name, code string) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		var per model.Permission
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).Where("id = ?", id).First(&per).Error; err != nil {
			return errors.ErrSystemBusy(err)
		}

		if name != "" {
			per.Name = name
		}

		if code != "" {
			per.Code = code
		}

		if err := tx.Updates(&per).Error; err != nil {
			return errors.ErrSystemBusy(err)
		}

		return nil
	})
}

func (r *UserRepo) RouteCreate(ctx context.Context, uid, method int32, url string) error {
	var route model.Route
	route.URL = url
	route.Method = int64(method)
	route.Commom = model.Commom{CreatorID: int64(uid), CreatedTime: time.Now().Unix()}
	if err := r.data.db.Create(&route).Error; err != nil {
		return errors.ErrSystemBusy(err)
	}

	return nil
}

func (r *UserRepo) RouteList(ctx context.Context, page, limit, method int32, url string) ([]model.Route, int64, error) {
	var list []model.Route
	var total int64
	tx := r.data.db.Table(model.RouteTablename)
	if url != "" {
		key := fmt.Sprintf("%%%s%%", url)
		tx.Where("url like ?", key)
	}

	if method > 0 {
		tx.Where("method = ?", method)
	}

	tx.Count(&total)
	tx.Offset((int(page) - 1) * int(limit)).Limit(int(limit))

	err := tx.Scan(&list).Error
	if err != nil {
		r.log.Errorf("[RoleList] get list err : %v", err)
	}

	return list, total, nil
}

func (r *UserRepo) RouteEdit(ctx context.Context, id, method int32, url string) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		var route model.Route
		if err := tx.Clauses(clause.Locking{
			Strength: "UPDATE",
		}).Where("id = ?", id).First(&route).Error; err != nil {
			return errors.ErrSystemBusy(err)
		}

		if url != "" {
			route.URL = url
		}

		if method > 0 {
			route.Method = int64(method)
		}

		if err := tx.Updates(&route).Error; err != nil {
			return errors.ErrSystemBusy(err)
		}

		return nil
	})
}

func (r *UserRepo) RouteDelete(ctx context.Context, ids []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if err := tx.Unscoped().Where("id = ?", id).Delete(&model.Route{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}

		return nil
	})
}

func (r *UserRepo) RouteRole(ctx context.Context, uid, route int32, role []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		var roles model.Role
		for _, v := range role {
			if err := tx.Model(&model.Role{}).Where("id = ?", v).First(&roles); err != nil {
				return errors.ErrRoleNotExit(err)
			}

			var routeRole model.RouteRole
			err := r.data.db.Table(model.RouteRoleTablename).Where("route = ? and role = ?", route, v).First(&routeRole).Error
			if err == nil || gorm.ErrRecordNotFound != err {
				return errors.ErrMuiltiRecord(err)
			}

			routeRole.Role = uint(v)
			routeRole.Route = int64(route)
			routeRole.Commom = model.Commom{CreatorID: int64(uid), CreatedTime: time.Now().Unix()}
			if err := r.data.db.Create(&routeRole).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

func (r *UserRepo) RouteRoleDelete(ctx context.Context, id int32, role []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range role {
			if err := tx.Unscoped().Where("`route` = ? and `role` = ?", id, v).Delete(&model.RouteRole{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

func (r *UserRepo) RoutePermission(ctx context.Context, uid, route int32, permission []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		var per model.Permission
		for _, v := range permission {
			if err := tx.Model(&model.Permission{}).Where("id = ?", v).First(&per).Error; err != nil {
				return errors.ErrPermissionNotExit(err)
			}

			var routePer model.RoutePermission
			err := r.data.db.Table(model.RoutePermissionTablename).Where("route = ? and permission = ?", route, v).First(&routePer).Error
			if err == nil || gorm.ErrRecordNotFound != err {
				return errors.ErrMuiltiRecord(err)
			}

			routePer.Permission = uint(v)
			routePer.Route = int64(route)
			routePer.Commom = model.Commom{CreatorID: int64(uid), CreatedTime: time.Now().Unix()}
			if err := r.data.db.Create(&routePer).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}

		return nil
	})
}

func (r *UserRepo) RoutePermissionDelete(ctx context.Context, id int32, permission []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range permission {
			if err := tx.Unscoped().Where("`route` = ? and `permission` = ?", id, v).Delete(&model.RoutePermission{}).Error; err != nil {
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

			if err := tx.Unscoped().Where("user = ? and permisson = ?", uid, v).Delete(&model.UserPermission{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

func (r *UserRepo) RouteDetails(ctx context.Context, routeID int32) (route model.Route, role []model.Role, permission []model.Permission, err error) {
	if err = r.data.db.Table(model.RouteTablename).Where("id = ?", routeID).First(&route).Error; err != nil {
		err = errors.ErrRouteNotExit(err)
		return
	}

	routeRoleSql := r.data.db.Model(&model.RouteRole{}).Where("route = ?", routeID).Select("id")
	if err = r.data.db.Model(&model.Role{}).Where("id in ( ? )", routeRoleSql).Scan(&role).Error; err != nil {
		err = errors.ErrSystemBusy(err)
		return
	}

	routePermisionSql := r.data.db.Model(&model.RoutePermission{}).Where("route = ?", routeID).Select("id")
	if err = r.data.db.Model(&model.Permission{}).Where("id in ( ? )", routePermisionSql).Scan(&permission).Error; err != nil {
		err = errors.ErrSystemBusy(err)
		return
	}

	return
}
