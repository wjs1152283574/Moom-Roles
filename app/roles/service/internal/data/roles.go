package data

import "context"

func (r *UserRepo) CreateUser(ctx context.Context, name, pass string, typ, status int64) error {
	return nil
}
