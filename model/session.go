package model

import (
	"database/sql"

	"github.com/jinzhu/gorm"
)

// Session struct represents row of DB table sessions
type Session struct {
	gorm.Model `json:"-"`
	User   User          `gorm:"ForeignKey:UserID"`
	UserID sql.NullInt64 `gorm:"index" json:"-"`

	AccessToken string `gorm:"type:char(36)" json:"access_token"`
}
