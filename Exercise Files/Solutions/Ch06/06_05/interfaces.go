package main

import "fmt"

type Animals interface {
	Speak() string
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
	return "Meow!"
}

type Cow struct {
}

func (c Cow) Speak() string {
	return "Moo!"
}

type Tree struct {
	
}

func main() {
	poodle := Animals(Dog{})
	fmt.Println(poodle.Speak())
	
	animals := []Animals{Dog{}, Cat{}, Cow{}}
	for _, animal := range animals {
		fmt.Println(animal.Speak())
	}
}
