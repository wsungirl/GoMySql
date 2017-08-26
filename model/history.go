package model

import "time"
import "database/sql"

type History struct {
	CreatedAt time.Time

	DB   Database      `json:"db"`
	DBID sql.NullInt64 `gorm:"index" json:"-"`

	User User `gorm:"index" json:"user"`

	Action string `gorm:"type:varchar(16)" json:"action"`
	Entity string `gorm:"type:varchar(16)" json:"entity"`
}
