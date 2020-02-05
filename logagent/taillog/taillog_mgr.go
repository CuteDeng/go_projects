package taillog

import (
	"demo/logagent/etcd"
	"fmt"
	"time"
)

var tskMgr *taillogMgr

type taillogMgr struct {
	logEntryConf []*etcd.LogEntry
	tskMap       map[string]*TailTask
	newConfChan  chan []*etcd.LogEntry
}

// Init ...
func Init(logEntryConf []*etcd.LogEntry) {
	tskMgr = &taillogMgr{
		logEntryConf: logEntryConf,
		tskMap:       make(map[string]*TailTask, 16),
		newConfChan:  make(chan []*etcd.LogEntry),
	}
	for _, logEntry := range logEntryConf {
		NewTailTask(logEntry.Path, logEntry.Topic)
	}
	go tskMgr.run()
}

func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			fmt.Println("新的配置：", newConf)
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan ...
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
