package model

type Action int

type Entity int

const (
	READ Action = iota
	CREATE
	UPDATE
)

const (
	DATABASE Entity = iota
	TABLE
	ROW
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

