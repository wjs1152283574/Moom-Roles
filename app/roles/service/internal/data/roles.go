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
