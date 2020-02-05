package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	cli *clientv3.Client
)

// LogEntry ...
type LogEntry struct {
	Path  string `json:"path"`
	Topic string `json:"topic"`
}

// Init ...
func Init(addrs []string, timeout time.Duration) (err error) {
	cli, err = clientv3.New(clientv3.Config{
		Endpoints:   addrs,
		DialTimeout: timeout,
	})
	if err != nil {
		fmt.Println("etcd connetc err:", err)
		return
	}
	return
}

// ReadConf ...
func ReadConf(key string) (logEntryConf []*LogEntry, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	resp, err := cli.Get(ctx, key)
	cancel()
	if err != nil {
		fmt.Println("get failed, err:", err)
		return
	}
	for _, ev := range resp.Kvs {
		err = json.Unmarshal(ev.Value, &logEntryConf)
		if err != nil {
			fmt.Println("unmarshal etcd value err:", err)
			return
		}
	}
	return
}
