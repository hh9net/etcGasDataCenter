package types

//单点
type KafKaMsg struct {
	Head KafKaHeader        `json:"head"`
	Data BillGasStationData `json:"data"`
}

type KafKaHeader struct {
	Topic       string `json:"topic"`       //消息类型 如exitdata
	Index       string `json:"index"`       //消息序号,自增
	Topicreply  string `json:"topicreply"`  // 消息回执的主题
	Id          string `json:"id"`          //数据ID
	Topictime   string `json:"topictime"`   // 入kafka的时间
	Lane_id     string `json:"lane_id"`     //车道id
	Parking_id  string `json:"parking_id"`  //停车场id
	Company_id  string `json:"company_id"`  //公司id
	Source_type string `json:"source_type"` // ddd
}

// 出口流水结构(exitdata)
type BillExitDatarrrrrrrrrrer struct {
	Bill_id          string `json:"bill_id"`          // 账单ID Y
	Programver       string `json:"programver"`       // Y	程序版本
	Programstarttime string `json:"programstarttime"` //	Y	程序启动时间
	Company_id       string `json:"company_id"`       // Y	公司/集团ID
	Parking_id       string `json:"parking_id"`       // Y	停车场编号，唯一【ok】
	Channel_id       string `json:"channel_id"`       // N	渠道ID
	Lane_id          string `json:"lane_id"`          // N	车道ID (必须唯一）
	Record_id        string `json:"record_id"`        // Y	商户内记录ID
	Record_no        string `json:"record_no"`        // Y	商户内记录的序号，每条必须加1，平台序号唯一。
	Lane_record_no   string `json:"lane_record_no"`   // Y	停车场的ETC交易序号（每条必须加1，不能重复）
	Etc_terminal_id  string `json:"etc_terminal_id"`  // Y	加密卡号
	Etc_termtrad_no  string `json:"etc_termtrad_no"`  // Y	加密卡交易序号
	Obu_serial       string `json:"obu_serial"`
	Obu_id           string `json:"obu_id"`           // Y	obu_id
	Obu_issuer       string `json:"obu_issuer"`       // Y	obu发行方
	Obu_plate        string `json:"obu_plate"`        // Y	obu内车牌
	Obu_plate_color  string `json:"obu_plate_color"`  // Y	obu车牌颜色
	Obu_expire_date  string `json:"obu_expire_date"`  // Y	obu有效期
	Card_id          string `json:"card_id"`          // Y	卡号
	Card_network     string `json:"card_network"`     // Y	卡网络号
	Card_trade_no    string `json:"card_trade_no"`    // Y	卡交易序号
	Card_issuer      string `json:"card_issuer"`      // Y	卡发行方
	Card_type        string `json:"card_type"`        // Y	卡类型
	Card_expired     string `json:"card_expired"`     // Y	卡到期时间
	Reset_money      string `json:"reset_money"`      // Y	交易后余额
	Money            string `json:"money"`            // Y	金额				【ok】
	Tac              string `json:"tac"`              // Y	TAC码
	Trade_time       string `json:"trade_time"`       // Y	交易时间
	Trade_type       string `json:"trade_type"`       // Y	交易类型
	Lane_key         string `json:"lane_key"`         // Y	车道接入授权
	Vehicle          string `json:"vehicle"`          // Y	车型
	Obu_status       string `json:"obu_status"`       // Y	OBU状态
	Plate_num        string `json:"plate_num"`        // Y	车牌号
	Plate_color      string `json:"plate_color"`      // Y	车牌颜色
	Algorithm_type   string `json:"algorithm_type"`   // Y	算法标识
	Black_ver        string `json:"black_ver"`        // Y	黑名单校验版本
	Notify           string `json:"notify"`           /// Y	记账通知回调
	Entry_time       string `json:"entry_time"`       // Y	用户入口时间
	Duration         string `json:"duration"`         // Y	用户停车时长（分）
	Bill_description string `json:"bill_description"` // Y	账单描述（给用户通知的信息）
	File0019         string `json:"file0019"`         // Y	0019文件
	File0015         string `json:"file0015"`         // Y	0015文件(卡信息)
	Obuinfo          string `json:"obuinfo"`          // Y	0bu信息
	Vehicileinfo     string `json:"vehicileinfo"`     // Y	车辆信息
	Feebackinfo      string `json:"feebackinfo"`      // Y	扣款信息
	Costtime         string `json:"costtime"`         // Y	毫秒
	Psamversion      string `json:"psamversion"`
	Card_version     string `json:"card_version"`
	Obu_type         string `json:"obu_type"`
	Key_type         string `json:"key_type"`
	Before_money     string `json:"before_money"`
	DeviceType       string `json:"devicetype"` // 车道还是读卡器
	SignDate         string `json:"sign_date"`  // 数据签名
}

