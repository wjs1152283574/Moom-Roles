package biz

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
)

func (r *RolesUseCase) PermissionCreate(ctx context.Context, req *v1.PermissionCreateRequest, uid int64) (*v1.PermissionCreateResponse, error) {
	if err := r.repo.PermissionCreate(ctx, int32(uid), req.Name, req.Code); err != nil {
		return &v1.PermissionCreateResponse{}, err
	}

	return &v1.PermissionCreateResponse{}, nil
}

func (r *RolesUseCase) PermissionList(ctx context.Context, req *v1.PermissionListRequest) (*v1.PermissionListResponse, error) {
	pers, total, err := r.repo.PermissionList(ctx, req.Page, req.Limit, req.Name, req.Code)
	if err != nil {
		return &v1.PermissionListResponse{}, err
	}

	var List []*v1.PermissionListResponse_PermissionListItem
	for _, v := range pers {
		List = append(List, &v1.PermissionListResponse_PermissionListItem{
			Id:    int32(v.ID),
			Name:  v.Name,
			Code:  v.Code,
			Ctime: int32(v.CreatedTime),
			Cname: v.CreatorName,
			Cid:   int32(v.CreatorID),
			Utime: int32(v.UpdatedTime),
		})
	}

	return &v1.PermissionListResponse{
		List:  List,
		Total: int32(total),
	}, nil
}

func (r *RolesUseCase) PermissionDelete(ctx context.Context, req *v1.PermissionDeleteRequest) (*v1.PermissionDeleteResponse, error) {
	err := r.repo.PermissionDelete(ctx, req.Id)
	if err != nil {
		return &v1.PermissionDeleteResponse{}, err
	}

	return &v1.PermissionDeleteResponse{}, nil
}

func (r *RolesUseCase) PermissionEdit(ctx context.Context, req *v1.PermissionEditRequest) (*v1.PermissionEditResponse, error) {
	err := r.repo.PermissionEdit(ctx, req.Id, req.Name, req.Code)
	if err != nil {
		return &v1.PermissionEditResponse{}, err
	}

	return &v1.PermissionEditResponse{}, nil
}
