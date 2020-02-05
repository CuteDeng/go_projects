package main

import (
	"demo/logagent/conf"
	"demo/logagent/etcd"
	"demo/logagent/kafka"
	"demo/logagent/taillog"
	"fmt"
	"sync"
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
	err = kafka.Init([]string{appconf.KafkaConf.Address}, appconf.KafkaConf.MaxSize)
	if err != nil {
		fmt.Println("kafka init err:", err)
		return
	}
	fmt.Println("kafka init success")
	err = etcd.Init([]string{appconf.EtcdConf.Address}, time.Duration(appconf.EtcdConf.Timeout)*time.Second)
	if err != nil {
		fmt.Println("etcd init err:", err)
		return
	}
	fmt.Println("etcd init success")
	//从etcd中读取配置项
	logEntryConf, err := etcd.ReadConf(appconf.EtcdConf.Key)
	if err != nil {
		fmt.Println("etcd read conf err:", err)
		return
	}
	// 监听配置
	var wg sync.WaitGroup
	taillog.Init(logEntryConf)
	NewConfChan := taillog.NewConfChan()
	wg.Add(1)
	go etcd.WatchConf(appconf.EtcdConf.Key, NewConfChan)
	wg.Wait()
}
