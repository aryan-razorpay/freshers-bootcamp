// package main
// import "fmt"
// func main() {
// 	fmt.Println(quote.Go())
// }

// func main(){
// 	fmt.Println()
// }

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// URL of the CSV file (from the GitHub link)
	url := "https://raw.githubusercontent.com/umbrae/reddit-top-2.5-million/master/data/quotes.csv"

	// Fetch the CSV file from the URL
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Check if the response is successful (200 OK)
	if resp.StatusCode != 200 {
		log.Fatalf("Failed to fetch data: %s", resp.Status)
	}

	// Create a CSV reader
	reader := csv.NewReader(resp.Body)

	// Read all records from the CSV
	var quotes [][]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file
		}
		if err != nil {
			log.Fatal(err)
		}
		quotes = append(quotes, record)
	}

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Randomly select a quote
	randomIndex := rand.Intn(len(quotes))
	randomQuote := quotes[randomIndex]

	// Print the random quote
	fmt.Println("Random Quote:")
	fmt.Println(randomQuote[4]) // Assuming the quote is in the first column
}

func getRandomQuote() string {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	return quotes[randomIndex][0] // Assuming the quote is in the first column
}


