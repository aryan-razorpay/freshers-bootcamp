package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("Hello, World!")
	mai()
	sqrt()
	vars()
}

func mai() {
	fmt.Println("Picking a random integer is", rand.Intn(10))
}

func sqrt() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

}

var c = 1
var python = true
var java bool

func vars() {
	var i int
	fmt.Println(i, c, python, java)
}
