package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var conffilepath = "./conf/config.toml" // go run main.go
//var conffilepath = "../conf/config.toml"

type Config struct { //配置文件要通过tag来指定配置文件中的名称
	//mysql 配置
	MHostname     string `ini:"mysql_hostname"`
	MPort         string `ini:"mysql_port"`
	MUserName     string `ini:"mysql_user"`
	MPass         string `ini:"mysql_pass"`
	Mdatabasename string `ini:"mysql_databasename"`
	MKeepalive    int    `ini:"mysql_keepalive"`
	MTimeout      int    `ini:"mysql_timeout"`

	//redis
	RedisAddr         string `ini:"redis_addr"`
	RedisPass         string `ini:"redis_pass"`
	Redisdatabasename string `ini:"redis_databasename"`
	//日志
	LogPath         string `ini:"log_Path"`
	LogMaxAge       int64  `ini:"log_maxAge"`
	LogRotationTime int64  `ini:"log_rotationTime"`
	LogFileName     string `ini:"log_FileName"`

	//外网id
	IpAddress string `ini:"ip_address"`

	//频率
	Frequency int `ini:"frequency"`

	//kafkaip
	KafkaIp string `ini:"kafka_ip"`

	KafkaIpa         string `ini:"kafka_ipa"`
	KafkaIpb         string `ini:"kafka_ipb"`
	KafkaIpc         string `ini:"kafka_ipc"`
	DdkafkaTopic     string `ini:"ddkafka_topic"`
	DdkafkaHourTopic string `ini:"ddkafka_hour_topic"` //ddkafka_hour_topic
}

//读取配置文件并转成结构体
func ReadConfig(path string) (Config, error) {
	var config Config
	conf, err := ini.Load(path) //加载配置文件
	if err != nil {
		log.Println("load config file fail!")
		return config, err
	}
	conf.BlockMode = false
	err = conf.MapTo(&config) //解析成结构体
	if err != nil {
		log.Println("mapto config file fail!")
		return config, err
	}
	return config, nil
}

//获取mysql 配置文件信息
func ConfigInit() *Config {
	//读配置文件
	config, err := ReadConfig(conffilepath) //也可以通过os.arg或flag从命令行指定配置文件路径
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(config)
	return &config
}
