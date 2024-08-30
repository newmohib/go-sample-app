package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Worker function that processes URLs
func worker(id int, jobs <-chan string, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range jobs {
		fmt.Printf("Worker %d processing URL: %s\n", id, url)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Second) // Simulate network delay
		results <- fmt.Sprintf("Processed by worker %d: %s", id, url)
	}
}

func main() {
	urls := []string{"https://example1.com", "https://example2.com", "https://example3.com"}
	jobs := make(chan string, len(urls))
	results := make(chan string, len(urls))

	var wg sync.WaitGroup

	// Start 3 worker goroutines
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send URLs as jobs to the workers
	for _, url := range urls {
		jobs <- url
	}
	close(jobs) // Close the jobs channel to indicate no more jobs

	wg.Wait() // Wait for all workers to finish

	close(results) // Close the results channel

	// Collect and print the results
	for res := range results {
		fmt.Println(res)
	}
}
