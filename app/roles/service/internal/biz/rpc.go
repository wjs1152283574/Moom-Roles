package biz

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
)

func (r *RolesUseCase) CheckRole(ctx context.Context, req *v1.CheckRoleRequest) (*v1.CheckRoleResponse, error) {
	if err := r.repo.CheckRole(ctx, req.Id, req.Code); err != nil {
		return &v1.CheckRoleResponse{Result: false}, err
	}

	return &v1.CheckRoleResponse{
		Result: true,
	}, nil
}

func (r *RolesUseCase) CheckPermission(ctx context.Context, req *v1.CheckPermissionRequest) (*v1.CheckPermissionResponse, error) {
	if err := r.repo.CheckPermission(ctx, req.Id, req.Code); err != nil {
		return &v1.CheckPermissionResponse{Result: false}, err
	}

	return &v1.CheckPermissionResponse{
		Result: true,
	}, nil
}
