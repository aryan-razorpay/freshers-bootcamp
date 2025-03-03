package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	defer fmt.Println("Welcome to files in golang")
	content := "This needs to be just a text"
	file, err := os.Create("./checkingFiles.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is",length)
	defer file.Close()
	readFile("checkingFiles.txt")
}

func readFile(fileName string){
	dataByte, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is:", string(dataByte))
}
