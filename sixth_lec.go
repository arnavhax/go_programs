//demo of io.WriteStrings

package main
import (
	"fmt"
	// "math"
	// "io"
	// "os"
)
func main() {
	// var dd int=20
	// var mm int=02
	// var yy int=2024
	// var str string
	// str=fmt.Sprintf("%02d-%02d-%04d\n",dd,mm,yy)
	// io.WriteString(os.Stdout,str)

	// var str1 string
	// //can use all three float string and integer
	// var integer1 int
	// fmt.Printf("Enter String: ")
	// fmt.Scan(&str1,&integer1)
	// fmt.Printf("Result: %s %d\n", str1,integer1) 

	// var a int=123
	// var b uint=0
	//assign the value of a to b
	//this is not possible in go as uint and int are not inter assignable

	// b=a

	//printing both values

	//this is not possible in go as uint and int are not inter assignable
	// a=b  not possible
	//b=a   not possible
	// we can typecast into same type for interassignment

	// b= uint(a)	
	// fmt.Printf("a=%d b=%d\n",a,b)

	// var a float64=39.123

	// fmt.Printf("%f",math.Abs(a))

	//typecast float32 to float64

	// var a float32=23.322
	// // fmt.Printf("%d",math.Sqrt(float64(a)))
	// r:=math.Sqrt(float64(a))
	// fmt.Printf("%.2f\n",r)

	// a:=23.2
	// fmt.Printf("%T\n",a)


	//Explicit Type conversions

	var x int =42

	var y float64=float64(x)

	var z uint=uint(y)

	fmt.Printf("Value of x is %d and type is %T\n",x,x)
	fmt.Printf("Value of y is %.2f and type is %T\n",y,y)
	fmt.Printf("Value of z is %d and type is %T\n",z,z)


}