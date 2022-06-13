package service

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func (r *RolesService) CheckRole(ctx context.Context, req *v1.CheckRoleRequest) (*v1.CheckRoleResponse, error) {
	if req.Id <= 0 || req.Code == "" {
		return &v1.CheckRoleResponse{}, errors.ErrSystemBusy
	}

	return r.uc.CheckRole(ctx, req)
}

func (r *RolesService) CheckPermission(ctx context.Context, req *v1.CheckPermissionRequest) (*v1.CheckPermissionResponse, error) {
	if req.Id <= 0 || req.Code == "" {
		return &v1.CheckPermissionResponse{}, errors.ErrSystemBusy
	}

	return r.uc.CheckPermission(ctx, req)
}

func (r *RolesService) CheckRouteRoleByToken(ctx context.Context, req *v1.CheckRouteRoleByTokenRequest) (*v1.CheckRouteRoleByTokenResponse, error) {
	if req.Token == "" || req.Code == "" {
		return &v1.CheckRouteRoleByTokenResponse{}, errors.ErrSystemBusy
	}

	return r.uc.CheckRouteRoleByToken(ctx, req)
}
