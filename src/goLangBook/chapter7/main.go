package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, z float64
}

type Person struct {
	Name string
}

func main() {
	c := Circle{x: 90, y: 10, z: 5}
	fmt.Println("mohan krishna reddy")
	fmt.Println(Area(c))

	p := Person{"mohan"}
	p.Speak()

	a :=  new(Android)
	a.Person.Talk()
}

func Area(c Circle) float64 {

	return math.Pi * c.x * c.y

}

func (p Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p Person) Speak() {
	fmt.Println("something to speak", p.Name)

}

type Android struct {
	Circle
	Person
}

type Shape interface {
	area() float64
}

//func totalArea(shapes Shape) float64{
//	var area float64
//	for _, s := range shapes {
//		area += s.area()
//	}
//	return area
//}