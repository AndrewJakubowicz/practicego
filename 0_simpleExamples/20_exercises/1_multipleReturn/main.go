// Write a func which takes an integer.
// Function will have two returns.
// First return should be the argument divided by 2.
// Second return should be a bool, whether or not argument was even.
package main

import "fmt"

func evenCheck(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func main() {
	fmt.Println(evenCheck(2))
	fmt.Println(evenCheck(1))
}
