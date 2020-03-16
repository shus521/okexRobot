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
	coin, genre := "BTC", "SWAP"
	args := os.Args
	if len(args) == 3 {
		coin, genre = args[1], args[2]
	}

	// 同时启动一个web服务，用来下单，撤单，获取订单列表
	go web()
	for {
		fmt.Println(time.Now().Local())
		service.GetCoinInfo(coin, genre)
		service.GetAllHolding(coin, genre)
		service.GetMyHolding(coin)
		time.Sleep(1000 * time.Millisecond)
	}
	//service.Order()
}

func web() {
	http.HandleFunc("/order", service.OrderServer)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
