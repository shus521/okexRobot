package service

import (
	"fmt"
	"net/http"
	util "okexRobot/utils"
)

/**
 * 永续合约下单
 * @param price 下单价格
 * @param num 下单数量(张)
 * @param uid 订单id
 * @param coin 币种
 * @param direction 交易方向 1开多 2开空 3平多 4平空
 */
func OrderServer(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	price := query.Get("price")
	num := query.Get("num")
	coin := query.Get("coin")
	direction := query.Get("direction")
	uid := util.GetGUid()
	result := Order(price, num, uid, coin, direction)
	fmt.Println(result)
	if result["error_code"] == "0" {
		fmt.Fprintf(w, "下单成功")
	} else {
		fmt.Fprintf(w, result["error_message"].(string))
	}
}

/**
 * 获取永续合约列表
 * @param coin 币种
 * @param state 订单状态 -2:失败 -1:撤单成功 0:等待成交 1:部分成交 2:完全成交 3:下单中
				4:撤单中 6: 未完成（等待成交+部分成交） 7:已完成（撤单成功+完全成交）
*/
func GetOrderListServer(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	coin := query.Get("coin")
	state := query.Get("state")
	result := GetOrderList(coin, state)
	list := result["order_info"].([]interface{})
	var tmp map[string]interface{}
	for _, v := range list {
		tmp = v.(map[string]interface{})
		fmt.Fprintf(w, "订单号："+tmp["order_id"].(string)+" 开单价："+
			tmp["price"].(string)+" 张数："+tmp["size"].(string)+" 方向："+tmp["type"].(string)+"\n")
		fmt.Printf("订单号：%s 开单价：%s 张数：%s 方向：%s\n", tmp["order_id"].(string),
			tmp["price"].(string), tmp["size"].(string), tmp["type"].(string))
	}
}
