package main

import (
	"fmt"
)
const(
	PI=3.14
)
func main() {
	var radius,area,circumference float32
	fmt.Println("Enter value of radius")
	fmt.Scanf("%f",&radius)
	area=radius*radius*PI
	circumference=PI*2*radius
	fmt.Printf("Area of circle is %.2f\n",area)
	fmt.Printf("Circumference of circle is %.2f\n",circumference)
}