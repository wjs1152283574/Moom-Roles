package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
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
