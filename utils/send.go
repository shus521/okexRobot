package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/**
 * 发送get请求
 */
func SendGet(url string) map[string]interface{} {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	data := map[string]interface{}{}
	json.Unmarshal(body, &data)
	return data
}
