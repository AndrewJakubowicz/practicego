package main

import "fmt"

// Prints out the number, the string and the bytes for that character.
// Works best on unix machine.
func main() {
	for i := 5000; i < 5020; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := 'a' // Rune because single quotes
	fmt.Println(foo)
	fmt.Printf("%T: %v \n", foo, foo)
	fmt.Printf("%T: %v \n", string(foo), string(foo))

}
