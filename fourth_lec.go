package main

import "fmt"
// out_var:=10 -- cannot declare like this outside function
var b = 10
func main() {
	// fmt.Println(out_var)
	fmt.Println(b)
	//multiple assign values of same datatypes
	var l,m,n,o int = 1,2,3,4
	fmt.Printf("data types are %T %T %T %T\n",l,m,n,o)
	//multiple assign values of diffrent datatypes
	var a,b,c,d =1,"2",3.1,4.112
	fmt.Printf("data types are %T %T %T %T\n",a,b,c,d)
	//when type keyword is not specified, you can declare diffrent types of variables in the same line

	// := multiple ele assignment

	intvar,strvar:=1,"arnav"
	fmt.Println(intvar,strvar)

	//declaring varibles in a block of diffrent types

	var(
		first int
		second int =1
		third string ="hello"
	)
	fmt.Println(first,second,third)

	fmt.Printf("%d\n%d\n%s\n",first,second,third)

}