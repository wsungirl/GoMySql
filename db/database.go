package db

import (
	"fmt"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetDB(id uint) (dbMod *model.Database, err error) {
	err = db.Preload("User").First(dbMod, id).Error
	return
}

func (db *DB) GetUserDatabases(user *model.User) ([]model.Database, error) {
	var dbs []model.Database
	err := db.Model(user).Related(&dbs).Error
	return dbs, err
}

func (db *DB) CreateDB(dbMod *model.Database) (err error) {

	if db.NewRecord(dbMod) {
		goto NOERR
	}

ERR:
	return fmt.Errorf("Can't create new database record: %v", err)

NOERR:
	err = db.Create(dbMod).Error
	if err != nil {
		goto ERR
	}

	err = db.Exec("CREATE SCHEMA " + dbMod.GetStoredName()).Error
	if err == nil {
		return
	}

	delErr := db.Delete(dbMod).Error
	if delErr != nil {
		return fmt.Errorf("Can't delete newly created db: <%v> while creating schema: %v", delErr, err)
	}

	return fmt.Errorf("Can't create schema: %v", err)
}
