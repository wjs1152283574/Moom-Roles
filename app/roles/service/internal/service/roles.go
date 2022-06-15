package service

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func (r *RolesService) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest) (*v1.RoleCreateResponse, error) {
	if req.Name == "" || req.Code == "" {
		return &v1.RoleCreateResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoleCreate(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoleList(ctx context.Context, req *v1.RoleListRequest) (*v1.RoleListResponse, error) {
	if req.Page <= 0 {
		req.Page = 1
	}

	if req.Limit <= 0 || req.Limit > 100 {
		req.Limit = 20
	}

	return r.uc.RoleList(ctx, req)
}

func (r *RolesService) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest) (*v1.RoleDeleteResponse, error) {
	if len(req.Id) <= 0 {
		return &v1.RoleDeleteResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoleDelete(ctx, req, r.GetUserID(ctx))
}

func (r *RolesService) RoleEdit(ctx context.Context, req *v1.RoleEditRequest) (*v1.RoleEditResponse, error) {
	if req.Id <= 0 || (req.Name == "" && req.Code == "") {
		return &v1.RoleEditResponse{}, errors.ErrInvalidParams()
	}

	return r.uc.RoleEdit(ctx, req, r.GetUserID(ctx))
}
