package main

import "fmt"

func main() {
	fmt.Println(42, "is a number.")
	fmt.Printf("%d - %b \n", 42, 42)             // first argument 42 is applied to the first format verb.
	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)   // decimal, binary, hexadecimal(with prefix)
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42) // notice change in hexadecimal format verb
	for i := 60; i <= 122; i++ {
		fmt.Printf("%d - %b - %x - %q \n", i, i, i, i) // Basic loop with utf-8
	}
	// Hi
}
