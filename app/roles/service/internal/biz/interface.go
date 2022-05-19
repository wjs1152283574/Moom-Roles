package biz

import "context"

type RolesRepo interface {
	CreateUser(ctx context.Context, name, pass string, typ, status int64) error
}
