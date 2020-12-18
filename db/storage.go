package db

import (
	"errors"
	"etcGasDataCenter/config"
	"etcGasDataCenter/types"
	"etcGasDataCenter/utils"
	"fmt"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

//结算监控平台数据层：数据的增删改查
func Newdb() {
	conf := config.ConfigInit() //初始化配置
	utils.InitLogrus(conf.LogPath, conf.LogFileName, time.Duration(24*conf.LogMaxAge)*time.Hour, time.Duration(conf.LogRotationTime)*time.Hour)
	mstr := conf.MUserName + ":" + conf.MPass + "@tcp(" + conf.MHostname + ":" + conf.MPort + ")/" + conf.Mdatabasename + "?charset=utf8&parseTime=true&loc=Local"
	utils.DBInit(mstr) //初始化数据库
}

//1、查询表是否存在
func QueryTable(tablename string) {
	db := utils.GormClient.Client
	is := db.HasTable(tablename)

	if is == false {
		log.Println("不存在", tablename)
		return
	}
	log.Println("表存在：", tablename, is)
}

//2、单点车道出口原始数据表插入
func InsertChedckyssjData(Chedckyssj *types.BDdChedckyssj) error {
	db := utils.GormClient.Client
	if err := db.Table("b_dd_chedckyssj").Create(Chedckyssj).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_dd_chedckyssj error:", err)
		return err
	}
	log.Println("单点车道出口原始数据表插入成功！", Chedckyssj.FVcJiaoyjlid)
	return nil
}

//2.1
func QueryChedckyssjData(jyjlid string) error {
	db := utils.GormClient.Client
	yssj := new(types.BDdChedckyssj)
	if err := db.Table("b_dd_chedckyssj").Where("F_VC_JIAOYJLID=?", jyjlid).First(yssj).Error; err != nil {
		// 错误处理...
		log.Println("Query b_dd_chedckyssj error:", err)
		if fmt.Sprint(err) == "record not found" {
			log.Println("QueryChedckyssjData err == `record not found`:", err)
			return nil
		}
		return err
	}
	log.Println("Query 单点车道出口原始数据表插入成功！", yssj.FVcJiaoyjlid)
	return errors.New("单点车道出口原始数据已经存在")
}

//3、单点车道出口原始数据加油明细插入
func InsertChedckyssjJYMXData(jymx *types.BJyzJiaymx) error {
	db := utils.GormClient.Client
	if err := db.Table("b_jyz_jiaymx").Create(jymx).Error; err != nil {
		// 错误处理...
		log.Println("Insert b_jyz_jiaymx error:", err)
		return err
	}
	log.Println("单点车道出口原始数据加油明细插入成功！")
	return nil
}

type Result struct {
	Rate int
}

//4 查询公司、停车场费率  b_tcc_tingcc     tp_gongsjt
func QueryRate(Parking_id, Company_id string) *Result {
	db := utils.GormClient.Client

	result := new(Result)

	sqlstr := `select F_NB_FEIL as rate  from b_tcc_tingcc where F_VC_TINGCCBH = ?`
	mydb := db.Raw(sqlstr, Parking_id).Scan(result)

	if mydb.Error != nil {
		Result := new(Result)
		log.Println(mydb.Error, "++++++++++++")
		if fmt.Sprint(mydb.Error) == "record not found" {
			log.Println("查询 停车场费率 为空")
		}

		sqlstr := `select F_VC_JIESFL as rate from tp_gongsjt where F_VC_GONGSJTID = ?`
		mydb := db.Raw(sqlstr, Company_id).Scan(Result)
		if mydb.Error != nil {
			log.Println(mydb.Error, "++++++++++++")

			if fmt.Sprint(mydb.Error) == "record not found" {
				log.Println("查询 公司费率 为空")
			}
			return nil
		}
		return Result
	}
	return result
}

