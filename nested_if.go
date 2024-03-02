// if(boolean ex){
// 	if(boolean ex 2){
// 		//condition executes
// 	}
// }

//nested if else and else if

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	x:=8
// 	y:=8

// 	if(x<y){
// 		fmt.Println("Y is Greater than X")

// 	}else if(x>y){
// 		fmt.Println("Y is Less than X")
// 	}else{
// 		fmt.Println("X is equal to Y")
// 	}
// }



// time based greetings

package main

import (
	"fmt"
)

func main() {
	var time int
	fmt.Println("Enter time in  24 hrs format , example 8 am is 0800")
	fmt.Scanf("%d",&time)
	if(time <0 || time >2359){
		fmt.Println("Invalid time")
	}
	if(time <=1000){
		fmt.Println("Good Morning")
	}else if(time<=2000){
		fmt.Println("Good Day")
	}else{
		fmt.Println("Good Evening")
	}
}