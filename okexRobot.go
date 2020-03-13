package main

import (
	service "okexRobot/service"
	util "okexRobot/utils"
	"os"
	"time"
)

func main() {
	coin, genre := "BTC", "SWAP"
	args := os.Args
	if len(args) == 3 {
		coin, genre = args[1], args[2]
	}
	conf := util.GlobalObject
	for {
		service.GetCoinInfo(conf.Url, coin, genre)
		service.GetAllHolding(conf.Url, coin, genre)
		time.Sleep(1000 * time.Millisecond)
	}
}
