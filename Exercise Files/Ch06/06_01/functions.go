package main

import "fmt"

func main() {

	sum :=addAllValues(23,45,67,89)

	fmt.Println(sum)

}

func addAllValues( values ...int)int  {

	sum := 0

	for i := range values{
		sum+= values[i]
	}

	return sum
}
