package main

import "fmt"

func main() {
	switch "Yo" {
	case "dope":
		fmt.Println("1")
	case "lol", "Yo":
		fmt.Println("You got it")
	default:
		fmt.Println("nothing")
	}
	// Using it as a lot of if statements
	switch {
	case 0 < 10 && false:
		fmt.Println("nothing")
	case 0 < 10:
		fmt.Println("This condition was true.")
	}
}
