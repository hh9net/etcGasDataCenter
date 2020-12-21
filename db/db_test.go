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

func TestHourDataStorage(t *testing.T) {
	Newdb()
	msg := `{
	"head": {
		"topic": "hoursGasStationStatisTopic",
		"index": "0",
		"topicreply": "SG_GATEWAY_mytopic_test",
		"id": "100200999911012020121408",
		"topictime": "",
		"lane_id": "1101",
		"parking_id": "1002009999",
		"company_id": "3202999999",
		"source_type": "ddd"
	},
	"data": {
		"bill_id": "100200999911012020121508",
		"company_id": "3202999999",
		"parking_id": "1002009999",
		"lane_id": "1101",
		"record_id": "1002009999110120201214110",
		"datetime_hour": "2020121507",
		"recordcnt": "0",
		"moneycnt": "0"
	}
}`

	data := new(types.KafKaBillHourMsg)
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		log.Println("执行  types.DdkafkaTopic  json.Unmarshal error:", err)
	}
	log.Println(data)
	log.Println(HourDataStorage(data))

}

func TestDataStorage(t *testing.T) {
	Newdb()

	msg := `{
	"head": {
		"topic": "billGasStationDataCollectTopic",
		"index": "101",
		"topicreply": "SG_GATEWAY_mytopic_test_gasstation",
		"id": "4201100020110120201209144627000030c17",
		"topictime": "2020-12-18 16:36:36",
		"lane_id": "1101",
		"parking_id": "3201111110",
		"company_id": "3201100020",
		"source_type": "ddd"
	},
	"data": {
		"algorithm_type": "00",
		"before_money": "1999816507",
		"bill_description": "",
		"bill_id": "4201100020110120201209144627000030c17",
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
		"programstarttime": "2020-12-18 14:46:23",
		"programver": "V1.0.20201118.release",
		"psamversion": "1",
		"record_id": "42011000201101202012091446270007",
		"record_no": "1",
		"reset_money": "1999816407",
		"tac": "224",
		"trade_time": "2020-12-18 14:46:27",
		"trade_type": "0",
		"unitprice": "2000",
		"vehicileinfo": "",
		"vehicle": "1"
	}
}`

	//s:=	"INSERT INTO `b_dd_chedckyssj` (`F_VC_JIAOYJLID`,`F_VC_JIAOYTJR`,`F_VC_JIAOYTJRS`,`F_VC_TINGCCBH`,`F_VC_CHEDID`,`F_VC_GONGSJTID`,`F_VC_SHANGHNJLID`,`F_VC_SHANGHNJLXH`,`F_VC_JIAMKH`,`F_VC_KAJMJYXH`,`F_VC_OBUID`,`F_VC_OBUFXF`,`F_VC_OBUCP`,`F_VC_OBUCPYS`,`F_DT_OBUYXQ`,`F_VC_KAH`,`F_VC_KAWLH`,`F_VC_KAJYXH`,`F_VC_KAFXF`,`F_NB_KALX`,`F_VC_KADQSJ`,`F_NB_JIAOYQYE`,`F_NB_JIAOYHYE`,`F_NB_JINE`,`F_VC_TACM`,`F_DT_JIAOYSJ`,`F_DT_JIAOYLX`,`F_VC_CHDJRSQ`,`F_VC_CHEX`,`F_VC_OBUZT`,`F_VC_CHEPH`,`F_VC_CHPYS`,`F_VC_SUANFBS`,`F_VC_HEIMDJYBB`,`F_VC_JIZTZHD`,`F_VC_ZHANGDMS`,`F_VC_SHUJQM`,`F_DT_CAIJSJ`,`F_VC_JIAOYZT`,`F_VC_HUIDTZSJ`,`F_VC_HUIDTZZT`,`F_VC_HUIDTZCS`,`F_VC_KABBH`,`F_VC_MIYBBH`,`F_VC_0019`,`F_VC_0015`,`F_VC_OBUYYXLH`,`F_VC_CHDJYXH`,`F_NB_SHIFTF`,`F_NB_SHOUXF`,`F_NB_OBULX`,`F_NB_MIYBS`,`F_VC_CHENGXBBH`,`F_DT_CHENGXQDSJ`,`F_NB_KOUKHS`,`F_NB_FUWSLRL`,`F_NB_FUWSLRE`,`F_NB_JIAOYZT`,`F_NB_SHUJHQFS`,`F_NB_TIANTJZT`,`F_NB_YINGYCJ`,`F_NB_SHEBLX`) VALUES ('3201100020110120201209144627000030a1','20201218','14','3201111110','1101','3201100020','32011000201101202012091446270002','1','01320002da12','000030f1','64c1773f','d5e3bdad33010001','浙AD65822','4','2029-11-02 00:00:00','1910230200395778','3301','0282','d5e3bdad33010001',23,'2029-11-02 00:00:00',1999816507,1999816407,100,'224','2020-12-18 14:46:27','0','nil','1','2200','浙AD65822','4','00','20201209','','','0','2020-12-18 18:11:35',0,'2020-12-18 18:11:35',1,1,'65','1','aa290000000000000000000000000000000000000000000000000000000000000000000000000000000000','d5e3bdad330100011741330119102302003957782019110220291102d5e341443635383232000000000401','3301611910470419','1',0,60,0,9,'V1.0.20201118.release','2020-12-18 14:46:23',664,0,0,0,0,0,31,0)"
	//msg1 :=`{billGasStationDataCollectTopic 21 SG_GATEWAY_mytopic_test_gasstation 200200999811012020121103 2020-12-11 16:00:05    ddd} {{200200999811012020121103       200200999811012020121103XXXXXXXX 0 0 2002009998  3202999999    0  V1.0.2020111811.release 1101   0  0 0   0 nil 0   0     0       0 0 0 0 0 0 0}     }}`
	data := new(types.KafKaMsg)
	err := json.Unmarshal([]byte(msg), &data)
	if err != nil {
		log.Println("执行  types.DdkafkaTopic  json.Unmarshal error:", err)
	}
	log.Println(data)
	log.Println(DataStorage(data))
}
