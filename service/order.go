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

func Order() map[string]interface{} {
	url := "/api/swap/v3/order"
	orderData := OrderData{
		ClientOid:    "test1",
		Size:         "100",
		OrderType:    "0",
		Type:         "1",
		MatchPrice:   "",
		Price:        "100",
		InstrumentID: "BTC-USDT-SWAP",
	}
	jsonStr, _ := json.Marshal(orderData)
	result := util.SendPost(url, string(jsonStr))
	return result
}
