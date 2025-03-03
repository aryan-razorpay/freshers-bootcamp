package main

import "fmt"

func main(){
	defer fmt.Println("Hello 1") 
	defer fmt.Println("Hello 2") 
	defer fmt.Println("Hello 3") 
	fmt.Println("About defer 1")
	fmt.Println("About defer 2")
	fmt.Println("About defer 3")
	fmt.Println("About defer 4")
	myDefer()
}

func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Println(i)
	}
	// for i:=range(5){
	// 	fmt.Println(i)
	// }
}