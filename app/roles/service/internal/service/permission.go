package service

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func (r *RolesService) PermissionCreate(ctx context.Context, req *v1.PermissionCreateRequest) (*v1.PermissionCreateResponse, error) {
	if req.Name == "" || req.Code == "" {
		return &v1.PermissionCreateResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.PermissionCreate(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) PermissionList(ctx context.Context, req *v1.PermissionListRequest) (*v1.PermissionListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	return r.uc.PermissionList(ctx, req)
}

func (r *RolesService) PermissionDelete(ctx context.Context, req *v1.PermissionDeleteRequest) (*v1.PermissionDeleteResponse, error) {
	if len(req.Id) <= 0 {
		return &v1.PermissionDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.PermissionDelete(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) PermissionEdit(ctx context.Context, req *v1.PermissionEditRequest) (*v1.PermissionEditResponse, error) {
	if req.Id <= 0 || (req.Name == "" && req.Code == "") {
		return &v1.PermissionEditResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.PermissionEdit(ctx, req, r.GetUserID(ctx))
}
