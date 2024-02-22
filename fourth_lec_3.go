//lvalues and rvalues

//literals are always r values
//lvalue expression refers to a memory location and may be used in both right and left hand sides

//rvalue refers to a data value that is stored in some memory location and only appear on th right side of a expression

//formats of integer literals
//  binary ,hexadecimal,decimal and octal

// hex format starts with 0x then it is refered as hexadecimal
//example 0xf

// octal starts with 0 0o or o0
// 0b is binary

// decimal starts without 0
package main

// if fmt not imported the println is a builtin function but we maybneed to import fmt for other operations like scanf

import (
	"fmt"
)

func main() {
	// println(15==017)
	// println(15==0xF)

	//only postincrement and post decrement are supported in go and preinc and predec not allowed
	// var a, b int = 1, 2
	// a++
	// b--
	// fmt.Printf("a is %d b is %d\n", a, b)

	//binary output
	// var prod int = 15
	// fmt.Printf("Binary of %d is %0b\n", prod, prod)

	// // octal output
	// fmt.Printf("Octal of %d is %b\n", prod, prod)
	//to print boolean value true or false %t,%T
	// fmt.Printf("Expression 15==0xf is %t\n",15==0xf)
	// //print type in case of boolean edpression
	// fmt.Printf("Expression 15==0xf is %T\n",15==0xf)

	//printing two values in the same format statement
	// var i,j string ="hello","world"
	// fmt.Printf("%s %s\n",i,j)

	// var l,m =20,"rajesh"

	// fmt.Print(l,m)

	//To print with quotes in the same format as previous statement
	// var test string="Arnav"
	// fmt.Printf("%#v\n",test)

	//   / can be used as an escape character for all chars except % and if you want to print % then add %% to example %s%%

	// var escape string = "hello"
	// fmt.Printf("'%s'\n", escape)

	var i = 15
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#X\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%02d\n", i)

}
