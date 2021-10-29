package main

import (
	"fmt"
	"math/rand"
	"time"
)

// delayClose closes a channel after a random amount of time.
func delayClose(c chan struct{}) {
	time.Sleep(time.Millisecond * time.Duration(rand.Int31n(10000)))
	fmt.Println("Closing channel after delay...")
	close(c)
	fmt.Println("Channel closed after delay.")
}

// sendRandomInts sends random ints to a channel in random intervals.
func sendRandomInts(c chan int) {
	for {
		c <- rand.Int()
		time.Sleep(time.Millisecond * time.Duration(rand.Int31n(1000)))

		// TODO what happens if this channel is closed; does the select's `x, ok ...` case become unblocking?
		if rand.Int31n(10) == 1 {
			fmt.Println("Stopping sending results; closing ints channel.")
			close(c)
			return
		}
	}
}

func main() {
	// seed random so repeated runs offer more insight into golang behaviour
	rand.Seed(time.Now().UnixNano())

	// quit the program after some time
	quit := make(chan struct{})
	go delayClose(quit)

	// do some work in parallel that generates results
	results := make(chan int)
	go sendRandomInts(results)

	// get results until program wants to quit or times out
	timeout := time.After(time.Second * 5)
	for {
		select {
		// these 3 cases would all be ready at the same time, so one is chosen pseudo-randomly
		case x, ok := <-results:
			fmt.Println("A", x, ok)
		case x := <-results:
			fmt.Println("B", x)
		case x := <-results:
			fmt.Println("C", x)

		// a closed channel will not block and can keep receiving
		case <-quit:
			fmt.Println("Quitting from closed channel.")
			return

		// can quit after a certain amount of time
		case <-timeout:
			fmt.Println("Quitting from timeout.")
			return

			// default case is optional; it executes if no other case is ready
			// default:
			// 	fmt.Println("Waiting...")
			// 	time.Sleep(time.Millisecond * 200)
		}

		// does adding delay cause potentially interrupt to be sent but not selected for a while while results are still selected? -> Yes!
		// time.Sleep(time.Second)
	}
}
