package main

import (
	"myblog_go/routers"
	_ "myblog_go/services"
)

func main() {
	r := routers.SetRoute()
	r.Run() // listen and serve on 0.0.0.0:8080
}