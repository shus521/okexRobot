package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func SetSign(timestamp string, method string, requestPath string, body string, secretKey string) string {
	message := timestamp + strings.ToUpper(method) + requestPath + body
	mac := hmac.New(sha256.New, []byte(secretKey))
	mac.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}

func GetOkTime(url string) interface{} {
	res, err := http.Get(url)
	if err != nil {
		return GetOkTime(url)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	data := map[string]interface{}{}
	json.Unmarshal(body, &data)
	return data["epoch"]
}
