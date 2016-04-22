// Example of defer.
package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

// Defer executes as a last in, first out style.
// Anything with defer will run when the function it is within (main()) in this case,
// is about to exit.

// First counting is printed.
// "done" is printed
// Now main() is about to exit.
// Defer functions run. Last defer call runs first.