package model

type IHistory interface {
	AddAction(hist *History) error
}

type History struct {
	DbID   int64  `json:"db_id,omitempty"`
	UId    int64  `json:"u_id,omitempty"`
	Action string `json:"action,omitempty"`
	Entity string `json:"entity,omitempty"`
}
