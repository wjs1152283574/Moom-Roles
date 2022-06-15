package biz

import (
	"context"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
)

func (r *RolesUseCase) RouteCreate(ctx context.Context, req *v1.RouteCreateRequest, uid int64) (*v1.RouteCreateResponse, error) {
	err := r.repo.RouteCreate(ctx, int32(uid), req.Method, req.Url)
	if err != nil {
		return &v1.RouteCreateResponse{}, err
	}

	return &v1.RouteCreateResponse{}, nil
}

func (r *RolesUseCase) RouteList(ctx context.Context, req *v1.RouteListRequest) (*v1.RouteListResponse, error) {
	list, total, err := r.repo.RouteList(ctx, req.Page, req.Limit, req.Method, req.Url)
	if err != nil {
		return &v1.RouteListResponse{}, err
	}

	var List []*v1.RouteListResponse_RouteListItem
	for _, v := range list {
		List = append(List, &v1.RouteListResponse_RouteListItem{
			Id:     int32(v.ID),
			Url:    v.URL,
			Method: int32(v.Method),
			Cid:    int32(v.CreatorID),
			Ctime:  int32(v.CreatedTime),
			Name:   v.CreatorName,
		})
	}

	return &v1.RouteListResponse{
		List:  List,
		Total: int32(total),
	}, nil
}

func (r *RolesUseCase) RouteEdit(ctx context.Context, req *v1.RouteEditRequest) (*v1.RouteEditResponse, error) {
	err := r.repo.RouteEdit(ctx, req.Id, req.Method, req.Url)
	if err != nil {
		return &v1.RouteEditResponse{}, err
	}

	return &v1.RouteEditResponse{}, nil
}

func (r *RolesUseCase) RouteDelete(ctx context.Context, req *v1.RouteDeleteRequest) (*v1.RouteDeleteResponse, error) {
	err := r.repo.RouteDelete(ctx, req.Id)
	if err != nil {
		return &v1.RouteDeleteResponse{}, err
	}

	return &v1.RouteDeleteResponse{}, nil
}

func (r *RolesUseCase) RouteRole(ctx context.Context, req *v1.RouteRoleRequest, uid int64) (*v1.RouteRoleResponse, error) {
	err := r.repo.RouteRole(ctx, int32(uid), req.Route, req.Role)
	if err != nil {
		return &v1.RouteRoleResponse{}, err
	}

	return &v1.RouteRoleResponse{}, nil
}

func (r *RolesUseCase) RouteRoleDelete(ctx context.Context, req *v1.RouteRoleDeleteRequest) (*v1.RouteRoleDeleteResponse, error) {
	err := r.repo.RouteRoleDelete(ctx, req.Id, req.Route)
	if err != nil {
		return &v1.RouteRoleDeleteResponse{}, err
	}

	return &v1.RouteRoleDeleteResponse{}, nil
}

func (r *RolesUseCase) RoutePermission(ctx context.Context, req *v1.RoutePermissionRequest, uid int64) (*v1.RoutePermissionResponse, error) {
	err := r.repo.RoutePermission(ctx, int32(uid), req.Route, req.Permisson)
	if err != nil {
		return &v1.RoutePermissionResponse{}, err
	}

	return &v1.RoutePermissionResponse{}, nil
}

func (r *RolesUseCase) RoutePermissionDelete(ctx context.Context, req *v1.RoutePermissionDeleteRequest) (*v1.RoutePermissionDeleteResponse, error) {
	err := r.repo.RoutePermissionDelete(ctx, req.Id, req.Permission)
	if err != nil {
		return &v1.RoutePermissionDeleteResponse{}, err
	}

	return &v1.RoutePermissionDeleteResponse{}, nil
}

func (r *RolesUseCase) RouteDetails(ctx context.Context, req *v1.RouteDetailsRequest) (*v1.RouteDetailsResponse, error) {
	route, roles, permisions, err := r.repo.RouteDetails(ctx, req.Id)
	if err != nil {
		return &v1.RouteDetailsResponse{}, err
	}

	var Roles []*v1.RouteDetailsResponse_Roles
	for _, v := range roles {
		Roles = append(Roles, &v1.RouteDetailsResponse_Roles{
			Id:   int32(v.ID),
			Name: v.Name,
			Code: v.Code,
		})
	}

	var Permission []*v1.RouteDetailsResponse_Permission
	for _, v := range permisions {
		Permission = append(Permission, &v1.RouteDetailsResponse_Permission{
			Id:   int32(v.ID),
			Name: v.Name,
			Code: v.Code,
		})
	}

	return &v1.RouteDetailsResponse{
		Id:         int32(route.ID),
		Url:        route.URL,
		Method:     int32(route.Method),
		Ctime:      int32(route.CreatedTime),
		Roles:      Roles,
		Permission: Permission,
	}, nil
}
