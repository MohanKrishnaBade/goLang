package main

import (
	"fmt"
	"sort"
)

func main() {

	var int = []int{1,209,3,4,5,6,56}

	fmt.Println(len(int))

	fmt.Println(append(int[5:]))

	fmt.Println(append(int[:len(int)-1]))

	sort.Ints(int)
	fmt.Println(int)



}