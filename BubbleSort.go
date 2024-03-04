package main

import (
	"fmt"
	// "sort"
)

func main() {
	arr := []int{31, 3, 31, 4, 1, 4, 2, 46, 6}
	// bubble sort logic
	// for i:=0;i<len(arr);i++{
	// 	for j:=0;j<len(arr)-i-1;j++{
	// 		if(arr[j]>arr[j+1]){
	// 			temp:=arr[j]
	// 			arr[j]=arr[j+1]
	// 			arr[j+1]=temp
	// 		}
	// 	}
	// }

	// alternatively we can use sort function
	// `sort.Ints(arr)` is a function call that sorts the elements of the integer slice `arr` in ascending
	// order. It uses the quicksort algorithm to efficiently rearrange the elements.
	// sort.Ints(arr)
	// // can also sort Float64,String and etc.
	// // assignment to sort in descending order
	// fmt.Println(arr)
	
	for i := 0; i < len(arr);i++ {
		for j:=i+1;i<len(arr);j++ {
			if(arr[i]==arr[j]){
				fmt.Printf("%d is the repeated element\n",arr[i])
				return
			}
		}
	}
	fmt.Println("No duplicate number found")

}
