// echo1 echoes the arguments given to it.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string // s, sep := "",  ""
	fmt.Println("File than ran called:", os.Args[0])
	for i := 1; i < len(os.Args); i++ {
		s += sep + strconv.Itoa(i) + ": " + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)
}
