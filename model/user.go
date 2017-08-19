package model

import (
	"golang.org/x/crypto/bcrypt"
)

// IUser is a public interface that should be implemented by DB
type IUser interface {
	AddUser(user *User) error
	GetUser(id int64) (*User, error)
	GetUserByName(name string) (*User, error)
}

// User struct represents row of DB table users
type User struct {
	ID           int64  `json:",omitempty"`
	Name         string `json:",omitempty"`
	Password     string `json:",omitempty"`
	PasswordHash string `json:"password_hash,omitempty"`
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
