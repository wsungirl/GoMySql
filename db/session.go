package db

import (
	"fmt"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetSession(token string) (*model.Session, error) {
	var session model.Session

	err := db.Preload("User").Where("access_token = ?", token).First(&session).Error

	return &session, err
}

// func (db *DB) GetSessionUser(session *model.Session) (user *model.User, err error) {
// 	err = db.Model(session).Related(user).Error
// 	return
// }

func (db *DB) AddSession(session *model.Session) (err error) {
	if !db.NewRecord(session) {
		goto ERR
	}

	err = db.Create(session).Error
	if err == nil {
		return
	}

ERR:
	return fmt.Errorf("Can't create session record: %v", err)
}

func (db *DB) RevokeSession(session *model.Session) (err error) {
	err = db.Delete(session).Error

	return
}
