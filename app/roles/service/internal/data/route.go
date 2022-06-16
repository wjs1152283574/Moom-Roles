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

func (r *roleRepo) RouteCreate(ctx context.Context, uid, method int32, url string) error {
	var route model.Route
	route.URL = url
	route.Method = int64(method)
	route.Commom = model.Commom{CreatorID: int64(uid), CreatedTime: time.Now().Unix()}
	if err := r.data.db.Create(&route).Error; err != nil {
		return errors.ErrSystemBusy(err)
	}

	return nil
}

func (r *roleRepo) RouteList(ctx context.Context, page, limit, method int32, url string) ([]model.Route, int64, error) {
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

func (r *roleRepo) RouteEdit(ctx context.Context, id, method int32, url string) error {
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

func (r *roleRepo) RouteDelete(ctx context.Context, ids []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, id := range ids {
			if err := tx.Unscoped().Where("id = ?", id).Delete(&model.Route{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
			// 删除关联表数据
			if err := tx.Unscoped().Where("route = ?", id).Delete(&model.RouteRole{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
			if err := tx.Unscoped().Where("route = ?", id).Delete(&model.RoutePermission{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}

		return nil
	})
}

func (r *roleRepo) RouteRole(ctx context.Context, uid, route int32, role []int32) error {
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

func (r *roleRepo) RouteRoleDelete(ctx context.Context, id int32, role []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range role {
			if err := tx.Unscoped().Where("`route` = ? and `role` = ?", id, v).Delete(&model.RouteRole{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

func (r *roleRepo) RoutePermission(ctx context.Context, uid, route int32, permission []int32) error {
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

func (r *roleRepo) RoutePermissionDelete(ctx context.Context, id int32, permission []int32) error {
	return r.data.db.Transaction(func(tx *gorm.DB) error {
		for _, v := range permission {
			if err := tx.Unscoped().Where("`route` = ? and `permission` = ?", id, v).Delete(&model.RoutePermission{}).Error; err != nil {
				return errors.ErrSystemBusy(err)
			}
		}
		return nil
	})
}

func (r *roleRepo) RouteDetails(ctx context.Context, routeID int32) (route model.Route, role []model.Role, permission []model.Permission, err error) {
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
