package main

import "fmt"

func main(){
	println("It's the Arrays")

	var fruits[4]string
	fruits[0] = "Apple"
	fruits[1] = "Mango"
	fruits[3] = "Peach"
	// fruits[2] = "Grapes"
	fmt.Println("Fruit list is:",fruits)
	fmt.Println("Fruit list is:",len(fruits))

	var veg = [3] string{"Potato", "Mushroom", "Beans"}
	fmt.Println("Fruit list is:",veg)

}