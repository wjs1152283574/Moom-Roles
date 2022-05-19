package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrAuthFail          = errors.New(401, "AuthFail", "Missing token or token incorrect")
	ErrInvalidParams     = errors.New(400, "InvalidParams", "invalid params")
	ErrInvalidUsername   = errors.New(400, "InvalidUsername", "invalid user name")
	ErrInvalidPass       = errors.New(400, "InvalidPass", "invalid password")
	ErrSystemBusy        = errors.New(500, "SystemBusy", "system busy")
	ErrMuiltpleUserName  = errors.New(400, "MuiltpleUserName", "user allready exit")
	ErrNoNeedCaptcha     = errors.New(401, "NoNeedCaptcha", "no need captcha")
	ErrInvalidVerifyCode = errors.New(401, "InvalidVerifyCode", "invalid verify code")
	ErrUserNotExit       = errors.New(404, "UserNotExit", "user not exit")

	ErrInvalidToken = errors.New(404, "InvalidToken", "invalid token")
)
