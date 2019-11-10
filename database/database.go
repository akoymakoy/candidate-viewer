package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func Connect() (*sql.DB, error) {
	// Configure the database connection (always check errors)
	db, err := sql.Open("mysql", "root:password@(127.0.0.1:3306)/makoy?parseTime=true")

	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	if err != nil {
		log.Fatal(err)
	}else{
		log.Print("SUCCESS DB")
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}else{
		log.Print("SUCCESS DB2")
	}
	return db,err
}