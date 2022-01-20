package main

import "fmt"

func main() {
	go func() {
		fmt.Println("Can this print before the program exits?...") // ...unlikely
	}()
}
