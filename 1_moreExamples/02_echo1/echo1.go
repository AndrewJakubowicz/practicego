// echo1 echoes the arguments given to it.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	fmt.Println("File than ran called:", os.Args[0])
	// s, sep := "",  ""  // This is an alternative of the above code.
	// os.Args[0] is the name of the command itself.
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "+"
	}
	fmt.Println(s)
}
