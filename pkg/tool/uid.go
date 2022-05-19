package tool

import (
	"strings"

	"github.com/nacos-group/nacos-sdk-go/inner/uuid"
)

// SimpleUUID 简洁UUID生成器
func SimpleUUID() string {
	u, err := uuid.NewV1()
	if err != nil {
		return ""
	}

	return strings.Replace(u.String(), "-", "", -1)
}
