package main

import (
	"fmt"
	"time"
	"math/rand"
)

const name  = "mohan"
func main() {

	fmt.Println("main works fine")
	c := fanOut(boaring("mohan"), boaring("bigBoss"))

	for i := 0; i < 5; i++ {
		fmt.Println("You say:", <-c)
	}

	fmt.Println("You're boring; I'm leaving "+name)

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Monday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func boaring(msg string) <-chan string {
	c := make(chan string)

	go func() {

		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanOut(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()

	return c
}
