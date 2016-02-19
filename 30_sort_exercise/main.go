/*
Use https://godoc.org/sort to sort the following (using different methods):

(1)
type people []string
studyGroup := people{"Zeno", "John", "Al", "Jenny"}

*/
package main

import (
	"fmt"
	"sort"
)

type people []string

// Setting up Interface interface for sort.
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

func main() {
	studyGroup := people{"Zuno", "Derren", "Jahn", "Albert", "Jenny", "Rachel"}
	fmt.Println(studyGroup)
	sort.Sort(sort.Reverse(studyGroup)) // Using reverse interface
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
