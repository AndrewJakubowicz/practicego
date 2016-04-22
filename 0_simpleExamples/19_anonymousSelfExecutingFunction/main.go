// This is more a functional programming example but it's good to be aware of it.
// Anonymous self executing functions essentially are anonymous functions with () at the end.
// This just means they execute.
package main

import "fmt"

func main() {
	func() {
		fmt.Println("WOW this executed!")
	}()
}
