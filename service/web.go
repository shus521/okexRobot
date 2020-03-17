package service

import (
	"fmt"
	"net/http"
	util "okexRobot/utils"
)

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
