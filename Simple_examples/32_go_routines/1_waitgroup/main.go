package main

import (
	"fmt"
	"sync"
)

// Setting up the wait group.
var wg sync.WaitGroup

func main() {
	// Add 2 to the wait Group
	wg.Add(2)
	go foo()
	go bar()
	wg.Wait() // Keeps func main running until it reaches zero.
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("\tfoo:", i)
	}
	wg.Done() // When this is triggered it subtracts 1 from the waitgroup.
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar:", i)
	}
	wg.Done() // subtracts 1 from wait group.
}
