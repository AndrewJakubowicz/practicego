// Create a file that asks for a smaller and larger number.
// Then divides the larger number by the smaller number.
package main

import "fmt"

var (
	num1 float64
	num2 float64
)

func main() {
	fmt.Print("Input a number: ")
	fmt.Scan(&num1)
	fmt.Print(" -- Thank you!\nInput a larger number: ")
	fmt.Scanln(&num2)
	fmt.Print(" -- Thank you!\n")
	fmt.Printf("Divide %v by %v = %v", num2, num1, num2/num1)

}
