package setting

import (
	"fmt"
	"github.com/Unknwon/goconfig"
	"os"
)

type Http struct {
	port string
}

var (
	cfg *goconfig.ConfigFile
)

func init() {
	config, err := goconfig.LoadConfigFile("conf/config.ini") //加载配置文件
	if err != nil {
		fmt.Println("get config file error")
		os.Exit(-1)
	}
	cfg = config
}

func GetConfig(sec string, key string) (val string) {
	if cfg == nil {
		return val;
	}

	val, _ = cfg.GetValue(sec, key)
	return val
}

func GetSessionKey() (key string) {
	return GetConfig("session", "key")
}
