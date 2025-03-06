package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	totalStudents = 200
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	var wg sync.WaitGroup
	ratings := make([]int, totalStudents)

	for i := range totalStudents {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			responseTime := time.Duration(rand.Intn(5)+1) * time.Second
			time.Sleep(responseTime)
			rating := rand.Intn(5) + 1
			ratings[i] = rating
			fmt.Printf("Student %d gave a rating of %d after %v\n", i+1, rating, responseTime)
		}(i)
	}

	wg.Wait()

	var totalRating int
	for _, rating := range ratings {
		totalRating += rating
	}
	average := float64(totalRating) / float64(totalStudents)

	fmt.Printf("\nThe average rating is: %.2f\n", average)
}
