package main

import (
	"fmt"
)

func calculate(x int, y int) (int, int) {
	return x + y, x - y
}
func main() {
	a := 10
	b := 20
	sum, diff := calculate(a, b)
	fmt.Printf("Sum is %d Diffrence is %d\n", sum, diff)
}
