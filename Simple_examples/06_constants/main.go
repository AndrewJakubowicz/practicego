package main

import "fmt"

const p string = "death & taxes"

// Constant is simply unchanging.

const (
	A = iota // 0
	B = iota // 1
	C        // assumes iota and makes it = 2
)

const (
	_  = iota
	KB = 1 << (iota * 10) // 1 << (1 * 10) (so bit shift 10 to the left)
	MB = 1 << (iota * 10) // 1 << (2 * 10) (bit shift 10 more times)
	GB = 1 << (iota * 10)
)

func main() {
	const q = 42

	fmt.Println("p - ", p)
	fmt.Println("q - ", q)

	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)

	fmt.Printf("%b\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
	fmt.Printf("%b\t%d\n", GB, GB)
}
