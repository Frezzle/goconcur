package main

import "fmt"

type PipelineStage[T any] func(<-chan T) <-chan T

func runPipeline[T any](input <-chan T, stages ...PipelineStage[T]) <-chan T {
	out := input
	for _, stage := range stages {
		out = stage(input)
		input = out
	}
	return out
}

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
	results := runPipeline(input, addN(1), double, addN(5))
	print(results)
}
