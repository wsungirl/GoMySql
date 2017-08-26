package model

const (
	READ_ACTION = "read"
	CREATE_ACTION = "create"
	UPDATE_ACTION = "update"
	DATABASE_ENTITY = "db"
	TABLE_ENTITY = "table"
	ROW_ENTITY = "row"
)

type IPermissions interface {
	SetPermissions(perm *Permissions) error
	CheckPermissions(perm *Permissions) (bool, error)
}

type Permissions struct {
	DbID   int64  `json:"db_id,omitempty"`
	UId    int64  `json:"u_id,omitempty"`
	Action string `json:"action,omitempty"`
	Entity string `json:"entity,omitempty"`
}

