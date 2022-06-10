package biz

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
)

func (r *RolesUseCase) CheckRole(ctx context.Context, req *v1.CheckRoleRequest) (*v1.CheckRoleResponse, error) {
	return &v1.CheckRoleResponse{}, nil
}

func (r *RolesUseCase) CheckPermission(ctx context.Context, req *v1.CheckPermissionRequest) (*v1.CheckPermissionResponse, error) {
	return &v1.CheckPermissionResponse{}, nil
}
