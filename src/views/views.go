package views

import (
	"github.com/gin-gonic/gin"

	controllers "github.com/Art0r/symmetrical-couscous/src/controllers"
)

func SetViews(r *gin.Engine) {
	v := r.Group("api/")

	{
		v.POST("", controllers.Default)
		v.GET("index/", controllers.DefaultHtml)
	}
}
