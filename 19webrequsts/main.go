package main

import (
	"fmt"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	fmt.Println("Hello, World!")

	respone, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	// fmt.Println("Response is of type: %T\n", respone)

	defer respone.Body.Close()
}