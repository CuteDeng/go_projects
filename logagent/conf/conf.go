package conf

// AppConf ...
type AppConf struct {
	KafkaConf `ini:"kafka"`
	EtcdConf  `ini:"etcd"`
}

// KafkaConf ...
type KafkaConf struct {
	Address string `ini:"address"`
	MaxSize int    `ini:"chan_max_size"`
}

// EtcdConf ...
type EtcdConf struct {
	Address string `ini:"address"`
	Key     string `ini:"log_collect_key"`
	Timeout int    `ini:"timeout"`
}

// TaillogConf ...
type TaillogConf struct {
	Topic    string `ini:"topic"`
	FileName string `ini:"filename"`
}
