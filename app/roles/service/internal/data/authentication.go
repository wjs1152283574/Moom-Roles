package data

import (
	"context"

	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func (r *roleRepo) CheckRole(ctx context.Context, uid int32, code string) error {
	var role model.Role
	if err := r.data.db.Table(model.RoleTableName).Where("code = ?", code).First(&role).Error; err != nil {
		return errors.ErrRoleNotExit(err)
	}

	var userRole model.UserRole
	if err := r.data.db.Table(model.UserRoleTablename).Where("user = ? and `role` = ?", uid, role.ID).First(&userRole); err != nil {
		return errors.ErrPermissionDeni(err)
	}

	return nil
}

func (r *roleRepo) CheckPermission(ctx context.Context, uid int32, code string) error {
	var permission model.Permission
	if err := r.data.db.Table(model.PermissionTableName).Where("code = ?", code).First(&permission).Error; err != nil {
		return errors.ErrRoleNotExit(err)
	}

	var userPermission model.UserPermission
	if err := r.data.db.Table(model.UserPermissionTablename).Where("user = ? and permission = ?", uid, permission.ID).First(&userPermission); err != nil {
		return errors.ErrPermissionDeni(err.Statement.Vars...)
	}

	return nil
}
