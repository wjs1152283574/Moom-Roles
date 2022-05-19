package service

import (
	"context"

	pb "github.com/it-moom/moom-roles/api/roles/service/v1"
)

// CreateSuperUser 生成超级用户
func (s *RolesService) CreateSuperUser(ctx context.Context, req *pb.CreateSuperUserRequest) (*pb.CreateSuperUserResponse, error) {
	return s.uc.CreateSuperUser(ctx)
}
