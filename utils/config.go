package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type GlobalObj struct {
	Url string `json:"url"`
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	confFile := "config/config.json"
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Println("reload conf err ", err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		log.Println(err)
	}
}
func init() {
	GlobalObject.Reload()
}

// 写入配置文件
func WriteConf(conf *GlobalObj) {
	config, _ := json.Marshal(conf)
	err := ioutil.WriteFile("config/config.json", config, 0666)
	if err != nil {
		fmt.Println("写入配置文件失败")
	}
}
