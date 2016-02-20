package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mutex sync.Mutex
var counter int

func main() {
	wg.Add(2)
	go increment("\tFoo:")
	go increment("Bar:")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

func increment(s string) {
	for i := 0; i < 75; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}
