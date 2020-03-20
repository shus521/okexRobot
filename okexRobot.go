package main

import (
	"fmt"
	"log"
	"net/http"
	service "okexRobot/service"
	"os"
	"time"
)

func main() {
	//var receiver []string = []string{"1780848749@qq.com"}
	coin, genre := "BTC", "SWAP"
	args := os.Args
	if len(args) == 3 {
		coin, genre = args[1], args[2]
	}

	// 同时启动一个web服务，用来下单，撤单，获取订单列表
	go web()
	var pricePool []float64
	for {
		fmt.Println(time.Now().Local())
		last := service.GetCoinInfo(coin, genre)
		if len(pricePool) < 5 {
			pricePool = append(pricePool, last)
		} else {
			pricePool = append(pricePool[1:], last)
		}

		fmt.Println(pricePool)
		//if last < 6200 {
		//	SendMail("btc价格提醒", "btc低于6200了", receiver)
		//} else if last > 6300 {
		//	SendMail("btc价格提醒", "btc高于6300了", receiver)
		//}
		service.GetAllHolding(coin, genre)
		service.GetMyHolding(coin, last)
		time.Sleep(1000 * time.Millisecond)
	}
}

func web() {
	http.HandleFunc("/order", service.OrderServer)
	http.HandleFunc("/getOrderList", service.GetOrderListServer)
	http.HandleFunc("/cancelOrder", service.CancelOrderServer)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
