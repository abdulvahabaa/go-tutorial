package main

import "fmt"

// jwtToken :=300000 //volorous operator only accept inside the method
// var jwtToken =300000 //valid
// var jwtToken int =300000 //valid

const LoginToken string = "lskdfgoisdglsdkf" //public


func main() {
	fmt.Println("variables")

	var username  string ="vahab"
	fmt.Println(username)
	fmt.Printf("Variale is of type: %T \n",username)

	// var isLoggedin  bool =false
	var isLoggedin  bool =true
	fmt.Println(isLoggedin)
	fmt.Printf("Variale is of type: %T \n",isLoggedin)

	// var smallVal  uint8 =256
	var smallVal  uint8 =255
	fmt.Println(smallVal)
	fmt.Printf("Variale is of type: %T \n",smallVal)

	// var smallFloat  float32 =255.34565655745747445365
	var smallFloat  float64 =255.34565655745747445365
	fmt.Println(smallFloat)
	fmt.Printf("Variale is of type: %T \n",smallFloat)

	// default values some aliases
	var anoterVariable  int 
	fmt.Println(anoterVariable)
	fmt.Printf("Variale is of type: %T \n",anoterVariable)

	// implicit type

	var website = "davinchicode.in"
	fmt.Println(website) 

	// no var style

	// numberOfUsers := 300000
	numberOfUsers := 300000.0
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variale is of type: %T \n",LoginToken)


}