package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// go's data race detector will detect the race here with `go build -race` and then running the binary

func init() {
	// data races will not happen with single-threaded program, though go's race detector does not take this into account
	// runtime.GOMAXPROCS(1)
}

func main() {
	var counter int32
	goroutines := 1000
	countTo := 1000

	var wg sync.WaitGroup
	wg.Add(goroutines)

	for g := 0; g < goroutines; g++ {
		go func() {
			for i := 0; i < countTo; i++ {
				atomic.AddInt32(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter, "Off by:", goroutines*countTo-int(counter))
}

// Output:
// Counter: 1000000 Off by: 0
