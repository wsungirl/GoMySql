package db

import (
	"fmt"

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
		&model.Permission{},
		&model.Session{},
	}

	for _, t := range tables {
		if db != nil {
			err = db.AutoMigrate(t).Error
			if err == nil {
				continue
			}
		}

		return nil, fmt.Errorf("Can't create tables: %v", err)
	}

	return &DB{db}, nil
}
