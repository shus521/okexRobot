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

/**
 * 获取永续合约币对信息
 * @param url ok平台地址
 * @param coin 要获取的币种
 * @param genre 查看类型 SWAP永续合约 futures交割合约 margin币币杠杆 spot币币 account资金账户
 */
func GetCoinInfo(url string, coin string, genre string) {
	url = url + "/api/swap/v3/instruments/" + coin + "-USD-" + genre + "/ticker"
	result := util.SendGet(url)
	fmt.Printf("[okex]%s上次交易价格: %s\n", coin, result["last"])

}

/**
 * 获取市场的总持仓量
 * @param url ok平台地址
 * @param coin 要获取的币种
 * @param genre 查看类型 SWAP永续合约 futures交割合约 margin币币杠杆 spot币币 account资金账户
 */
func GetAllHolding(url string, coin string, genre string) {
	url = url + "/api/swap/v3/instruments/" + coin + "-USD-" + genre + "/open_interest"
	result := util.SendGet(url)
	fmt.Printf("[okex]%s平台总持仓量: %s\n", coin, result["amount"])
}
