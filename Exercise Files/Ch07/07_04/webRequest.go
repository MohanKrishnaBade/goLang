package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {

	url := "http://services.explorecalifornia.org/json/tours.php"

	resp, err := http.Get(url)
	errorCheck(err)

	fmt.Printf("response type %T\n", resp)

	defer resp.Body.Close()

	bites, err := ioutil.ReadAll(resp.Body)

	errorCheck(err)

	string := string(bites)

	fmt.Println(string)

}

func errorCheck(err error)  {
	if err != nil {
		panic(err)
	}
}
