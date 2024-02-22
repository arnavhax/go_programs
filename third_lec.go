package main

import "fmt"
//Global constant which can be used anywhere and is immutable
//untyped constant
const PI=3.14
const pi=3.14
//typed constant
const A int=20
//typed and untyped declaration of constants
const(
	L int =1
	M=3.14
	N="arnav"
)
func main() {
	fmt.Println(PI)
	// PI=PI+20
	fmt.Println(PI)
	fmt.Println(pi)
	fmt.Println(L)
	fmt.Println(M)
	fmt.Println(N)

	const LENGTH int =10
	const WIDTH int =20
	var area int
	area=LENGTH*WIDTH
	fmt.Printf("The Area of rectangle is %d\n",area)
	//int is not a keyword
	// var int int
	// int=10
	// fmt.Println(int)
	//case sensitive const is a keyword but still accepts Const
	const Const =10.2
	fmt.Println(Const)
	//uint range 0 to 255
	// var a uint8=300
	// fmt.Println(a)

	var student1 string ="jane"
	var student2 ="john"//type is infered 
	x:=9.1//type is infered
	//three ways of printing a string
	fmt.Printf("name of first student is %s\n",student1)//printf means we spefify format
	fmt.Printf("hello %s\n",student2)//%d is a integer number
	fmt.Printf("X is %T",x)// %T prints type


	// not definig value garbage collection demo

	var a string 
	var b int
	var c bool 
	//default value of bool is false
	// c=true
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	


}