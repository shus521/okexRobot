package service

/**
 * @author ty
 * @desc 下单
 * @version 1.0
 */

import (
	"encoding/json"
	util "okexRobot/utils"
)

type OrderData struct {
	ClientOid    string `json:"client_oid"`
	Size         string `json:"size"`
	OrderType    string `json:"order_type"`
	Type         string `json:"type"`
	MatchPrice   string `json:"match_price"`
	Price        string `json:"price"`
	InstrumentID string `json:"instrument_id"`
}

/**
 * 永续合约下单
 * @param price 下单价格
 * @param num 下单数量(张)
 * @param uid 订单id
 * @param coin 币种
 * @param direction 交易方向 1开多 2开空 3平多 4平空
 */
func Order(price string, num string, uid string, coin string, direction string) map[string]interface{} {
	url := "/api/swap/v3/order"
	orderData := OrderData{
		ClientOid:    uid,
		Size:         num,
		OrderType:    "0",
		Type:         direction,
		MatchPrice:   "",
		Price:        price,
		InstrumentID: coin + "-USDT-SWAP",
	}
	jsonStr, _ := json.Marshal(orderData)
	result := util.SendPost(url, string(jsonStr))
	return result
}
