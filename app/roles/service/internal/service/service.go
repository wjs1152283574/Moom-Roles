package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRolesService)

type RolesService struct {
	v1.UnimplementedRolesServer
	uc  *biz.RolesUseCase
	log *log.Helper
}

func NewRolesService(uc *biz.RolesUseCase, logger log.Logger) *RolesService {
	return &RolesService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/roles"))}

}
