package main

import (
	"fmt"
)

func main() {

	states := make(map[string]string)

	states["AL"] = "Alabama"
	states["OH"] = "ohio"
	states["CA"] ="california"


	fmt.Println(states)
	delete(states,"OH")

	fmt.Println(states)


	for k,v := range states{
		fmt.Printf("%v => %v\n",k,v )
	}
}