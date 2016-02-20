package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

// Imagine a channel like a relay race. Passing a baton.
func main() {
	c := make(chan int)
	wg.Add(1)
	go func() {
		for i := 0; i < 20; i++ {
			c <- i
			// Code stops until the value is taken off the channel
		}
		wg.Done()
	}()
	go func() {
		for {
			fmt.Println(<-c) // Takes value off the channel.
		}
	}()
	wg.Wait()
}
