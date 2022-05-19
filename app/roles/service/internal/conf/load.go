package conf

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/config"
)

// keys 这个key需要对应json文件里面的最外层的键，即监听这个键，而不是这个json文件名
var keys = []string{
	"super_user",
}

var SU SuperUser

func LoadConf(conf config.Config) (err error) {
	for _, key := range keys {
		err = conf.Watch(key, func(key string, value config.Value) {
			switch key {
			case "temp":
				err = value.Scan(&SU) // conf.Scan(&UConf)
			default:
				fmt.Println("load empty configs")
			}
		})
		if err != nil {
			panic(err)
		}
	}

	return conf.Scan(&SU)
}
