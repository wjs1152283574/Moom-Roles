package biz

import "github.com/go-kratos/kratos/v2/log"

type RolesRepo interface {
}

type RolesUseCase struct {
	repo RolesRepo
	log  *log.Helper
}

func NewRolesUseCase(repo RolesRepo, logger log.Logger) *RolesUseCase {
	return &RolesUseCase{repo: repo, log: log.NewHelper(log.With(logger, "module", "biz/roles"))}
}
