package main

import (
	"fmt"
	"strings"
)

func main()  {
	callInMain()
}
func callInMain() {

	fmt.Println(strings.ToUpper("Hello world"))
}
