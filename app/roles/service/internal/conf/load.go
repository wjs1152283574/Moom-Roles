package conf

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/config"
)

// keys 这个key需要对应json文件里面的最外层的键，即监听这个键，而不是这个json文件名
var keys = []string{
	"super_user",
	"global",
}

var SU SuperUser
var GB Global

func LoadConf(conf config.Config) (err error) {
	for _, key := range keys {
		err = conf.Watch(key, func(key string, value config.Value) {
			switch key {
			case "super_user":
				err = value.Scan(&SU) // conf.Scan(&UConf)
			case "global":
				err = value.Scan(&GB) // conf.Scan(&UConf)
			default:
				fmt.Println("load empty configs")
			}
		})
		if err != nil {
			panic(err)
		}
	}

	err = conf.Scan(&SU)
	if err != nil {
		panic(err)
	}

	err = conf.Scan(&GB)
	if err != nil {
		panic(err)
	}
	fmt.Println(GB, SU)
	return
}
