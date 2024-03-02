package main

import (
	"fmt"
)

func main() {
	// var x interface{}
	var x interface{} = "RKNEC"

	switch i := x.(type) {
	case nil:
		fmt.Printf("type of i is %T\n", i)

	case int:
		fmt.Println("type int")
	case float64:
		fmt.Println("x is a float64")
	case func(int) float64:
		fmt.Println("x is a func(int)")
	case bool, string:
		fmt.Println("x is a bool or string")
	default:
		fmt.Println("dont know the type")
	}
}
