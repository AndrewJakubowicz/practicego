package main

import "fmt"

func zero(z *int) {
	*z = 0
}

type Vertex struct {
	X int
	Y int
}

func main() {

	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a // the type is a pointer to an int. Points to a memory address where an int is stored.

	fmt.Println(b)
	fmt.Println(*b) // de-referenced so returns 43

	*b = 42 // reassign value of b to 42 which changes a.
	fmt.Println(a)

	// Check out this amazing function.
	x := 11
	fmt.Println(x)
	zero(&x)
	fmt.Println(x)

	// Struct example
	// reference struct through a pointer
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)

}
