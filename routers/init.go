package routers

import (
	"github.com/gin-gonic/gin"
	"myblog_go/controllers/admin/index"
)

func SetRoute() (r *gin.Engine) {
	r = gin.Default();
	admin := r.Group("/admin");
	admin.Use()
	{
		admin.GET("/index/index", index.Index);
	}
	return r;
}
