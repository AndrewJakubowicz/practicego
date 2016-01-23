package main

import "fmt"

func main() {
	fmt.Println("First loop!")
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("i = %d\tj = %d\n", i, j)
		}
	}

	fmt.Println("Second loop!")
	n := 0
	for n < 5 {
		fmt.Println(n)
		n++
	}

	fmt.Println("Third loop!")
	i := 0
	for {
		i++ // called incrementer
		if i%2 == 1 {
			continue
		}
		if i >= 10 {
			break
		}
		fmt.Println(i)
		// if i++ is here then it's called a post.

	}
}
