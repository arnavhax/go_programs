package main

import (
	"fmt"
)
func main() {

	var farheniet,celcius float64
	fmt.Println("Enter value in celsius")
	fmt.Scanf( "%f",&celcius )
	// formula to convert farhenheit to celsius is Celsius = (Farheinet-32)/1.8
	farheniet= (celcius*1.8)+32
	fmt.Printf("Value in farhenheit is %.2f\n",farheniet)
}