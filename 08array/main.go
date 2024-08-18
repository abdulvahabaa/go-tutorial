package main

import "fmt"

func main() {
	fmt.Println("Welocome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit list is: ", len(fruitList))

	var vegList = [3]string {"potato","carrot","onion"}
	fmt.Println("Vegy list",vegList)
	fmt.Println("Vegy list",len(vegList))
}