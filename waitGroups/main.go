package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals = []string{"test"}
var wg sync.WaitGroup
var mut sync.Mutex

func main(){
	fmt.Println("Wait Groups")
	websites:=[]string{"https://google.com", "https://stackoverflow.com/", "https://go.dev/", "https://geeksforgeeks.org/"}
	for _, web:= range(websites){
		wg.Add(1)
		getStatusCode(web)
	}
	wg.Wait()
	fmt.Println(signals)
}
func getStatusCode(endpoint string){
	defer wg.Done()
	res, err:= http.Get(endpoint)
	if err!=nil{
		fmt.Println("Wrong Endpoint")
	} else{
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Println(res.StatusCode)
	}
}