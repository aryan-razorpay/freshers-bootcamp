package main

import "fmt"
func main(){
	fmt.Println("Welcome to pointers")
	// var ptr *int
	// fmt.Println(ptr)

	myNumber:=23
	var ptr = &myNumber
	var ptr1 *int = &myNumber
	fmt.Println("Value of pointer is ",ptr)
	fmt.Println("Value of pointer is ",ptr1)
	fmt.Println(*ptr)
	*ptr = *ptr+2
	fmt.Println("Value of pointer is ",*ptr)

} 