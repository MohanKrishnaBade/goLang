package main

import (
	"path/filepath"
	"fmt"
)

func main() {

	root, _ := filepath.Abs(".")

	fmt.Printf("file patth is: %v", root)
}
