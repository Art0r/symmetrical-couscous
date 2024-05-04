package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World")
}

func DefaultHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
