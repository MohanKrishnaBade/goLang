package main

import "fmt"

type Dog struct {
}

type Animal interface {
	speak() string
}

func (d Dog) speak() string {

	return "woooh!"
}

type cow struct {
}

func (c cow) speak() string {

	return "amba!!!"
}

func main() {
	poodle := Dog{}
	fmt.Println(poodle)

	animals := []Animal{Dog{},cow{}}
	for _, animal := range animals {
		fmt.Println(animal.speak())
	}

}
