package dataServer

import (
	log "github.com/sirupsen/logrus"
	"time"
)

//从kafka 消费数据

func BillGasStationDataCollect() {

KafkaI:
	log.Println("执行处理kafka数据+++++++【kafka执行】++++处理kafka数据")
	//处理kafka数据
	err := ConsumerGroup()
	if err != nil {
		log.Println("+++++++++++++++++【执行ConsumerGroup() 处理kafka数据】 error :", err)
		time.Sleep(time.Second * 30)
		goto KafkaI
	}

}
