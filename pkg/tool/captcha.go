package tool

import (
	"fmt"

	"github.com/mojocn/base64Captcha"
)

// Store xx
var Store = base64Captcha.DefaultMemStore

// GetCaptcha 获取验证码
func GetCaptcha() (string, string, error) {
	// 生成默认数字
	driver := base64Captcha.DefaultDriverDigit
	// 生成base64图片
	res := base64Captcha.NewCaptcha(driver, Store)

	// 获取
	id, b64s, err := res.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return "", "", err
	}

	return id, b64s, nil
}

// Verify 验证验证码
func Verify(id string, val string) bool {
	if id == "" || val == "" {
		return false
	}
	// 同时在内存清理掉这个图片
	return Store.Verify(id, val, true)
}
