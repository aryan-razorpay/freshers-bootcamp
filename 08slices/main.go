package main

import (
	"fmt"
	"sort"
)

func main(){
	fmt.Println("Welcome to slices ")
	// var fruits1 = []string{}
	var fruits = []string{"Apple", "Tomato"}
	fmt.Printf("Type of fruits are %T\n", fruits)
	fmt.Println(fruits)
	fruits = append(fruits, "Mango", "Banana")
	fmt.Println(fruits)
	fruits = append(fruits[1:])
	fmt.Println(fruits)
	fruits = append(fruits[2:3])
	fmt.Println(fruits)
	fruits = append(fruits, "Mango", "Banana")
	fruits = append(fruits[:2])
	fmt.Println(fruits)

	highscores:= make([]int, 4)
	highscores[0]=32
	highscores[1]=34
	highscores[2]=35
	highscores[3]=326
	// highscores[4]=326
	highscores = append(highscores, 21, 45)
	fmt.Println(highscores)
	sort.Ints(highscores)
	fmt.Println(highscores)

	var courses = []string{"ruby", "javascript", "swift", "golang", "reactjs"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}