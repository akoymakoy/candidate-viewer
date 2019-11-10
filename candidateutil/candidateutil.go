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

func Get(db *sql.DB,id int) Candidate {


	// Query a single user
	var (
		Id        int
		Name,Email  string
	)

	query := "SELECT id, name,email FROM candidates WHERE id = ?"
	if err := db.QueryRow(query, id).Scan(&Id, &Name, &Email); err != nil {
		log.Fatal(err)
	}

	details := Candidate{
		Id:   (Id),
		Name: (Name),
		Email: (Email),
	}

	return details
}