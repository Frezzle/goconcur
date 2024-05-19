package main

import (
	"fmt"
	"sync"
	"time"
)

func repeat(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		repeat("hello")
		wg.Done()
	}()

	go func() {
		repeat("hi")
		wg.Done()
	}()

	wg.Wait()
}
