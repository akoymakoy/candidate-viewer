package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func Connect() (*sql.DB, error) {

	// Set this in app.yaml when running in production.
	datastoreName := "root:G0Str0ng@(127.0.0.1:3306)/hrdb"

	log.Print("CONNECTING TO: ")
	log.Print(datastoreName)
	var err error
	db, err = sql.Open("mysql", datastoreName)
	if err != nil {
		log.Print("CANT CONNECT")
		log.Fatal(err)
	}

	return db,err
}