package main

import "fmt"

func main(){
	fmt.Println("Learning functions")
	greeter();
	// func greeter2(){
	// 	fmt.Println("Another hello")
	// }
	result:= adder(3, 5)
	fmt.Println("The addition is", result)
	// fmt.Println("The result is ", proAdder(2, 3, 4, 5, 6))
	resultpro, message := proAdder(2, 3, 4, 5, 6)
	fmt.Println("The result is ", resultpro)
	fmt.Println("The message is ", message)
}

func adder(val1 int, val2 int) int{
	return val1+val2
}

func proAdder(arg ...int) (int, string){
	total:=0
	for _, val :=range arg{
		total+=val
	}
	return total, "Hello coders"
}

func greeter(){
	fmt.Println("Hello everybody")
}
// func (){
// 	fmt.Println("Hello everybody")
// }()