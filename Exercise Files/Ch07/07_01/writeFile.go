package main

import (
	"fmt"
	"os"
	io2 "io"
)

func main() {

	content := "hello world, this is best programming language nothignkfnkfd"

	file, err := os.Create("hello_world.txt")
	checkErrors(err)
	defer file.Close()

	io, err := io2.WriteString(file, content)

	checkErrors(err)
	fmt.Printf("create a file and append somethig %v", io)

}

func checkErrors(err error) {

	if err != nil {
		panic(err)
	}

}
