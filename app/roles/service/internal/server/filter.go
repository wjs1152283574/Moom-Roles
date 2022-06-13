package server

// 不需要鉴权的路由
var NotneedAuth = []string{
	"v1/auth/role/superuser/create",
	"/v1/role/captcha",
}
