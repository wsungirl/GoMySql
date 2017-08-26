package model

import "github.com/jinzhu/gorm"

// Session struct represents row of DB table sessions
type Session struct {
	gorm.Model
	User        User   `gorm:"index"`
	AccessToken string `gorm:"type:char(36)" json:"access_token"`
}
