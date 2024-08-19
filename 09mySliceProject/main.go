package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welocme to class of slices")

	// var fruitList = []string{}
	var fruitList = []string{"Apple","Tomato","Peach"}
	fmt.Printf("Type of fruitlist  is %T\n",fruitList)

	fruitList=append(fruitList, "Mango","Banana") // add values to array
	fmt.Println(fruitList)
	
	// delete values

	// fruitList=append(fruitList[1:]) // remove first element(0th index) form the array
	// fruitList=append(fruitList[:3]) // remove  element form the array after 3rd elemnet
	fruitList=append(fruitList[1:3]) // only shows second and 3 rd elemnt form the areey
	fmt.Println(fruitList)

	highScores := make([]int,4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	// highScores[4] = 777 //shows error : panic: runtime error: index out of range [4] with length 4

	highScores = append(highScores, 555,666,321)

	fmt.Println(highScores)
	
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	
	// how to remove a vlaue form slices based on index
	var courses = []string{"reactjs","javascript","swift","python","ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

	
}