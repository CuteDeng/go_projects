package main

import (
	"fmt"
	"go_projects/logtransfer/conf"
	"go_projects/logtransfer/es"
	"go_projects/logtransfer/kafka"

	"gopkg.in/ini.v1"
)

func main() {
	var cfg = new(conf.LogTransferCfg)
	err := ini.MapTo(cfg, "./conf/cfg.ini")
	if err != nil {
		fmt.Println("ini map err:", err)
		return
	}
	fmt.Println(cfg)
	err = es.Init(cfg.ESCfg.Address)
	if err != nil {
		fmt.Println("es init err:", err)
		return
	}
	fmt.Println("es init success")
	err = kafka.Init([]string{cfg.KafkaCfg.Address}, cfg.KafkaCfg.Topic)
	if err != nil {
		fmt.Println("kafka init err:", err)
		return
	}
	fmt.Println("init kafka success")
	select {}
}
