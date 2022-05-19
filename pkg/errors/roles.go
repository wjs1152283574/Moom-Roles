package errors

import "github.com/go-kratos/kratos/v2/errors"

var (
	ErrAuthFail         = errors.New(401, "AuthFail", "Missing token or token incorrect")
	ErrSystemBusy       = errors.New(500, "SystemBusy", "system busy")
	ErrMuiltpleUserName = errors.New(400, "MuiltpleUserName", "user allready exit")
	ErrNoNeedCaptcha    = errors.New(401, "NoNeedCaptcha", "no need captcha")
)
