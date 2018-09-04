package main

import "fmt"

func main() {

	x := 56

	fmt.Println(x)
	pointer(&x)
	fmt.Println(x)
}

//pointer example
func pointer(x *int)  {
	*x =234

}