package main

import "fmt"

type Dog struct {
	Breed string
	Weight int
	Sound string
}

//use functions as methods here..
func main() {

	values := Dog{"somethng",50,"wooo"}
	values.speak()
}

func (d Dog) speak()  {
	fmt.Println(d.Sound)
}