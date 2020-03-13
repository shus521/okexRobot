package main

import (
	service "okexRobot/service"
	util "okexRobot/utils"
	"os"
)

func main() {
	coin := "btc-usdt"
	args := os.Args
	if len(args) == 2 {
		coin = args[1]
	}
	conf := util.GlobalObject
	service.GetCoinInfo(conf.Url, coin)

}
