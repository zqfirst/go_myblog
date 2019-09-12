package main

import (
	"myblog_go/routers"
	"myblog_go/services/setting"
)

func main() {
	r := routers.SetRoute()
	port := setting.GetConfig("HTTP", "port")
	r.Run("127.0.0.1:" + port) // listen and serve on 0.0.0.0:8080
}
