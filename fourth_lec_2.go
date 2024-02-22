package main

// import multiple libraries

import (
	"fmt"
)

func main() {
	//accepting string input
	// var name string
	// fmt.Println("Enter your name")
	// fmt.Scanf("%s",&name)
	// fmt.Printf("Your name is:%s\n",name)

	// var name1,name2 string
	// taking first name and lastname in second line
	// fmt.Println("Enter yout first name and last name")
	// fmt.Scanf("%s %s",&name1,&name2)
	// fmt.Printf("Your name is:%s %s\n",name1,name2)

	//taking input usinf scanln

	// var first string

	// fmt.Println("Enter first string")
	// fmt.Scanln(&first)
	// fmt.Println("Enter second string")
	// var second string
	// fmt.Scanln(&second)
	// fmt.Print(first+" "+second+"\n")

	//addition in printf
	// var a, b int = 10, 20
	// fmt.Printf("%d %d\n", a, b)
	// fmt.Printf("%d\n", a+b)

	const name, dept = "ghello", "CS"
	fmt.Printf("%s is a %s portal\n", name, dept)

	//get binary value using printf
	var test int = 100
	fmt.Printf("This is binary of %d is %b\n", test, test)
	var test2 float32 = 0.201
	//get floating point value using printf
	fmt.Printf("This is floating point of %f is %g\n", test2, test2)

	//if you dont know var type you can use %v
	fmt.Printf("%s is a %s portal\n", name, dept)
	
	//complex numbers in go

	var comp = 3+9i
	fmt.Printf("Value of complex number is %e",comp)


}
