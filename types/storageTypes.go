package types

import (
	"time"
)

var KafkaIp string

var DdkafkaTopic string
var DdkafkaHourTopic string

//1、单点车道出口原始数据表
//CREATE TABLE `b_dd_chedckyssj` 单点车道出口原始数据-U (
type BDdChedckyssj struct {
	FVcJiaoyjlid   string    `gorm:"column:F_VC_JIAOYJLID"`                  //  `F_VC_JIAOYJLID` varchar(128) NOT NULL COMMENT '交易记录ID 停车场ID+车道ID+年月日时分秒',
	FVcJiaoytjr    string    `gorm:"column:F_VC_JIAOYTJR; default:'NULL'"`   //  `F_VC_JIAOYTJR` varchar(32) DEFAULT NULL COMMENT '交易统计日',
	FVcJiaoytjrs   string    `gorm:"column:F_VC_JIAOYTJRS; default:'NULL'"`  //  `F_VC_JIAOYTJRS` varchar(32) DEFAULT NULL COMMENT '交易统计时',
	FVcTingccbh    string    `gorm:"column:F_VC_TINGCCBH"`                   //  `F_VC_TINGCCBH` varchar(32) NOT NULL COMMENT '停车场编号',
	FVcChedid      string    `gorm:"column:F_VC_CHEDID"`                     //  `F_VC_CHEDID` varchar(32) NOT NULL COMMENT '车道ID',
	FVcGongsjtid   string    `gorm:"column:F_VC_GONGSJTID"`                  //  `F_VC_GONGSJTID` varchar(32) NOT NULL COMMENT '公司/集团ID',
	FVcQudid       string    `gorm:"column:F_VC_QUDID; default:'NULL'"`      //  `F_VC_QUDID` varchar(32) DEFAULT NULL COMMENT '渠道ID',
	FVcShanghnjlid string    `gorm:"column:F_VC_SHANGHNJLID"`                //  `F_VC_SHANGHNJLID` varchar(128) NOT NULL COMMENT '商户内记录ID',
	FVcShanghnjlxh string    `gorm:"column:F_VC_SHANGHNJLXH"`                //  `F_VC_SHANGHNJLXH` varchar(32) NOT NULL COMMENT '商户内记录的序号',
	FVcJiamkh      string    `gorm:"column:F_VC_JIAMKH"`                     //  `F_VC_JIAMKH` varchar(32) NOT NULL COMMENT '加密卡号',
	FVcKajmjyxh    string    `gorm:"column:F_VC_KAJMJYXH"`                   //  `F_VC_KAJMJYXH` varchar(32) NOT NULL COMMENT '加密卡交易序号',
	FVcObuid       string    `gorm:"column:F_VC_OBUID"`                      //  `F_VC_OBUID` varchar(32) NOT NULL COMMENT 'Obuid',
	FVcObufxf      string    `gorm:"column:F_VC_OBUFXF"`                     //  `F_VC_OBUFXF` varchar(32) NOT NULL COMMENT 'obu发行方',
	FVcObucp       string    `gorm:"column:F_VC_OBUCP"`                      //  `F_VC_OBUCP` varchar(32) NOT NULL COMMENT 'obu内车牌',
	FVcObucpys     string    `gorm:"column:F_VC_OBUCPYS"`                    //  `F_VC_OBUCPYS` varchar(32) NOT NULL COMMENT 'obu车牌颜色',
	FDtObuyxq      time.Time `gorm:"column:F_DT_OBUYXQ"`                     //  `F_DT_OBUYXQ` datetime NOT NULL COMMENT 'obu有效期',
	FVcKah         string    `gorm:"column:F_VC_KAH"`                        //  `F_VC_KAH` varchar(32) NOT NULL COMMENT '卡号',
	FVcKawlh       string    `gorm:"column:F_VC_KAWLH"`                      //  `F_VC_KAWLH` varchar(32) NOT NULL COMMENT '卡网络号',
	FVcKajyxh      string    `gorm:"column:F_VC_KAJYXH"`                     //  `F_VC_KAJYXH` varchar(32) NOT NULL COMMENT '卡交易序号',
	FVcKafxf       string    `gorm:"column:F_VC_KAFXF"`                      //  `F_VC_KAFXF` varchar(32) NOT NULL COMMENT '卡发行方',
	FNbKalx        int       `gorm:"column:F_NB_KALX; default:'NULL'"`       //  `F_NB_KALX` int(11) DEFAULT NULL COMMENT '卡类型',
	FVcKadqsj      time.Time `gorm:"column:F_VC_KADQSJ"`                     //  `F_VC_KADQSJ` datetime NOT NULL COMMENT '卡到期时间',
	FNbJiaoyqye    int64     `gorm:"column:F_NB_JIAOYQYE"`                   //  `F_NB_JIAOYQYE` bigint(20) NOT NULL COMMENT '交易前余额',
	FNbJiaoyhye    int64     `gorm:"column:F_NB_JIAOYHYE"`                   //  `F_NB_JIAOYHYE` bigint(20) NOT NULL COMMENT '交易后余额',
	FNbJine        int64     `gorm:"column:F_NB_JINE"`                       //  `F_NB_JINE` int(11) NOT NULL COMMENT '金额',
	FVcTacm        string    `gorm:"column:F_VC_TACM"`                       //  `F_VC_TACM` varchar(32) NOT NULL COMMENT 'TAC码',
	FDtJiaoysj     time.Time `gorm:"column:F_DT_JIAOYSJ"`                    //  `F_DT_JIAOYSJ` datetime NOT NULL COMMENT '交易时间',
	FDtJiaoylx     string    `gorm:"column:F_DT_JIAOYLX"`                    //  `F_DT_JIAOYLX` varchar(32) NOT NULL COMMENT '交易类型',
	FVcChdjrsq     string    `gorm:"column:F_VC_CHDJRSQ"`                    //  `F_VC_CHDJRSQ` varchar(32) NOT NULL COMMENT '车道接入授权',[???????]
	FVcChex        string    `gorm:"column:F_VC_CHEX"`                       //  `F_VC_CHEX` varchar(32) NOT NULL COMMENT '车型',
	FVcObuzt       string    `gorm:"column:F_VC_OBUZT"`                      //  `F_VC_OBUZT` varchar(32) NOT NULL COMMENT 'OBu状态',
	FVcCheph       string    `gorm:"column:F_VC_CHEPH"`                      //  `F_VC_CHEPH` varchar(32) NOT NULL COMMENT '车牌号',
	FVcChpys       string    `gorm:"column:F_VC_CHPYS"`                      //  `F_VC_CHPYS` varchar(32) NOT NULL COMMENT '车牌颜色',[????????]
	FVcSuanfbs     string    `gorm:"column:F_VC_SUANFBS"`                    //  `F_VC_SUANFBS` varchar(32) NOT NULL COMMENT '算法标识',
	FVcHeimdjybb   string    `gorm:"column:F_VC_HEIMDJYBB"`                  //  `F_VC_HEIMDJYBB` varchar(32) NOT NULL COMMENT '黑名单校验版本',
	FVcJiztzhd     string    `gorm:"column:F_VC_JIZTZHD"`                    //  `F_VC_JIZTZHD` varchar(1024) NOT NULL COMMENT '记账通知回调',
	FDtYonghrksj   time.Time `gorm:"column:F_DT_YONGHRKSJ; default:'NULL'"`  //  `F_DT_YONGHRKSJ` datetime DEFAULT NULL COMMENT '用户入口时间',
	FNbYonghtcsc   int       `gorm:"column:F_NB_YONGHTCSC"`                  //  `F_NB_YONGHTCSC` int(11) DEFAULT NULL COMMENT '用户停车时长(分)',
	FVcZhangdms    string    `gorm:"column:F_VC_ZHANGDMS"`                   //  `F_VC_ZHANGDMS` varchar(512) NOT NULL COMMENT '账单描述（给用户通知的信息）',
	FVcShujqm      string    `gorm:"column:F_VC_SHUJQM"`                     //  `F_VC_SHUJQM` varchar(32) NOT NULL COMMENT '数据签名',
	FDtCaijsj      time.Time `gorm:"column:F_DT_CAIJSJ"`                     //  `F_DT_CAIJSJ` datetime NOT NULL COMMENT '采集时间',
	FVcJiaoyzt     int       `gorm:"column:F_VC_JIAOYZT"`                    //  `F_VC_JIAOYZT` int(11) NOT NULL COMMENT '交易状态',
	FVcYiclx       int       `gorm:"column:F_VC_YICLX; default:'NULL'"`      //  `F_VC_YICLX` int(11) DEFAULT NULL COMMENT '异常类型',
	FVcYicyy       string    `gorm:"column:F_VC_YICYY; default:'NULL'"`      //  `F_VC_YICYY` varchar(32) DEFAULT NULL COMMENT '异常原因',
	FVcHuidtzsj    time.Time `gorm:"column:F_VC_HUIDTZSJ"`                   //  `F_VC_HUIDTZSJ` datetime NOT NULL COMMENT '回调通知时间',
	FVcHuidtzzt    int       `gorm:"column:F_VC_HUIDTZZT"`                   //  `F_VC_HUIDTZZT` int(11) NOT NULL COMMENT '回调通知状态 1:表示已通知;0:表示未通知',
	FVcHuidtzcs    int       `gorm:"column:F_VC_HUIDTZCS"`                   //  `F_VC_HUIDTZCS` int(11) NOT NULL COMMENT '回调通知次数',
	FVcZuofbj      int       `gorm:"column:F_VC_ZUOFBJ; default:'NULL'"`     //  `F_VC_ZUOFBJ` int(11) DEFAULT NULL COMMENT '作废标记',
	FVcZuofsj      time.Time `gorm:"column:F_VC_ZUOFSJ; default:'NULL'"`     //  `F_VC_ZUOFSJ` datetime DEFAULT NULL COMMENT '作废时间',
	FVcYicblbj     int       `gorm:"column:F_VC_YICBLBJ; default:'NULL'"`    //  `F_VC_YICBLBJ` int(11) DEFAULT NULL COMMENT '异常补录标记 1:表示异常补录',
	FVcYicblsj     time.Time `gorm:"column:F_VC_YICBLSJ; default:'NULL'"`    //  `F_VC_YICBLSJ` datetime DEFAULT NULL COMMENT '异常补录时间',
	FVcKabbh       string    `gorm:"column:F_VC_KABBH; default:'NULL'"`      //  `F_VC_KABBH` varchar(32) DEFAULT NULL COMMENT '卡版本号',
	FVcMiybbh      string    `gorm:"column:F_VC_MIYBBH; default:'NULL'"`     //  `F_VC_MIYBBH` varchar(32) DEFAULT NULL COMMENT '密钥版本号',
	FVc_0019       string    `gorm:"column:F_VC_0019; default:'NULL'"`       //  `F_VC_0019` varchar(128) DEFAULT NULL COMMENT '0019文件',
	FVc_0015       string    `gorm:"column:F_VC_0015; default:'NULL'"`       //  `F_VC_0015` varchar(128) DEFAULT NULL COMMENT '0015文件',
	FVcObuxx       string    `gorm:"column:F_VC_OBUXX; default:'NULL'"`      //  `F_VC_OBUXX` varchar(128) DEFAULT NULL COMMENT 'OBU信息',
	FVcChelxx      string    `gorm:"column:F_VC_CHELXX; default:'NULL'"`     //  `F_VC_CHELXX` varchar(512) DEFAULT NULL COMMENT '车辆信息',
	FVcKkfhxx      string    `gorm:"column:F_VC_KKFHXX; default:'NULL'"`     //  `F_VC_KKFHXX` varchar(128) DEFAULT NULL COMMENT '扣款返回信息',
	FDtTongbsj     time.Time `gorm:"column:F_DT_TONGBSJ; default:'NULL'"`    //  `F_DT_TONGBSJ` datetime DEFAULT NULL COMMENT '同步时间',
	FNbTongbzt     int       `gorm:"column:F_NB_TONGBZT; default:'0'"`       //  `F_NB_TONGBZT` int(11) NOT NULL DEFAULT '0' COMMENT '同步状态 0、未同步，1、同步中，2、已同步，3、失败',
	FVcObuyyxlh    string    `gorm:"column:F_VC_OBUYYXLH"`                   //  `F_VC_OBUYYXLH` varchar(32) NOT NULL COMMENT 'obu应用序列号',
	FVcChdjyxh     string    `gorm:"column:F_VC_CHDJYXH"`                    //  `F_VC_CHDJYXH` varchar(32) NOT NULL COMMENT '车道交易序号',
	FNbShiftf      int       `gorm:"column:F_NB_SHIFTF"`                     //  `F_NB_SHIFTF` int(11) NOT NULL DEFAULT '0' COMMENT '是否退费 0、否，1、是',
	FVcTuifrq      string    `gorm:"column:F_VC_TUIFRQ; default:'NULL'"`     //  `F_VC_TUIFRQ` varchar(32) DEFAULT NULL COMMENT '退费日期',
	FNbFeil        int       `gorm:"column:F_NB_FEIL;"`                      //  `F_NB_FEIL` int(11) DEFAULT NULL COMMENT '费率-NEW 万分比',
	FNbShouxf      int64     `gorm:"column:F_NB_SHOUXF;"`                    //  `F_NB_SHOUXF` int(11) DEFAULT NULL COMMENT '手续费-NEW 单位分',
	FNbObulx       int       `gorm:"column:F_NB_OBULX"`                      //  `F_NB_OBULX` int(11) NOT NULL DEFAULT '1' COMMENT 'OBU类型-NEW 1、单片式；2、双片式',
	FNbMiybs       int       `gorm:"column:F_NB_MIYBS"`                      //  `F_NB_MIYBS` int(11) NOT NULL DEFAULT '0' COMMENT '秘钥标识-NEW B5里面的密钥标识0：3DES ;4:  SM4',
	FVcChengxbbh   string    `gorm:"column:F_VC_CHENGXBBH; default:'NULL'"`  //  `F_VC_CHENGXBBH` varchar(32) DEFAULT NULL COMMENT '程序版本号-NEW',
	FDtChengxqdsj  time.Time `gorm:"column:F_DT_CHENGXQDSJ; default:'NULL'"` //  `F_DT_CHENGXQDSJ` datetime DEFAULT NULL COMMENT '程序启动时间-NEW',
	FNbKoukhs      int       `gorm:"column:F_NB_KOUKHS"`                     //  `F_NB_KOUKHS` int(11) DEFAULT NULL COMMENT '扣款耗时-NEW 毫秒',
	FNbFuwslrl     int       `gorm:"column:F_NB_FUWSLRL"`                    //  `F_NB_FUWSLRL` int(11) DEFAULT NULL COMMENT '服务商利润率-NEW 万分比；停车场公司签约费率与服务商签约费率的差额；举例：千分之六-千分之四=千分之二；',
	FNbFuwslre     int       `gorm:"column:F_NB_FUWSLRE"`                    //  `F_NB_FUWSLRE` int(11) DEFAULT NULL COMMENT '服务商利润额-NEW单位分；停车场公司手续费与服务商手续费的差额，及本平台需要给服务商的钱；',
	FNbJiaoyzt     int       `gorm:"column:F_NB_JIAOYZT"`                    //  `F_NB_JIAOYZT` int(11) NOT NULL DEFAULT '0' COMMENT '校验状态 0：初始:、1：已接受',
	FNbShujhqfs    int       `gorm:"column:F_NB_SHUJHQFS"`                   //  `F_NB_SHUJHQFS` int(11) NOT NULL DEFAULT '0' COMMENT 数据获取方式 0：停车场标准方式；1：停车场刷卡方式；',
	FNbTiantjzt    int       `gorm:"column:F_NB_TIANTJZT"`                   //  `F_NB_TIANTJZT` int(11) DEFAULT '0' COMMENT '天统计状态 0：未统计、1：已统计',
	FNbYingycj     int       `gorm:"column:F_NB_YINGYCJ"`                    //`F_NB_YINGYCJ` int(11) NOT NULL DEFAULT '1' COMMENT '应用场景 1、单点停车场；31、单点加油站',
	FNbSheblx      int       `gorm:"column:F_NB_SHEBLX"`                     //`F_NB_SHEBLX` int(11) NOT NULL DEFAULT '0' COMMENT '设备类型 0、标准设备；1、手持机',
	//  PRIMARY KEY (`F_VC_JIAOYJLID`),
	//  KEY `IDX_TONGJRQ` (`F_VC_JIAOYTJR`),
	//  KEY `IDX_TONGJXS` (`F_VC_JIAOYTJRS`),
	//  KEY `IDX_TONGBZT` (`F_NB_TONGBZT`),
	//  KEY `IDX_JIAOYZT` (`F_NB_JIAOYZT`),
	//  KEY `IDX_WANGLH` (`F_VC_KAWLH`)
	//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='车道出口原始数据-U';
}

