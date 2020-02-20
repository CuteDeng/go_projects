package session

import "sync"

import "errors"

// MemorySession ...
type MemorySession struct {
	sessionID string
	data      map[string]interface{}
	rwlock    sync.RWMutex
}

// NewMemorySession ...
func NewMemorySession(id string) (session *MemorySession) {
	session = &MemorySession{
		sessionID: id,
		data:      make(map[string]interface{}, 16),
	}
	return
}

// Set ...
func (m *MemorySession) Set(key string, value interface{}) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	m.data[key] = value
	return
}

// Get ...
func (m *MemorySession) Get(key string) (value interface{}, err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value, ok := m.data[key]
	if !ok {
		err = errors.New("key not exists in session")
		return
	}
	return
}

// Del ...
func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data, key)
	return nil
}

// Save ...
func (m *MemorySession) Save() error {
	return nil
}
