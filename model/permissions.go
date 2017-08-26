package model

import (
	"database/sql"
	"time"
)

type PermAction uint8

//go:generate enumer -type PermAction -transform=snake -autotrimprefix
const (
	PermActRead PermAction = iota
	PermActCreate
	PermActUpdate
)

type PermEntity uint8

//go:generate enumer -type PermAction -transform=snake -autotrimprefix
const (
	PermEntDatabase PermEntity = iota
	PermEntTable
	PermEntRow
)

type Permissions struct {
	ID        uint       `gorm:"primary_key" json:"-"`
	CreatedAt time.Time  `json:"-"`
	UpdatedAt time.Time  `json:"-"`
	DeletedAt *time.Time `sql:"index" json:"-"`

	DB   Database      `json:"db"`
	DBID sql.NullInt64 `gorm:"index" json:"-"`

	User User `gorm:"index" json:"user"`

	Action string `json:"action"`
	Entity string `json:"entity"`
}
