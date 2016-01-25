package main

import "fmt"

func wrapper() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

// Observe how the addScore and secondScore can keep track of themselves.
// This is a silly example so play with it until you understand it.
// Read about closures: http://www.howtogolang.com/post/closure/
func main() {
	oneScore := wrapper()
	secondScore := wrapper()
	fmt.Println("1: ", oneScore())
	fmt.Println("2: ", secondScore())
	for i := 0; i < 10; i++ {
		fmt.Println("1: ", oneScore())
	}
	fmt.Println("2: ", secondScore())
}
