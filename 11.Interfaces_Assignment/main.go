//Write a program that creates two custom structs type called 'triangle' and 'rectangle' that implement a method called 'area' which returns the area of the shape.
//The rectange should have a method that takes two parameters for the length and width, while the triangle should have a method that takes one parameter for the base and one for the height.
//The rectangle should return the area as length * width, while the triangle should return the area as 0.5 * base * height.
//The triangle and rectangle sides are type float64.
//Add a shape interface that has a method called 'area' that returns a float64.

package main

import "fmt"

type triangle struct {
	base   float64
	height float64
}

type rectange struct {
	base   float64
	height float64
}

type shape interface {
	area() float64
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}
func (r rectange) area() float64 {
	return r.base * r.height
}
func main() {
	r := rectange{2, 3}
	t := triangle{2, 3}

	// Option 1: Call area method directly
	fmt.Println(r.area())
	fmt.Println(t.area())

	// Option 2: Use a function that accepts the shape interface
	// printArea(r)
	// printArea(t)
}

// Uncomment this function if you want to use Option 2
// func printArea(s shape) {
// 	fmt.Println(s.area())
// }
