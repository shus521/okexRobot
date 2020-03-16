package service

import (
	"fmt"
	"net/http"
)

func OrderServer(w http.ResponseWriter, req *http.Request) {
	result := Order()
	fmt.Println(result)
	if result["error_code"] == "0" {
		fmt.Fprintf(w, "下单成功")
	} else {
		fmt.Fprintf(w, result["error_message"].(string))
	}
}
