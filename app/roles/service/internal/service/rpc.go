package service

import (
	"context"

	pb "github.com/it-moom/moom-roles/api/roles/service/v1"
)

func (r *RolesService) CheckRole(ctx context.Context, req *pb.CheckRoleRequest) (*pb.CheckRoleResponse, error) {
	return &pb.CheckRoleResponse{}, nil
}

func (r *RolesService) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	return &pb.CheckPermissionResponse{}, nil
}
