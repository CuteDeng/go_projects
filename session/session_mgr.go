package session

type SessionMgr interface {
	Init(address string, options ...string) (err error)
	CreateSession() (session Session, err error)
	GetSession(sessionId string) (session Session, err error)
}
