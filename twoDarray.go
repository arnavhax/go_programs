package main

import (
	"fmt"
)

func main() {
	var matrix [3][3]int

	fmt.Printf("enter Matrix Elements: \n")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Enter matrix element %d%d\n", i, j)
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	fmt.Printf("Matrix is\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
	var leftSum int = 0
	fmt.Printf("Left Diagonal is\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == j {
				leftSum += matrix[i][j]
				fmt.Printf("%d ", matrix[i][j])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Left Diagonal sum is %d\n", leftSum)
	var rightSum int = 0
	fmt.Printf("Right Diagonal is\n")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i+j == 2 {
				rightSum += matrix[i][j]
				fmt.Printf("%d ", matrix[i][j])
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Printf("Right Diagonal sum is%d\n", rightSum)
}

// add two matrix assignment
// fibonacci using goto
// power using go to
// check palindrome using go to
// prime no or not using go to and loop
// implement binary search using go

