package people

import (
	"log"

	"github.com/go-resty/resty/v2"
)

var (
	server_address = "http://163.18.110.100/api/EmployeeJobtitleTypes"
)

type POST struct {
	JobtitleId  int    `json:jobtitleId`
	Name        string `json:name`
	CompanyHash string `json:companyHash`
}

func CompanyLogin(code string) string {
	client := resty.New()
	url := server_address + "Companies/" + code
	resp, err := client.R(). //點要連接要最後,表示方法延續
					Get(url)
	if err != nil {
		log.Fatal(err) //log也可print;fatal 會做以下兩點
		//1.打印輸出內容 2.退出應用程序
	}
	return string(resp.Body())
}
func POST_P(body string) string {
	client := resty.New()
	url := server_address //+ "EmployeeJobtitleTypes/add_jobtitle"
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(body).
		Post(url)
	if err != nil {
		log.Fatal(err)
	}
	return string(resp.Body())
}
func PUT_P() string {
	client := resty.New()
	url := server_address // + "EmployeeJobtitleTypes/UpdateJobtitle"
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`[
			{
			"jobtitleId": 3012,
			"name": "string2"
			}
			]`).
		Put(url)
	if err != nil {
		log.Fatal(err)
	}
	return string(resp.Body())

}

func DELETE_P(id string) string {
	client := resty.New()
	url := server_address + "EmployeeJobtitleTypes/DeleteJobtitle/" + id
	resp, err := client.R().
		Delete(url)
	if err != nil {
		log.Fatal(err)

	}
	return string(resp.Body())
}
func GET_P() []byte {
	client := resty.New()
	url := server_address // + "EmployeeJobtitleTypes/add_jobtitle"
	resp, err := client.R().
		Get(url)
	if err != nil {
		log.Fatal(err)
	}
	return resp.Body()
}
