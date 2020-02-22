package session

import "fmt"

var (
	sessionMgr SessionMgr
)

// Init ...
func Init(provider, address string, options ...string) (err error) {
	switch provider {
	case "memory":
		sessionMgr = NewMemorySessionMgr()
	case "redis":
		sessionMgr = NewRedisSessionMgr(address)
	default:
		fmt.Errorf("不支持")
		return
	}
	err = sessionMgr.Init(address)
	return
}
