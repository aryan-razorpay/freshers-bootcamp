package main

import (
	"fmt"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")
	fmt.Fprintf(w, "<h1>Welcome to the Hello Page!</h1>")
	select {
	case <-time.After(2 * time.Second):
		fmt.Fprintf(w, "hello\n")
	case <-ctx.Done(): 
		err := ctx.Err()
		fmt.Println("server:", err)

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("server: starting server on :8060")
	if err := http.ListenAndServe(":8060", nil); err != nil {
		fmt.Println("server: error starting server:", err)
	}
}
