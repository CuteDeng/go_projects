package session

import (
	"errors"
	"sync"

	uuid "github.com/satori/go.uuid"
)

// MemorySessionMgr ...
type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// NewMemorySessionMgr ...
func NewMemorySessionMgr() (msm *MemorySessionMgr) {
	msm = &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1000),
	}
	return
}

// Init ...
func (msm *MemorySessionMgr) Init(address string, options ...string) (err error) {
	return
}

// CreateSession ...
func (msm *MemorySessionMgr) CreateSession() (session Session, err error) {
	msm.rwlock.Lock()
	defer msm.rwlock.Unlock()
	uuid, err := uuid.NewV4()
	if err != nil {
		return
	}
	sessionID := uuid.String()
	session = NewMemorySession(sessionID)
	msm.sessionMap[sessionID] = session
	return
}

// GetSession ...
func (msm *MemorySessionMgr) GetSession(sessionID string) (session Session, err error) {
	msm.rwlock.Lock()
	defer msm.rwlock.Unlock()
	session, ok := msm.sessionMap[sessionID]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
