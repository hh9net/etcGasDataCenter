package db

import (
	"encoding/json"
	"etcGasDataCenter/types"

	log "github.com/sirupsen/logrus"
	"testing"
)

func TestNewTables(t *testing.T) {
	Newdb()
	//NewTables()
}

func TestQueryRate(t *testing.T) {
	Newdb()
	//数据入库
	//log.Println(QueryRate("00ff25000010", "051a7000001"))
}

func TestChedckyssjDataUpdate(t *testing.T) {
	Newdb()

	//log.Println(ChedckyssjDataUpdate("320102000111032020042916134000000386"))
}

func TestDataStorage(t *testing.T) {
	Newdb()

	msg := `{
	"head": {
		"topic": "billGasStationDataCollectTopic",
		"index": "101",
		"topicreply": "SG_GATEWAY_mytopic_test_gasstation",
		"id": "3201100020110120201209144627000030f5",
		"topictime": "2020-12-09 16:36:36",
		"lane_id": "1101",
		"parking_id": "3201111110",
		"company_id": "3201100020",
		"source_type": "ddd"
	},
	"data": {
		"algorithm_type": "00",
		"before_money": "1999816507",
		"bill_description": "",
		"bill_id": "3201100020110120201209144627000030f5",
		"black_ver": "20201209",
		"card_expired": "20291102",
		"card_id": "1910230200395778",
		"card_issuer": "d5e3bdad33010001",
		"card_network": "3301",
		"card_trade_no": "0282",
		"card_type": "23",
		"card_version": "65",
		"channel_id": "",
		"company_id": "3201100020",
		"costtime": "664",
		"devicetype": "30",
		"duration": "0",
		"entry_time": "",
		"etc_terminal_id": "01320002da12",
		"etc_termtrad_no": "000030f1",
		"feebackinfo": "",
		"file0015": "d5e3bdad330100011741330119102302003957782019110220291102d5e341443635383232000000000401",
		"file0019": "aa290000000000000000000000000000000000000000000000000000000000000000000000000000000000",
		"key_type": "9",
		"lane_id": "1101",
		"lane_key": "nil",
		"lane_record_no": "1",
		"litre": "10",
		"money": "100",
		"notify": "",
		"obu_expire_date": "20291102",
		"obu_id": "64c1773f",
		"obu_issuer": "d5e3bdad33010001",
		"obu_plate": "浙AD65822",
		"obu_plate_color": "4",
		"obu_serial": "3301611910470419",
		"obu_status": "2200",
		"obu_type": "0",
		"obuinfo": "",
		"oilavailable": "1000",
		"oilgunnumber": "1号油枪",
		"oiltype": "02",
		"parking_id": "3201111110",
		"plate_color": "4",
		"plate_num": "浙AD65822",
		"programstarttime": "2020-12-09 14:46:23",
		"programver": "V1.0.20201118.release",
		"psamversion": "1",
		"record_id": "32011000201101202012091446270000",
		"record_no": "1",
		"reset_money": "1999816407",
		"tac": "224",
		"trade_time": "2020-12-09 14:46:27",
		"trade_type": "0",
		"unitprice": "2000",
		"vehicileinfo": "",
		"vehicle": "1"
	}
}`

	data := new(types.KafKaMsg)
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		log.Println("执行  types.DdkafkaTopic  json.Unmarshal error:", err)
	}
	log.Println(data)
	log.Println(DataStorage(data))
}