//CREATE TABLE `b_jyz_jiaymx` (
type BJyzJiaymx struct {
	FVcJiaoyjlid string `gorm:"column:F_VC_JIAOYJLID"` // `F_VC_JIAOYJLID` varchar(128) NOT NULL COMMENT '交易记录ID 停车场ID+车道ID+年月日时分秒',
	FVcYouplx    string `gorm:"column:F_VC_YOUPLX"`    //	`F_VC_YOUPLX` varchar(32) DEFAULT NULL COMMENT '油品类型 01:97号汽油;02:95号汽油;03:93号汽油;04:92号汽油;05:90号汽油;06:98 号汽油;;11:5号柴油;12:0号柴油 ;13:10号柴油;14:20号柴油;15:35号柴油;16:50号柴油',
	FNbYoupdj    int    `gorm:"column:F_NB_YOUPDJ"`    //	`F_NB_YOUPDJ` int(11) DEFAULT NULL COMMENT '油品单价 单位：分',
	FNbJiayl     string `gorm:"column:F_NB_JIAYL"`     //	`F_NB_JIAYL` varchar(32) DEFAULT NULL COMMENT '加油量 单位：升',
	FVcYouqh     string `gorm:"column:F_VC_YOUQH"`     //	`F_VC_YOUQH` varchar(32) DEFAULT NULL COMMENT '油枪号',
	FNbKejyl     string `gorm:"column:F_NB_KEJYL"`     //	`F_NB_KEJYL` varchar(32) DEFAULT NULL COMMENT '可加油量 单位：升',
	//	PRIMARY KEY (`F_VC_JIAOYJLID`)
	//) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COMMENT = '加油明细'
}

