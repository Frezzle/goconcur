// TODO put this into github.com/frezzle/goconcur
package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// How many Go features can I cram into one place?
//
// 1. Generics.
// 2. Variadic functions.
// 3. Concurrency (goroutines + channels).
//
// Any more?
func fanIn[T any](channels ...chan T) chan T {
	out := make(chan T)
	for _, in := range channels {
		go func() {
			for {
				out <- <-in
			}
		}()
	}
	return out
}

func startTalker(name string, silence chan any, wg *sync.WaitGroup) chan string {
	ch := make(chan string)
	go func() {
		defer func() {
			fmt.Println(name, "calling wg.Done()")
			wg.Done()
		}()
		for {
			select {
			case <-silence:
				ch <- fmt.Sprintf("%s: Bye!", name)
				return
			default:
				ch <- fmt.Sprintf("%s: %s.", name, randString(10))
				time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			}
		}
	}()
	return ch
}

func main() {
	silence := make(chan any)
	var wg sync.WaitGroup
	wg.Add(3) // tricky: better pattern is to Add before the possibility of Done's being called, but then we can't use dynamic len(talkers) here :(
	talkers := []chan string{
		startTalker("Alice", silence, &wg),
		startTalker("Bob", silence, &wg),
		startTalker("Charlie", silence, &wg),
	}
	inbox := fanIn(talkers...)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan any)
	go func() {
		initiatedStoppingTalkers := false
		for {
			select {
			case message := <-inbox:
				fmt.Printf("Received: %s\n", message)
			case signal := <-signals:
				if initiatedStoppingTalkers {
					fmt.Printf("Received signal %s; already stopped talkers, just waiting for them now...\n", signal)
					break
				}
				initiatedStoppingTalkers = true
				fmt.Printf("Received signal %s; stopping talkers...\n", signal)
				close(silence)
				go func() {
					wg.Wait()
					done <- 0
				}()
			}
		}
	}()
	<-done
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
