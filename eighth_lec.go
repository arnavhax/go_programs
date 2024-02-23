package main

import (
	"fmt"
)
func sum(number int)(int){
	if number<=0{
		return 0
	}
	return number+sum(number-1),
}

func main() {
	var num=874
	var result=sum(num)
	fmt.Println("Sum:", result)
}