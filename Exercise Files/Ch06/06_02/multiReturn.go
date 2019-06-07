package main

import "fmt"

func main() {

	s,l := fullStringReturn("mohan","bade")
	fmt.Printf("fomratted string: %v,%v\n", s,l)
}

func fullStringReturn(f, l string) ( fullName string, length int) {

	fullName = f + " " + l

	length = len(fullName)

	return

}
