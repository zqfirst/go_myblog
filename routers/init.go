package routers

import (
	"github.com/gin-gonic/gin"
	"myblog_go/controllers/admin/index"
)

func SetRoute() (r *gin.Engine) {

	r = gin.Default();

	//引入模板文件
	r.LoadHTMLGlob("templates/**/*/*")

	//设置静态资源模板路径
	r.Static("/public", "static/")

	admin := r.Group("/admin");
	admin.Use()
	{
		admin.GET("/index/index", index.Index);
		admin.GET("/login", index.Login);
	}
	return r;
}

