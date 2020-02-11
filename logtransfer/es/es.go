package es

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/olivere/elastic"
)

var (
	client    *elastic.Client
	logDataCh chan *LogData
)

type LogData struct {
	Topic string `json:"topic"`
	Data  string `json:"data"`
}

func Init(address string) (err error) {
	if !strings.HasPrefix(address, "http://") {
		address = "http://" + address
	}
	client, err = elastic.NewClient(elastic.SetURL(address))
	if err != nil {
		fmt.Println("init es err:", err)
		return
	}
	logDataCh = make(chan *LogData, 10000)
	go SendToEs()
	return
}

func SendToChan(ld *LogData) {
	logDataCh <- ld
}

func SendToEs() {
	for {
		select {
		case ld := <-logDataCh:
			put1, err := client.Index().Index(ld.Topic).Type("web_log_data").BodyJson(ld).Do(context.Background())
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("indexed data %s to index %s,type %s\n", put1.Id, put1.Index, put1.Type)
		default:
			time.Sleep(time.Second)
		}
	}
}
