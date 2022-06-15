package service

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func (r *RolesService) RouteCreate(ctx context.Context, req *v1.RouteCreateRequest) (*v1.RouteCreateResponse, error) {
	if req.Method <= 0 || req.Url == "" {
		return &v1.RouteCreateResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteCreate(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RouteList(ctx context.Context, req *v1.RouteListRequest) (*v1.RouteListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	return r.uc.RouteList(ctx, req)
}

func (r *RolesService) RouteEdit(ctx context.Context, req *v1.RouteEditRequest) (*v1.RouteEditResponse, error) {
	if req.Id <= 0 || (req.Method <= 0 && req.Url == "") {
		return &v1.RouteEditResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteEdit(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RouteDelete(ctx context.Context, req *v1.RouteDeleteRequest) (*v1.RouteDeleteResponse, error) {
	if len(req.Id) <= 0 {
		return &v1.RouteDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteDelete(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RouteRole(ctx context.Context, req *v1.RouteRoleRequest) (*v1.RouteRoleResponse, error) {
	if len(req.Role) <= 0 || req.Route <= 0 {
		return &v1.RouteRoleResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteRole(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RouteRoleDelete(ctx context.Context, req *v1.RouteRoleDeleteRequest) (*v1.RouteRoleDeleteResponse, error) {
	if req.Id <= 0 || len(req.Route) <= 0 {
		return &v1.RouteRoleDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteRoleDelete(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoutePermission(ctx context.Context, req *v1.RoutePermissionRequest) (*v1.RoutePermissionResponse, error) {
	if len(req.Permisson) <= 0 || req.Route <= 0 {
		return &v1.RoutePermissionResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoutePermission(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoutePermissionDelete(ctx context.Context, req *v1.RoutePermissionDeleteRequest) (*v1.RoutePermissionDeleteResponse, error) {
	if req.Id <= 0 || len(req.Permission) <= 0 {
		return &v1.RoutePermissionDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoutePermissionDelete(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RouteDetails(ctx context.Context, req *v1.RouteDetailsRequest) (*v1.RouteDetailsResponse, error) {
	if req.Id <= 0 {
		return &v1.RouteDetailsResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RouteDetails(ctx, req)
}
