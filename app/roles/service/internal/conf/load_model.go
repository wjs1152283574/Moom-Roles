package conf

type SuperUser struct {
	SuperUser []SuperUserItem `json:"super_user"`
}

type SuperUserItem struct {
	Name string
	Pass string
}

type Global struct {
	Verify      bool   // 是否开启图片验证码
	TokenTtl    int64  // token 有效时间 秒
	TokenScreat string // 密钥
	Issuer      string // 颁发机构
}
