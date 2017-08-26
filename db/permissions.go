package db

import (
	"github.com/wsungirl/GoMySql/model"
	"errors"
)

func (db *DB) SetPermissions(perm *model.Permissions) error {
	_, err := db.Exec("INSERT INTO permissions VALUES(?,?,?,?);", perm.DbID, perm.UId, perm.Action, perm.Entity)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) CheckPermissions(perm * model.Permissions) (bool, error) {
	row := db.QueryRow("SELECT id FROM permissions WHERE db_id=? and user_id=? and action=? and entity=?;", perm.DbID, perm.UId, perm.Action, perm.Entity)
	var idDb int64
	err := row.Scan(&idDb)
	if err != nil {
		return false, errors.New("error while checking permissions: " + err.Error())
	}
	return true, nil
}