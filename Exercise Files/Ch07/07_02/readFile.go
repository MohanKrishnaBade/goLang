package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	fileName := "./hello_world.txt"

	read, err := ioutil.ReadFile(fileName)
	checkError(err)
	stringConersion := string(read)
	fmt.Println("file content", stringConersion)

}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
