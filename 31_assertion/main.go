package main

import "fmt"

func main() {
	// Creating an empty interface with underlying string type.
	var name interface{} = "Sydney"
	str, ok := name.(string)
	if ok {
		fmt.Printf("%T\n", str)
	} else {
		fmt.Println("Not a string!")
	}
}
