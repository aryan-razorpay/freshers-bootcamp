package main

import (
	"fmt"

	"rsc.io/quote/v3"
	"rsc.io/sampler"
)

func Hello() string {
	return sampler.Hello()
}
func main(){
	fmt.Println(Hello());
	fmt.Println(quote.HelloV3());
}