//单点车道出口原始数据入库
func DataStorage(data *types.KafKaMsg) error {

	ysdata := new(types.BDdChedckyssj)

	log.Println("交易时间:", data.Data.Trade_time)
	if data.Data.Trade_time == "" {
		return errors.New("交易时间不能为空")
	}

	JYsj := strings.Split(data.Data.Trade_time, " ") //2020-04-29 16:28:55
	if len(JYsj) != 2 {
		return errors.New("交易时间,交易时间格式不正确")
	}

	JYTJY := strings.Split(JYsj[0], "-") //2020-04-29 16:28:55
	if len(JYTJY) != 3 {
		return errors.New("交易统计日,交易时间格式不正确")
	}

	qyssjerr := QueryChedckyssjData(ysdata.FVcJiaoyjlid)
	if qyssjerr != nil {
		if fmt.Sprint(qyssjerr) == "单点车道出口原始数据已经存在" {
			log.Println(qyssjerr)
			return qyssjerr
		} else {
			return qyssjerr
		}
	}

	//数据赋值
	ysdata.FVcJiaoytjr = JYTJY[0] + JYTJY[1] + JYTJY[2] //  `F_VC_JIAOYTJR` varchar(32) DEFAULT NULL COMMENT '交易统计日',
	log.Println("数据赋值交易统计日:", ysdata.FVcJiaoytjr)

	JYTJS := strings.Split(JYsj[1], ":")
	if len(JYTJS) != 3 {
		return errors.New("交易统计时,交易时间格式不正确")
	}
	ysdata.FVcJiaoytjrs = JYTJS[0] //  `F_VC_JIAOYTJRS` varchar(32) DEFAULT NULL COMMENT '交易统计时',
	log.Println("数据赋值交易统计时:", ysdata.FVcJiaoytjrs)
	ysdata.FVcJiaoyjlid = data.Data.Bill_id //  `F_VC_JIAOYJLID` varchar(128) NOT NULL COMMENT '交易记录ID 停车场ID+车道ID+年月日时分秒+加密卡交易序号',
	log.Println("数据赋值交易记录ID:", ysdata.FVcJiaoyjlid)

	ysdata.FVcTingccbh = data.Data.Parking_id      //  `F_VC_TINGCCBH` varchar(32) NOT NULL COMMENT '停车场编号',
	ysdata.FVcChedid = data.Data.Lane_id           //  `F_VC_CHEDID` varchar(32) NOT NULL COMMENT '车道ID',
	ysdata.FVcGongsjtid = data.Data.Company_id     //  `F_VC_GONGSJTID` varchar(32) NOT NULL COMMENT '公司/集团ID',
	ysdata.FVcQudid = data.Data.Channel_id         //  `F_VC_QUDID` varchar(32) DEFAULT NULL COMMENT '渠道ID',
	ysdata.FVcShanghnjlid = data.Data.Record_id    //  `F_VC_SHANGHNJLID` varchar(128) NOT NULL COMMENT '商户内记录ID',
	ysdata.FVcShanghnjlxh = data.Data.Record_no    //  `F_VC_SHANGHNJLXH` varchar(32) NOT NULL COMMENT '商户内记录的序号',
	ysdata.FVcJiamkh = data.Data.Etc_terminal_id   //  `F_VC_JIAMKH` varchar(32) NOT NULL COMMENT '加密卡号',
	ysdata.FVcKajmjyxh = data.Data.Etc_termtrad_no //  `F_VC_KAJMJYXH` varchar(32) NOT NULL COMMENT '加密卡交易序号',
	ysdata.FVcObuid = data.Data.Obu_id             //  `F_VC_OBUID` varchar(32) NOT NULL COMMENT 'Obuid',
	ysdata.FVcObufxf = data.Data.Obu_issuer        //  `F_VC_OBUFXF` varchar(32) NOT NULL COMMENT 'obu发行方',
	ysdata.FVcObucp = data.Data.Obu_plate          //  `F_VC_OBUCP` varchar(32) NOT NULL COMMENT 'obu内车牌',
	ysdata.FVcObucpys = data.Data.Obu_plate_color  //  `F_VC_OBUCPYS` varchar(32) NOT NULL COMMENT 'obu车牌颜色',

	ysdata.FDtObuyxq = utils.StrDATETimeTotime(data.Data.Obu_expire_date) //  `F_DT_OBUYXQ` datetime NOT NULL COMMENT 'obu有效期',

	ysdata.FVcKah = data.Data.Card_id          //  `F_VC_KAH` varchar(32) NOT NULL COMMENT '卡号',
	ysdata.FVcKawlh = data.Data.Card_network   //  `F_VC_KAWLH` varchar(32) NOT NULL COMMENT '卡网络号',
	ysdata.FVcKajyxh = data.Data.Card_trade_no //  `F_VC_KAJYXH` varchar(32) NOT NULL COMMENT '卡交易序号',
	ysdata.FVcKafxf = data.Data.Card_issuer    //  `F_VC_KAFXF` varchar(32) NOT NULL COMMENT '卡发行方',
	clx, _ := strconv.Atoi(data.Data.Card_type)
	ysdata.FNbKalx = clx //  `F_NB_KALX` int(11) DEFAULT NULL COMMENT '卡类型',

	ysdata.FVcKadqsj = utils.StrDATETimeTotime(data.Data.Card_expired) //  `F_VC_KADQSJ` datetime NOT NULL COMMENT '卡到期时间',

	hm, _ := strconv.Atoi(data.Data.Reset_money) //之后
	ysdata.FNbJiaoyhye = int64(hm)               // `F_NB_JIAOYHYE` bigint(20) NOT NULL COMMENT '交易后余额',

	qm, _ := strconv.Atoi(data.Data.Before_money) //之前
	ysdata.FNbJiaoyqye = int64(qm)                // //  `F_NB_JIAOYQYE` bigint(20) NOT NULL COMMENT '交易前余额',

	m, _ := strconv.Atoi(data.Data.Money)
	ysdata.FNbJine = int64(m) //  `F_NB_JINE` int(11) NOT NULL COMMENT '金额',

	ysdata.FVcTacm = data.Data.Tac                                //  `F_VC_TACM` varchar(32) NOT NULL COMMENT 'TAC码',
	ysdata.FDtJiaoysj = utils.StrTimeTotime(data.Data.Trade_time) //  `F_DT_JIAOYSJ` datetime NOT NULL COMMENT '交易时间',

	ysdata.FDtJiaoylx = data.Data.Trade_type //  `F_DT_JIAOYLX` varchar(32) NOT NULL COMMENT '交易类型',

	if data.Data.Lane_key != "" {
		ysdata.FVcChdjrsq = data.Data.Lane_key //  `F_VC_CHDJRSQ` varchar(32) NOT NULL COMMENT '车道接入授权',[???????]
	} else {
		ysdata.FVcChdjrsq = ""
	}

	ysdata.FVcChex = data.Data.Vehicle           //  `F_VC_CHEX` varchar(32) NOT NULL COMMENT '车型',
	ysdata.FVcObuzt = data.Data.Obu_status       //  `F_VC_OBUZT` varchar(32) NOT NULL COMMENT 'OBu状态',
	ysdata.FVcCheph = data.Data.Plate_num        //  `F_VC_CHEPH` varchar(32) NOT NULL COMMENT '车牌号',
	ysdata.FVcChpys = data.Data.Plate_color      //  `F_VC_CHPYS` varchar(32) NOT NULL COMMENT '车牌颜色',[????????]
	ysdata.FVcSuanfbs = data.Data.Algorithm_type //  `F_VC_SUANFBS` varchar(32) NOT NULL COMMENT '算法标识',
	ysdata.FVcHeimdjybb = data.Data.Black_ver    //  `F_VC_HEIMDJYBB` varchar(32) NOT NULL COMMENT '黑名单校验版本',
	ysdata.FVcJiztzhd = data.Data.Notify         //  `F_VC_JIZTZHD` varchar(1024) NOT NULL COMMENT '记账通知回调',

	if data.Data.Entry_time != "" {
		ysdata.FDtYonghrksj = utils.StrTimeTotime(data.Data.Entry_time) //  `F_DT_YONGHRKSJ` datetime DEFAULT NULL COMMENT '用户入口时间',

	}

	yhtcsc, _ := strconv.Atoi(data.Data.Duration)
	ysdata.FNbYonghtcsc = yhtcsc //  `F_NB_YONGHTCSC` int(11) DEFAULT NULL COMMENT '用户停车时长(分)',

	ysdata.FVcZhangdms = data.Data.Bill_description //  `F_VC_ZHANGDMS` varchar(512) NOT NULL COMMENT '账单描述（给用户通知的信息）',

	ysdata.FVcShujqm = "0" //  `F_VC_SHUJQM` varchar(32) NOT NULL COMMENT '数据签名',

	ysdata.FDtCaijsj = utils.StrTimeTotime(time.Now().Format("2006-01-02 15:04:05")) //  `F_DT_CAIJSJ` datetime NOT NULL COMMENT '采集时间',入库时间 2020-04-30 15:23:45

	ysdata.FVcJiaoyzt = 0 //  `F_VC_JIAOYZT` int(11) NOT NULL COMMENT '交易状态',
	ysdata.FVcYiclx = 0   //  `F_VC_YICLX` int(11) DEFAULT NULL COMMENT '异常类型',
	//ysdata.FVcYicyy = //  `F_VC_YICYY` varchar(32) DEFAULT NULL COMMENT '异常原因',

	ysdata.FVcHuidtzsj = utils.StrTimeTotime(time.Now().Format("2006-01-02 15:04:05")) //  `F_VC_HUIDTZSJ` datetime NOT NULL COMMENT '回调通知时间',

	ysdata.FVcHuidtzzt = 1 //  `F_VC_HUIDTZZT` int(11) NOT NULL COMMENT '回调通知状态 1:表示已通知;0:表示未通知',
	ysdata.FVcHuidtzcs = 1 //  `F_VC_HUIDTZCS` int(11) NOT NULL COMMENT '回调通知次数',

	//ysdata.FVcZuofbj = //  `F_VC_ZUOFBJ` int(11) DEFAULT NULL COMMENT '作废标记',
	//ysdata.FVcZuofsj = //  `F_VC_ZUOFSJ` datetime DEFAULT NULL COMMENT '作废时间',
	//ysdata.FVcYicblbj = //  `F_VC_YICBLBJ` int(11) DEFAULT NULL COMMENT '异常补录标记 1:表示异常补录',
	//ysdata.FVcYicblsj = //  `F_VC_YICBLSJ` datetime DEFAULT NULL COMMENT '异常补录时间',

	ysdata.FVcKabbh = data.Data.Card_version //  `F_VC_KABBH` varchar(32) DEFAULT NULL COMMENT '卡版本号',

	ysdata.FVcMiybbh = data.Data.Psamversion //  `F_VC_MIYBBH` varchar(32) DEFAULT NULL COMMENT '密钥版本号',
	ysdata.FVc_0019 = data.Data.File0019     //  `F_VC_0019` varchar(128) DEFAULT NULL COMMENT '0019文件',
	ysdata.FVc_0015 = data.Data.File0015     //  `F_VC_0015` varchar(128) DEFAULT NULL COMMENT '0015文件',

	ysdata.FVcObuxx = data.Data.Obuinfo       //  `F_VC_OBUXX` varchar(128) DEFAULT NULL COMMENT 'OBU信息',
	ysdata.FVcChelxx = data.Data.Vehicileinfo //  `F_VC_CHELXX` varchar(512) DEFAULT NULL COMMENT '车辆信息',
	ysdata.FVcKkfhxx = data.Data.Feebackinfo  //  `F_VC_KKFHXX` varchar(128) DEFAULT NULL COMMENT '扣款返回信息',
	//ysdata.FDtTongbsj = //  `F_DT_TONGBSJ` datetime DEFAULT NULL COMMENT '同步时间',
	//	ysdata.FNbTongbzt =                      //  `F_NB_TONGBZT` int(11) NOT NULL DEFAULT '0' COMMENT '同步状态 0、未同步，1、同步中，2、已同步，3、失败',
	ysdata.FVcObuyyxlh = data.Data.Obu_serial    //  `F_VC_OBUYYXLH` varchar(32) NOT NULL COMMENT 'obu应用序列号',
	ysdata.FVcChdjyxh = data.Data.Lane_record_no //  `F_VC_CHDJYXH` varchar(32) NOT NULL COMMENT '车道交易序号',
	ysdata.FNbShiftf = 0                         //  `F_NB_SHIFTF` int(11) NOT NULL DEFAULT '0' COMMENT '是否退费 0、否，1、是',
	//ysdata.FVcTuifrq      = //  `F_VC_TUIFRQ` varchar(32) DEFAULT NULL COMMENT '退费日期',

	//查询停车场费率，查询公司费率
	Rate := QueryRate(data.Data.Parking_id, data.Data.Company_id)

	if Rate != nil {
		ysdata.FNbFeil = Rate.Rate //  `F_NB_FEIL` int(11) DEFAULT NULL COMMENT '费率-NEW 万分比',
		//金额*费率/10000
		ysdata.FNbShouxf = int64(ysdata.FNbJine * int64(Rate.Rate) / 10000) //  `F_NB_SHOUXF` int(11) DEFAULT NULL COMMENT '手续费-NEW 单位分',

	} else {
		ysdata.FNbShouxf = 60
	}

	obulx, _ := strconv.Atoi(data.Data.Obu_type)
	ysdata.FNbObulx = obulx //  `F_NB_OBULX` int(11) NOT NULL DEFAULT '1' COMMENT 'OBU类型-NEW 1、单片式；2、双片式',

	mybs, _ := strconv.Atoi(data.Data.Key_type)
	ysdata.FNbMiybs = mybs                     //  `F_NB_MIYBS` int(11) NOT NULL DEFAULT '0' COMMENT '秘钥标识-NEW B5里面的密钥标识  0：3DES ;4:  SM4',
	ysdata.FVcChengxbbh = data.Data.Programver //  `F_VC_CHENGXBBH` varchar(32) DEFAULT NULL COMMENT '程序版本号-NEW',

	ysdata.FDtChengxqdsj = utils.StrTimeTotime(data.Data.Programstarttime) //  `F_DT_CHENGXQDSJ` datetime DEFAULT NULL COMMENT '程序启动时间-NEW',

	hs, _ := strconv.Atoi(data.Data.Costtime)
	ysdata.FNbKoukhs = hs //  `F_NB_KOUKHS` int(11) DEFAULT NULL COMMENT '扣款耗时-NEW 毫秒',

	ysdata.FNbFuwslrl = 0 //  `F_NB_FUWSLRL` int(11) DEFAULT NULL COMMENT '服务商利润率-NEW 万分比；停车场公司签约费率与服务商签约费率的差额；举例：千分之六-千分之四=千分之二；',
	ysdata.FNbFuwslre = 0 //  `F_NB_FUWSLRE` int(11) DEFAULT NULL COMMENT '服务商利润额-NEW单位分；停车场公司手续费与服务商手续费的差额，及本平台需要给服务商的钱；',

	//ysdata.FNbJiaoyzt       = //  `F_NB_JIAOYZT` int(11) NOT NULL DEFAULT '0' COMMENT '校验状态 0：初始:、1：已接受',

	ysdata.FNbShujhqfs = 0 //  `F_NB_SHUJHQFS` int(11) NOT NULL DEFAULT '0' COMMENT 数据获取方式 0：停车场标准方式；1：停车场刷卡方式；',
	//ysdata.FNbTiantjzt =0 //  `F_NB_TIANTJZT` int(11) DEFAULT '0' COMMENT '天统计状态 0：未统计、1：已统计',

	if data.Data.DeviceType == "31" {
		log.Println("应用场景:", ysdata.FNbYingycj)
		ysdata.FNbYingycj = 31 //`F_NB_YINGYCJ` int(11) NOT NULL DEFAULT '1' COMMENT '应用场景 1、单点停车场；31、单点加油站',
	} else {
		log.Println("应用场景不对:", data.Data.DeviceType, "已经改成31", data.Data.Bill_id)
		ysdata.FNbYingycj = 31 //`F_NB_YINGYCJ` int(11) NOT NULL DEFAULT '1' COMMENT '应用场景 1、单点停车场；31、单点加油站',
	}

	//ysdata.FNbSheblx = 0 //`F_NB_SHEBLX` int(11) NOT NULL DEFAULT '0' COMMENT '设备类型 0、标准设备；1、手持机',

	//单点车道出口原始交易数据入库
	inerr := InsertChedckyssjData(ysdata)
	if inerr != nil {
		log.Println("+++++++++单点车道出口原始交易数据入库失败++++++", inerr)
		return inerr
	}

	//加油明细赋值
	jymx := new(types.BJyzJiaymx)
	jymx.FVcJiaoyjlid = data.Data.Bill_id // `F_VC_JIAOYJLID` varchar(128) NOT NULL COMMENT '交易记录ID 停车场ID+车道ID+年月日时分秒',
	jymx.FVcYouplx = data.Data.Oil_type   //	`F_VC_YOUPLX` varchar(32) DEFAULT NULL COMMENT '油品类型 01:97号汽油;02:95号汽油;03:93号汽油;04:92号汽油;05:90号汽油;06:98 号汽油;;11:5号柴油;12:0号柴油 ;13:10号柴油;14:20号柴油;15:35号柴油;16:50号柴油',
	Unit_price, _ := strconv.Atoi(data.Data.Unit_price)
	jymx.FNbYoupdj = Unit_price //	`F_NB_YOUPDJ` int(11) DEFAULT NULL COMMENT '油品单价 单位：分',
	Litre, _ := strconv.Atoi(data.Data.Litre)
	jymx.FNbJiayl = Litre                  //	`F_NB_JIAYL` varchar(32) DEFAULT NULL COMMENT '加油量 单位：升',
	jymx.FVcYouqh = data.Data.Oilgunnumber //	`F_VC_YOUQH` varchar(32) DEFAULT NULL COMMENT '油枪号',
	Oilavailable, _ := strconv.Atoi(data.Data.Oilavailable)
	jymx.FNbKejyl = Oilavailable //	`F_NB_KEJYL` varchar(32) DEFAULT NULL COMMENT '可加油量 单位：升',

	//新增车道出口原始交易数据加油明细
	inmxerr := InsertChedckyssjJYMXData(jymx)
	if inmxerr != nil {
		log.Println("+++++++++++++++++++++++新增车道出口原始交易数据加油明细失败+++++inmxerr:", inmxerr)
		return inmxerr
	}
	return nil
}

