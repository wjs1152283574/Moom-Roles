package biz

import (
	"context"
	"strconv"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/pkg/errors"
	"github.com/it-moom/moom-roles/pkg/tool"
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

func (r *RolesUseCase) CheckRouteRoleByToken(ctx context.Context, req *v1.CheckRouteRoleByTokenRequest) (*v1.CheckRouteRoleByTokenResponse, error) {
	tokenCliam, err := tool.NewJWT(conf.GB.TokenScreat).ParseToken(req.Token)
	if err != nil {
		return &v1.CheckRouteRoleByTokenResponse{Result: false}, errors.ErrInvalidToken
	}

	uid, err := strconv.Atoi(tokenCliam.Subject)
	if err != nil {
		return &v1.CheckRouteRoleByTokenResponse{Result: false}, errors.ErrInvalidToken
	}

	if err := r.repo.CheckRole(ctx, int64(uid), req.Code); err != nil {
		return &v1.CheckRouteRoleByTokenResponse{Result: false}, err
	}

	return &v1.CheckRouteRoleByTokenResponse{Result: true}, nil
}

func (r *RolesUseCase) CheckRouteRoleByID(ctx context.Context, req *v1.CheckRouteRoleByIDRequest) (*v1.CheckRouteRoleByIDResponse, error) {
	if err := r.repo.CheckRole(ctx, req.Id, req.Code); err != nil {
		return &v1.CheckRouteRoleByIDResponse{Result: false}, err
	}

	return &v1.CheckRouteRoleByIDResponse{Result: true}, nil
}