// 出口流水结构(exitdata)
type BillExitData struct {
	Bill_id          string `json:"bill_id"`          // 账单ID Y
	Programver       string `json:"programver"`       // Y	程序版本
	Programstarttime string `json:"programstarttime"` //	Y	程序启动时间
	Company_id       string `json:"company_id"`       // Y	公司/集团ID
	Parking_id       string `json:"parking_id"`       // Y	停车场编号，唯一
	Channel_id       string `json:"channel_id"`       // N	渠道ID
	Lane_id          string `json:"lane_id"`          // N	车道ID (必须唯一）
	Record_id        string `json:"record_id"`        // Y	商户内记录ID
	Record_no        string `json:"record_no"`        // Y	商户内记录的序号，每条必须加1，平台序号唯一。
	Lane_record_no   string `json:"lane_record_no"`   // Y	停车场的ETC交易序号（每条必须加1，不能重复）
	Etc_terminal_id  string `json:"etc_terminal_id"`  // Y	加密卡号
	Etc_termtrad_no  string `json:"etc_termtrad_no"`  // Y	加密卡交易序号
	Obu_serial       string `json:"obu_serial"`       //obu应用序列号
	Obu_id           string `json:"obu_id"`           // Y	obu_id
	Obu_issuer       string `json:"obu_issuer"`       // Y	obu发行方
	Obu_plate        string `json:"obu_plate"`        // Y	obu内车牌
	Obu_plate_color  string `json:"obu_plate_color"`  // Y	obu车牌颜色
	Obu_expire_date  string `json:"obu_expire_date"`  // Y	obu有效期
	Card_id          string `json:"card_id"`          // Y	卡号
	Card_network     string `json:"card_network"`     // Y	卡网络号
	Card_trade_no    string `json:"card_trade_no"`    // Y	卡交易序号
	Card_issuer      string `json:"card_issuer"`      // Y	卡发行方
	Card_type        string `json:"card_type"`        // Y	卡类型
	Card_expired     string `json:"card_expired"`     // Y	卡到期时间
	Reset_money      string `json:"reset_money"`      // Y	交易后余额
	Money            string `json:"money"`            // Y	金额
	Tac              string `json:"tac"`              // Y	TAC码
	Trade_time       string `json:"trade_time"`       // Y	交易时间
	Trade_type       string `json:"trade_type"`       // Y	交易类型
	Lane_key         string `json:"lane_key"`         // Y	车道接入授权
	Vehicle          string `json:"vehicle"`          // Y	车型
	Obu_status       string `json:"obu_status"`       // Y	OBU状态
	Plate_num        string `json:"plate_num"`        // Y	车牌号
	Plate_color      string `json:"plate_color"`      // Y	车牌颜色
	Algorithm_type   string `json:"algorithm_type"`   // Y	算法标识
	Black_ver        string `json:"black_ver"`        // Y	黑名单校验版本
	Notify           string `json:"notify"`           /// Y	记账通知回调
	Entry_time       string `json:"entry_time"`       // Y	用户入口时间
	Duration         string `json:"duration"`         // Y	用户停车时长（分）
	Bill_description string `json:"bill_description"` // Y	账单描述（给用户通知的信息）
	File0019         string `json:"file0019"`         // Y	0019文件
	File0015         string `json:"file0015"`         // Y	0015文件(卡信息)
	Obuinfo          string `json:"obuinfo"`          // Y	0bu信息
	Vehicileinfo     string `json:"vehicileinfo"`     // Y	车辆信息
	Feebackinfo      string `json:"feebackinfo"`      // Y	扣款信息
	Costtime         string `json:"costtime"`         // Y	扣费耗时 毫秒
	Psamversion      string `json:"psamversion"`      //密钥版本好
	Card_version     string `json:"card_version"`     //卡版本号
	Obu_type         string `json:"obu_type"`         //OBU类型 OBU类型-NEW 1、单片式；2、双片式',
	Key_type         string `json:"key_type"`         //密钥类型[密钥标识]
	Before_money     string `json:"before_money"`     //扣费后余额
	DeviceType       string `json:"devicetype"`       // 车道还是读卡器 0车道 1读卡器  31 加油
}

type BillGasStationData struct {
	BillExitData
	Oil_type     string `json:"oiltype"`      // 油品类型:油品类型 01:97号汽油;02:95号汽油;03:93号汽油;04:92号汽油;05:90号汽油;06:98 号汽油;;11:5号柴油;12:0号柴油 ;13:10号柴油;14:20号柴油;15:35号柴油;16:50号柴油
	Unit_price   string `json:"unitprice"`    // 单位:分
	Litre        string `json:"litre"`        // 加油量 单位：升
	Oilgunnumber string `json:"oilgunnumber"` // 油枪号
	Oilavailable string `json:"oilavailable"` // 可加油余
}

