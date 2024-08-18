package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	creationDate := time.Date(2020, time.August, 10,23,23,0,0,time.UTC)
	fmt.Println(creationDate)
	fmt.Println(creationDate.Format("01-02-2006 Monday"))

}