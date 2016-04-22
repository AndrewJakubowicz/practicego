// Create a program that asks for a name and then prints "Hello <NAME>"
package main

import (
	"fmt"
	"bufio"
	"os"
)

var name string

func main() {
	fmt.Print("Enter your name: ")
	/*
	fmt.Scan(&name) // Doesn't work for many space seperated things because needs that many variables.
	fmt.Println("Hello", name)
	*/

	// This allows multiple words.
	in := bufio.NewReader(os.Stdin)
	line, _ := in.ReadString('\n')
	fmt.Println("Hello " + line)

}
