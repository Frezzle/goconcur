package main

import (
	"fmt"
	"time"
)

func playCatch(player string, c chan interface{}) {
	for {
		x := <-c
		fmt.Println(player, "caught it; throwing it back...")
		time.Sleep(time.Millisecond * 100)
		c <- x
	}
}

func main() {
	c := make(chan interface{})

	// summon players to wait for the first throw
	go playCatch("B", c)
	go playCatch("C", c)
	go playCatch("D", c)

	// throw a ball
	c <- 1

	// one player must play within the main goroutine so program keeps running
	playCatch("A", c)

	// Q1: What happens when multiple goroutines are waiting to receive from the same channel?
	// Q2: In what order do the players catch the ball? Does the order stay the same or not during execution?
}
