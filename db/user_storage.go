package db

import (
	"etcGasDataCenter/types"
	"etcGasDataCenter/utils"
	"github.com/sirupsen/logrus"
)

func QueryUserLoginmsg(username string) (error, *types.BSysYongh) {
	db := utils.GormClient.Client
	user := new(types.BSysYongh)
	if err := db.Table("b_sys_yongh").Where("F_VC_ZHANGH = ?", username).First(user).Error; err != nil {
		logrus.Println("查询用户登录信息失败！")
		return err, nil
	}
	logrus.Println("查询用户登录信息 ok:", user.FVcMingc, user.FVcZhangh)

	return nil, user

}
