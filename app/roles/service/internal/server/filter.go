package server

// 不需要鉴权的路由
var NotneedAuth = []string{
	"/superuser/create",
	"/captcha",
}
