package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func SendGet(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return string(body)
}
