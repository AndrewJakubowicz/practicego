// Alter the code to use a Function Expression.
package main

import "fmt"

func main() {
	half := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	fmt.Println(half(2))
	fmt.Println(half(1))
}
