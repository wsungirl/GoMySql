package db

import (
	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetDBInfo(dbObj *model.Database) (*model.Database, error) {
	var mDb model.Database
	row := db.QueryRow("SELECT id, user_id, db_name FROM dbs WHERE id=?", dbObj.ID)
	err := row.Scan(&mDb.ID, &mDb.Uid, &mDb.DBname)
	if err != nil {
		return nil, err
	}
	return &mDb, nil
}

func (db *DB) GetDBList(user *model.User) ([]model.Database, error) {
	var tmpDB model.Database
	var databases []model.Database
	rows, err := db.Query("SELECT id, user_id, db_name  FROM dbs WHERE user_id=?", user.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tmpDB.ID, &tmpDB.Uid, &tmpDB.DBname)
		if err != nil {
			return nil, err
		}
		databases = append(databases, tmpDB)
	}
	return databases, nil
}

func (db *DB) CreateDB(sDb *model.Database) error {
	_, err := db.Exec("INSERT INTO dbs VALUES(?,?,?);", sDb.ID, sDb.Uid, sDb.DBname)
	if err != nil {
		return err
	}
	return nil
}