//4 用户表
//CREATE TABLE `b_sys_yongh` (
type BSysYongh struct {
	FVcId          string    `gorm:"column:F_VC_ID"`         //`F_VC_ID` varchar(32) NOT NULL COMMENT 'ID 由系统生成的唯一标识；',
	FVcZhangh      string    `gorm:"column:F_VC_ZHANGH"`     //`F_VC_ZHANGH` varchar(32) NOT NULL COMMENT '账号 idx-',
	FVcMingc       string    `gorm:"column:F_VC_MINGC"`      //`F_VC_MINGC` varchar(32) NOT NULL COMMENT '姓名',
	FVcNic         string    `gorm:"column:F_VC_NIC"`        //`F_VC_NIC` varchar(32) DEFAULT NULL COMMENT '昵称',
	FNbNiann       int       `gorm:"column:F_NB_NIANN"`      //`F_NB_NIANN` int(11) DEFAULT NULL COMMENT '年龄',
	FNbXingb       int       `gorm:"column:F_NB_XINGB"`      //`F_NB_XINGB` int(11) DEFAULT NULL COMMENT '性别 0：男性；1：女性；',
	FVcDianh       string    `gorm:"column:F_VC_DIANH"`      //`F_VC_DIANH` varchar(32) NOT NULL COMMENT '电话',
	FVcYouj        string    `gorm:"column:F_VC_YOUJ"`       //`F_VC_YOUJ` varchar(32) DEFAULT NULL COMMENT '邮件',
	FVcToux        string    `gorm:"column:F_VC_TOUX"`       //`F_VC_TOUX` varchar(1024) DEFAULT NULL COMMENT '头像',
	FVcMim         string    `gorm:"column:F_VC_MIM"`        //`F_VC_MIM` varchar(32) NOT NULL COMMENT '密码 md5加密',
	FVcGongsid     string    `gorm:"column:F_VC_GONGSID"`    //`F_VC_GONGSID` varchar(32) DEFAULT NULL COMMENT '公司ID',
	FVcZuzid       string    `gorm:"column:F_VC_ZUZID"`      //`F_VC_ZUZID` varchar(32) DEFAULT 'root' COMMENT '组织ID 为root则表示处于公司根组织',
	FNbZhuangt     int       `gorm:"column:F_NB_ZHUANGT"`    //`F_NB_ZHUANGT` int(11) DEFAULT '1' COMMENT '状态 1：正常；2：已禁用；',
	FNbShenfzyyzzt int       `gorm:"column:F_NB_SHENFZYZZT"` //`F_NB_SHENFZYZZT` int(11) DEFAULT '0' COMMENT '身份证验证状态 0：待提交；1：待审核；2：审核通过；3：审核驳回，需修改信息；4：审核拒绝；',
	FVcShenfzhm    string    `gorm:"column:F_VC_SHENFZHM"`   //`F_VC_SHENFZHM` varchar(32) DEFAULT NULL COMMENT '身份证号码',
	FVcShenfzzp    string    `gorm:"column:F_VC_SHENFZZP"`   //`F_VC_SHENFZZP` varchar(1024) DEFAULT NULL COMMENT '身份证照片',
	FDtChuangjsj   time.Time `gorm:"column:F_DT_CHUANGJSJ"`  //`F_DT_CHUANGJSJ` datetime NOT NULL COMMENT '创建时间',
	FDtDenglsj     time.Time `gorm:"column:F_DT_DENGLSJ"`    //`F_DT_DENGLSJ` datetime DEFAULT NULL COMMENT '登录时间',
	FVcDenglip     string    `gorm:"column:F_VC_DENGLIP"`    //`F_VC_DENGLIP` varchar(32) DEFAULT NULL COMMENT '登录IP',
	FDtGengxsj     time.Time `gorm:"column:F_DT_GENGXSJ"`    //`F_DT_GENGXSJ` datetime DEFAULT NULL COMMENT '更新时间',
	FVcWeixid      string    `gorm:"column:F_VC_WEIXID"`     //`F_VC_WEIXID` varchar(32) DEFAULT NULL COMMENT '微信ID',
	FVcQqid        string    `gorm:"column:F_VC_QQID"`       //`F_VC_QQID` varchar(32) DEFAULT NULL COMMENT 'QQ ID',
	FNbGuanlylx    int       `gorm:"column:F_NB_GUANLYLX"`   //`F_NB_GUANLYLX` int(11) DEFAULT '0' COMMENT ' 0：普通用户；1：系统超级管理员；10：单点公司管理员；11：单点停车场管理员；20：总对总公司管理员；21: 总对总停车场管理员31：服务商管理员；',
	FNbYindzt      int       `gorm:"column:F_NB_YINDZT"`     //`F_NB_YINDZT` int(11) DEFAULT '0' COMMENT '引导状态 0：未引导； 1：已引导',
}

