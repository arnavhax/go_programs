package main

import(
	"fmt"
)

func calculate() func() int{ //func()int is retun type of calculate as it rerturs a anonymous function
	//returns func() int function
	num:=1
	return func() int {
		num=num+2
		return num
	}
}
func main() {
	//call calculate in main

	odd:=calculate()
	// anonymous function can be used to retain value in multiple function calls
	fmt.Println(odd())
	fmt.Println(odd())
	fmt.Println(odd())
	//call calculate again
	// closure helps in data isolation and preserving the state of variables even after returning from functions
	odd2:=calculate()
	fmt.Println(odd2())
}