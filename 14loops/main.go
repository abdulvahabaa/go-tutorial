package main

import "fmt"

func main() {
	fmt.Println("Welocme to loops in golang")

	days :=[]string{"Sunday","Tuesday","Wednesday","Friday","Saturday"}

	fmt.Println(days)

	// for i:=0;i<len(days);i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

// 	for index,day := range days{
// 		fmt.Printf("Index is %v and vlaue is %v\n", index,day) //
// 	}

	// for _,day := range days{
	// 	fmt.Printf("Index is and vlaue is %v\n",day) 
	// }

	rougeValue:=0

	for rougeValue <10{

		if rougeValue ==2{
			goto codyinger 
		}

		if rougeValue == 5{
			// break

			rougeValue++
			continue
		}
		fmt.Println("Value is: ",rougeValue)
		rougeValue++
	}

	codyinger:
	fmt.Println("Jumping at learn golang ")
}