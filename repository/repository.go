package repository

import "database/sql"

// Your repository functions go here

type Database struct {
	MySQLDB *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		MySQLDB: db,
	}
}

// Add more repository functions as needed
