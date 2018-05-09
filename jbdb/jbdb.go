package jbdb

import (
	"database/sql"
)

type JobBoardDB struct {
	DB *sql.DB
}


// Set up connection to DB
func New(db *sql.DB) *JobBoardDB {
	return &JobBoardDB{
		DB: db,
	}
}


