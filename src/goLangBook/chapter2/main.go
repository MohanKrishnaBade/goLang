
package calculations

import "fmt"

const (
	name   string = "mohan"
	age    int    = 27
	weight int    = 177
)

var (
	id      int    = 2
	lapType string = "mac"
)

func cal() {

	for i := 1; i <= 15; i++ {
		switch true {
		case multiplesOfFiveAndThree(i):
			fmt.Println("FIZZBUZZ")
		case multiplesOfThree(i):
			fmt.Println("FIZZ")
		case multipleOfFive(i):
			fmt.Println("BUZZ")
		}
	}
	fmt.Println(array())
	//dump := []int{1, 2,3}
	//arryData := make(map[int]int)
	//vgvgh := []int{}
	////var x [] int
	////copy(arryData, dump)
	//fmt.Println(arryData)
	//fmt.Println(vgvgh)

	y := make(map[string]string)
	y["mohan"] = "nothing to down"

	fmt.Println(y["mohan"])

	elements := map[string]int{
		"first":  12,
		"second": 3,
		"third":  23,
		"fourth": 25,
		"fifith": 28,
		"sixth":  30,
	}

	fmt.Println(elements["second"])

	slice := make([]int, 3, 9)
	fmt.Println(len(slice))

	something := [6]string{"a", "b", "c", "d", "e", "f"}

	foo := something[2:5]

	for _, value := range foo {

		fmt.Println(value)
	}

	f := []int{1, 2, 3, 4, 4, 5, 6, 6, 7, 7, 8, 8, 99, 90, 100, 98, 101, 120, 105}
	fmt.Println(maxNumber(f))
}

func multiplesOfThree(i int) bool {
	return i%3 == 0
}

func multipleOfFive(i int) bool {
	return i%5 == 0
}
func multiplesOfFiveAndThree(i int) bool {
	if multipleOfFive(i) && multiplesOfThree(i) {
		return true
	}
	return false
}

func array() float64 {

	x := []float64{10, 25, 33, 42, 5, 45}

	var total float64 = 0

	for _, value := range x {
		total += value
	}

	return total / float64(len(x))
}

func maxNumber(f []int) (y int) {

	x := 0
	for _, value := range f {

		if value > x {
			y = value
		}
		x = value
	}

	return
}
