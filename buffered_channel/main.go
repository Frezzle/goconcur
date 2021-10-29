package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	go func() {
		c <- 1
		c <- 2

		// Adding just this would block this goroutine because no more space in buffered channel.
		// c <- 3
	}()

	fmt.Println("Received", <-c)
	fmt.Println("Received", <-c)

	// Adding just this would block this goroutine because buffered channel is empty, also eventually deadlocking.
	// fmt.Println("Received", <-c)
}
