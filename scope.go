package main

import "fmt"
// declare global variables before main function

var sum int

func addNumbers(){
	//local variables
	sum=12+21

}
func main() {
	addNumbers()
	//cannot access sum in main
	fmt.Println("Sum is",sum)
}