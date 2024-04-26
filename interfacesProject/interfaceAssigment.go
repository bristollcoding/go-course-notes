package main

import "fmt"

// struct types
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sidelength float64
}

// interface type
type shape interface {
	getArea() float64
}

func interfaceMain() {

	t := triangle{base: 1, height: 1}
	s := square{sidelength: 1}

	printArea(t)
	printArea(s)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sidelength * s.sidelength
}
