package main

import "fmt"

func main() {
	fmt.Println("Maps in goLang")

	languages :=make(map[string]string)

	languages["JS"]="JavaScript"
	languages["RB"]="Ruby"
	languages["PY"]="Python"

	fmt.Println("List of all languages: ",languages)
	fmt.Println("JS shorts for : ",languages["JS"])
	
	// delete(languages,"RB")
	fmt.Println("List of all languages: ",languages)

	// loops in golang
	for key,value := range languages {
		fmt.Printf("For key %v, value is %v\n ", key,value)
	}
	
	// comma ok syntax
	for _,value := range languages {
		fmt.Printf("For key v, value is %v\n ",value)
	}

}