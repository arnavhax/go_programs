package main

import (
	"fmt"
)

type AreaFunc func(int, int) int
type Rectangle struct {
	length  int
	breadth int
	color   string
	area    AreaFunc
}

func main() {
	result := Rectangle{
		length:  10,
		breadth: 20,
		color:   "Red",
		area: func(l int, b int) int {
			return l * b
		},
	}
	fmt.Println("Area of rectangle is ", result.area(result.length, result.breadth))
	fmt.Println("Color of rectangle is ", result.color)

}
