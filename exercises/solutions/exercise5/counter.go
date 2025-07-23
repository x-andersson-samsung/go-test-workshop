package exercise5

import (
	"log"
	"sync"
	"time"
)

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
