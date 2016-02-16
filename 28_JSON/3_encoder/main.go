// Example of encoding of JSON using Encode.
package main

import (
	"encoding/json"
	"os"
)

// Person struct to be encoded into JSON.
// Repetition of the tags concept.
type Person struct {
	First       string
	Last        string
	Age         int `json:"my aGe"`
	notExported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 007}
	// Note that NewEncoder takes a writer interface and returns a pointer to Encoder.
	json.NewEncoder(os.Stdout).Encode(p1)
}
