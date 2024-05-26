package main

import "fmt"

func channelFromSlice[T any](s []T) <-chan T {
	out := make(chan T)
	go func() {
		for i := range s {
			out <- s[i]
		}
		close(out)
	}()
	return out
}

func double(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range input {
			out <- 2 * num
		}
		close(out)
	}()
	return out
}

// sprinkle some generator func in there
func addN(n int) func(<-chan int) <-chan int {
	return func(input <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for num := range input {
				out <- n + num
			}
			close(out)
		}()
		return out
	}
}

func print(ch <-chan int) {
	for res := range ch {
		fmt.Println("Result:", res)
	}
}

func main() {
	input := channelFromSlice([]int{1, 2, 3, 4, 5})
	doubled := double(input)
	added := addN(5)(doubled)
	print(added)
}
