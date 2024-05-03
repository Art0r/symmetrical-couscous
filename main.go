package main

import (
	"math/rand"

	"time"

	"github.com/Art0r/symmetrical-couscous/src/classes"
	"github.com/Art0r/symmetrical-couscous/src/database"
	"github.com/Art0r/symmetrical-couscous/src/models"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func populateDatabase() {

	for i := 0; i < 10; i++ {
		
		id := uuid.NewString()
		now := time.Now()
		name := faker.Name()
		email := faker.Email()
		telephone := faker.Phonenumber()
		apto := rand.Int63n(2000)

		resident := classes.Resident{
			Id:        id,
			Name:      name,
			Email:     email,
			Telephone: telephone,
			Apto:      apto,
			Created:   now,
		}

		models.CreateResident(resident)
	}

}

func main() {

	database.StartDb()

	populateDatabase()

	// r := gin.Default()

	// r.GET("", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "Hello World")
	// })

	// r.Run("192.168.18.12:8000")
}
