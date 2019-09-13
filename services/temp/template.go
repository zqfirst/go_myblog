package temp

import (
	"github.com/gin-gonic/gin/ginS"
	"html/template"
	"myblog_go/models/consts"
)

func SetTep(view string) {
	ginS.SetHTMLTemplate(template.Must(template.ParseFiles(consts.ADMIN_BASE_MAIN, consts.ADMIN_TEMLATES + view + ".html")))
}
