package db

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/wsungirl/GoMySql/model"
)

// DB is custom wrapper for gorm.DB
type DB struct {
	*gorm.DB
}

// InitDB initializes SQLite database
func InitDB(driver, connString string) (*DB, error) {
	db, err := gorm.Open(driver, connString)
	if err != nil {
		return nil, err
	}

	tables := []interface{}{
		&model.User{},
		&model.Database{},
		&model.History{},
		&model.Permissions{},
		&model.Session{},
	}

	for t := range tables {
		if db != nil && !db.HasTable(t) {
			db.CreateTable(t)
			continue
		}

		return nil, errors.New("Can't create tables")
	}

	return &DB{db}, nil
}
