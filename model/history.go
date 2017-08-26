package model

import "time"
import "database/sql"

type History struct {
	CreatedAt time.Time

	DB   Database      `form:"ForeignKey:DBID" json:"db"`
	DBID sql.NullInt64 `gorm:"index" json:"-"`

	User   User          `gorm:"ForeignKey:UserID" json:"user"`
	UserID sql.NullInt64 `gorm:"index" json:"-"`

	Action string `gorm:"type:varchar(16)" json:"action"`
	Entity string `gorm:"type:varchar(16)" json:"entity"`
}

func (*History) TableName() string {
	return "history"
}
