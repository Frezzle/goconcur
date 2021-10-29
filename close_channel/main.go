package main

import (
	"fmt"
)

func main() {
	numbers := 10

	even := make(chan int, numbers)
	odd := make(chan int, numbers)

	go func() {
		for x := 0; x < numbers; x++ {
			if x%2 == 0 {
				even <- x
			} else {
				odd <- x
			}
		}

		// Close the channels, otherwise the range loop in main goroutine will keep waiting and cause deadlock.
		close(even)
		close(odd)

		// even <- 100 // panic: send on closed channel
	}()

	// keep reading each channel until closed...

	// ...can use fancy range loop
	for x := range even {
		fmt.Println("Received even", x)
	}

	// ...or can use bool manually (which is what the range loop is doing behind the scenes)
	for {
		x, open := <-odd
		if !open {
			break
		}
		fmt.Println("Received odd", x)
	}
}
