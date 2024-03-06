package main

import (
	"fmt"
	"sort"
)

func main() {
	// primeNumbers:=[] int {2,3,5,7}

	// numbers:= [] int{1,2,3}
	// 	// source destination
	// copy(numbers,primeNumbers)
	// // copies 3 elements only from primenumbers since size of numbers is 3
	// fmt.Print(numbers)


	// numbers:=[]int{1,2,3,4,5,6,7}

	// for i:=0;i<len(numbers);i++ {
	// 	fmt.Printf("%d\n",numbers[i])  // traversing elements in slice
	// }

	// traversing slice using range()

	// This code snippet in Go is demonstrating the concept of slices in Go. Here's a breakdown of what it
	// does:
	// arr:=[5] int{2,3,4,5,6}
	// //array with fixed length but slice is flexibl e

	// intSlice:=arr[1:4]

	// fmt.Println("Slice Elements")
	// for index,ele :=range intSlice[1:] {// iteration using range function
	// 	fmt.Printf("Element at %d is %d\n",index,ele)
	// }

	slice:= [] string{"I","Am","Superman"}
	sort.Strings(slice)

	fmt.Println("Sorted String slice")

	for _,item:=range slice{
		fmt.Printf("%s\n",item)
	}



}