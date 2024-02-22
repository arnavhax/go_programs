package main
import (
	"fmt"
	// "unsafe"
	"reflect"
)
func main() {
	//arithimatic operators +,-,*,/,%,++,--
	
	//assignment operators += -= *= /=

	// var num int=10

	// num+=1
	// fmt.Printf( "Value of num after adding 1 is %d\n",num )
	// num-=1
	// fmt.Printf( "Value of num after subtracting 1 is %d\n",num )
	// num*=2
	// fmt.Printf( "Value of num after multiplying 2 is %d\n",num )
	// num/=2
	// fmt.Printf( "Value of num after dividing by 2 is %d\n",num )

	// var a,b int
	// fmt.Printf("Enter a and b\n")
	// fmt.Scanf("%d %d\n",&a,&b)
	// fmt.Printf("A modulo B is %d\n",a%b)

	// binary operators

	// a:=5
	// b:=3
	// c:=a&b
	// fmt.Printf("And of %d and %d is %d\n",a,b,c)
	// c=a|b
	// fmt.Printf("Or of %d and %d is %d\n",a,b,c)

	//comparison operators

	// a:=10;

	// fmt.Println(a>16)
	// fmt.Println(a==10)
	// fmt.Println(a<112)

	//logical operators

	// b:=5
	// fmt.Println(a && b)
	// fmt.Println(a || 16)
	// fmt.Println(! a && 16)


	//sizeof operator is in the unsafe library

	var num1 int=10
	// var num2 byte=252
	// var float1 float32=2931.2311
	// var str string="rajesh"

	// fmt.Printf("size of %d is %d\n",num1,unsafe.Sizeof(num1))
	// fmt.Printf("size of %d is %d\n",num2,unsafe.Sizeof(num2))
	// fmt.Printf("size of %f is %d\n",float1,unsafe.Sizeof(float1))
	// fmt.Printf("size of %s is %d\n",str,unsafe.Sizeof(str))
	
	//typeof operator is in the reflect library

	// fmt.Println("type of num1 is",reflect.TypeOf(num1))
	// fmt.Println("type of num2 is",reflect.TypeOf(num2))
	// fmt.Println("type of float1 is",reflect.TypeOf(float1))
	// fmt.Println("type of str is",reflect.TypeOf(str))

	//valueof operator

	fmt.Println("Type of num1 is",reflect.ValueOf(num1).Kind())


}