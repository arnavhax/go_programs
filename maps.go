package main

import (
	"fmt"
)

func main() {
	// creating a map
	subjectMarks := map[string]int{"maths": 80, "science": 90, "EVS": 100}

	fmt.Println(subjectMarks) // printing map

	// accessing the values of map

	fmt.Println("Marks in Maths", subjectMarks["maths"])

	subjectMarks["maths"] = 90 // update key value in map
	fmt.Println("Updated Marks in maths", subjectMarks["maths"])

	// access map using variables

	var sub1 string = "EVS"
	fmt.Println("Marks of sub 1 is", subjectMarks[sub1])

	// adding key value pairs to map
	fmt.Println(subjectMarks)
	subjectMarks["Hindi"] = 75
	fmt.Println(subjectMarks)

	// delete key value from map

	delete(subjectMarks, "Hindi")
	fmt.Println(subjectMarks)

	// iterating dictionary for both keys and values

	for subject, marks := range subjectMarks {
		fmt.Println("Subject", subject, "marks", marks)
	}

	// iterating keys only
	for subject, _ := range subjectMarks {
		fmt.Println("Subject", subject)
	}

	// iterating values only
	fmt.Println(subjectMarks)
	for _, marks := range subjectMarks {
		fmt.Println("Marks", marks)
	}

}
