package main

import (
	"fmt"
)

func main() {
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 1))
}

func CountDigits(i int) (count int) {
	for i != 0 {

		i = i / 10
		count = count + 1
	}
	return
}