//更新回调时间
func ChedckyssjDataUpdate(Jiaoyjlid string) error {
	db := utils.GormClient.Client
	sj := time.Now().Format("2006-01-02 15:04:05")
	if err := db.Table("b_dd_chedckyssj").Where("F_VC_JIAOYJLID= ?", Jiaoyjlid).Update("F_VC_HUIDTZSJ", sj).Error; err != nil {
		log.Println("Jiaoyjlid", Jiaoyjlid, "更新更新回调时间 error:", err)
		return err
	}
	log.Println("Jiaoyjlid", Jiaoyjlid, "更新更新回调时间  成功")
	return nil
}

//单点车道出口原始数据入库
func HourDataStorage(data *types.KafKaBillHourMsg) error {
	db := utils.GormClient.Client
	xstjdata := new(types.BDdXiaostj)
	//先查询

	if qerr := db.Table("b_dd_xiaostj").Where("F_VC_JILID = ?", data.Data.Record_id).First(xstjdata).Error; qerr != nil {
		// 错误处理...
		log.Println("单点车道出口小时统计插入 query b_dd_xiaostj error:", qerr)
		if fmt.Sprint(qerr) == "record not found" {

		} else {
			return qerr
		}
	} else {
		log.Println("单点车道出口小时统计存在:", xstjdata)
		return errors.New("单点车道出口小时统计已经存在，不可以插入")
	}
	log.Println("单点车道出口小时统计数据库不存在，可以插入:", data.Data.Record_id)

	//赋值
	xstjdata.FVcGongsjtid = data.Data.Company_id //`F_VC_GONGSJTID` varchar(32) NOT NULL COMMENT '公司/集团ID',
	xstjdata.FVcTingccbh = data.Data.Parking_id  //`F_VC_TINGCCBH` varchar(32) NOT NULL COMMENT '停车场编号',
	xstjdata.FVcChedid = data.Data.Lane_id       //`F_VC_CHEDID` varchar(32) NOT NULL COMMENT '车道ID',
	xstjdata.FVcJilid = data.Data.Record_id      //`F_VC_JILID` varchar(128) DEFAULT NULL COMMENT '记录ID',
	//sj := utils.StrDATETimeToHourtime(data.Data.Recordcnt)
	xstjdata.FVcTongjxs = data.Data.Datetime_hour //`F_VC_TONGJXS` varchar(32) NOT NULL COMMENT '统计小时 yyMMddhh',

	c, _ := strconv.Atoi(data.Data.Recordcnt)
	xstjdata.FNbJilzs = c //`F_NB_JILZS` int(11) DEFAULT NULL COMMENT '记录总数',
	m, _ := strconv.Atoi(data.Data.Moneycnt)
	xstjdata.FNbJinezs = m //`F_NB_JINEZS` int(11) DEFAULT NULL COMMENT '金额总数',

	xstjdata.FDtShangcsj = utils.StrTimeTotime(time.Now().Format("2006-01-02 15:04:05")) //`F_DT_SHANGCSJ` datetime DEFAULT NULL COMMENT '上传时间',

	xstjdata.FNbChedlx = 2 //`F_NB_CHEDLX` int(11) NOT NULL DEFAULT '1' COMMENT '车道类型 1：入口；2：出口；',

	//插入
	if err := db.Table("b_dd_xiaostj").Create(xstjdata).Error; err != nil {
		// 错误处理...
		log.Println("单点车道出口小时统计插入 Insert b_dd_xiaostj error:", err)
		return err
	}

	log.Println("单点车道出口小时统计插入成功！")

	return nil
}
