package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

/**
 * 发送get请求
 */
func SendGet(url string, needHeader string) map[string]interface{} {
	client := &http.Client{}
	config := GlobalObject
	newUrl := config.Url + url
	req, _ := http.NewRequest("GET", newUrl, nil)
	//非公用方法需要设置请求头
	if needHeader != "public" {

		req.Header.Add("OK-ACCESS-KEY", config.Apikey)
		req.Header.Add("OK-ACCESS-PASSPHRASE", config.Passphrase)
		time := GetOkTime(config.Url + "/api/general/v3/time").(string)

		req.Header.Add("OK-ACCESS-TIMESTAMP", time)
		sign := SetSign(time, "GET", url, "", config.Secrekey)
		req.Header.Add("OK-ACCESS-SIGN", sign)

	}
	resp, _ := client.Do(req)
	data := map[string]interface{}{}
	if resp == nil {
		return data
	}
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &data)
	return data
}
