package main

import "fmt"
const LoginToken string = "hello"
func main() {
	var username string = "Aryan"
	fmt.Println(username)
	fmt.Printf("Variable is of type %T", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type %T", isLoggedIn)

	var smallVal uint = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type %T", smallVal)

	//no var style
	numberofUsers := 23 
	fmt.Println(numberofUsers)
	println(LoginToken)
}
