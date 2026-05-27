package exercise5

import (
	"log"
	"sync"
	"time"
)

//1. Create and test a thread-safe counter with the following requirements:
//
//Implement a Counter struct with following methods:
//Increment() int - adds 1 and returns new value
//Decrement() int - subtracts 1 and returns new value
//Get() int - returns current value
//Reset() - sets value back to 0
//
//2. Write tests that verify:
//Basic operations work correctly
//Counter works correctly when called from multiple goroutines
//Test using the race detector
//
//Tips:
//
//Use sync.WaitGroup to wait for goroutines in test (we'll provide example)
//Start with basic tests before adding concurrent ones
//Use sync.Mutex to protect shared state

type Counter struct{}

// Using WaitGroup Example
func UsingWaitGroup() {
	// Create a waitGroup
	wG := sync.WaitGroup{}

	// Add tokens as needed
	routineCount := 2
	wG.Add(routineCount)

	// Run goroutines
	for range routineCount {
		go func() {
			time.Sleep(1 * time.Second)
			log.Printf("routine done")

			// Free a token
			wG.Done()
		}()
	}

	// Wait until all routines are done
	wG.Wait()
	log.Printf("everything done")
}
