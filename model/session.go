package model

type ISessionUserGet interface {
	GetSessionUser(access_token string) (*User, error)
}

type Session struct {
	UserId      int64  `json:"user_id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
