package main

import "fmt"
import "math"

func main() {
	//Declare a variable of type int and assign it the values

	var a,b int
	a=20
	b=10
	//adding two numbers

	var sum int
	sum=a+b
	fmt.Println("The value of sum is: ",sum)
	// product of two numbers
	var prod int
	prod=a*b
	fmt.Println("The value of prod is: ",prod)
	//divide by two numbers
	var div int
	div=a/b
	fmt.Println("The value of div is: ",div)
	// difference between two numbers
	var diff int
	diff=a-b
	fmt.Println("The value of diff is: ",diff)
	//modulo of two numbers
	var modulo int
	modulo=a%b
	fmt.Println("The value of modulo is: ",modulo)
	// exponent of two numbers
	var expo int
	expo=int(math.Pow(float64(a),float64(b)))
	fmt.Println("The value of exponent is: ",expo)
}