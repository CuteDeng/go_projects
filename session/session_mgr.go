package session

// SessionMgr ...
type SessionMgr interface {
	Init(address string, options ...string) (err error)
	CreateSession() (session Session, err error)
	GetSession(sessionID string) (session Session, err error)
}
