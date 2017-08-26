package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetUser(id uint) (*model.User, error) {
	var user model.User

	err := db.First(&user, id).Error
	return &user, err
}

func (db *DB) GetUserByName(name string) (*model.User, error) {
	var user model.User

	err := db.Where("name=?", name).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return &user, err
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
