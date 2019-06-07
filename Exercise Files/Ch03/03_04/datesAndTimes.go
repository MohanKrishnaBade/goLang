package main

import (
	"fmt"
	"time"
)

func main() {
	time := time.Date(2009,time.November,10,23,0,0,0,time.UTC)

	fmt.Printf("time : %v\n",time)

	now := time.UTC()

	fmt.Printf("now: %v\n", now)

	fmt.Printf("month is: %v\n", time.Month())
	fmt.Printf("day is: %v\n", time.Day())
	//fmt.Printf("date is: %v\n", time.Date())

	tomarrow := now.AddDate(1,4,5)

	fmt.Printf("tomarrow is %v, %v, %v, %v\n", tomarrow,tomarrow.Day(),tomarrow.Month(),tomarrow.Year())

}