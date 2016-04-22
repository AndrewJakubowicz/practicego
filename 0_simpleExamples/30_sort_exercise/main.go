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

// Question 1
type people []string

// Setting up Interface interface for sort.
func (p people) Len() int           { return len(p) }
func (p people) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool { return p[i] < p[j] }

// End question 1

func main() {
	fmt.Println("\nQuestion 1 results using my own type methods:")
	studyGroup := people{"Zuno", "Derren", "Jahn", "Albert", "Jenny", "Rachel"}
	sort.Sort(sort.Reverse(studyGroup)) // Using reverse interface
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)

	/*
		(2)
		s := []string{"Zeno", "John", "Al", "Jenny"}
	*/
	fmt.Println("\nQuestion 2 using StringSlice.")
	s := []string{"Zeno", "John", "Al", "Jenny"}
	// sort.Sort(sort.StringSlice(s))
	sort.Strings(s) // Shorthand for above.
	fmt.Println("Sorted:", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("Reversed sort:", s)

	/*
	  (3)
	  n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	*/
	fmt.Println("\nQuestion 3, int slice sorting:")
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	sort.Ints(n) // Shorthand for using IntSlice - look at doc.
	fmt.Println("Sorted using Ints, shorthand for IntSlice:", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("Reversed Ints:", n)
}
