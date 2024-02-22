package main
// this is a program to demonstrate definition of functions in go
import(
	"fmt"
	"math"
)
// always start with print keyword

// funciton without arguements

// func greet(){
// 	fmt.Println("Good Morning")
// }

// function returns nothing
// func addNumbers(){
// 	n1:=1;n2:=2;n3:=3
// 	sum:=n1+n2+n3

// 	fmt.Println("Sum :",sum)
// }

// function declaration with arguments
// func addNumbers(n1 int,n2 int){
// 	sum:=n1+n2
// 	fmt.Println("Sum :",sum)
// }

//function with return arguments

// func addNumbers(n1 int,n2 int)int{
// 	sum:=n1+n2
	

// 	// code after reuturn statements
// 	//return statement must be the last statement
// 	fmt.Printf("After return statement\n")
// 	return sum
// }

// function with two return arguments

func calculate(n1 int,n2 int)(int,int){
	sum:=n1+n2
	diff:=int(math.Abs(float64(n1)-float64(n2)))

	return sum,diff
}

func square(n1 int) int{
	square:=n1*n1
	return square
}

func main() {
	// fmt.Println("Enter two numbers to sum")
	// var a,b int
	// fmt.Scan(&a,&b)
	//pass variables rather than values
	//store return value in sumi
	// sumi:=addNumbers(a,b)
	// // addNumbers(12,21)
	// fmt.Println("Sum is ",sumi)

	// sum,diff:=calculate(a,b)

	// fmt.Println("Sum ",sum,"Diffrence ",diff)
	var inp int
	fmt.Println("Enter num to calculate square")
	fmt.Scan(&inp)
	fmt.Printf("Square of %d is %d \n",inp,square(inp))

}