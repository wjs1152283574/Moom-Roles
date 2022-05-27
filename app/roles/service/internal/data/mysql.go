package data

import (
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	log := log.NewHelper(log.With(logger, "module", "user-service/data/gorm"))
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}

	// TODO: remove debug mode
	sqlDB, err := db.Debug().DB() // 维护链接池
	if err != nil {
		db.Statement.ReflectValue.Close()
		log.Fatalf("failed making connection-pool: %v", err)
	}

	sqlDB.SetMaxIdleConns(int(conf.Database.MaxIdle))                                    // 空闲最大数量
	sqlDB.SetMaxOpenConns(int(conf.Database.MaxOpen))                                    // 最大链接数
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(conf.Database.ConnLifeTimeHours)) // 最大可复用时间

	if err := db.AutoMigrate(&model.User{}, &model.Role{}, &model.Permission{}, &model.UserRole{}, &model.RolePermission{}, &model.UserPermission{}, &model.Route{}, &model.RouteRole{}, &model.RoutePermission{}); err != nil {
		log.Fatal(err)
	}
	return db
}
