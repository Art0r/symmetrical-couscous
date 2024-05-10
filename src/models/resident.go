package models

import (
	"log"

	"github.com/Art0r/symmetrical-couscous/src/classes"
	"github.com/Art0r/symmetrical-couscous/src/database"
	"github.com/Art0r/symmetrical-couscous/src/utils"
)

func CreateResident(resident classes.Resident) {

	db, err := database.InitConnection()
	if err != nil {
		log.Fatal(err)
	}

	query, err := utils.GetQueryAsString("resident/create")
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(
		resident.Id,
		resident.Email,
		resident.Name,
		resident.Apto,
		resident.Hash,
		resident.Telephone,
		resident.Created,
	)

	if err != nil {
		log.Fatal(err)
	}
}
