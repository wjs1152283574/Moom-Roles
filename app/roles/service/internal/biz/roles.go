package biz

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/pkg/errors"
)

func (r *RolesUseCase) RoleCreate(ctx context.Context, req *v1.RoleCreateRequest, creator int64) (*v1.RoleCreateResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, creator) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.RoleCreate(ctx, int32(creator), req.Name, req.Code); err != nil {
		return &v1.RoleCreateResponse{}, err
	}

	return &v1.RoleCreateResponse{}, nil
}

func (r *RolesUseCase) RoleList(ctx context.Context, req *v1.RoleListRequest) (*v1.RoleListResponse, error) {
	roles, total, err := r.repo.RoleList(ctx, req.Page, req.Limit, req.Name, req.Code)
	if err != nil {
		return &v1.RoleListResponse{}, err
	}

	var List []*v1.RoleListResponse_RoleListItem
	for _, v := range roles {
		List = append(List, &v1.RoleListResponse_RoleListItem{
			Id:    int32(v.ID),
			Name:  v.Name,
			Code:  v.Code,
			Cid:   int32(v.CreatorID),
			Cname: v.CreatorName,
			Ctime: int32(v.CreatedTime),
			Utime: int32(v.UpdatedTime),
		})
	}

	return &v1.RoleListResponse{
		List:  List,
		Total: int32(total),
	}, nil
}

func (r *RolesUseCase) RoleDelete(ctx context.Context, req *v1.RoleDeleteRequest, uid int64) (*v1.RoleDeleteResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, uid) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.RoleDelete(ctx, req.Id); err != nil {
		return &v1.RoleDeleteResponse{}, err
	}

	return &v1.RoleDeleteResponse{}, nil
}

func (r *RolesUseCase) RoleEdit(ctx context.Context, req *v1.RoleEditRequest, uid int64) (*v1.RoleEditResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, uid) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.RoleEdit(ctx, req.Id, req.Name, req.Code); err != nil {
		return &v1.RoleEditResponse{}, err
	}

	return &v1.RoleEditResponse{}, nil
}
