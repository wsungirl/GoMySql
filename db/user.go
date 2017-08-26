package db

import (
	"fmt"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetUser(id uint) (user *model.User, err error) {
	err = db.First(user, id).Error
	return
}

func (db *DB) GetUserByName(name string) (user *model.User, err error) {
	err = db.Where("name=?", name).First(user).Error
	return
}

func (db *DB) AddUser(user *model.User) (err error) {
	if user.PasswordHash == "" {
		if err = user.GenPassHash(); err != nil {
			return fmt.Errorf("Can't gen password hash: %v", err)
		}
	}

	if !db.NewRecord(user) {
		goto ERR
	}

	err = db.Create(user).Error
	if err == nil {
		return
	}

ERR:
	return fmt.Errorf("Can't create history item: %v", err)
}
