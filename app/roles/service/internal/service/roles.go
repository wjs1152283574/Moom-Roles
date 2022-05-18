package service

import (
	"context"

	pb "github.com/it-moom/moom-roles/api/roles/service/v1"
)

func (s *RolesService) CreateRoles(ctx context.Context, req *pb.CreateRolesRequest) (*pb.CreateRolesReply, error) {
	return &pb.CreateRolesReply{}, nil
}
