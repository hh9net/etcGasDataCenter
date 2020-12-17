package main

import (
	"etcGasDataCenter/config"
	"etcGasDataCenter/dataServer"
	"etcGasDataCenter/db"
	"etcGasDataCenter/types"
	"etcGasDataCenter/utils"

	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	conf := config.ConfigInit() //初始化配置
	log.Println("配置文件信息：", *conf)
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)

	//初始化数据库
	db.Newdb()

	types.KafkaIp = conf.KafkaIp
	types.DdkafkaTopic = conf.DdkafkaTopic
	types.DdkafkaHourTopic = conf.DdkafkaHourTopic
	log.Println("conf.KafkaIp:", conf.KafkaIp)
	log.Println("conf.DdkafkaTopic:", conf.DdkafkaTopic)
	log.Println("DdkafkaHourTopic:", conf.DdkafkaHourTopic)

	dataServer.BillGasStationDataCollect()
}