//3停车场表
type BTccTingcc struct {
	FVcTingccbh     string    `gorm:"column:F_VC_TINGCCBH"`      //`F_VC_TINGCCBH` varchar(32) NOT NULL COMMENT '停车场编号',
	FVcGongsbh      string    `gorm:"column:F_VC_GONGSBH"`       //`F_VC_GONGSBH` varchar(32) DEFAULT NULL COMMENT '公司/集团编号 idx-',
	FVcQudbh        string    `gorm:"column:F_VC_QUDBH"`         //`F_VC_QUDBH` varchar(32) DEFAULT NULL COMMENT '渠道编号',
	FVcTingccwlbh   string    `gorm:"column:F_VC_TINGCCWLBH"`    //`F_VC_TINGCCWLBH` varchar(32) DEFAULT NULL COMMENT '停车场网络编号 由于前期要与旧平台同步，改字段请用数字表示',
	FNbTingcclx     int64     `gorm:"column:F_NB_TINGCCLX"`      //`F_NB_TINGCCLX` int(11) NOT NULL DEFAULT '1' COMMENT '停车场类型 1：单点；2：总对总；',
	FVcMingc        string    `gorm:"column:F_VC_MINGC"`         //`F_VC_MINGC` varchar(32) DEFAULT NULL COMMENT '名称-NEW',
	FVcDiz          string    `gorm:"column:F_VC_DIZ"`           //`F_VC_DIZ` varchar(512) DEFAULT NULL COMMENT '地址',
	FVcJingd        string    `gorm:"column:F_VC_JINGD"`         //`F_VC_JINGD` decimal(32,10) DEFAULT NULL COMMENT '经度',
	FVcWeid         string    `gorm:"column:F_VC_WEID"`          //`F_VC_WEID` decimal(32,10) DEFAULT NULL COMMENT '维度',
	FVcGuanlyid     string    `gorm:"column:F_VC_GUANLYID"`      //`F_VC_GUANLYID` varchar(32) NOT NULL COMMENT '管理员ID-NEW',
	FDtChuangjsj    time.Time `gorm:"column:F_DT_CHUANGJSJ"`     //`F_DT_CHUANGJSJ` datetime DEFAULT NULL COMMENT '创建时间',
	FVcChuangjr     string    `gorm:"column:F_VC_CHUANGJR"`      //`F_VC_CHUANGJR` varchar(32) DEFAULT NULL COMMENT '创建人',
	FNbZhuangt      int       `gorm:"column:F_NB_ZHUANGT"`       //`F_NB_ZHUANGT` int(11) DEFAULT '1' COMMENT '状态-U 1：正常；2：待审核；3：停用；',
	FVcVerifyStatus int       `gorm:"column:F_VC_VERIFY_STATUS"` //`F_VC_VERIFY_STATUS` int(11) DEFAULT NULL COMMENT '审核结果-NEW 1：审核通过；2：待审核；3：审核驳回，需修改信息；4：审核拒绝；',
	FVcFuzrdh       string    `gorm:"column:F_VC_FUZRDH"`        //`F_VC_FUZRDH` varchar(32) DEFAULT NULL COMMENT '负责人电话-D',
	FVcFuzrxm       string    `gorm:"column:F_VC_FUZRXM"`        //`F_VC_FUZRXM` varchar(32) DEFAULT NULL COMMENT '负责人姓名-D',
	FVcShengdm      string    `gorm:"column:F_VC_SHENGDM"`       //`F_VC_SHENGDM` varchar(32) DEFAULT NULL COMMENT '省代码',
	FVcShengmc      string    `gorm:"column:F_VC_SHENGMC"`       //`F_VC_SHENGMC` varchar(32) DEFAULT NULL COMMENT '省名称',
	FVcShidm        string    `gorm:"column:F_VC_SHIDM"`         //`F_VC_SHIDM` varchar(32) DEFAULT NULL COMMENT '市代码',
	FVcShimc        string    `gorm:"column:F_VC_SHIMC"`         //`F_VC_SHIMC` varchar(32) DEFAULT NULL COMMENT '市名称',
	FVcQudm         string    `gorm:"column:F_VC_QUDM"`          //`F_VC_QUDM` varchar(32) DEFAULT NULL COMMENT '区代码',
	FVcQumc         string    `gorm:"column:F_VC_QUMC"`          //`F_VC_QUMC` varchar(32) DEFAULT NULL COMMENT '区名称',
	FNbFeil         int       `gorm:"column:F_NB_FEIL"`          //`F_NB_FEIL` int(11) DEFAULT NULL COMMENT '费率 万分比',
}

