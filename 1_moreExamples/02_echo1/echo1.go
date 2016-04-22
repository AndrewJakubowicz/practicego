// echo1 echoes the arguments given to it.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	// os.Args[0] is the name of the command itself.
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "+"
	}
	fmt.Println(s)
}
