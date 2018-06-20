package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "An implectly typed string"

	int1 := 56
	fmt.Printf("str1: %v:%T\n", str1, str1)

	fmt.Printf("int1: %v:%T\n", int1, int1)

	fmt.Println(strings.ToUpper(str1))
	fmt.Println(strings.Title(str1))

	string1 := "hello world"
	string2 := "HELLO WORLD"

	fmt.Println("Equal?", (string1 == string2))

	fmt.Println("non case sesisitive comparision:", strings.EqualFold(string1,string2))

	fmt.Println("substring check:", strings.Contains(str1,"imp"))
}
