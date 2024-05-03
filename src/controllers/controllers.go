package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Default(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Hello World")
}
