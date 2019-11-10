// contains utility functions for a candidates competence matrix
package candidateutil

import (
	"database/sql"
	"log"
)

type Candidate struct {
	Id int
	Name,Email   string
}

func get(db *sql.DB) Candidate {


	// Query a single user
	var (
		Id        int
		Name,Email  string
	)

	query := "SELECT id, name,email FROM candidates WHERE id = ?"
	if err := db.QueryRow(query, 1).Scan(&Id, &Name, &Email); err != nil {
		log.Fatal(err)
	}

	details := Candidate{
		Id:   (Id),
		Name: (Name),
		Email: (Email),
	}

	return details
}