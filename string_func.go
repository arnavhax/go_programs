package main

import (
	"fmt"
	"strings" //used for all string functions
)

func main() {
	// var college_name string="Shri Ramdeobaba College of Engineering and Management"
	// // var college_type string="Management"
	// var college_type string="Arts"

	// if(strings.Contains(college_name, college_type)){
	// 	fmt.Printf("String (%s) contains sub-string(%s)\n",college_name, college_type)
	// }else{
	// 	fmt.Printf("String (%s) does not contain sub-string(%s)\n",college_name, college_type)
	// }

	// strings.ToUpper(str)- to capitalise strings
	// strings.ToLower(str)-

	var str string = "Hello World"
	var ind int
	ind = strings.Index(str, "W")
	fmt.Println("index is",ind)
	fmt.Println("Upper of string is",strings.ToUpper(str))
	fmt.Println("Lower of string is",strings.ToLower(str))

}
