package main

import (
	"fmt"
	"sync"
)

// var mu sync.RWMutex // the read/write mutex allows reads to happen simultaneously if no write is happening.

func main() {
	var counter int
	var mu sync.Mutex
	goroutines := 1000
	countTo := 1000

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < countTo; i++ {
				// lock the mutex; block until it's this goroutine's turn to lock it
				mu.Lock()

				// some non-trivial operation on shared resource...
				val := counter
				val++
				counter = val

				// give other goroutines a chance to access shared resource
				mu.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter, "Off by:", countTo*goroutines-counter)
}

// Output:
// Counter: 1000000 Off by: 0
