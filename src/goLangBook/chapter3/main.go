package main

import "fmt"

func main() {

	x, y := twoReturnValues()

	fmt.Println(x, y)
	defer fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9))
	fmt.Println(closure())
	fmt.Println(factorial(5))
	fmt.Println(makeEvenGenerator())


	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("mohan go and sleep")
}

func twoReturnValues() (int, int) {

	return 1, 2
}

//variadic functions

func add(args ...int) (y int) {

	for _, value := range args {
		y += value
	}

	return
}

//writing a function inside a function is  called closure.

func closure() (x int) {

	increment := func() (x int) {
		x++
		return
	}
	return increment()
}

//not working for some reasons.
func makeEvenGenerator() func() uint {

	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func factorial(x uint) uint {

	if x == 0 {
		return 1
	}

	return x * factorial(x-1)
}
