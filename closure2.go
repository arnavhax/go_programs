package main

import(
	"fmt"
)
// outer function

func greet() func() string {
	//varirable defined outside the inner function
	name := "John"

	//returns a nested anonymous function

	return func() string{
		name:="Hi "+name
		return name
	}
}

func main() {
	//call the outer function

	message:=greet()
	//call the inner function
	fmt.Println(message()) //prints Hi John
}