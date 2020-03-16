package main

import (
	"fmt"
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
	for {
		fmt.Println(time.Now().Local())
		service.GetCoinInfo(coin, genre)
		service.GetAllHolding(coin, genre)
		service.GetMyHolding(coin)
		time.Sleep(1000 * time.Millisecond)
	}
	//service.Order()
}
