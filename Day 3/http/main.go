package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	url := "https://www.google.com"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v", err)
	}
	defer resp.Body.Close()

	fmt.Printf("Response status for %s: %s\n", url, resp.Status)

	scanner := bufio.NewScanner(resp.Body)

	for i := 0; scanner.Scan(); i++ {
		fmt.Println(scanner.Text())
		if i >= 3 { 
			break
		}
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading the response body: %v", err)
	}
}
