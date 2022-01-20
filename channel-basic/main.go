package main

import "fmt"

func main() {
	ch := make(chan string)
	go func() {
		ch <- "Hello!"
	}()
	fmt.Println("Message received:", <-ch)
}
