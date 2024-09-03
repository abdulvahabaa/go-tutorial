package main

import "fmt"

func main() {
	// greeter()
	// fmt.Println("Welcome to funcions in golang")

	// greeter()

	// greeterTwo() 

	result := adder(3,4)

	fmt.Println("Result is:",result)

	proRes := proAdder(1,2,3,4,5,6,7,8,9)

	fmt.Println("proRes is:",proRes)	



}

// func greeterTwo(){
// 	fmt.Println("Another method")
// }

// func greeter(){
// 	fmt.Println("Namastey from Golang")
// }

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) int {

	total := 0

	for _, value := range values {
		total += value
	}

	return total,"Hi pro result fuction"
}


	