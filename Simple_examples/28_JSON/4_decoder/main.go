// Showing decoding JSON.
package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Person is the type we will decoding JSON into.
type Person struct {
	First       string
	Last        string `json:"last name"`
	Age         int
	notExported int
}

func main() {
	var p1 Person
	rdr := strings.NewReader(`{"First":"Andrew", "last name":"Cloter", "Age":22}`)
	// NewDecoder requires a reader as input.
	json.NewDecoder(rdr).Decode(&p1) // In the documentation notice that it says 'pointed to v'.

	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
