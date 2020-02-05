package taillog

import (
	"fmt"
	"time"

	"demo/logagent/kafka"

	"github.com/hpcloud/tail"
)

// TailTask ...
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

// NewTailTask ...
func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.init()
	return
}

// Init ...
func (t *TailTask) init() {
	config := tail.Config{
		Follow:    true,
		ReOpen:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file err:", err)
	}
	go t.run()
}

// run ...
func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines:
			kafka.SendChan(t.topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
