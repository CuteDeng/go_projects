package taillog

import (
	"demo/logagent/etcd"
)

var tskMgr *taillogMgr

type taillogMgr struct {
	logEntryConf []*etcd.LogEntry
}

// Init ...
func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntryConf: logEntryConf,
	}
	for _, logEntry := range logEntryConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
}
