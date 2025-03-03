package main

import (
	"fmt"
)

func main(){
	fmt.Println("Loops")
	days:=[]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	fmt.Println(days)
	// for d:=0; d<len(days); d++{
	// 	fmt.Println(days[d])
	// }
	// for i:=range(days){
	// 	fmt.Println(days[i])
	// }
	for index, day:=range(days){
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	rougeVlaue:=1
	for rougeVlaue<10{
		if(rougeVlaue==2){
			goto lco
		}

		if(rougeVlaue==5){
			// break
			rougeVlaue++
			continue
		}
		fmt.Println("Value is:", rougeVlaue);
		rougeVlaue++
	}
	lco:{
		fmt.Println("This is an example of I guess go to statements  ")
	}
}