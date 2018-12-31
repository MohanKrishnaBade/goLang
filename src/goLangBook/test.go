package main

import "d"
import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

)

// function to add an array of numbers.
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// writes the sum to the go routines.
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c1 := make(chan int)
	c2 := make(chan int)
	// spin up a goroutine.

	d := s[:len(s)/3];
	f := s[len(s)/3:];
	go sum(d, c1)
	// spin up a goroutine.
	go sum(f, c2)
	x, y := <-c1, <-c2 // receive from c1 aND C2

	fmt.Println(x, y, x+y)

	diff()

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:8889)/mohan")

	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT email, phone FROM data1 WHERE id>=1")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer rows.Close()

	for rows.Next() {

		var (
			email string
			phone int32
		)
		rows.Scan(&email, &phone);
		fmt.Println(email, phone)
	}

	db.Close();

	//time.Sleep(3000*time.Second)
}

//db, err := sql.Open("mysql", "<username>:<pw>@tcp(<HOST>:<port>)/<dbname>"

func diff() {
	for i := 1; i <= 10; i++ {
		go fmt.Println(i);
	}
}