//CREATE TABLE `b_dd_xiaostj` (
type BDdXiaostj struct {
	FVcGongsjtid string    //`F_VC_GONGSJTID` varchar(32) NOT NULL COMMENT '公司/集团ID',
	FVcTingccbh  string    //`F_VC_TINGCCBH` varchar(32) NOT NULL COMMENT '停车场编号',
	FVcChedid    string    //`F_VC_CHEDID` varchar(32) NOT NULL COMMENT '车道ID',
	FVcJilid     string    //`F_VC_JILID` varchar(128) DEFAULT NULL COMMENT '记录ID',
	FVcTongjxs   string    //`F_VC_TONGJXS` varchar(32) NOT NULL COMMENT '统计小时 yyMMddhh',
	FNbJilzs     int       //`F_NB_JILZS` int(11) DEFAULT NULL COMMENT '记录总数',
	FNbJinezs    int       //`F_NB_JINEZS` int(11) DEFAULT NULL COMMENT '金额总数',
	FDtShangcsj  time.Time //`F_DT_SHANGCSJ` datetime DEFAULT NULL COMMENT '上传时间',
	FNbChedlx    int       //`F_NB_CHEDLX` int(11) NOT NULL DEFAULT '1' COMMENT '车道类型 1：入口；2：出口；',
	//PRIMARY KEY (`F_VC_GONGSJTID`,`F_VC_TINGCCBH`,`F_VC_CHEDID`,`F_VC_TONGJXS`),
	//KEY `IDX_CHEDLX` (`F_NB_CHEDLX`),
	//KEY `IDX_TINGCC` (`F_VC_TINGCCBH`),
	//KEY `IDX_GONGS` (`F_VC_GONGSJTID`),
	//KEY `IDX_CHED` (`F_VC_CHEDID`)
	//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='单点汇总（按小时） 由车道传上来的数据记录';
}
