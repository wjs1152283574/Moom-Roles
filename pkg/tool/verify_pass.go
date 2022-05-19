package tool

import (
	"bytes"
	"regexp"
	"unicode"
)

func VerifyPassFormat(pass string) (res bool) {
	rune := []rune(pass)
	if len(pass) != len(rune) { // 验证是否包含中文
		return false
	}

	if len(pass) < 8 || len(pass) > 16 { // 长度 8 - 16
		return false
	}

	// 不包含字母
	if !IncludeLetter(pass) {
		return false
	}

	var validID = regexp.MustCompile(`\d`)
	if !validID.MatchString(pass) { // 未包含数字
		return false
	}

	// 包含特殊字符
	var validIDw = regexp.MustCompile(`[*~!@#$%&-+]`)

	return !validIDw.MatchString(pass)
}

func IncludeLetter(str string) (result bool) {
	r := bytes.Runes([]byte(str))
	for _, v := range r {
		if unicode.IsLetter(v) {
			return true
		}
	}

	return false
}
