package views

import (
	"github.com/gin-gonic/gin"

	controllers "github.com/Art0r/symmetrical-couscous/src/controllers"
)

func SetViews(r *gin.Engine) {
	r.LoadHTMLGlob("templates/*")
	r.Static("/static", "static")

	v := r.Group("")

	{
		v.GET("", controllers.Default)
		v.GET("index/", controllers.DefaultHtml)
	}
}
