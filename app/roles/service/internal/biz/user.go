package biz

import (
	"context"
	"strconv"
	"time"

	v1 "github.com/it-moom/moom-roles/api/roles/service/v1"
	"github.com/it-moom/moom-roles/app/roles/service/internal/conf"
	"github.com/it-moom/moom-roles/app/roles/service/internal/model"
	"github.com/it-moom/moom-roles/pkg/errors"
	"github.com/it-moom/moom-roles/pkg/key"
	"github.com/it-moom/moom-roles/pkg/tool"
)

func (r *RolesUseCase) CreateSuperUser(ctx context.Context) (*v1.CreateSuperUserResponse, error) {
	// 组装用户数据
	var data []model.User
	for _, v := range conf.SU.SuperUser {
		data = append(data, model.User{
			Name:     v.Name,
			Pass:     tool.Base64Md5(v.Pass),
			Type:     2,
			Status:   1,
			ReadOnly: true,
			Commom:   model.Commom{CreatedTime: time.Now().Unix()},
		})

	}
	// 新建用户
	if err := r.repo.CreateUser(ctx, data); err != nil {
		return &v1.CreateSuperUserResponse{}, err
	}

	return &v1.CreateSuperUserResponse{}, nil
}

func (r *RolesUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// 需要检验验证码
	if conf.GB.Global.Verify {
		if !tool.Verify(req.Key, req.Val) {
			return &v1.LoginResponse{}, errors.ErrInvalidVerifyCode()
		}
	}
	// 检验用户合法性
	user, err := r.repo.CheckUser(ctx, req.Name)
	if err != nil {
		return &v1.LoginResponse{}, err
	}
	// 检测用户状态
	if user.Status == 2 {
		return &v1.LoginResponse{}, errors.ErrUserLockUp()
	}
	// 检验用户密码
	if tool.Base64Md5(req.Pass) != user.Pass {
		return &v1.LoginResponse{}, errors.ErrInvalidPass()
	}
	// 生成token
	token, err := tool.NewJWT(conf.GB.Global.TokenScrect).CreateToken(strconv.Itoa(int(user.ID)), conf.GB.Global.Issuer, conf.GB.Global.TokenTtl)
	if err != nil {
		return &v1.LoginResponse{}, err
	}
	// 存入redis
	err = r.repo.RedisSet(ctx, key.TokenKey(uint64(user.ID)), token, int32(conf.GB.Global.TokenTtl))
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		Token: token,
	}, nil
}

func (r *RolesUseCase) CreateAdminUser(ctx context.Context, req *v1.CreateAdminUserRequest, uid int64) (*v1.CreateAdminUserResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, uid) {
		errors.ErrPermissionDeni("only for super user")
	}
	// 组装用户数据
	var data []model.User
	data = append(data, model.User{
		Name:   req.Name,
		Pass:   tool.Base64Md5(req.Pass),
		Icon:   req.Icon,
		Type:   1,
		Status: 1,
		Commom: model.Commom{CreatedTime: time.Now().Unix(), CreatorID: uid},
	})

	// 新建用户
	if err := r.repo.CreateUser(ctx, data); err != nil {
		return &v1.CreateAdminUserResponse{}, err
	}

	return &v1.CreateAdminUserResponse{}, nil
}

func (r *RolesUseCase) AdminUserList(ctx context.Context, req *v1.AdminUserListRequest) (*v1.AdminUserListResponse, error) {
	users, total, err := r.repo.UserList(ctx, req.Name, req.Cname, req.Page, req.Limit, req.Type, req.Status)
	if err != nil {
		return &v1.AdminUserListResponse{}, err
	}

	var List []*v1.AdminUserListResponse_List
	for _, v := range users {
		List = append(List, &v1.AdminUserListResponse_List{
			Id:     int32(v.ID),
			Cid:    int32(v.CreatorID),
			Cname:  v.CreatorName,
			Ctime:  int32(v.CreatedTime),
			Name:   v.Name,
			Pass:   v.Pass,
			Type:   int32(v.Type),
			Status: int32(v.Status),
			Icon:   v.Icon,
		})
	}

	return &v1.AdminUserListResponse{
		Total: int32(total),
		List:  List,
	}, nil
}

func (r *RolesUseCase) AdminUserInfos(ctx context.Context, uid int64) (*v1.AdminUserInfosResponse, error) {
	user, err := r.repo.UserBaseInfos(ctx, int32(uid))
	if err != nil {
		return &v1.AdminUserInfosResponse{}, err
	}

	var Role []*v1.Role
	roles, err := r.repo.UserRoleList(ctx, int32(uid))
	if err != nil {
		return &v1.AdminUserInfosResponse{}, err
	}
	for _, v := range roles {
		Role = append(Role, &v1.Role{
			Id:   int32(v.ID),
			Code: v.Code,
		})
	}

	var Per []*v1.Permissions
	permissions, err := r.repo.UserPermissionList(ctx, int32(uid))
	if err != nil {
		return &v1.AdminUserInfosResponse{}, err
	}
	for _, v := range permissions {
		Per = append(Per, &v1.Permissions{
			Id:   int32(v.ID),
			Code: v.Code,
		})
	}

	return &v1.AdminUserInfosResponse{
		Id:          int32(user.ID),
		Name:        user.Name,
		Icon:        user.Icon,
		Roles:       Role,
		Permossions: Per,
	}, nil
}

