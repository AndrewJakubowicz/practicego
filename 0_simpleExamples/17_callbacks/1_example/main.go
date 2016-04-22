// Example of a callback
package main

import "fmt"

// This function takes a slice of numbers and a function that takes in an int.
// Then when I call this function I can call it with any function that takes
// an int.
func example(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	example([]int{1, 2, 6, 5, 8}, func(n int) {
		fmt.Println(n)
	})
}
