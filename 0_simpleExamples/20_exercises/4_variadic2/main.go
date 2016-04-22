// Write a variadic function that can take in the following inputs:
// foo(1, 2)
// foo(1, 2, 3)
// aSlice := []int{1, 2, 3, 4}
// foo(aSlice...)
// foo()
package main

import "fmt"

func foo(n ...int) {
	fmt.Println(n)
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	foo(aSlice...)
	foo()
}
