package main

import (
	"fmt"
	"strings"
)

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment()) // 1
	fmt.Println(increment()) // 2

	// Create a board
	game := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Some turns
	game[1][0] = "1"
	game[0][2] = "2"
	game[1][2] = "3"

	printBoard(game)
}

func printBoard(s [][]string) {
	for i := 0; i < len(s); i++ {
		fmt.Printf("%s\n", strings.Join(s[i], " "))
	}
}
