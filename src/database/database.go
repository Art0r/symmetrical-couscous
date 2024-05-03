package database

import (
	"database/sql"
	"log"

	"github.com/Art0r/symmetrical-couscous/src/utils"
)


func InitConnection() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "db.sqlite3")

	if err != nil {
		return nil, err
	}

	return db, nil
}


func StartDb() {

	db, err := InitConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("SQLITE VERSION: ", version)

	query, err := utils.GetQueryAsString("init")
	if err != nil {
		panic(err.Error())
	}

	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal("Error executing initial routine: ", err.Error())
	}

    _, err = stmt.Exec()
    if err != nil {
        log.Fatal("Error executing statement: ", err.Error())
    }
	defer stmt.Close()
	
	log.Println("Executed initial routine")

}
