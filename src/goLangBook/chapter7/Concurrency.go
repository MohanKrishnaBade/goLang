package main

import (
	"fmt"
	"time"
	"math/rand"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		time.Sleep(time.Second)
	}
}

func main() {
	c := make(chan string)
	go boring("boring!!!!", c)
	go tooBoring(c, "what happend")
	go toomuchBoring(c, "what are you doing")
	for i := 0; i < 12; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value.
	}
	fmt.Println("You're boring; I'm leaving.")
	fmt.Println(rand.Intn(1e3));
}

func boring(msg string, c chan string) {
	for i := 0; i <= 12; i++ {
		c <- fmt.Sprintf("%s %d", msg, i) // Expression to be sent can be any suitable value.
		time.Sleep(time.Duration(318) * time.Microsecond)
	}
}

func tooBoring(c chan string, str string) {
	for i := 0; i <= 12; i++ {
		c <- fmt.Sprintf("%s %s", "something coming from the channel::::", str)
		time.Sleep(time.Duration(318) * time.Microsecond)
	}
}

func toomuchBoring(c chan string, string string) {
	for i := 0; i <= 12; i++ {
		c <- fmt.Sprintf("%s %s", "something coming from the tooMuchBoring channel", string)
		time.Sleep(time.Duration(318) * time.Microsecond)
	}
}
