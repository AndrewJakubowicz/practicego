package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(res.Body)
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Move word by word.
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
