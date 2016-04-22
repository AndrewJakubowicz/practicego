// FizzBuzz problem
// Write a program that prints the numbers between 1 to 100.
// But for multiples of 3 it prints "Fizz"
// and for multiples of 5 it prints "Buzz".
// For multiples of both it prints "FizzBuzz".
package main

import "fmt"

func main() {
	// Using if and if else.
	/*
		for i := 1; i <= 100; i++ {
			if i%5 == 0 && i%3 == 0 {
				fmt.Println("FizzBuzz")
			} else if i%5 == 0 {
				fmt.Println("Buzz")
			} else if i%3 == 0 {
				fmt.Println("Fizz")
			} else {
				fmt.Println(i)
			}
		}
	*/

	// Using a switch statement.
	for i := 1; i <= 100; i++ {
		switch {
		case i%5 == 0 && i%3 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
