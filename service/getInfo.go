package service

/**
 * @author ty
 * @desc 获取各种数据
 * @version 1.0
 */

import (
	"fmt"
	util "okexRobot/utils"
	"strconv"
)

/**
 * 获取永续合约币对信息
 * @param url ok平台地址
 * @param coin 要获取的币种
 * @param genre 查看类型 SWAP永续合约 futures交割合约 margin币币杠杆 spot币币 account资金账户
 */
func GetCoinInfo(coin string, genre string) float64 {
	url := "/api/swap/v3/instruments/" + coin + "-USDT-" + genre + "/ticker"
	result := util.SendGet(url, "public")
	fmt.Printf("[okex-USDT永续合约]%s上次交易价格: %s\n", coin, result["last"])
	last, _ := strconv.ParseFloat(result["last"].(string), 64)
	return last

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
	fmt.Printf("[okex-USDT永续合约]%s平台总持仓量: %s\n", coin, result["amount"])
}

/**
 * 获取我的USDT永续合约某个币种的持仓量
 * @param url ok平台地址
 * @param coin 要获取的币种
 */
func GetMyHolding(coin string, last float64) {
	url := "/api/swap/v3/" + coin + "-USDT-SWAP/position"
	result := util.SendGet(url, "private")
	if result["holding"] == nil {
		fmt.Println("获取异常")
	} else {
		holding := result["holding"].([]interface{})[0].(map[string]interface{})
		avg, _ := strconv.ParseFloat(holding["avg_cost"].(string), 64)
		num, _ := strconv.ParseFloat(holding["position"].(string), 64)
		pnl := (avg - last) * 0.0001 * num
		fmt.Printf("[okex-USDT永续合约]开单方向:%s;持仓均价:%s;预估强平价:%s;盈亏:%f\n", holding["side"],
			holding["avg_cost"], holding["liquidation_price"], pnl)
	}
}

/**
 * 获取某币种的永续合约列表
 * @param coin 币种
 * @param state 订单状态 -2:失败 -1:撤单成功 0:等待成交 1:部分成交 2:完全成交 3:下单中
				4:撤单中 6: 未完成（等待成交+部分成交） 7:已完成（撤单成功+完全成交）
*/
func GetOrderList(coin string, state string) map[string]interface{} {
	url := "/api/swap/v3/orders/" + coin + "-USDT-SWAP?state=" + state
	result := util.SendGet(url, "private")
	return result
}
