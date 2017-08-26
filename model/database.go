package model

import (
	"fmt"
	"time"
)

type TableRow struct {
	ID     uint     `json:"id"`
	Values []string `json:"values"`
	Table  *DBTable `json:"table"`
}

type TableColumn struct {
	Field string `gorm:"column:field" json:"field"`
	Type  string `gorm:"column:type" json:"type"`
	// Null  bool     `json:"null"`
	// Key   string   `json:"key"`
	// Default string `json:"default"`
	// Extra string	  `json:"extra"`
}

type DBTable struct {
	Name    string        `json:"name"`
	Columns []TableColumn `json:"columns"`

	DB *Database `json:"db"`
}

type Database struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time `json:"-"`

	User User   `gorm:"index" json:"user,omitmpty"`
	Name string `json:"name,omitempty"`
}

func (db *Database) GetStoredName() string {
	return fmt.Sprintf("db_%d_%d", db.User.ID, db.ID)
}
