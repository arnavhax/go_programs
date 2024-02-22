//if statement
// if condition {

//}

package main

import (
	"fmt"
	"os"
)

func main() {
	// if 201>12 {
	// 	fmt.Println("201 is greater than 12")
	// }

	// var a,b int=12,312

	// if b>a {
	// 	fmt.Println("b is greater than a")
	// }

	// if true {
	// 	fmt.Println("if true will always enter the if block")
	// }

	// if !(18>20) {
	// 	fmt.Println("not 20 is greater than 18")
	// }

	// initialise variables in if itself

	// x:=8
	// if y:=12;x<y{ // note that this initialisation of y is local to the if block only
	// 	fmt.Println("x is less than y")
	// }

	// if statement error handling using inline statement

	// var name string
	// var age int
	// fmt.Println("Enter name and age")
	// // we use _ here before , as we dont want to accept or store the return value of fmt.Scan() in if block
	// if _, err := fmt.Scan(&name,&age);err !=nil {
	// 	fmt.Println(err);
	// 	os.Exit(1)
	// }
	// fmt.Printf("Name is %s\n",name)
	// fmt.Printf("Age is %d\n",age)
	//if we want to store the return value of fmt.Scan we can use this code alternatively
	if vari, err := fmt.Scan(&name,&age);err !=nil {
		fmt.Println(vari,"Error code")
		fmt.Println(err);
		os.Exit(1)
	}
	fmt.Printf("Name is %s\n",name)
	fmt.Printf("Age is %d\n",age)

}