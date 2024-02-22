package main

import (
	"fmt"
)
// recursive function in go

func recurse(number int){
	// display the number
	if(number < 0){
		fmt.Println("CountDown Stops")
		return
	}
	fmt.Println(number)

	recurse(number-1)
}
func main() {
	fmt.Println("Enter Countdown timing:")
	var countdown int
	fmt.Scan(&countdown)
	recurse(countdown)

}