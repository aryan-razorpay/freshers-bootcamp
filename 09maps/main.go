package main

import "fmt"

func main(){
	fmt.Println("Maps in Golang")
	languages := make(map[string]string)
	languages["JS"] = "JS"
	languages["JS"] = "JS"
	languages["JS"] = "JS"
	languages["ruby"] ="ruby"
	languages["bill"] ="ruby"
	languages["python"] ="python"

	fmt.Println("List of languages: ", languages)
	delete(languages, "ruby")
	fmt.Println("List of languages: ", languages)

	for key, value:= range languages{
		fmt.Println("For key %v, value is %v", key, value)
	}
}

