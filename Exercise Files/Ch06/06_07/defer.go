package main

import "fmt"

func main() {

	fmt.Println("first one")
	 defer fmt.Println("last one")

	fmt.Println("middle one")

}

func myFunc() {

}
