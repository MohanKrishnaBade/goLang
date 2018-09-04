package main

import "fmt"

func main() {

	x := []int{1, 2, 3, 4, 5, 6, 6}
	fmt.Println(sum(x))
	fmt.Println(half(1))
	fmt.Println(variadic(1, 2, 4, 5, 6, 7, 70, 9))

	//fmt.Println(makeOddGenerator(2))
	fmt.Println(fib(10))

	p := 1.5
	//q := 9

	//fmt.Println(swap(&p, &q))

	fmt.Println(test(&p))
}

func sum(x []int) (y int) {

	for _, value := range x {
		y += value
	}
	return
}

func half(x int) (int, bool) {

	if x%2 == 0 {
		return 1, true
	}

	return 0, false
}

//variadic  function and parameter example

func variadic(p ...int) (x int) {

	for _, value := range p {
		if value > x {
			x = value
		}
	}
	return
}

func makeOddGenerator(r int) func() int {

	x := func() int {

		if r%2 == 0 {
			return r + 3
		}
		return r + 2
	}

	return x
}

//caliculatinf febabanci number
// fib returns a function that returns
// successive Fibonacci numbers.
//func fib() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}

func swap(x, y *int) (int, int) {

	return *y, *x

}

func test(x *float64) float64 {
	*x = *x * *x

	return *x
}

func fib(number int) int {
	if number == 0 || number == 1{
		return number
	}

	return fib(number - 2) + fib(number - 1)
}