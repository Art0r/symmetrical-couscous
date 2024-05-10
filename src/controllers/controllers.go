package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Art0r/symmetrical-couscous/src/classes"
	"github.com/gin-gonic/gin"
)

func Default(ctx *gin.Context) {
	var response classes.ResidentReponse

	dec := json.NewDecoder(ctx.Request.Body)
	if dec == nil {
		ctx.String(http.StatusBadRequest, "Corpo da requisição inválido")
		return
	}

	dec.Decode(&response)
	defer ctx.Request.Body.Close()

	if (response.Apto == "") || (response.Email == "") || (response.Telephone == "") {
		ctx.String(http.StatusBadRequest, "Corpo da requisição inválido")
		return
	}

	

	ctx.JSON(http.StatusOK, response)
}

func DefaultHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
