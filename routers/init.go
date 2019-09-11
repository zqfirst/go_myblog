package routers

import (
	"github.com/gin-gonic/gin"
	"myblog_go/controllers/admin/index"
)

func SetRoute() (r *gin.Engine) {
	r = gin.Default();
	//router.GET("/ping", func(context *gin.Context) {
	//	context.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	admin := r.Group("/admin");
	admin.Use()
	{
		admin.GET("/index/index", index.Index);
	}
	return r;
}
