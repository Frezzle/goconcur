package main

import (
	"fmt"
	"time"
)

func repeat(message string) {
	for i := 0; i < 10; i++ {
		fmt.Println(message)
		time.Sleep(time.Second)
	}
}

func main() {
	go repeat("hello")
	go repeat("hi")
	time.Sleep(5 * time.Second)
}
