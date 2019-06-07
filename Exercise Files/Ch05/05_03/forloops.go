package main

import "fmt"

func main() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i <= 8; i++ {
		sum += i
		fmt.Println("some", sum)
	}

	for i := 0; i <= len(colors)-1; i++ {
		fmt.Println("colors:", colors[i])
	}

	for i := range colors{
		fmt.Println(colors[i])

	}

	sum = 1

	for sum < 1000 {
		sum += sum

		if sum >500 {
			break
		} else if sum >200 {
			goto endOfProgram
		}
	}

	fmt.Println(sum)

	endOfProgram : fmt.Println("end of the program")

}
