package dataServer

import "log"

//从kafka 消费数据

func BillGasStationDataCollect() {

KafkaI:
	log.Println("执行处理kafka数据++++++++++++++++++++++++【kafka执行】+++++++++++++++++++++++++++++++++处理kafka数据")
	//处理kafka数据

	err := ConsumerGroup()
	if err != nil {
		log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++【执行go程 处理kafka数据】 error :", err)
		goto KafkaI
	}

}
