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
