package model

type ISession interface {
	GetSessionUser(accessToken string) (*User, error)
	RevokeSession(session *Session) error
	AddSession(session *Session) error
}

type Session struct {
	UserID      int64  `json:"user_id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
