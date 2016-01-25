package main

import "fmt"

func main() {
	data := []float64{2, 6, 3, 5}
	fmt.Println(average(data...)) // Notice the 3 dots after. Splits the slice.

}

func average(sf ...float64) float64 { // The dots before allows for an unlimited number of float64's.
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
