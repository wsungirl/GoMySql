package db

import (
	"fmt"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) AddHistory(hist *model.History) (err error) {
	if !db.NewRecord(hist) {
		goto ERR
	}

	err = db.Create(hist).Error
	if err == nil {
		return
	}

ERR:
	return fmt.Errorf("Can't create history item: %v", err)
}
