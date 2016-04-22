package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	fmt.Println("Atomic adding!")
	wg.Add(2)
	go increment("\tFoo:")
	go increment("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func increment(s string) {
	for i := 0; i < 75; i++ {
		// race:
		// counter++
		// no race:
		atomic.AddInt64(&counter, 1)
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

// go run main.go
// go run -race main.go
