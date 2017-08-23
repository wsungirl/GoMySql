package db

import (
	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) AddAction(hist *model.History) error {
	_, err := db.Exec("INSERT INTO history VALUES(?,?,?,?);", hist.DbID, hist.UId, hist.Action, hist.Entity)
	if err != nil {
		return err
	}
	return nil
}
