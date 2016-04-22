// This shows you how a slice grows.
package main

import "fmt"

func main() {
	mySlice := make([]int, 0, 5)

	fmt.Println("---------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println("---------")

	for i := 0; i < 81; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Cap:", cap(mySlice), "Value:", mySlice[i])
	}

	mySlice[0]++
	fmt.Println(mySlice[0])
}
