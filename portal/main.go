package main

import (
	"time"
)

func main() {
	c := make(chan int)

	go func() {
		c <- 1 // Does this cause a deadlock?
	}()

	time.Sleep(time.Millisecond * 1000)
}
