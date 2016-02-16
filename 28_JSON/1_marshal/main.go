package main

import (
	"encoding/json"
	"fmt"
)

// Person is the struct to convert to JSON.
// Note what is and isn't exported.
type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := Person{"Andrew", "Bond", 20, 007}
	bs, _ := json.Marshal(p1) // Turning struct into byte slice.
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs)) // Wow that's a json string now.
}
