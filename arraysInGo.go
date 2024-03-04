package main
import(
	"fmt"
)

func main() {
	//declate array of type integer
	// defined size[size=5]
	// var arrayOfInt=[5]int{1,2,3,4,5}
	// fmt.Println(arrayOfInt)
	// fmt.Printf("%T",arrayOfInt)


	//array declaration and defination in go

	// var arrayOfIntegers[3] int

	// arrayOfIntegers[0] = 1
	// arrayOfIntegers[1] = 2
	// arrayOfIntegers[2] = 3

	// arrayOfIntegers:=[5] int{0:7,3:9} //initailise only specific indices
	// fmt.Println(arrayOfIntegers)

	// declare array of undefined size

	var arrayOfString = [...] string{"Hello","Rajesh"}
	arrayOfString[0]="Bye" //put into array of string
	fmt.Println(arrayOfString)

	// slide 282
}