package main

import (
	"etcGasDataCenter/config"
	"etcGasDataCenter/utils"
	"log"
	"time"
)

func main() {
	conf := config.ConfigInit() //初始化配置
	log.Println("配置文件信息：", *conf)
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)

}
