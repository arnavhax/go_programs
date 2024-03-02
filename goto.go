// Goto is  used to navigate executinon inside a function

package main

import (
	"fmt"
)

func main() {
	i := 1

loop:// this is goto reference statement
	fmt.Println(i)
	i++
	if i <= 5 {
		goto loop
	}
	fmt.Println("Loop ends")
}
