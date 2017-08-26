package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// User struct represents row of DB table users
type User struct {
	ID        uint      `gorm:"primary_key" json:"id,omitempty"`
	CreatedAt time.Time `json:"-"`

	Name         string `json:"name,omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
	Password     string `gorm:"-" json:"password,omitempty"`

	Databases []Database `json:"-"`
}

// GenPassHash generates bcrypt hash and rewrites PasswordHash field of User struct
func (user *User) GenPassHash() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hash)

	return nil
}
