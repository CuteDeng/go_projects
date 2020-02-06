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
		tailObj := NewTailTask(logEntry.Path, logEntry.Topic)
		mk := fmt.Sprintf("%s_%s", logEntry.Path, logEntry.Topic)
		tskMgr.tskMap[mk] = tailObj
	}
	go tskMgr.run()
}

func (t *taillogMgr) run() {
	for {
		select {
		case newConf := <-t.newConfChan:
			for _, conf := range newConf {
				mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
				_, ok := t.tskMap[mk]
				fmt.Println("ok:", ok)
				if ok {
					continue
				} else {
					tailObj := NewTailTask(conf.Path, conf.Topic)
					t.tskMap[mk] = tailObj
				}
			}
			for _, conf := range t.logEntryConf {
				isDelete := true
				for _, nconf := range newConf {
					if conf.Path == nconf.Path && conf.Topic == nconf.Topic {
						isDelete = false
						continue
					}
				}
				if isDelete {
					mk := fmt.Sprintf("%s_%s", conf.Path, conf.Topic)
					t.tskMap[mk].cancelFunc()
				}
			}
		default:
			time.Sleep(time.Second)
		}
	}
}

// NewConfChan ...
func NewConfChan() chan<- []*etcd.LogEntry {
	return tskMgr.newConfChan
}
