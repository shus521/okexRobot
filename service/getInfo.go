package service

/**
 * @author ty
 * @desc 获取各种数据
 * @version 1.0
 */

import (
	"encoding/json"
	"fmt"
	util "okexRobot/utils"
)

/**获取永续合约币对信息
 * @param url ok平台地址
 * @param coin 要获取的币种
 * @param genre 查看类型 SWAP永续合约 futures交割合约 margin币币杠杆 spot币币 account资金账户
 */
func GetCoinInfo(url string, coin string, genre string) {
	url = url + "/api/swap/v3/instruments/" + coin + "-USD-" + genre + "/ticker"
	fmt.Println(url)
	result := util.SendGet(url)
	data := map[string]interface{}{}
	json.Unmarshal(result, &data)
	fmt.Println(data["last"])

}
