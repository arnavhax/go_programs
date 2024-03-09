package main
import (
	"fmt"
)
func main() {
	numbers:=[]int {1,2,3,4,5}
				//  0 1 2 3 4
	fmt.Println("Sliced Numbers",numbers[0:3])
	fmt.Println("Sliced Numbers",numbers[:3])
	fmt.Println("Sliced Numbers",numbers[2:4])
	fmt.Println("Sliced Numbers",numbers[0:])

	// make function for slices

	var num =make([]int,3,5) // here 3 is length and 5 is capacity
	fmt.Println(num)
	fmt.Println(cap(num),len(num))
	
}