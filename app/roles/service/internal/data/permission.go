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
