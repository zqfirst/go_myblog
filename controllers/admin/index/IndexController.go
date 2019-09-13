package index

import (
	"github.com/gin-gonic/gin"
	"myblog_go/services/temp"
	"net/http"
)

func Index(c *gin.Context) {
	temp.SetTep("index/index")
	c.HTML(http.StatusOK, "base", gin.H{})
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index/login.tpl", "{}");
}
