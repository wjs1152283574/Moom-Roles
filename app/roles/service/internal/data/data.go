package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"github.com/it-moom/moom-roles/app/roles/service/internal/biz"
	"gorm.io/gorm"
)

var _ biz.RolesRepo = (*roleRepo)(nil)

type roleRepo struct {
	data *Data
	log  *log.Helper
}

// Data .
type Data struct {
	rd  *redis.Client
	db  *gorm.DB
	log *log.Helper
}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDB, NewRd, NewroleRepo)

func NewroleRepo(data *Data, logger log.Logger) biz.RolesRepo {
	return &roleRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/roles")),
	}
}

// NewData .
func NewData(db *gorm.DB, rd *redis.Client, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(log.With(logger, "module", "roles-service/data"))

	d := &Data{
		rd:  rd,
		db:  db,
		log: log,
	}

	return d, func() {

	}, nil
}
