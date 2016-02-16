package main

import "fmt"

/*
// Simple Example
type foo int

func main() {
	var myAge foo // myAge is of type foo. Foo has underlying type int.
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge) // Output will be of type main.foo
}
*/

// Struct is a composite type.
// Each field in a struct is declared with known types.
// fields can have tags.
type person struct { // type has fields
	first string // field
	last  string // fields can have tags as shown below
	age   int    // Can add tags to fields. Useful for JSON.
}

// Showing how to do some simple inheritance
type doubleZero struct {
	person
	first         string
	licenseToKill bool
}

// Methods
// func (receiver) (name of Function) (return) {}
// receiver attaches this function to this type.
func (p person) fullName() string {
	return p.first + " " + p.last + " the regular human."
}

// Outer type method overwrites inner type.
func (dz doubleZero) fullName() string {
	return dz.first + " " + dz.last + " pretending to be regular."

}

func main() {
	p1 := person{"Andrew", "Jak", 22}
	p2 := doubleZero{
		person: person{
			first: "Jake",
			last:  "James",
			age:   19,
		},
		first:         "Bond",
		licenseToKill: true,
	}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age, p2.licenseToKill)
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	fmt.Println(p2.person.fullName()) // can still reach the inner type method.
}
