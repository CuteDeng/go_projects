package conf

// AppConf ...
type AppConf struct {
	KafkaConf   `ini:"kafka"`
	TaillogConf `ini:"taillog"`
}

// KafkaConf ...
type KafkaConf struct {
	Address string `ini:"address"`
}

// TaillogConf ...
type TaillogConf struct {
	Topic    string `ini:"topic"`
	FileName string `ini:"filename"`
}
