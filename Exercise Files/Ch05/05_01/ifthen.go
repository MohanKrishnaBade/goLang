package main

import "fmt"

func main() {


	if x := -20; x > 0 {
		fmt.Println("greater than zero")
	} else if x == 0 {
		fmt.Println("equal to zero")
	} else {
		fmt.Println("less than zero")
	}

}
