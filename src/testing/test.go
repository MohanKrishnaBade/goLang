package main

import (
	"fmt"
)

var c, java, php bool = true,false,true


func add(a, b int) int {

	return a + b;
}

func swap(x, y string) (string, string) {

	return y, x;
}

func addSub(x, y int) (a, b int) {
	a = x + y
	b = x - y
	return
}

func loop()(int)  {
	x := 10

	//it act's like while loop
	for x<=30 {
		x += x;

	}

	return x
}

//no while loops in go, while also spelled as for.

func tryif(x int)(int) {

	if x <10 {
		return x
	} else{
		return 0
	}
}

func main() {
	i := 2
	defer fmt.Println("deffferered value")
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	s := string{1,2,3}

	defer fmt.Print(s.x)

	a := [6]int{1,23,4,56,6}
	fmt.Println(a)

}

type string struct {
	x int
	y int
	z int
}




