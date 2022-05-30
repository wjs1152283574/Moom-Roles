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
			Name:   v.Name,
			Pass:   tool.Base64Md5(v.Pass),
			Type:   2,
			Status: 1,
			Commom: model.Commom{CreatedTime: time.Now().Unix()},
		})

	}
	// 新建用户
	if err := r.repo.CreateUser(ctx, data); err != nil {
		return &v1.CreateSuperUserResponse{}, errors.ErrSystemBusy
	}

	return &v1.CreateSuperUserResponse{}, nil
}

func (r *RolesUseCase) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	// 需要检验验证码
	if conf.GB.Verify {
		if !tool.Verify(req.Key, req.Val) {
			return &v1.LoginResponse{}, errors.ErrInvalidVerifyCode
		}
	}
	// 检验用户合法性
	user, err := r.repo.CheckUser(ctx, req.Name)
	if err != nil {
		return &v1.LoginResponse{}, err
	}
	// 检验用户密码
	if tool.Base64Md5(req.Pass) != user.Pass {
		return &v1.LoginResponse{}, errors.ErrInvalidPass
	}
	// 生成token
	token, err := tool.NewJWT(conf.GB.TokenScreat).CreateToken(strconv.Itoa(int(user.ID)), conf.GB.Issuer, conf.GB.TokenTtl)
	if err != nil {
		return &v1.LoginResponse{}, err
	}
	// 存入redis
	err = r.repo.RedisSet(ctx, key.TokenKey(uint64(user.ID)), token, conf.GB.TokenTtl)
	if err != nil {
		return &v1.LoginResponse{}, err
	}

	return &v1.LoginResponse{
		Token: token,
	}, nil
}

func (r *RolesUseCase) CreateAdminUser(ctx context.Context, req *v1.CreateAdminUserRequest, uid int64) (*v1.CreateAdminUserResponse, error) {
	// 组装用户数据
	var data []model.User
	for _, v := range conf.SU.SuperUser {
		data = append(data, model.User{
			Name:   v.Name,
			Pass:   tool.Base64Md5(v.Pass),
			Type:   1,
			Status: 1,
			Commom: model.Commom{CreatedTime: time.Now().Unix(), CreatorID: uid},
		})

	}
	// 新建用户
	if err := r.repo.CreateUser(ctx, data); err != nil {
		return &v1.CreateAdminUserResponse{}, errors.ErrSystemBusy
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
			Id:     int64(v.ID),
			Cid:    v.CreatorID,
			Cname:  v.CreatorName,
			Ctime:  v.CreatedTime,
			Name:   v.Name,
			Pass:   v.Pass,
			Type:   v.Type,
			Status: v.Status,
			Icon:   v.Icon,
		})
	}

	return &v1.AdminUserListResponse{
		Total: total,
		List:  List,
	}, nil
}

func (r *RolesUseCase) AdminUserInfos(ctx context.Context, uid int64) (*v1.AdminUserInfosResponse, error) {
	user, err := r.repo.UserBaseInfos(ctx, uid)
	if err != nil {
		return &v1.AdminUserInfosResponse{}, err
	}

	var Role []*v1.Role
	roles, err := r.repo.UserRoleList(ctx, uid)
	if err != nil {
		return &v1.AdminUserInfosResponse{}, err
	}
	for _, v := range roles {
		Role = append(Role, &v1.Role{
			Id:   int64(v.ID),
			Code: v.Code,
		})
	}

	var Per []*v1.Permissions
	permissions, err := r.repo.UserPermissionList(ctx, uid)
	if err != nil {
		return &v1.AdminUserInfosResponse{}, err
	}
	for _, v := range permissions {
		Per = append(Per, &v1.Permissions{
			Id:   int64(v.ID),
			Code: v.Code,
		})
	}

	return &v1.AdminUserInfosResponse{
		Id:          int64(user.ID),
		Name:        user.Name,
		Icon:        user.Icon,
		Roles:       Role,
		Permossions: Per,
	}, nil
}

func (r *RolesUseCase) AdminUserEdit(ctx context.Context, req *v1.AdminUserEditRequest) (*v1.AdminUserEditResponse, error) {
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
	if err := r.repo.SetRoles(ctx, req.Uid, creator, req.Rid); err != nil {
		return &v1.SetRolesResponse{}, err
	}

	return &v1.SetRolesResponse{}, nil
}
