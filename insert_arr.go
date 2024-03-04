package main

import(
	"fmt"
)
func main(){
	originalArr:=[] int{1,2,4,5,6,7}
	
	newItem:=3
	index:=2
	newArray:=insertIntoArray(originalArr,newItem,index)

	fmt.Println("Original Array: ", originalArr)
	fmt.Println("New Array: ", newArray)

}

// The `insertIntoArray` function inserts an element into a specified index of an integer array in Go.
func insertIntoArray(arr[] int,item int,index int)[]int{
	newArr:=make([]int,len(arr)+1)

	// make function makes a slice of a given length

	copy(newArr[:index],arr[:index])
	// copy copies into arguement 1 element of arguement 2, we have used array slicing here

	newArr[index]=item

	copy(newArr[index+1:],arr[index:])

	return newArr
}