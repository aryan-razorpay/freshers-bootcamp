package main

import "fmt"

func main(){
	fmt.Println("Channels")
	myCh:= make(chan int)
	go func ()  {
		fmt.Println(<-myCh)
	}()
	myCh <- 5
}