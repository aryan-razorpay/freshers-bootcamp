package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://portfolioaryan.vercel.app/"

func main(){
	fmt.Println("Web requests")
	response, err:=http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close()
	dataBytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(dataBytes)
	fmt.Println(content)
}