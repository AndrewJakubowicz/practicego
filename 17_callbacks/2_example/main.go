// A cool example of a callback.
package main

import "fmt"

// A function that allows us to use it while using whatever filter we want.
func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	xs := filter([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(xs) // 2, 4
}
