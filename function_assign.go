package main

import (
	"fmt"
)

// func test(x int) {
// 	fmt.Println("hello", x)
// }
// test:= func(x int){
// 	fmt.Println(x)
// }
// anonymous function cannot be declared outside main or any other function
func main() {
	//assigning function to a variable
	// x := test
	// x(5)

	//anonymous funtion creation insinde a function 
	// test:= func(x int){
	// 	fmt.Println(x)
	// }
	// test(5)

	test:= func(x int)int {
		return x*x
	}(8)// this is sent as a arguement to the function statically
	fmt.Println(test)


}
