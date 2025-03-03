package main

import "fmt"

func main() {
	fmt.Println("Structs")
	hitesh := User{"Aryan", "aryan@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("Aryan details are: %+v\n", hitesh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
