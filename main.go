// Simple interface examples. Remember that Interfaces are a type.
package main

import (
	"fmt"
	"math"
)

// Square is the struct we are playing with.
type Square struct {
	side float64
}

// Circle is another struct that will be another Shape.
type Circle struct {
	radius float64
}

// Method on the Square struct.
func (z Square) area() float64 {
	return z.side * z.side
}

// Method on Circle struct.
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

// Shape is of type interface. Interface defines functionality.
// Both Circle and Square have an area() function that matches what it takes in and returns.
// takes no arguments and returns a float64. Thus they are Shapes.
type Shape interface {
	area() float64
}

// This function takes the interface Shape.
// Because Square has method area(). That means that Square is a Shape.
// Kind of makes sense!
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{12}
	c := Circle{5}
	// Check out the new power. Polymorphism.
	info(s) // The concrete type implementing the interface Shape is a Square
	info(c) // The concrete type implementing the interface Shape is a Circle
}
