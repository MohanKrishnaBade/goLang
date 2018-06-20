package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	//dow := rand.Intn(6) + 1
	//fmt.Println("Day", dow)

	switch dow := rand.Intn(6) + 1; dow {

	case 4:
		fmt.Println("this is thirsday")
	fallthrough
	case 7:
		fmt.Println("this is sunday")
	fallthrough
	default:
		fmt.Println("this is some random weekday")
	}

	switch dow := rand.Intn(6) + 1; {
	case dow == 4:
		fmt.Println("this is thirsday")
	fallthrough
	case dow == 7:
		fmt.Println("this is sunday")
	default:
		fmt.Println("this is some random weekday")
	}
}
