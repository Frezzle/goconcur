package main

import (
	"fmt"
	"sync"
)

// go's data race detector will detect the race here with `go build -race` and then running the binary

func init() {
	// data races will not happen with single-threaded program, though go's race detector does not take this into account
	// runtime.GOMAXPROCS(1)
}

func main() {
	var counter int
	goroutines := 1000
	countTo := 1000

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < countTo; i++ {
				counter++ // increment operator is not atomic
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter, "Off by:", goroutines*countTo-counter)
}

// Output example:
// Counter: 342242 Off by: 657758
