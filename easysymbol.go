package main 

import (
 "fmt"	
) 


func main() {
 var a int =1 
 var b =&a 
fmt.Println("a的值",a)
fmt.Println("a的址",&a)
fmt.Println("a的值",*b)
}
