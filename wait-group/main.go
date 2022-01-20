package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("This prints before the program exits!")
		wg.Done()
	}()
	wg.Wait()
}