func (r *RolesUseCase) AdminUserEdit(ctx context.Context, req *v1.AdminUserEditRequest, uid int64) (*v1.AdminUserEditResponse, error) {
	// 权限认证 (非超级用户&&不是修改自己的信息)
	if !r.repo.IsSuperUser(ctx, uid) && req.Uid != int32(uid) {
		errors.ErrPermissionDeni("only for super user")
	}
	user, err := r.repo.UserBaseInfos(ctx, req.Uid)
	if err != nil {
		return &v1.AdminUserEditResponse{}, err
	}

	if req.Name != "" {
		user.Name = req.Name
	}

	if req.Pass != "" {
		user.Pass = tool.Base64Md5(req.Pass)
	}

	if req.Icon != "" {
		user.Icon = req.Icon
	}

	user.UpdatedTime = time.Now().Unix()

	return &v1.AdminUserEditResponse{}, r.repo.UserBaseEdit(ctx, user)
}

func (r *RolesUseCase) SetRoles(ctx context.Context, req *v1.SetRolesRequest, creator int64) (*v1.SetRolesResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, creator) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.SetRoles(ctx, req.Uid, int32(creator), req.Rid); err != nil {
		return &v1.SetRolesResponse{}, err
	}

	return &v1.SetRolesResponse{}, nil
}

func (r *RolesUseCase) SetRolesDelete(ctx context.Context, req *v1.SetRolesDeleteRequest, creator int64) (*v1.SetRolesDeleteResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, creator) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.SetRolesDelete(ctx, req.Id, req.Role); err != nil {
		return &v1.SetRolesDeleteResponse{}, err
	}

	return &v1.SetRolesDeleteResponse{}, nil
}

func (r *RolesUseCase) SetPermission(ctx context.Context, req *v1.SetPermissionRequest, creator int64) (*v1.SetPermissionResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, creator) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.SetRoles(ctx, req.Uid, int32(creator), req.Pid); err != nil {
		return &v1.SetPermissionResponse{}, err
	}

	return &v1.SetPermissionResponse{}, nil
}

func (r *RolesUseCase) SetPermissionDelete(ctx context.Context, req *v1.SetPermissionDeleteRequest, uid int64) (*v1.SetPermissionDeleteResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, uid) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.SetPermissionDelete(ctx, req.Id, req.Permission); err != nil {
		return &v1.SetPermissionDeleteResponse{}, err
	}

	return &v1.SetPermissionDeleteResponse{}, nil
}

func (r *RolesUseCase) AdminUserDelete(ctx context.Context, req *v1.AdminUserDeleteRequest, uid int64) (*v1.AdminUserDeleteResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, uid) {
		errors.ErrPermissionDeni("only for super user")
	}

	if err := r.repo.UserDelete(ctx, req.Id); err != nil {
		return &v1.AdminUserDeleteResponse{}, err
	}

	return &v1.AdminUserDeleteResponse{}, nil
}

func (r *RolesUseCase) AdminUserLock(ctx context.Context, req *v1.AdminUserLockRequest, uid int64) (*v1.AdminUserLockResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, uid) {
		errors.ErrPermissionDeni("only for super user")
	}
	_, err := r.repo.CheckUserByID(ctx, int64(req.Id))
	if err != nil {
		return &v1.AdminUserLockResponse{}, err
	}

	if err := r.repo.UpdateUserStatus(ctx, int64(req.Id), 2); err != nil {
		return &v1.AdminUserLockResponse{}, err
	}

	return &v1.AdminUserLockResponse{}, nil
}

func (r *RolesUseCase) AdminUserUnLock(ctx context.Context, req *v1.AdminUserUnLockRequest, uid int64) (*v1.AdminUserUnLockResponse, error) {
	// 权限认证
	if !r.repo.IsSuperUser(ctx, uid) {
		errors.ErrPermissionDeni("only for super user")
	}
	_, err := r.repo.CheckUserByID(ctx, int64(req.Id))
	if err != nil {
		return &v1.AdminUserUnLockResponse{}, err
	}

	if err := r.repo.UpdateUserStatus(ctx, int64(req.Id), 1); err != nil {
		return &v1.AdminUserUnLockResponse{}, err
	}

	return &v1.AdminUserUnLockResponse{}, nil
}
