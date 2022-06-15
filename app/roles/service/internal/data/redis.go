package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func NewRd(conf *conf.Data, logger log.Logger) *redis.Client {
	log.NewHelper(log.With(logger, "module", "roles-service/data/redis"))
	opts := redis.Options{
		Addr:         conf.Redis.Addr,
		Username:     conf.Redis.Auth,
		Password:     conf.Redis.Password,
		ReadTimeout:  conf.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: conf.Redis.WriteTimeout.AsDuration(),
		PoolSize:     int(conf.Redis.Pool),
		DB:           int(conf.Redis.Db),
	}

	return redis.NewClient(&opts)
}

func (r *UserRepo) RedisSet(ctx context.Context, key, val string, ttl int32) error {
	if err := r.data.rd.Set(ctx, key, val, time.Duration(ttl)).Err(); err != nil {
		return errors.ErrSystemBusy(err)
	}

	return nil
}
