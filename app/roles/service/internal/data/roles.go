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

func (r *UserRepo) RoleEdit(ctx context.Context, id int32, name, code string) error {
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
