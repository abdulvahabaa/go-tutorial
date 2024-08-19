package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritence in golang; No super or parent

	vahab :=User{"Vahab","vahab@go.dev",true,18}

	fmt.Println(vahab) // {Vahab vahab@go.dev true 18}
	fmt.Printf("vahab details are : %+v\n", vahab) // vahab details are : {Name:Vahab Email:vahab@go.dev Status:true Age:18}
	fmt.Printf("Name is %v and email is %v", vahab.Name,vahab.Email) // Name is Vahab and email is vahab@go.dev
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}