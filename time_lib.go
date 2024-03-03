// time function for fetching current dates and times
package main

import (
	"fmt"
	"time"
)

func main() {
	// yyyy, mm, dd := time.Now().Date()

	// // fmt.Printf("Todays date is %d/%d/%d\n", dd, mm, yyyy)
	// currentDateTime := time.Now()
	// hour, min, sec := currentDateTime.Hour(), currentDateTime.Minute(), currentDateTime.Second()

	// fmt.Printf("Current DateTime is %d/%d/%d %d:%d:%d\n", dd, mm, yyyy, hour, min, sec)

	//sleep function in go

	fmt.Println("Hello World")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Good Morning")
	time.Sleep(5 * time.Second)
	fmt.Println("This is CS50")
}
