package main

import "fmt"

func main() {
	var name, rollNo, test1, test2, test3, total, percentage = "Arnav Balpande", "31", 78, 87, 92, 0, 0
	var result bool = true
	fmt.Printf("Name: %s\t Roll No: %s\n", name, rollNo)
	total = test1 + test2 + test3
	percentage = total / 3
	fmt.Printf("Test 1: %d\n", test1)
	fmt.Printf("Test 2: %d\n", test2)
	fmt.Printf("Test 3: %d\n", test3)
	fmt.Printf("Total Marks: %d ", total)
	fmt.Printf("Percentage: %d%% ", percentage)
	if test1 < 40 || test2 < 40 || test3 < 40 {
		result = false
	}
	if result == false {
		fmt.Printf("Result:FAIL\n")
	} else {
		fmt.Printf("Result:PASS\n")
	}
}
