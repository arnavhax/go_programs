// package main

// import (
// 	"fmt"
// 	"time"// time package to fetch date
// )

// func main() {

// 	today:=time.Now()
// 	thisMonth := today.Month()

// 	switch thisMonth {
// 	case 1:
// 		fmt.Println("January")
// 	case 2:
// 		fmt.Println("February")
// 	case 3:
// 		fmt.Println("March")
// 	case 4:
// 		fmt.Println("April")
// 	case 5:
// 		fmt.Println("May")
// 	case 6:
// 		fmt.Println("June")
// 	default:
// 		fmt.Println("Unknown month")
// }
// }

//conditional switch

package main

import (
	"fmt"
)

func main() {
	thisMonth := 6

	switch thisMonth {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31 Days")// note that we dont require break in go to stop swithch execution
	case 4, 6, 9, 11:
		fmt.Println("30 Days")

	case 2:
		fmt.Println("28 or 29 days")
	}
}
