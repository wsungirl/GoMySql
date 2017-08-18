package model

type Session struct {
	UserId  int `json:"user_id"`
	AccessTocken string  `json:"access_tocken"`
}

