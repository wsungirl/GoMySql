package db

import (
	"database/sql"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetUser(id int64) (*model.User, error) {
	var user model.User

	row := db.QueryRow("SELECT * FROM users WHERE id=?", id)

	if err := row.Scan(&user.ID, &user.Name, &user.PasswordHash); err != nil {
		return nil, err
	}

	return &user, nil
}

func (db *DB) GetUserByName(name string) (*model.User, error) {
	var user model.User

	row := db.QueryRow("SELECT id, name, password_hash FROM users WHERE name=?", name)

	if err := row.Scan(&user.ID, &user.Name, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (db *DB) AddUser(user *model.User) error {
	if user.PasswordHash == "" {
		if err := user.GenPassHash(); err != nil {
			return err
		}
	}

	_, err := db.Exec(
		"INSERT INTO users(name, password_hash) VALUES(?,?)",
		user.Name, user.PasswordHash,
	)

	return err
}
