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

func (p people) Len() int {
	// fmt.Println("length is", len(p))
	return len(p)
}
func (p people) Swap(i, j int) {
	// fmt.Println("swapping", p[i], "with", p[j])
	p[i], p[j] = p[j], p[i]
}
func (p people) Less(i, j int) bool {
	// fmt.Println("is", p[i], "less than", p[j])
	return p[i] < p[j]
}

// Reverse sort by pretty much creating another type that has different methods.
type reverse people

// These three methods are a requirement of Interface interface for sort.
func (l reverse) Len() int {
	return len(l)
}

func (l reverse) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l reverse) Less(i, j int) bool { return l[i] > l[j] }

func main() {
	studyGroup := people{"Zuno", "Jahn", "Albert", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(reverse(studyGroup))
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
