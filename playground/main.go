package main

import (
	"fmt"
	"reflect"
	"time"
)

func say(msg string) {
	fmt.Println(msg)
}

func maybeSay(msg string) error {
	fmt.Println(msg)
	return nil
}

func main() {
	fmt.Println("start")
	defer fmt.Println("end")

	go say("1")
	go fmt.Println("2")

	// TODO: what's the right way to return values from goroutine?
	go maybeSay("3")

	// execute lambda with arg
	go func(x string) {
		say(x)
	}("4")

	// basic channel
	var c chan int // channels start as nil
	fmt.Println("channel", c, reflect.TypeOf(c))
	c = make(chan int) // new unbuffered channel
	sum := func(nums []int, destination chan int) {
		s := 0
		for _, num := range nums {
			s += num
		}
		destination <- s // will succeed even if destination channel is nil
	}
	go sum([]int{1, 2, 3}, c)
	go sum([]int{4, 5, 6}, c)
	sum1, sum2 := <-c, <-c
	fmt.Println("Sums:", sum1, sum2)

	// TODO: right way to exit only after all goroutines are done?
	time.Sleep(time.Millisecond * 100)

	ch := make(chan int, 1)
	sendStuff(ch)
	select {
	case val := <-ch:
		fmt.Println("Received from channel:", val)
	case <-time.After(time.Millisecond * 100):
		fmt.Println("Timed out waiting for receive")
		// default: // uncomment this to cause default case to execute, assuming no value was sent to channel
		// 	fmt.Println("Default cause because all others are blocked")
	}
}

func sendStuff(c chan<- int) {
	// c <- 0 // uncomment this to prevent the timeout
}
