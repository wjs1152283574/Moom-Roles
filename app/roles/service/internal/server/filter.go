package server

// 不需要鉴权的路由
var Notneed = []string{
	"/v1/auth/role/superuser",
	"/v1/role/captcha",
}
