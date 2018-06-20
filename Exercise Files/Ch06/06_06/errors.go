package main

import (
	"os"
	"fmt"
	"errors"
)

func main() {
	f, err := os.Open("mohan.ext")

	//err != nil ? fmt.Println(f) : fmt.Println(err)

	myError := errors.New("my new error")
	if err == nil {
		fmt.Println(f)
	} else {
		fmt.Println(myError)
	}

	attaendabce := map[string]bool{
		"mohan": true,
		"likki": false}


		attend, ok := attaendabce["m"]

		if ok {
			fmt.Println(errors.New("key found"))
			fmt.Printf("%v", attend)
		} else{
			fmt.Println(errors.New("key not found"))
		}


}
