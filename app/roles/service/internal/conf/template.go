package conf

type SuperUser struct {
	SuperUser []SuperUserItem `json:"super_user"`
}

type SuperUserItem struct {
	Name string
	Pass string
}

type Global struct {
	Global Globals `json:"global"`
}

type Globals struct {
	Verify      bool   `json:"verify"`       // 是否开启图片验证码
	TokenTtl    int64  `json:"token_ttl"`    // token 有效时间 秒
	TokenScrect string `json:"token_screct"` // 密钥
	Issuer      string `json:"issuer"`       // 颁发机构
	Http        bool   `json:"http"`         // 是否支持http接口鉴权
}
