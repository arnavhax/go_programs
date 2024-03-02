package main

import "fmt"

func displayNumbers() func() int {
	number := 0

	return func() int {
		number += 1
		return number
	}
}
func main() {
	fmt.Println(displayNumbers()())
	fmt.Println(displayNumbers()())

	num2 := displayNumbers()
	fmt.Println(num2())
	fmt.Println(num2())
}
