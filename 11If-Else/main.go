package main

import "fmt"

func main(){
	fmt.Println("If else in Golang")
	loginCount:=23
	var result string
	if (loginCount<23) {
		result = "Regular User"
	} else if(loginCount>23){
		result="Watch Out"
	} else {
		result = "Exactly 23"
	}
	fmt.Println(result)

	if num:=11 ; num==11{
		println("num is less")
	}
}