package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/Art0r/symmetrical-couscous/src/classes"
	"github.com/Art0r/symmetrical-couscous/src/database"
	"github.com/Art0r/symmetrical-couscous/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	database.StartDb()

	res1 := classes.Resident{
		Id:        uuid.NewString(),
		Name:      "Maria",
		Email:     "maria@maria.com",
		Telephone: "11923456789",
		Apto:      123,
		Created:   time.Now(),
	}

	res2 := classes.Resident{
		Id:        uuid.NewString(),
		Name:      "Joao",
		Email:     "joao@joao.com",
		Telephone: "11923456789",
		Apto:      124,
		Created:   time.Now(),
	}

	res3 := classes.Resident{
		Id:        uuid.NewString(),
		Name:      "Jose",
		Email:     "jose@jose.com",
		Telephone: "11923456789",
		Apto:      125,
		Created:   time.Now(),
	}

	fmt.Println(res1)
	models.CreateResident(res1)
	models.CreateResident(res2)
	models.CreateResident(res3)

	r := gin.Default()

	r.GET("", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello World")
	})

	r.Run("192.168.18.12:8000")
}
