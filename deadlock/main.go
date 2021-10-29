package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1                        // deadlocks here, but if this line were gone...
	fmt.Println("Received:", <-c) // ...then it would deadlock here
}
