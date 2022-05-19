package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewRolesUseCase)

type RolesUseCase struct {
	repo RolesRepo
	log  *log.Helper
}

func NewRolesUseCase(repo RolesRepo, logger log.Logger) *RolesUseCase {
	return &RolesUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "biz/roles"))}
}
