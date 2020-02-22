package session

import (
	"encoding/json"
	"errors"
	"sync"

	"github.com/garyburd/redigo/redis"
)

// SessionFlagNone ...
const (
	SessionFlagNone = iota
	SessionFlagModify
)

// RedisSession ...
type RedisSession struct {
	sessionID string
	pool      *redis.Pool
	data      map[string]interface{}
	rwlock    sync.RWMutex
	flag      int
}

// NewRedisSession ...
func NewRedisSession(id string, pool *redis.Pool) (session *RedisSession) {
	session = &RedisSession{
		sessionID: id,
		data:      make(map[string]interface{}, 16),
		pool:      pool,
		flag:      SessionFlagNone,
	}
	return
}

// Set ...
func (r *RedisSession) Set(key string, value interface{}) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.data[key] = value
	r.flag = SessionFlagModify
	return
}

// Get ...
func (r *RedisSession) Get(key string) (value interface{}, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	value, ok := r.data[key]
	if !ok {
		err = r.loadFromRedis()
		if err != nil {
			return
		}
		value, ok = r.data[key]
		if !ok {
			err = errors.New("key not exists")
			return
		}
	}
	return
}

// loadFromRedis ...
func (r *RedisSession) loadFromRedis() (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	conn := r.pool.Get()
	data, err := redis.String(conn.Do("GET", r.sessionID))
	if err != nil {
		return
	}
	err = json.Unmarshal([]byte(data), &r.data)
	if err != nil {
		return
	}
	return
}

// Del ...
func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	delete(r.data, key)
	r.flag = SessionFlagModify
	return nil
}

// Save ...
func (r *RedisSession) Save() (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	if r.flag != SessionFlagModify {
		return
	}
	data, err := json.Marshal(r.data)
	if err != nil {
		return
	}
	conn := r.pool.Get()
	_, err = conn.Do("SET", r.sessionID, string(data))
	if err != nil {
		return
	}
	r.flag = SessionFlagNone
	return
}
