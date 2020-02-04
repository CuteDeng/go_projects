package main

import (
	"demo/logagent/conf"
	"demo/logagent/kafka"
	"demo/logagent/taillog"
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

var (
	appconf = new(conf.AppConf)
)

func main() {
	err := ini.MapTo(appconf, "./conf/conf.ini")
	if err != nil {
		fmt.Println("ini mapto err:", err)
		return
	}
	fmt.Println(appconf)
	err = kafka.Init([]string{appconf.Address})
	if err != nil {
		fmt.Println("kafka init err:", err)
		return
	}
	fmt.Println("kafka init success")
	err = taillog.Init(appconf.FileName)
	if err != nil {
		fmt.Println("taillog init err:", err)
		return
	}
	fmt.Println("taillog init success")
	run()
}

func run() {
	for {
		select {
		case line := <-taillog.ReadLog():
			kafka.SendKafkaMsg(appconf.Topic, line.Text)
		default:
			time.Sleep(time.Second)
		}
	}
}
