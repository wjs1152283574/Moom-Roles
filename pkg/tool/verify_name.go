package tool

import "regexp"

func VerifyNameFormat(pass string) (res bool) {
	if len(pass) <= 1 || len(pass) > 10 { // 长度 1 - 10
		return false
	}
	// 包含特殊字符
	var validIDw = regexp.MustCompile(`[*~!@#$%&-+]`)

	return !validIDw.MatchString(pass)
}
