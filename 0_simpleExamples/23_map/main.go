package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Hello!",
		1: "rar!",
		2: "How do you do?!",
	}

	for k, v := range myGreeting {
		fmt.Println(k, "-", v)

	}
}
