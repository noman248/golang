package api_test

import (
	"log"

	"github.com/go-resty/resty/v2"
)

var (
	server_address = "http://163.18.110.100" + "/api/"
)

func CampanyLogin(code string) string {
	client := resty.New()
	url := server_address + "Companies/" + code
	resp, err := client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
		//log 輸出有日期
		//打印輸出內容,退出應用程式
	}
	return string(resp.Body())
}
