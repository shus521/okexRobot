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
func GetCoinInfo(coin string, genre string) {
	url := "/api/swap/v3/instruments/" + coin + "-USDT-" + genre + "/ticker"
	result := util.SendGet(url, "public")
	fmt.Printf("[okex]%s上次交易价格: %s\n", coin, result["last"])

}

/**
 * 获取市场的总持仓量
 * @param url ok平台地址
 * @param coin 要获取的币种
 * @param genre 查看类型 SWAP永续合约 futures交割合约 margin币币杠杆 spot币币 account资金账户
 */
func GetAllHolding(coin string, genre string) {
	url := "/api/swap/v3/instruments/" + coin + "-USDT-" + genre + "/open_interest"
	result := util.SendGet(url, "public")
	fmt.Printf("[okex]%s平台总持仓量: %s\n", coin, result["amount"])
}

/**
 * 获取我的永续合约某个币种的持仓量
 * @param url ok平台地址
 * @param coin 要获取的币种
 */
func GetMyHolding(coin string) {
	url := "/api/swap/v3/" + coin + "-USDT-SWAP/position"
	result := util.SendGet(url, "private")
	//json.Unmarshal([]byte(result["holding"].(string)), data)
	fmt.Println(result["holding"])
}
