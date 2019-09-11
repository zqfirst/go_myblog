package setting

import (
	"flag"
	"fmt"
	"github.com/larspensjo/config"
)

type Http struct {
	port string
}

type Db struct {
	host string
	port string
	user string
	pass string
}

var (
	configFile = flag.String("configfile", "config.ini", "General configuration file")
)

func Init() {
	SetConfig(*configFile);
}

func SetConfig(c string) {
	configInfo, _ := config.ReadDefault(c);
	fmt.Print(configInfo);
}
