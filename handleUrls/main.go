package main

import (
	"fmt"
	"net/url"
)

const myUrl string= "https://www.freecodecamp.org:3000/news/learn-react-2024/di?shi=hi"

func main(){
	fmt.Println("Url Hadling")
	fmt.Println(myUrl)
	result, _ :=url.Parse(myUrl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Print(result.RawQuery)
	fmt.Println(result.Port())
	fmt.Printf("\n")
	qparams := result.Query()
	fmt.Printf("The type of query params is %T\n", qparams)
	fmt.Println(qparams["shi"])

	for _, val:=range(qparams){
		fmt.Println("Params is:", val)
	}
	partsofUrl := &url.URL{
		Scheme: "https",
		Host: "www.freecodecamp.org",
		Path:"/news",
		RawQuery: "user=Aryan",
	}
	anotherUrl := partsofUrl.String()
	fmt.Println(anotherUrl)
}