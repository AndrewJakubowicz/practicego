// Write a function with one variadic parameter that funds the greatest number in a list of numbers.
package main

import "fmt"

func greatest(n ...int) int {
	largest := 0
	for _, v := range n {
		if largest < v {
			largest = v
		}
	}
	return largest

}

func main() {
	numbers := []int{1, 5, 8, 9, 14}
	fmt.Println("answer: ", greatest(numbers...))
}
