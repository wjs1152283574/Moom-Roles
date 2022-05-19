package conf

type SuperUser struct {
	SuperUser []SuperUserItem `json:"super_user"`
}

type SuperUserItem struct {
	Name string
	Pass string
}

type Global struct {
	Verify bool // 是否开启图片验证码
}
