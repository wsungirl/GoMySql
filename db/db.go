package db

import (
	"database/sql"
)

// DB is custom wrapper for sql.DB
type DB struct {
	*sql.DB
}

// InitDB initializes SQLite database
func InitDB(driver, connString string) (*db.DB, error) {
	db, err := sql.Open(driver, connString)
	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
