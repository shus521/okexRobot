package service

/**
 * @author ty
 * @desc 获取各种数据
 * @version 1.0
 */

import (
	"fmt"
	util "okexRobot/utils"
)

//获取币对信息
func GetCoinInfo(url string, coin string) {
	url = url + "/api/spot/v3/instruments/" + coin + "/ticker"
	result := util.SendGet(url)
	fmt.Println(result)

}
