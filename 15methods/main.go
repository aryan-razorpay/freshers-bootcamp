package main

import "fmt"

func main(){
	fmt.Println("This is about methods")
	Aryan := User{"Aryan", "aryan@go.dev", true, 16}
	fmt.Println(Aryan)
	fmt.Printf("Aryan details are: %+v\n", Aryan)
	Aryan.GetStatus()
	Aryan.NewMail()
	fmt.Printf("Aryan details are: %+v\n", Aryan)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus(){
	fmt.Println("Is user active", u.Status)
}
func (u User) NewMail(){
	u.Email = "hello@go.dev"
	fmt.Println("Email of this user is",u.Email)
}