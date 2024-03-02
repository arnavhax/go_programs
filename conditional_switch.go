// package main

// import(
// 	"fmt"
// )

// func main() {
// 	x:=30
// 	y:=30

// 	switch {
// 	case x>y:
// 		fmt.Println("x is greater than y")

// 	case x<y:
// 		fmt.Println("y is greater than x")

// 	default:
// 		fmt.Println("x and y are equal")

// }
// }

//switch initialiser statement and fallthrough

package main

import (
	"fmt"
)

func main() {
	switch x:=3;x {// this is switch initialiser statement
	case 1:
		fmt.Println("1")
		// fallthrough
	case 3:
		fmt.Println("3")
		fallthrough// makes break void and directly executes next case
	case 5:
		fmt.Println("5")
	}
}
