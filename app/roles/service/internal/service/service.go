package service

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRolesService)
var MockUID int64 = 1024

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

// GetUserID 从上下文中获取userid；取出的为interface，需要断言
func (s *RolesService) GetUserID(ctx context.Context) int64 {
	inter_id := ctx.Value("userid")
	if inter_id == nil {
		return MockUID
	}
	id, err := strconv.Atoi(inter_id.(string))
	if err != nil {
		return MockUID
	}
	return int64(id)
}
