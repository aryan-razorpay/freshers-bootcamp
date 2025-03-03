package main

import (
	"fmt"
	"time"
)

func task(id int){
	fmt.Println("Doing tasks", id)
}

func main() {
	fmt.Println("Goroutines")
	for i:=0; i<=10;i++{
		go task(i)
	}
	time.Sleep(time.Second*2)
}
