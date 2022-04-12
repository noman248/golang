package main
import (
	"fmt"
	"encoding/json"
)

	
	type student struct{
Name string `json:name`
Number string `json:number`
	}
 	
func main() {
	//d := 6
	amy := student{
Name:"amy",
Number:"C109118248",
	}
	andy := student{}
	andy.Name="andy"
	andy.Number="C108118248"
	class := []student{amy,andy}
	//class2 :=[]student{} 
	/*for index,value := range class{
		class2 =append(class,value)
	fmt.Println(index,class2)
	}*/
	marshal,err :=json.Marshal(class)
	if err != nil{
fmt.Println(err)

	}

	class3:=[]student{}
	err = json.Unmarshal(marshal,&class3)
	if err != nil{
		fmt.Println(err)
			}
 fmt.Println(class3) 

}