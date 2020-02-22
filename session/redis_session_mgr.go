package session

import (
	"errors"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

// RedisSessionMgr ...
type RedisSessionMgr struct {
	sessionMap map[string]Session
	address    string
	password   string
	rwlock     sync.RWMutex
	pool       *redis.Pool
}

// NewRedisSessionMgr ...
func NewRedisSessionMgr(address string) (rsm *RedisSessionMgr) {
	rsm = &RedisSessionMgr{
		sessionMap: make(map[string]Session, 1000),
		address:    address,
	}
	return
}

// Init ...
func (rsm *RedisSessionMgr) Init(address string, options ...string) (err error) {
	var password string
	if len(options) > 0 {
		password = options[0]
	}
	rsm.pool = myPool(address, password)
	rsm.address = address
	return
}

func myPool(address, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     64,
		MaxActive:   1000,
		IdleTimeout: time.Second * 60,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			// 验证密码是否正确
			_, err = conn.Do("AUTH", password)
			if err != nil {
				conn.Close()
				return nil, err
			}
			return conn, nil
		},
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},
	}
}

// CreateSession ...
func (rsm *RedisSessionMgr) CreateSession() (session Session, err error) {
	rsm.rwlock.Lock()
	defer rsm.rwlock.Unlock()
	uuid, err := uuid.NewV4()
	if err != nil {
		return
	}
	sessionID := uuid.String()
	session = NewRedisSession(sessionID, rsm.pool)
	rsm.sessionMap[sessionID] = session
	return
}

// GetSession ...
func (rsm *RedisSessionMgr) GetSession(sessionID string) (session Session, err error) {
	rsm.rwlock.Lock()
	defer rsm.rwlock.Unlock()
	session, ok := rsm.sessionMap[sessionID]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}
