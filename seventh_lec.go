package main
// golang program to demonstrate type conversion and assignment issues during assignment to diffrent types
import (
	"fmt"
)

func main() {
	var num1,num2 int

	num1=10
	num2=5

	//doesnt give correct results due to diff types
	// fmt.Printf("Average of %d and %d is %.2f\n", num1, num2,(num1+num2)/2)

	// use %d instead

	// fmt.Printf("Average of %d and %d is %d\n", num1, num2,(num1+num2)/2)



}