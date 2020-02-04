package taillog

import "github.com/hpcloud/tail"

var (
	tailObj *tail.Tail
)

// Init ...
func Init(fileName string) (err error) {
	tailObj, err = tail.TailFile(fileName, tail.Config{
		Follow:    true,
		ReOpen:    true,
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2},
		MustExist: false,
		Poll:      true,
	})
	return
}

// ReadLog ...
func ReadLog() chan *tail.Line {
	return tailObj.Lines
}