//入口
type BillEntryData struct {
	Bill_id          string `json:"bill_id"`          // Y	账单ID
	Programver       string `json:"programver"`       // Y	程序版本
	Programstarttime string `json:"programstarttime"` // Y	程序启动时间
	Company_id       string `json:"company_id"`       // Y	公司/集团ID
	Parking_id       string `json:"parking_id"`       // Y	停车场编号，唯一
	Channel_id       string `json:"channel_id"`       // N	渠道ID
	Lane_id          string `json:"lane_id"`          // N	车道ID (必须唯一）
	Record_id        string `json:"record_id"`        // Y	商户内记录ID
	Record_no        string `json:"record_no"`        // Y	商户内记录的序号，每条必须加1，平台序号唯一。
	Lane_record_no   string `json:"lane_record_no"`   // Y	停车场的ETC交易序号（每条必须加1，不能重复）
	Etc_terminal_id  string `json:"etc_terminal_id"`  // Y	加密卡号
	Etc_termtrad_no  string `json:"etc_termtrad_no"`  // Y	加密卡交易序号
	Obu_serial       string `json:"obu_serial"`
	Obu_id           string `json:"obu_id"`           // Y	obu_id
	Obu_issuer       string `json:"obu_issuer"`       // Y	obu发行方
	Obu_plate        string `json:"obu_plate"`        // Y	obu内车牌
	Obu_plate_color  string `json:"obu_plate_color"`  // Y	obu车牌颜色
	Obu_expire_date  string `json:"obu_expire_date"`  // Y	obu有效期
	Card_id          string `json:"card_id"`          // Y	卡号
	Card_network     string `json:"card_network"`     // Y	卡网络号
	Card_trade_no    string `json:"card_trade_no"`    // Y	卡交易序号
	Card_issuer      string `json:"card_issuer"`      // Y	卡发行方
	Card_type        string `json:"card_type"`        // Y	卡类型
	Card_expired     string `json:"card_expired"`     // Y	卡到期时间
	Reset_money      string `json:"reset_money"`      // Y	交易后余额
	Money            string `json:"money"`            // Y	金额
	Tac              string `json:"tac"`              // Y	TAC码
	Trade_time       string `json:"trade_time"`       // Y	交易时间
	Trade_type       string `json:"trade_type"`       // Y	交易类型
	Lane_key         string `json:"lane_key"`         // Y	车道接入授权
	Vehicle          string `json:"vehicle"`          // Y	车型
	Obu_status       string `json:"obu_status"`       // Y	OBU状态
	Plate_num        string `json:"plate_num"`        // Y	车牌号
	Plate_color      string `json:"plate_color"`      // Y	车牌颜色
	Algorithm_type   string `json:"algorithm_type"`   // Y	算法标识
	Black_ver        string `json:"black_ver"`        // Y	黑名单校验版本
	Rotify           string `json:"notify"`           // Y	记账通知回调
	Entry_time       string `json:"entry_time"`       // Y	用户入口时间
	Duration         string `json:"duration"`         // Y	用户停车时长（分）
	Bill_description string `json:"bill_description"` // Y	账单描述（给用户通知的信息）
	File0019         string `json:"file0019"`         // Y	0019文件
	File0015         string `json:"file0015"`         // Y	0015文件(卡信息)
	Obuinfo          string `json:"obuinfo"`          // Y	0bu信息
	Vehicileinfo     string `json:"vehicileinfo"`     //  Y    车辆信息
	Feebackinfo      string `json:"feebackinfo"`      // Y	扣款信息
	Costtime         string `json:"costtime"`         // Y	毫秒
	Psamversion      string `json:"psamversion"`
	Card_version     string `json:"card_version"`
	Obu_type         string `json:"obu_type"`
	Key_type         string `json:"key_type"`
	Before_money     string `json:"before_money"`
	DeviceType       string `json:"devicetype"` // 车道还是读卡器
}

type KafKaReplyData struct {
	Id string `json:"id"`
}

//
type BillHourData struct {
	Bill_id       string `json:"bill_id"`
	Company_id    string `json:"company_id"`    // Y	公司/集团ID
	Parking_id    string `json:"parking_id"`    // Y	停车场编号，唯一
	Lane_id       string `json:"lane_id"`       // Y	车道ID
	Record_id     string `json:"record_id"`     //	Y	记录ID
	Datetime_hour string `json:"datetime_hour"` //	Y	YYYYMMDDHH
	Recordcnt     string `json:"recordcnt"`     // Y	记录总数
	Moneycnt      string `json:"moneycnt"`      //	Y	金额总数
}
