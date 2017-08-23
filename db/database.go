package db

import (
	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetDBInfo(idDb int64) (*model.Database, error) {
	var mDb model.Database
	row := db.QueryRow("SELECT * FROM dbs WHERE id=?", idDb)
	err := row.Scan(&mDb.ID, &mDb.Uid, &mDb.DBname)
	if err != nil {
		return nil, err
	}
	return &mDb, nil
}

func (db *DB) GetDBList(idUser int64) ([]string, error) {
	var str string
	var names []string
	rows, err := db.Query("SELECT db_name FROM dbs WHERE user_id=?", idUser)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&str)
		if err != nil {
			return nil, err
		}
		names = append(names, str)
	}
	return names, nil
}

func (db *DB) CreateDB(sDb *model.Database) error {
	_, err := db.Exec("INSERT INTO dbs VALUES(?,?,?);", sDb.ID, sDb.Uid, sDb.DBname)
	if err != nil {
		return err
	}
	return nil
}
