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
	user, err := r.repo.CheckUserByID(ctx, int64(req.Id))
	if err != nil {
		return &v1.CheckRoleResponse{}, err
	}

	if user.Status == 2 {
		return &v1.CheckRoleResponse{}, errors.ErrUserLockUp()
	}

	if err := r.repo.CheckRole(ctx, req.Id, req.Code); err != nil {
		return &v1.CheckRoleResponse{Result: false}, err
	}

	return &v1.CheckRoleResponse{
		Result: true,
	}, nil
}

func (r *RolesUseCase) CheckPermission(ctx context.Context, req *v1.CheckPermissionRequest) (*v1.CheckPermissionResponse, error) {
	user, err := r.repo.CheckUserByID(ctx, int64(req.Id))
	if err != nil {
		return &v1.CheckPermissionResponse{}, err
	}

	if user.Status == 2 {
		return &v1.CheckPermissionResponse{}, errors.ErrUserLockUp()
	}

	if err := r.repo.CheckPermission(ctx, req.Id, req.Code); err != nil {
		return &v1.CheckPermissionResponse{Result: false}, err
	}

	return &v1.CheckPermissionResponse{
		Result: true,
	}, nil
}

func (r *RolesUseCase) CheckRouteRoleByToken(ctx context.Context, req *v1.CheckRouteRoleByTokenRequest) (*v1.CheckRouteRoleByTokenResponse, error) {
	tokenCliam, err := tool.NewJWT(conf.GB.Global.TokenScrect).ParseToken(req.Token)
	if err != nil {
		return &v1.CheckRouteRoleByTokenResponse{Result: false}, errors.ErrInvalidToken()
	}

	uid, err := strconv.Atoi(tokenCliam.Subject)
	if err != nil {
		return &v1.CheckRouteRoleByTokenResponse{Result: false}, errors.ErrInvalidToken()
	}

	user, err := r.repo.CheckUserByID(ctx, int64(uid))
	if err != nil {
		return &v1.CheckRouteRoleByTokenResponse{}, err
	}

	if user.Status == 2 {
		return &v1.CheckRouteRoleByTokenResponse{}, errors.ErrUserLockUp()
	}

	if err := r.repo.CheckRole(ctx, int32(uid), req.Code); err != nil {
		return &v1.CheckRouteRoleByTokenResponse{Result: false}, err
	}

	return &v1.CheckRouteRoleByTokenResponse{Result: true}, nil
}

func (r *RolesUseCase) CheckRouteRoleByID(ctx context.Context, req *v1.CheckRouteRoleByIDRequest) (*v1.CheckRouteRoleByIDResponse, error) {
	user, err := r.repo.CheckUserByID(ctx, int64(req.Id))
	if err != nil {
		return &v1.CheckRouteRoleByIDResponse{}, err
	}

	if user.Status == 2 {
		return &v1.CheckRouteRoleByIDResponse{}, errors.ErrUserLockUp()
	}

	if err := r.repo.CheckRole(ctx, req.Id, req.Code); err != nil {
		return &v1.CheckRouteRoleByIDResponse{Result: false}, err
	}

	return &v1.CheckRouteRoleByIDResponse{Result: true}, nil
}

func (r *RolesUseCase) CheckRoutePermissionByToken(ctx context.Context, req *v1.CheckRoutePermissionByTokenRequest) (*v1.CheckRoutePermissionByTokenResponse, error) {
	tokenCliam, err := tool.NewJWT(conf.GB.Global.TokenScrect).ParseToken(req.Token)
	if err != nil {
		return &v1.CheckRoutePermissionByTokenResponse{Result: false}, errors.ErrInvalidToken()
	}

	uid, err := strconv.Atoi(tokenCliam.Subject)
	if err != nil {
		return &v1.CheckRoutePermissionByTokenResponse{Result: false}, errors.ErrInvalidToken()
	}

	user, err := r.repo.CheckUserByID(ctx, int64(uid))
	if err != nil {
		return &v1.CheckRoutePermissionByTokenResponse{}, err
	}

	if user.Status == 2 {
		return &v1.CheckRoutePermissionByTokenResponse{}, errors.ErrUserLockUp()
	}

	if err := r.repo.CheckPermission(ctx, int32(uid), req.Code); err != nil {
		return &v1.CheckRoutePermissionByTokenResponse{Result: false}, err
	}

	return &v1.CheckRoutePermissionByTokenResponse{Result: true}, nil
}

func (r *RolesUseCase) CheckRoutePermissionByID(ctx context.Context, req *v1.CheckRoutePermissionByIDRequest) (*v1.CheckRoutePermissionByIDResponse, error) {
	user, err := r.repo.CheckUserByID(ctx, int64(req.Id))
	if err != nil {
		return &v1.CheckRoutePermissionByIDResponse{}, err
	}

	if user.Status == 2 {
		return &v1.CheckRoutePermissionByIDResponse{}, errors.ErrUserLockUp()
	}

	if err := r.repo.CheckPermission(ctx, req.Id, req.Code); err != nil {
		return &v1.CheckRoutePermissionByIDResponse{Result: false}, err
	}

	return &v1.CheckRoutePermissionByIDResponse{Result: true}, nil
}
