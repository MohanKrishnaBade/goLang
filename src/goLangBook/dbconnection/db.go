package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main() {

	Db("select email, first_name from data1 where id>0", "mohan")
}

func Db(query string, dbName string) {

	var (
		userName   string = "root"
		passCode   string = "root"
		hostServer string = "127.0.0.1:8889"
	)

	//building a sql query to make a connection.
	var source string = userName + ":" + passCode + "@tcp(" + hostServer + ")/" + dbName;
	db, err := sql.Open("mysql", source);
	if err != nil {
		fmt.Println(err.Error());
	}

	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	defer rows.Close()

	for rows.Next() {

		var (
			email    string
			lastName string
		)
		rows.Scan(&email, &lastName);
		fmt.Println(email, lastName)
	}

	defer db.Close();
}
