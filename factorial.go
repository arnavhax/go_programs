package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}
func main() {
	fmt.Println("Enter number to find factorial of")
	var n int
	fmt.Scan(&n)
	fmt.Printf("Factorial of %d is %d\n", n, fact(n))
}
