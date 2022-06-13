package service

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func (r *RolesService) CheckRole(ctx context.Context, req *v1.CheckRoleRequest) (*v1.CheckRoleResponse, error) {
	if req.Id <= 0 || req.Code == "" {
		return &v1.CheckRoleResponse{}, errors.ErrInvalidParams
	}

	return r.uc.CheckRole(ctx, req)
}

func (r *RolesService) CheckPermission(ctx context.Context, req *v1.CheckPermissionRequest) (*v1.CheckPermissionResponse, error) {
	if req.Id <= 0 || req.Code == "" {
		return &v1.CheckPermissionResponse{}, errors.ErrInvalidParams
	}

	return r.uc.CheckPermission(ctx, req)
}

func (r *RolesService) CheckRouteRoleByToken(ctx context.Context, req *v1.CheckRouteRoleByTokenRequest) (*v1.CheckRouteRoleByTokenResponse, error) {
	if req.Token == "" || req.Code == "" {
		return &v1.CheckRouteRoleByTokenResponse{}, errors.ErrInvalidParams
	}

	return r.uc.CheckRouteRoleByToken(ctx, req)
}

func (r *RolesService) CheckRouteRoleByID(ctx context.Context, req *v1.CheckRouteRoleByIDRequest) (*v1.CheckRouteRoleByIDResponse, error) {
	if req.Id <= 0 || req.Code == "" {
		return &v1.CheckRouteRoleByIDResponse{}, errors.ErrInvalidParams
	}

	return r.uc.CheckRouteRoleByID(ctx, req)
}

func (r *RolesService) CheckRoutePermissionByToken(ctx context.Context, req *v1.CheckRoutePermissionByTokenRequest) (*v1.CheckRoutePermissionByTokenResponse, error) {
	if req.Token == "" || req.Code == "" {
		return &v1.CheckRoutePermissionByTokenResponse{}, errors.ErrInvalidParams
	}

	return r.uc.CheckRoutePermissionByToken(ctx, req)
}
