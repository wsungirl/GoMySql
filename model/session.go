package model

// ISession is a public interface that should be implemented by DB
type ISession interface {
	GetSessionUser(accessToken string) (*User, error)
	RevokeSession(session *Session) error
	AddSession(session *Session) error
}

// Session struct represents row of DB table sessions
type Session struct {
	UserID      int64  `json:"user_id,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}
