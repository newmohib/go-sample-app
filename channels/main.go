package main

import (
	"fmt"
	"time"
)

// Producer function that generates numbers and sends them to the channel
func producer(numbers chan int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producing:", i)
		numbers <- i            // Send numbers to the channel
		time.Sleep(time.Second) // Simulate work with sleep 1s
		fmt.Println("time.Second:", time.Second)
	}
	close(numbers) // Close the channel when done producing
}

// Consumer function that receives numbers and processes them
func consumer(numbers chan int) {
	for num := range numbers { // Range over the channel to receive values
		square := num * num
		fmt.Println("Consumed:", num, "Squared:", square)
	}
}

func main() {
	numbers := make(chan int) // Create a channel for integers

	go producer(numbers) // Start the producer goroutine
	consumer(numbers)    // Start the consumer goroutine (in main)
}
