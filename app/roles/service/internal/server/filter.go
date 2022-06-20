package server

// 不需要鉴权的路由
var NotneedAuth = []string{
	"/superuser/create",
	"/captcha",
	"/login",
}

var HttpCheck = []string{
	"/role/check",
	"/permission/check",
	"/role/check/token",
	"/role/check/id",
	"/permission/check/token",
	"/permission/check/id",
}
