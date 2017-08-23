package db

import (
	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) SetPermissions(perm *model.Permissions) error {
	_, err := db.Exec("INSERT INTO permissions VALUES(?,?,?,?);", perm.DbID, perm.UId, perm.Action, perm.Entity)
	if err != nil {
		return err
	}
	return nil
}
