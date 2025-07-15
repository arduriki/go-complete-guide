package main

import (
	"fmt"
	"time"
)

// These functions are going to accept a goroutine with a channel.
func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	// Send data to a channel - in this case a bool value.
	doneChan <- true
	// Close the channel once it's done, assign it to the slowest function.
	close(doneChan)
}

func main() {
	// Create a channel (like a communication device or "channel") for the goroutines.
	done := make(chan bool)
	// dones := make([]chan bool, 4)

	// dones[0] = make(chan bool)
	go greet("Nice to meet you!", done)
	// dones[1] = make(chan bool)
	go greet("How are you?", done)
	// dones[2] = make(chan bool)
	go slowGreet("How ... are ... you ...?", done)
	// dones[3] = make(chan bool)
	go greet("I hope you're liking this course!", done)

	// read from the channels
	// for _, done := range dones {
	// 	<-done
	// }

	for range done {
		// fmt.Println(doneChan)
	}
}
