package middleware

import (
	"github.com/gin-gonic/gin"
	"myblog_go/services/setting"
	"net/http"
)

func IsLogin(c *gin.Context) {
	loginuser := setting.GetLoginSession(c)
	if !loginuser {
		c.Redirect(http.StatusFound, "/admin/login")
	}
	c.Next()
}
