package setting

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SetLoginSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Set(GetSessionKey(), "123")
	session.Save()
}

func GetLoginSession(c *gin.Context) (r bool) {
	session := sessions.Default(c)
	loginuser := session.Get(GetSessionKey())
	return loginuser == "123"
}
