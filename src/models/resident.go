package models

import (
	"log"

	"github.com/Art0r/symmetrical-couscous/src/classes"
	"github.com/Art0r/symmetrical-couscous/src/database"
	"github.com/Art0r/symmetrical-couscous/src/utils"
)

func RetrieveResidentETA(resident classes.ResidentReponse) (map[string]string, error) {
	// retrieve resident filtering by email AND telephone AND apto

	var retrievedResident map[string]string = map[string]string{
		"Id":   "",
		"Hash": "",
	}

	db, err := database.InitConnection()
	if err != nil {
		log.Fatal(err)
	}

	query, err := utils.GetQueryAsString("resident/retrieve")
	if err != nil {
		log.Fatal(err)
	}

	row := db.QueryRow(query, resident.Email)

	var Id string
	var Hash string

	err = row.Scan(&Id, &Hash)
	if err != nil {
		return retrievedResident, err
	}

	retrievedResident["Id"] = Id
	retrievedResident["Hash"] = Hash

	return retrievedResident, nil
}

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
