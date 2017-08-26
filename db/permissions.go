package db

import (
	"fmt"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) SetPermission(perm *model.Permission) (err error) {
	if db.NewRecord(perm) {
		err = fmt.Errorf("Can't create permission: %v",
			db.Create(perm).Error)
	} else {
		err = fmt.Errorf("Can't update permission: %v",
			db.Save(perm).Error)
	}

	return
}
