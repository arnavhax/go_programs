//python like range function in go

package main

import (
	"fmt"
)

func main() {
	colors := [4]string{"red", "green", "blue", "yellow"}
	for index, val := range colors {
		fmt.Println(index, "-", val)
	}
	for _, val := range colors { //use _ if you dont want index or val while iterating range
		fmt.Println(val)
	}
}
