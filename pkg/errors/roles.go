package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrAuthFail          = errors.New(401, "AuthFail", "Missing token or token incorrect")
	ErrInvalidParams     = errors.New(400, "InvalidParams", "invalid params")
	ErrInvalidUID        = errors.New(400, "InvalidUID", "invalid uid")
	ErrInvalidUsername   = errors.New(400, "InvalidUsername", "invalid user name")
	ErrInvalidPass       = errors.New(400, "InvalidPass", "invalid password")
	ErrMuiltiRecord      = errors.New(400, "MuiltiRecord", "muiltiple record")
	ErrSystemBusy        = errors.New(500, "SystemBusy", "system busy")
	ErrMuiltpleUserName  = errors.New(400, "MuiltpleUserName", "user allready exit")
	ErrNoNeedCaptcha     = errors.New(401, "NoNeedCaptcha", "no need captcha")
	ErrInvalidVerifyCode = errors.New(401, "InvalidVerifyCode", "invalid verify code")
	ErrPermissionDeni    = errors.New(401, "PermissionDeni", "ckech permission fail")
	ErrUserNotExit       = errors.New(404, "UserNotExit", "user not exit")
	ErrRouteNotExit      = errors.New(404, "RouteNotExit", "route not exit")
	ErrRoleNotExit       = errors.New(404, "RoleNotExit", "role not exit")

	ErrInvalidToken = errors.New(404, "InvalidToken", "invalid token")
)
