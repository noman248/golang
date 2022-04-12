package main

import (
	"encoding/json"
	"fmt"
	"git-/people"
)

func main() {
	 fmt.Println(people.PUT_P())

	/*data := people.GET_P()
	str := []people.POST{}
	if err := json.Unmarshal(data, &str); err != nil {
		fmt.Println(err)
	}
	fmt.Println(str)*/
	/*job := people.POST{
		Name:        "hello",
		CompanyHash: "17C99EADD5A77BD5117719B12478E7",
	}
	joblist := []people.POST{job}
	marshal, err := json.Marshal(joblist)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(marshal))
	fmt.Println(people.POST_P(string(marshal)))*/

	// fmt.Println(people.DELETE_P("3013"))

}
