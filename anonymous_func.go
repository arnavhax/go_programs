package main

// use of anonymous function in go
import (
	"fmt"
)
func findSquare(num int)int {
	square:=num*num
	return square
}
func main() {
	// declaring and calling an anonymous function using a variable without paremeters
	// var greet=func(){
	// 	fmt.Println("Hello this is anonymous function")
	// }
	// welcome:=greet
	// // greet:=func(){
	// // 	fmt.Println("this is inferred anonymmous func")
	// // }
	// greet()
	// welcome()

	//calling anonymous function without parameters

	// var sum = func(n1 , n2 int) {//both arguments can be given in by writing once in the function
	// 	result := n1 + n2
	// 	fmt.Println("Sum is ", result)
	// }
	// sum(5, 10)
	// fmt.Printf("Type of function is %T\n", sum)


	// anonymous function with return type


	// var sum = func(n1 int, n2 int)int{
	// 	return n1+n2
	// }

	// result:=sum(6, 10)
	// fmt.Printf("result of sum is %d\n",result)

	//anonymous function as an arguement

	//only an anonymous function can be declared inside a function

	sum:= func(num1 int,num2 int)int{
		return num1 + num2
	}

	result:=findSquare(sum(5,1))// declared before main
	fmt.Printf("Result is %d\n", result)
}
