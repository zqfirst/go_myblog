package routers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"myblog_go/controllers/admin/index"
	"myblog_go/controllers/admin/middleware"
	"myblog_go/services/setting"
)

func SetRoute() (r *gin.Engine) {

	r = gin.Default();

	//设置session midddleware
	store := cookie.NewStore([]byte(setting.GetSessionKey()))
	r.Use(sessions.Sessions("mysession", store))

	//引入模板文件
	r.LoadHTMLGlob("templates/**/*/*")

	//设置静态资源模板路径
	r.Static("/public", "static/")

	//设置默认路由
	r.NoRoute(index.NotFound)

	admin := r.Group("/admin");

	//登录页不添加中间层校验
	admin.GET("/login", index.Login);
	admin.POST("/login", index.Login);
	//登录校验
	admin.Use(middleware.IsLogin)
	{
		admin.GET("/", index.Index);
		admin.GET("/index/index", index.Index);
	}

	return r;
}
