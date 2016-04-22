// Implement the square root function using Newton's method.
// In this case, Newton's method is to approximate Sqrt(x)
// by picking a starting point z and then repeating:
// z = z - (z^2 - x) / (2*z)
// a) repeat that calculation 10 times and see how close you
// get to the answer for various values (1, 2, 3, ...).

// b) Next, change the loop condition to stop once the value has
// stopped changing (or only changes by a very small delta).
// See if that's more or fewer iterations. How close are you to the math.Sqrt?
package main

import (
	"fmt"
	"math"
)

// Finds the square root of the number.
func Sqrt(x float64) float64 {
	z, c, epsilon := float64(1), float64(0), float64(0.0000000001)
	i := 0
	for i <= 50 {
		i++
		c = z - ((z*z)-x)/(2*z)
		if z < c+epsilon && z > c-epsilon {
			z = c
			break
		} else {
			z = c
		}
		fmt.Println(z)
	}
	fmt.Println("Took", i, "repeats.")
	return z

}

func main() {
	for i := 100; i < 110; i++ {
		fmt.Println("Guess: ", Sqrt(float64(i)))
		fmt.Println("Check: ", math.Sqrt(float64(i)))
	}
}
