package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Art0r/symmetrical-couscous/src/classes"
	"github.com/Art0r/symmetrical-couscous/src/models"
	"github.com/Art0r/symmetrical-couscous/src/utils"

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

	resident, err := models.RetrieveResidentETA(response)
	if err != nil {
		ctx.String(http.StatusNotFound, err.Error())
		return
	}

	hash := utils.CreateHash(response.Email, response.Telephone, response.Apto)

	fmt.Println(resident["Hash"])
	fmt.Println(hash)
	
	if hash == resident["Hash"] {
		ctx.JSON(http.StatusOK, resident)
		return
	}

	ctx.JSON(http.StatusForbidden, resident)
}

func DefaultHtml(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
