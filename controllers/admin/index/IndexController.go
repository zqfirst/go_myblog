package index

import (
	"github.com/gin-gonic/gin"
	"myblog_go/models/mg_user"
	"myblog_go/services/setting"
	"myblog_go/services/temp"
	"net/http"
)

func Index(c *gin.Context) {
	temp.SetTep("index/index")
	c.HTML(http.StatusOK, "base", gin.H{})
}

func Login(c *gin.Context) {
	username := c.PostForm("username");
	password := c.PostForm("password");
	if username != "" {
		user := mg_user.GetByUsername(username)
		if mg_user.CheckPwd(password, user) {
			//设置cookie
			//c.SetCookie(setting.GetSessionKey(), "123", 60*60*24, "/", "", false, true)
			//设置session
			setting.SetLoginSession(c)
		}

		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"msg":    "登录成功",
		})
	} else {
		c.HTML(http.StatusOK, "login", "{}");
	}
}

func NotFound(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404NotFound", "{}")
}
