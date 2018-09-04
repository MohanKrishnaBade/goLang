package main

import (
	"fmt"
	"math"
)

func main() {

	c := Circle1{3}
	r := Rectangle{2,5}

	measure(c)
	measure(r)
}

type geometry interface {
	area() float64
	perim() float64
}

type Circle1 struct {
	radius float64
}
type Rectangle struct {
	width, height float64
}

//circle methods
func (c Circle1) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle1) perim() float64 {
	return 2 * math.Pi * c.radius
}

//rectangle methods.

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func (r Rectangle) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
