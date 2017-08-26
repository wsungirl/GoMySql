package db

import (
	"github.com/wsungirl/GoMySql/model"
	"errors"
	"fmt"
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

	_, err := db.Exec("INSERT INTO dbs (`user_id`, `db_name`) VALUES(?,?);", sDb.Uid, sDb.DBname)

	if err != nil {
		return errors.New("Error inserting DB info: "+err.Error() )
	}

	newDbId := db.QueryRow("SELECT LAST_INSERT_ID();")

	if newDbId == nil {
		return errors.New("Cant get last ID" )
	}


	var schemaName string
	var dbIdText string
	err = newDbId.Scan(&dbIdText)
	if err != nil{
		return errors.New("Error getting DbID" + err.Error())
	}

	schemaName = fmt.Sprintf("db_%d_%s", sDb.Uid, dbIdText)

	_, err = db.Exec("INSERT INTO permissions (`user_id`, `db_id`,`action`, `entity`) VALUES(?,?,'read', 'db'),(?,?,'update', 'db');", sDb.Uid, dbIdText, sDb.Uid, dbIdText)
	if err != nil {
		return errors.New("Error inserting DB permissions info: "+err.Error() )
	}

	_, err = db.Exec("CREATE SCHEMA " + schemaName )
	if err != nil {
		return errors.New("cant create schema: "+ err.Error())
	}

	return nil
}
