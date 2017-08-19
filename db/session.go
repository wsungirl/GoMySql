package db

import (
	"database/sql"

	"github.com/wsungirl/GoMySql/model"
)

func (db *DB) GetSessionUser(access_token string) (*model.User, error) {
	var user model.User

	row := db.QueryRow(
		"SELECT id, name, password_hash FROM users WHERE id=(SELECT user_id FROM sessions WHERE access_token=?)",
		access_token,
	)

	if err := row.Scan(&user.ID, &user.Name, &user.PasswordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (db *DB) AddSession(session *model.Session) error {

	if _, err := db.Exec(
		"INSERT INTO sessions (user_id, access_token) VALUES (?,?)",
		session.UserID, session.AccessToken,
	); err != nil {
		return err
	}

	return nil
}

func (db *DB) RevokeSession(session *model.Session) error {
	if _, err := db.Exec(
		"DELETE FROM sessions WHERE access_token=?",
		session.AccessToken,
	); err != nil {
		return err
	}

	return nil
}
